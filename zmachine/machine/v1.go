// V1 story file specific information.
package machine

func Version1(mem VolatileMemoryData) Version {
	ops := assembleOpCodes(1)
	return &v1Version{
		mem:    mem,
		header: &v1Header{mem: mem, marked: mem},
		ops: &OpDecodeV1_4{
			variableOpcodes:  ops.variableOpcodes,
			shortOpcodes:     ops.shortOpcodes,
			longOpCodes:      ops.longOpCodes,
			doubleVarOpCodes: ops.doubleVarOpCodes,
			zscii:            NewZsciiV1(),
		},
	}
}

type v1Version struct {
	mem    VolatileMemoryData
	header *v1Header
	ops    OpDecode
}

func (v *v1Version) Header() Header {
	return v.header
}

func (v *v1Version) Opcodes() OpDecode {
	return v.ops
}

func (v *v1Version) InitialRoutineState() *RoutineCallState {
	return v1InitialRoutineState(v.mem)
}

// v1InitialRoutineState is an all-but-v6 routine state creation.
func v1InitialRoutineState(mem MemoryData) *RoutineCallState {
	startAddress := asJoinedByteAddress(mem.ByteAt(0x06), mem.ByteAt(0x07))
	return NewRoutineCallState(
		startAddress,

		// VM starting environment contains no local variables.
		make([]uint16, 0),
	)
}

// ---- V1 Header Implementation - provided things by the version.
type v1Header struct {
	mem    VolatileMemoryData
	marked MemoryData
}

func (v *v1Header) VersionNumber() int {
	return int(v.mem.ByteAt(0x00))
}

func (v *v1Header) HighMemoryBaseAddress() AbsAddr {
	return asJoinedByteAddress(v.mem.ByteAt(0x04), v.mem.ByteAt(0x05))
}

func (v *v1Header) DictionaryAddress() AbsAddr {
	return asJoinedByteAddress(v.mem.ByteAt(0x08), v.mem.ByteAt(0x09))
}

func (v *v1Header) ObjectTableAddress() AbsAddr {
	return asJoinedByteAddress(v.mem.ByteAt(0x0a), v.mem.ByteAt(0x0b))
}

func (v *v1Header) GlobalVariableTableAddress() AbsAddr {
	return asJoinedByteAddress(v.mem.ByteAt(0x0c), v.mem.ByteAt(0x0d))
}

func (v *v1Header) StaticMemoryBaseAddress() AbsAddr {
	return asJoinedByteAddress(v.mem.ByteAt(0x0e), v.mem.ByteAt(0x0f))
}

func (v *v1Header) TranscriptEnabled() bool {
	return IsMemoryBitSet(v.mem, 0x10, 0)
}

func (v *v1Header) SetTranscriptEnabled(val bool) error {
	f := v.mem.ByteAt(0x10) & 0b11111110
	if val {
		f = f | 0b00000001
	}
	return v.mem.WriteByteAt(f, 0x10)
}

func (v *v1Header) RevisionNumber() uint16 {
	return asWord(v.mem.ByteAt(0x32), v.mem.ByteAt(0x33))
}

func (v *v1Header) ValueSettable(pos AbsAddr, val uint8) bool {
	// The only settable thing in v1 is the transcript enabled bit.
	if pos == 0x10 {
		base := v.mem.ByteAt(0x10) & 0b11111110
		if val == base || val == (base|0b00000001) {
			return true
		}
	}
	return false
}

func (v *v1Header) MarkInterpreterStart() {
	v.marked = v.mem.Clone()
}

func (v *v1Header) OnReset() {
	// Reset all "RST" values.

	// Update flag 1
	v.mem.WriteByteAt(v.marked.ByteAt(0x01), 0x01)

	// Update flag 2
	// Note: preserve bit 0
	flag2 := v.marked.ByteAt(0x10)
	bit0 := flag2 & 0b00000001
	v.mem.WriteByteAt((flag2&0b11111110)|bit0, 0x10)

	// Standard revision number
	v.mem.WriteByteAt(v.marked.ByteAt(0x32), 0x32)
}

func (v *v1Header) OnRestart() {
	// Flags 2 is preserved, but all other header changes are reset.
}

// ---- V1 Header Implementation - stuff outside the version.

func (v *v1Header) FileChecksum() uint16 {
	return 0
}
func (v *v1Header) SetInterpreter(number uint8, version uint8) error {
	return nil
}
func (v *v1Header) SetScreenHeight(lines uint8) error {
	return nil
}
func (v *v1Header) SetScreenWidth(chars uint8) error {
	return nil
}
func (v *v1Header) SetScreenHeightUnits(units uint16) error {
	return nil
}
func (v *v1Header) SetScreenWidthUnits(units uint16) error {
	return nil
}
func (v *v1Header) SetFontWidthUnits(units uint8) error {
	return nil
}
func (v *v1Header) SetFontHeightUnits(units uint8) error {
	return nil
}
func (v *v1Header) SetOutputStream3TextSentPixelWidth(pixels uint16) error {
	return nil
}
func (v *v1Header) StatusLineType() StatusLineType {
	return ScoreTurnsStatusLine
}
func (v *v1Header) TwoDiscStory() bool {
	return false
}
func (v *v1Header) StatusLineAvailable() bool {
	return true
}
func (v *v1Header) SetStatusLineAvailable(val bool) error {
	return nil
}
func (v *v1Header) ScreenSplitAvailable() bool {
	return false
}
func (v *v1Header) SetScreenSplitAvailable(val bool) error {
	return nil
}
func (v *v1Header) DefaultVariablePitchFont() bool {
	return false
}
func (v *v1Header) SetDefaultVariablePitchFont(val bool) error {
	return nil
}
func (v *v1Header) ColorsAvailable() bool {
	return false
}
func (v *v1Header) SetColorsAvailable(val bool) error {
	return nil
}
func (v *v1Header) PictureDisplayAvailable() bool {
	return false
}
func (v *v1Header) SetPictureDisplayAvailable(val bool) error {
	return nil
}
func (v *v1Header) BoldfaceAvailable() bool {
	return false
}
func (v *v1Header) SetBoldfaceAvailable(val bool) error {
	return nil
}
func (v *v1Header) ItalicAvailable() bool {
	return false
}
func (v *v1Header) SetItalicAvailable(val bool) error {
	return nil
}
func (v *v1Header) FixedSpaceFontAvailable() bool {
	return false
}
func (v *v1Header) SetFixedSpaceFontAvailable(val bool) error {
	return nil
}
func (v *v1Header) SoundEffectsAvailable() bool {
	return false
}
func (v *v1Header) SetSoundEffectsAvailable(val bool) error {
	return nil
}
func (v *v1Header) TimedKeyboardInputAvailable() bool {
	return false
}
func (v *v1Header) SetTimedKeyboardInputAvailable(val bool) error {
	return nil
}
func (v *v1Header) DefaultColors() (uint8, uint8) {
	return 0, 0
}
func (v *v1Header) SetDefaultColors(foreground uint8, background uint8) error {
	return nil
}
func (v *v1Header) AbbreviationsTableAddress() AbsAddr {
	return 0
}
func (v *v1Header) TerminatingCharactersTableAddress() AbsAddr {
	return 0
}
func (v *v1Header) UnicodeTranslationTableAddress() AbsAddr {
	return 0
}
func (v *v1Header) FileLength() int {
	return int(v.mem.Size())
}
func (v *v1Header) ForcedFixedPitchFontEnabled() bool {
	return false
}
func (v *v1Header) RequestsStatusLineRedraw() bool {
	return false
}
func (v *v1Header) SetRequestStatusLineRedraw() error {
	return nil
}
func (v *v1Header) GameRequestsPictures() bool {
	return false
}
func (v *v1Header) SetPicturesNotAvailable() error {
	return nil
}
func (v *v1Header) GameRequestsUndoOpcodes() bool {
	return false
}
func (v *v1Header) SetUndoNotAvailable() error {
	return nil
}
func (v *v1Header) GameRequestsMouseSupport() bool {
	return false
}
func (v *v1Header) SetMouseNotAvailable() error {
	return nil
}
func (v *v1Header) RequestsColors() bool {
	return false
}
func (v *v1Header) GameRequestsSounds() bool {
	return false
}
func (v *v1Header) SetSoundsNotAvailable() error {
	return nil
}
func (v *v1Header) GameRequestsMenus() bool {
	return false
}
func (v *v1Header) SetMenusNotAvailable() error {
	return nil
}
func (v *v1Header) SetMouseClickPos(x uint16, y uint16) error {
	return nil
}
