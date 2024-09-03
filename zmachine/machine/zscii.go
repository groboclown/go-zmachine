// Performs ZSCII encoding and decoding.
package machine

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
)

// ZsciiChar defines a 10-bit ZSCII character.
//
// Transforming this character into Unicode requires knowledge of the story file, and can also depend upon the
// story version.
type ZsciiChar uint16

// ZsciiTranslation converts to and from ZSCII and Unicode.
type ZsciiTranslation interface {
	ZsciiToUnicode(out ZsciiChar) rune
	UnicodeToZscii(in rune) ZsciiChar
	InputToZscii(in UserInput) ZsciiChar
}

// Zscii allows for encoding and decoding ZSCII characters.
type Zscii interface {
	// DecodeZscii turns the z-character encoded text starting at the position, up to the maximum length byte count, into Unicode characters.
	//
	// Returns the position of the pointer after the last byte read for the string.
	DecodeString(memory MemoryData, pos AbsAddr, maxLength int) ([]rune, AbsAddr, error)

	// DecodeZscii turns the z-character encoded text starting at the position, up to the maximum length byte count, into ZSCII characters.
	//
	// Returns the position of the pointer after the last byte read for the string.
	DecodeZscii(memory MemoryData, pos AbsAddr, maxLength int) ([]ZsciiChar, AbsAddr, error)

	// EncodeZscii allows turning user-entered text into encoded ZSCII at the given memory address.
	EncodeZscii(text []ZsciiChar) ([]uint8, error)

	EncodeInput(special unicode.SpecialCase, in []UserInput) ([]uint8, error)
}

// zsciiV1 decodes v1 story files.
//
// It has no state.
type zsciiV1 struct{}

func NewZsciiV1() Zscii {
	return zsciiV1{}
}

func (z zsciiV1) DecodeString(memory MemoryData, pos AbsAddr, maxLength int) ([]rune, AbsAddr, error) {
	decoded, pos, err := z.DecodeZscii(memory, pos, maxLength)
	if err != nil {
		return nil, pos, err
	}
	return ZsciiToUnicodeString(decoded, stdZsciiTranslation{}), pos, nil
}

func (zsciiV1) DecodeZscii(memory MemoryData, pos AbsAddr, maxLength int) ([]ZsciiChar, AbsAddr, error) {
	top := pos + AbsAddr(maxLength)
	if top > AbsAddr(memory.Size()) {
		top = AbsAddr(memory.Size())
	}
	ret := make([]ZsciiChar, 0)
	shift := 0
	alpha := 0
	tenBitStatus := 0
	tenBit := 0
	idx := pos
	for ; idx < top; idx += 2 {
		w0, w1, w2, last := zsciiWords(memory.ByteAt(idx), memory.ByteAt(idx+1))
		if last {
			top = idx
		}
		for _, w := range []int{w0, w1, w2} {
			if tenBitStatus == 1 {
				tenBit = w << 5
				tenBitStatus = 2
				continue
			}
			if tenBitStatus == 2 {
				ret = append(ret, ZsciiChar(tenBit|w))
				tenBitStatus = 0
				continue
			}
			switch w {
			case 0:
				// Always printed as a space
				ret = append(ret, 32)
				shift = 0
			case 1:
				ret = append(ret, 13)
				shift = 0
			case 2:
				// Shift next character.
				shift = 1
			case 3:
				// Shift next character.
				shift = 2
			case 4:
				// Shift-lock
				shift = 0
				alpha = (alpha + 1) % 3
			case 5:
				// Shift-lock
				shift = 0
				alpha = (alpha + 2) % 3
			default:
				tAlpha := (alpha + shift) % 3
				shift = 0
				if tAlpha == 2 && w == 6 {
					tenBitStatus = 1
					continue
				}
				ret = append(ret, v1ConvTable[tAlpha][w])
			}
		}
	}
	return ret, idx, nil
}

func (zsciiV1) EncodeZscii(text []ZsciiChar) ([]uint8, error) {
	return encodeZscii(text, 6, v1ConvTable)
}

func (z zsciiV1) EncodeInput(special unicode.SpecialCase, in []UserInput) ([]uint8, error) {
	en, err := UserInputToZscii(special, in, stdZsciiTranslation{})
	if err != nil {
		return nil, err
	}
	return z.EncodeZscii(en)
}

// zsciiV2 decodes v2 story files.
type zsciiV2 struct {
	abbreviations [][]ZsciiChar
}

func NewZsciiV2(abbreviations [][]ZsciiChar) (Zscii, error) {
	if len(abbreviations) > 32 {
		return nil, errors.New("maximum of 32 abbreviations available for v2 stories")
	}
	return zsciiV2{abbreviations: abbreviations}, nil
}

func (z zsciiV2) DecodeString(memory MemoryData, pos AbsAddr, maxLength int) ([]rune, AbsAddr, error) {
	decoded, pos, err := z.DecodeZscii(memory, pos, maxLength)
	if err != nil {
		return nil, pos, err
	}
	return ZsciiToUnicodeString(decoded, stdZsciiTranslation{}), pos, nil
}

func (z zsciiV2) DecodeZscii(memory MemoryData, pos AbsAddr, maxLength int) ([]ZsciiChar, AbsAddr, error) {
	top := pos + AbsAddr(maxLength)
	if top > AbsAddr(memory.Size()) {
		top = AbsAddr(memory.Size())
	}
	ret := make([]ZsciiChar, 0)
	shift := 0
	alpha := 0
	abbrev := false
	tenBitStatus := 0
	tenBit := 0
	idx := pos
	for ; idx < top; idx += 2 {
		w0, w1, w2, last := zsciiWords(memory.ByteAt(idx), memory.ByteAt(idx+1))
		if last {
			top = idx
		}
		for _, w := range []int{w0, w1, w2} {
			if abbrev {
				if w >= len(z.abbreviations) {
					return nil, idx, fmt.Errorf("zscii requested out-of-range abbreviation %d", w)
				}
				ret = append(ret, z.abbreviations[w]...)
				abbrev = false
				continue
			}
			if tenBitStatus == 1 {
				tenBit = w << 5
				tenBitStatus = 2
				continue
			}
			if tenBitStatus == 2 {
				ret = append(ret, ZsciiChar(tenBit|w))
				tenBitStatus = 0
				continue
			}

			if w == 1 {
				abbrev = true
				continue
			}
			if w == 2 || w == 3 {
				// Shift next character.
				shift = w - 1
				continue
			}
			if w == 4 || w == 5 {
				// Shift-lock.
				alpha = (alpha + w - 3) % 3
				continue
			}
			tAlpha := (alpha + shift) % 3
			shift = 0

			if tAlpha == 2 && w == 6 {
				tenBitStatus = 1
				continue
			}
			ret = append(ret, v2ConvTable[tAlpha][w])
		}
	}
	return ret, pos, nil
}

func (zsciiV2) EncodeZscii(text []ZsciiChar) ([]uint8, error) {
	return encodeZscii(text, 6, v2ConvTable)
}

func (z zsciiV2) EncodeInput(special unicode.SpecialCase, in []UserInput) ([]uint8, error) {
	en, err := UserInputToZscii(special, in, stdZsciiTranslation{})
	if err != nil {
		return nil, err
	}
	return z.EncodeZscii(en)
}

// zsciiV3_plus decodes v3+ story files.
type zsciiV3_plus struct {
	abbreviations [][]ZsciiChar
	alphabetTable [][]ZsciiChar
	encodeLength  int
	txn           ZsciiTranslation
}

func NewZsciiV3(abbreviations [][]ZsciiChar) (Zscii, error) {
	if len(abbreviations) > 96 { // 3 * 32
		return nil, errors.New("maximum of 96 abbreviations available for v3+ stories")
	}
	return zsciiV3_plus{abbreviations: abbreviations, alphabetTable: v2ConvTable, encodeLength: 6, txn: NewZsciiTranslationV1_4()}, nil
}

func NewZsciiV4(abbreviations [][]ZsciiChar) (Zscii, error) {
	if len(abbreviations) > 96 { // 3 * 32
		return nil, errors.New("maximum of 96 abbreviations available for v3+ stories")
	}
	return zsciiV3_plus{abbreviations: abbreviations, alphabetTable: v2ConvTable, encodeLength: 9, txn: NewZsciiTranslationV1_4()}, nil
}

// NewZsciiV5_Plus uses a custom alphabet table for the story file.
//
// Word 0x34 in the header, if non-zero, points to the byte address of the alphabet table.
func NewZsciiV5_Plus(abbreviations [][]ZsciiChar, rawAlphabetTable []uint8, txn ZsciiTranslation) (Zscii, error) {
	if len(abbreviations) > 96 { // 3 * 32
		return nil, errors.New("maximum of 96 abbreviations available for v5+ stories")
	}
	if len(rawAlphabetTable) != 78 {
		return nil, errors.New("alphabet table must contain 78 octets")
	}
	// Turn the alphabet table into 3 sets of 32 values.
	alphabetTable := [][]ZsciiChar{make([]ZsciiChar, 32), make([]ZsciiChar, 32), make([]ZsciiChar, 32)}
	idx := 0
	for major := 0; major < 3; major++ {
		for pos := 6; pos < 32; pos++ {
			alphabetTable[major][pos] = ZsciiChar(rawAlphabetTable[idx])
			idx++
		}
	}
	alphabetTable[2][6] = 0  // must be a 10-bit encoding
	alphabetTable[2][7] = 13 // must be a newline

	return zsciiV3_plus{abbreviations: abbreviations, alphabetTable: alphabetTable, encodeLength: 9, txn: txn}, nil
}

func (z zsciiV3_plus) DecodeString(memory MemoryData, pos AbsAddr, maxLength int) ([]rune, AbsAddr, error) {
	decoded, pos, err := z.DecodeZscii(memory, pos, maxLength)
	if err != nil {
		return nil, pos, err
	}
	return ZsciiToUnicodeString(decoded, stdZsciiTranslation{}), pos, nil
}

func (z zsciiV3_plus) DecodeZscii(memory MemoryData, pos AbsAddr, maxLength int) ([]ZsciiChar, AbsAddr, error) {
	top := pos + AbsAddr(maxLength)
	if top > AbsAddr(memory.Size()) {
		top = AbsAddr(memory.Size())
	}
	ret := make([]ZsciiChar, 0)
	shift := 0
	alpha := 0
	abbrev := 0
	tenBitStatus := 0
	tenBit := 0
	idx := pos
	for ; idx < top; idx += 2 {
		w0, w1, w2, last := zsciiWords(memory.ByteAt(idx), memory.ByteAt(idx+1))
		if last {
			top = idx
		}
		for _, w := range []int{w0, w1, w2} {
			if abbrev > 0 {
				aIdx := (32 * (abbrev - 1)) + w
				if aIdx >= len(z.abbreviations) {
					return nil, pos, fmt.Errorf("zscii requested out-of-range abbreviation %d", aIdx)
				}
				ret = append(ret, z.abbreviations[aIdx]...)
				abbrev = 0
				continue
			}
			if tenBitStatus == 1 {
				tenBit = w << 5
				tenBitStatus = 2
				continue
			}
			if tenBitStatus == 2 {
				ret = append(ret, ZsciiChar(tenBit|w))
				tenBitStatus = 0
				continue
			}

			if w >= 1 && w <= 3 {
				shift = 0
				abbrev = w
				continue
			}
			if w == 4 || w == 5 {
				shift = w - 3
				continue
			}
			tAlpha := alpha + shift
			shift = 0

			if tAlpha == 2 && w == 6 {
				tenBitStatus = 1
				continue
			}
			ret = append(ret, z.alphabetTable[tAlpha][w])
		}
	}
	return ret, pos, nil
}

func (z zsciiV3_plus) EncodeZscii(text []ZsciiChar) ([]uint8, error) {
	return encodeZscii(text, z.encodeLength, z.alphabetTable)
}

func (z zsciiV3_plus) EncodeInput(special unicode.SpecialCase, in []UserInput) ([]uint8, error) {
	en, err := UserInputToZscii(special, in, z.txn)
	if err != nil {
		return nil, err
	}
	return z.EncodeZscii(en)
}

// zsciiWords decodes the 16-bit word into 3 z-characters + a marker for 'last character'.
func zsciiWords(b0, b1 uint8) (int, int, int, bool) {
	i0 := int(b0)
	i1 := int(b1)
	z0 := (i0 >> 2) & 0b0011111
	z1 := ((i0 & 0b00000011) << 3) | ((i1 >> 5) & 0b00000111)
	z2 := i1 & 0b00011111
	return z0, z1, z2, (b0 & 0b10000000) != 0
}

func encodeZscii(text []ZsciiChar, size int, convTable [][]ZsciiChar) ([]uint8, error) {
	// Total string length must be 6 or 9
	if size%3 != 0 {
		return nil, fmt.Errorf("encoded size must be a multiple of 3, found %d", size)
	}
	bare := make([]int, size)
	// Characters must be '5' padded.
	for i := 0; i < size; i++ {
		bare[i] = 5
	}

	// Turn the ZSCII characters into z-characters.
	for i := 0; i < size && i < len(text); i++ {
		c := text[i]

		// Text must be lower case.
		// This uses standard ASCII rules.
		if c >= 'A' && c <= 'Z' {
			c = c - 'A' + 'a'
		}

		// Characters allowed should only be from the first row - a-z ?
		found := 100
		for j := 0; j < len(convTable[0]); j++ {
			if convTable[0][j] == ZsciiChar(c) {
				found = j
				break
			}
		}
		if found == 100 {
			return nil, fmt.Errorf("encountered invalid character to encode: '%v' position %d", text, i)
		}
		bare[i] = found
	}

	ret := make([]uint8, size/3)
	idx := 0
	for i := 0; i < size; i += 3 {
		z0 := bare[i] & 0xff
		z1 := bare[i+1] & 0xff
		z2 := bare[i+2] & 0xff
		ret[idx] = uint8(((z0 << 2) & 0b01111100) | ((z1 >> 3) & 0b00000011))
		idx++
		ret[idx] = uint8(((z1 << 5) & 0b11100000) | (z2 & 0b00011111))
		idx++
	}
	return ret, nil
}

var v1ConvTable = [][]ZsciiChar{
	// A0 lookup
	{0, 0, 0, 0, 0, 0, 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'},
	// A1 lookup
	{0, 0, 0, 0, 0, 0, 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'},
	// A2 lookup
	{0, 0, 0, 0, 0, 0, 0, '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', ',', '!', '?', '_', '#', '\'', '"', '/', '\\', '<', '-', ':', '(', ')'},
}

var v2ConvTable = [][]ZsciiChar{
	// A0 lookup
	{0, 0, 0, 0, 0, 0, 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'},
	// A1 lookup
	{0, 0, 0, 0, 0, 0, 'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z'},
	// A2 lookup
	{0, 0, 0, 0, 0, 0, 0, 13, '0', '1', '2', '3', '4', '5', '6', '7', '8', '9', '.', ',', '!', '?', '_', '#', '\'', '"', '/', '\\', '-', ':', '(', ')'},
}

const NullChar = ZsciiChar(0)           // Output
const DelChar = ZsciiChar(8)            // Input
const TabChar = ZsciiChar(9)            // Output, For V6, if at the start of a line, it's an indented paragraph, middle of a screen line is a space.
const SentenceSpaceChar = ZsciiChar(11) // Output, V6 only - a long space between sentences.
const NewlineChar = ZsciiChar(13)       // Input & Output
const EscChar = ZsciiChar(27)           // Input
const UpChar = ZsciiChar(129)           // Input
const DownChar = ZsciiChar(130)         // Input
const LeftChar = ZsciiChar(131)         // Input
const RightChar = ZsciiChar(132)        // Input
const F1Char = ZsciiChar(133)           // Input
const F2Char = ZsciiChar(134)           // Input
const F3Char = ZsciiChar(135)           // Input
const F4Char = ZsciiChar(136)           // Input
const F5Char = ZsciiChar(137)           // Input
const F6Char = ZsciiChar(138)           // Input
const F7Char = ZsciiChar(139)           // Input
const F8Char = ZsciiChar(140)           // Input
const F9Char = ZsciiChar(141)           // Input
const F10Char = ZsciiChar(142)          // Input
const F11Char = ZsciiChar(143)          // Input
const F12Char = ZsciiChar(144)          // Input
const Keypad0Char = ZsciiChar(145)      // Input
const Keypad1Char = ZsciiChar(146)      // Input
const Keypad2Char = ZsciiChar(147)      // Input
const Keypad3Char = ZsciiChar(148)      // Input
const Keypad4Char = ZsciiChar(149)      // Input
const Keypad5Char = ZsciiChar(150)      // Input
const Keypad6Char = ZsciiChar(151)      // Input
const Keypad7Char = ZsciiChar(152)      // Input
const Keypad8Char = ZsciiChar(153)      // Input
const Keypad9Char = ZsciiChar(154)      // Input
const MenuClickChar = ZsciiChar(252)    // Input
const DoubleClickChar = ZsciiChar(253)  // Input
const SingleClickChar = ZsciiChar(254)  // Input

var ZsciiStdLookup = map[ZsciiChar]rune{
	NullChar:          0,
	TabChar:           9,
	SentenceSpaceChar: '\u2003',
	NewlineChar:       '\n',
	32:                ' ',
	33:                '!',
	34:                '"',
	35:                '#',
	36:                '$',
	37:                '%',
	38:                '&',
	39:                '\u2019', // '\'', right single quote
	40:                '(',
	41:                ')',
	42:                '*',
	43:                '+',
	44:                ',',
	45:                '-',
	46:                '.',
	47:                '/',
	48:                '0',
	49:                '1',
	50:                '2',
	51:                '3',
	52:                '4',
	53:                '5',
	54:                '6',
	55:                '7',
	56:                '8',
	57:                '9',
	58:                ':',
	59:                ';',
	60:                '<',
	61:                '=',
	62:                '>',
	63:                '?',
	64:                '@',
	65:                'A',
	66:                'B',
	67:                'C',
	68:                'D',
	69:                'E',
	70:                'F',
	71:                'G',
	72:                'H',
	73:                'I',
	74:                'J',
	75:                'K',
	76:                'L',
	77:                'M',
	78:                'N',
	79:                'O',
	80:                'P',
	81:                'Q',
	82:                'R',
	83:                'S',
	84:                'T',
	85:                'U',
	86:                'V',
	87:                'W',
	88:                'X',
	89:                'Y',
	90:                'Z',
	91:                '[',
	92:                '\\',
	93:                ']',
	94:                '^',
	95:                '_',
	96:                '\u2018', // '`', left single quote
	97:                'a',
	98:                'b',
	99:                'c',
	100:               'd',
	101:               'e',
	102:               'f',
	103:               'g',
	104:               'h',
	105:               'i',
	106:               'j',
	107:               'k',
	108:               'l',
	109:               'm',
	110:               'n',
	111:               'o',
	112:               'p',
	113:               'q',
	114:               'r',
	115:               's',
	116:               't',
	117:               'u',
	118:               'v',
	119:               'w',
	120:               'x',
	121:               'y',
	122:               'z',
	123:               '{',
	124:               '|',
	125:               '}',
	126:               '~',

	155: '\u00e4', // a-diaeresis
	156: '\u00f6', // o-diaeresis
	157: '\u00fc', // u-diaeresis
	158: '\u00c4', // A-diaeresis
	159: '\u00d6', // O-diaeresis
	160: '\u00dc', // U-diaeresis
	161: '\u00df', // sz-ligature
	162: '\u00bb', // quotation
	163: '\u00ab', // marks
	164: '\u00eb', // e-diaeresis
	165: '\u00ef', // i-diaeresis
	166: '\u00ff', // y-diaeresis
	167: '\u00cb', // E-diaeresis
	168: '\u00cf', // I-diaeresis
	169: '\u00e1', // a-acute
	170: '\u00e9', // e-acute
	171: '\u00ed', // i-acute
	172: '\u00f3', // o-acute
	173: '\u00fa', // u-acute
	174: '\u00fd', // y-acute
	175: '\u00c1', // A-acute
	176: '\u00c9', // E-acute
	177: '\u00cd', // I-acute
	178: '\u00d3', // O-acute
	179: '\u00da', // U-acute
	180: '\u00dd', // Y-acute
	181: '\u00e0', // a-grave
	182: '\u00e8', // e-grave
	183: '\u00ec', // i-grave
	184: '\u00f2', // o-grave
	185: '\u00f9', // u-grave
	186: '\u00c0', // A-grave
	187: '\u00c8', // E-grave
	188: '\u00cc', // I-grave
	189: '\u00d2', // O-grave
	190: '\u00d9', // U-grave
	191: '\u00e2', // a-circumflex
	192: '\u00ea', // e-circumflex
	193: '\u00ee', // i-circumflex
	194: '\u00f4', // o-circumflex
	195: '\u00fb', // u-circumflex
	196: '\u00c2', // A-circumflex
	197: '\u00ca', // E-circumflex
	198: '\u00ce', // I-circumflex
	199: '\u00d4', // O-circumflex
	200: '\u00db', // U-circumflex
	201: '\u00e5', // a-ring
	202: '\u00c5', // A-ring
	203: '\u00f8', // o-slash
	204: '\u00d8', // O-slash
	205: '\u00e3', // a-tilde
	206: '\u00f1', // n-tilde
	207: '\u00f5', // o-tilde
	208: '\u00c3', // A-tilde
	209: '\u00d1', // N-tilde
	210: '\u00d5', // O-tilde
	211: '\u00e6', // ae-ligature
	212: '\u00c6', // AE-ligature
	213: '\u00e7', // c-cedilla
	214: '\u00c7', // C-cedilla
	215: '\u00fe', // Icelandic thorn
	216: '\u00f0', // Icelandic eth
	217: '\u00de', // Icelandic Thorn
	218: '\u00d0', // Icelandic Eth
	219: '\u00a3', // pound symbol
	220: '\u0153', // oe-ligature
	221: '\u0152', // OE-ligature
	222: '\u00a1', // inverted !
	223: '\u00bf', // inverted ?
}

func ZsciiToUnicodeString(out []ZsciiChar, txn ZsciiTranslation) []rune {
	ret := make([]rune, 0, len(out))
	for _, z := range out {
		c := txn.ZsciiToUnicode(z)
		if c != 0 {
			ret = append(ret, c)
		}
	}
	return ret
}

func UnicodeStringToLowerZscii(special unicode.SpecialCase, text string, txn ZsciiTranslation) ([]ZsciiChar, error) {
	lower := strings.ToLowerSpecial(special, text)
	return UnicodeStringToZscii(lower, txn)
}

func UnicodeStringToZscii(text string, txn ZsciiTranslation) ([]ZsciiChar, error) {
	ret := make([]ZsciiChar, 0)
	for _, c := range text {
		z := txn.UnicodeToZscii(c)
		if z != NullChar {
			ret = append(ret, z)
		}
	}
	return ret, nil
}

func UserInputToZscii(special unicode.SpecialCase, input []UserInput, txn ZsciiTranslation) ([]ZsciiChar, error) {
	ret := make([]ZsciiChar, 0, len(input))
	for _, in := range input {
		lowered := UserInput{Ctrl: in.Ctrl, Key: in.Key}
		if in.Ctrl == 0 {
			lowered.Key = special.ToLower(lowered.Key)
		}
		z := txn.InputToZscii(lowered)
		if z != NullChar {
			ret = append(ret, z)
		}
	}
	return ret, nil
}

func NewZsciiTranslationV1_4() ZsciiTranslation {
	return stdZsciiTranslation{}
}

// NewZsciiTranslationV5Lookup returns a ZSCII to Unicode translation using the optional lookup table.
//
// In version 5 or later, if header word 3 is non-zero, then this is the byte address of the Unicode translation table.
// The first word is the number of 2-byte words in the translation table, and that is what must be passed in.
func NewZsciiTranslationV5Lookup(lookupRange []uint16) (ZsciiTranslation, error) {
	if len(lookupRange) <= 0 {
		return NewZsciiTranslationV1_4(), nil
	}
	ext := make(map[ZsciiChar]rune)
	idx := 155
	for _, uni := range lookupRange {
		ext[ZsciiChar(idx)] = rune(uni)
		idx++
	}
	if idx > 251 {
		return nil, errors.New("stories can have at most 97 custom unicode character lookups")
	}
	return translationTableZscii{ext}, nil
}

type stdZsciiTranslation struct{}

func (stdZsciiTranslation) ZsciiToUnicode(out ZsciiChar) rune {
	r, ok := ZsciiStdLookup[out]
	if !ok {
		return 0
	}
	return r
}

func (stdZsciiTranslation) UnicodeToZscii(in rune) ZsciiChar {
	return basicUnicodeToZscii(in)
}

func (stdZsciiTranslation) InputToZscii(in UserInput) ZsciiChar {
	if in.Key != 0 {
		return basicUnicodeToZscii(in.Key)
	}
	return in.Ctrl
}

type translationTableZscii struct {
	extendedTranslate map[ZsciiChar]rune
}

func (t translationTableZscii) ZsciiToUnicode(out ZsciiChar) rune {
	r, ok := t.extendedTranslate[out]
	if ok {
		return r
	}
	r, ok = ZsciiStdLookup[out]
	if !ok {
		return 0
	}
	return r
}

func (t translationTableZscii) UnicodeToZscii(in rune) ZsciiChar {
	// Ugh.
	for z, c := range t.extendedTranslate {
		if c == in {
			return z
		}
	}
	return basicUnicodeToZscii(in)
}

func (t translationTableZscii) InputToZscii(in UserInput) ZsciiChar {
	if in.Key != 0 {
		return t.UnicodeToZscii(in.Key)
	}
	return in.Ctrl
}

func basicUnicodeToZscii(in rune) ZsciiChar {
	// Ugh.
	for z, c := range ZsciiStdLookup {
		if c == in {
			return z
		}
	}
	// Other characters from the keyboard.
	switch in {
	case 10:
		return NewlineChar
	case '\'':
		return 39
	case '`':
		return 96
	}
	return NullChar
}

// DecodeAbbreviationsTable performs the necessary transformations on the ZSCII strings in the abbreviations table.
//
// It takes a simplified zscii decoder, which does not have any abbreviations in it.
func DecodeAbbreviationsTable(mem MemoryData, tablePos AbsAddr, maxEntries int, zscii Zscii) ([][]ZsciiChar, error) {
	// The abbreviations table is a list of "word addresses" to abbreviation strings.
	maxSize := AbsAddr(mem.Size())
	ret := make([][]ZsciiChar, 0, maxEntries)
	for i := 0; i < maxEntries; i++ {
		pos := asWordAddress(
			asWord(mem.ByteAt(tablePos), mem.ByteAt(tablePos+1)),
		)
		if pos >= maxSize {
			return nil, fmt.Errorf("abbreviation table (abbrev %d) points outside memory space (%x)", i, pos)
		}
		z, _, err := zscii.DecodeZscii(mem, pos, int(maxSize))
		if err != nil {
			return nil, err
		}
		ret = append(ret, z)
		tablePos += 2
	}
	return ret, nil
}
