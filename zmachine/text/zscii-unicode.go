// ZSCII to Unicode translation.

package text

func uChar(v int64) string {
	return string(rune(v))
}

var ZsciiSpecialUnicode = map[int]string{
	0:   "",            // char 0 is output as no-text.
	9:   " ",           // char 9 should only be printed as a tab if it's at the start of a line, otherwise just one space.
	11:  uChar(0x2001), // char 11 is a sentence space, V6 only.
	13:  uChar(10),     // newline
	34:  "\"",          // 0x22 - neutral double quote
	39:  uChar(0x2019), // 0x27 - right single quote
	96:  uChar(0x2018), // 0x60 - left single quote
	155: uChar(228),    // 0e4 - a-diaeresis
	156: uChar(246),    // 0f6 - o-diaeresis
	157: uChar(252),    // 0fc - u-diaeresis
	158: uChar(196),    // 0c4 - A-diaeresis
	159: uChar(214),    // 0d6 - O-diaeresis
	160: uChar(220),    // 0dc - U-diaeresis
	161: uChar(223),    // 0df - sz-ligature
	162: uChar(187),    // 0bb - quotation (<< or ")
	163: uChar(171),    // 0ab - marks (>> or ")
	164: uChar(235),    // 0eb - e-diaeresis
	165: uChar(239),    // 0ef - i-diaeresis
	166: uChar(255),    // 0ff - y-diaeresis
	167: uChar(203),    // 0cb - E-diaeresis
	168: uChar(207),    // 0cf - I-diaeresis
	169: uChar(225),    // 0e1 - a-acute
	170: uChar(233),    // 0e9 - e-acute
	171: uChar(237),    // 0ed - i-acute
	172: uChar(243),    // 0f3 - o-acute
	173: uChar(250),    // 0fa - u-acute
	174: uChar(253),    // 0fd - y-acute
	175: uChar(193),    // 0c1 - A-acute
	176: uChar(201),    // 0c9 - E-acute
	177: uChar(205),    // 0cd - I-acute
	178: uChar(211),    // 0d3 - O-acute
	179: uChar(218),    // 0da - U-acute
	180: uChar(221),    // 0dd - Y-acute
	181: uChar(224),    // 0e0 - a-grave
	182: uChar(232),    // 0e8 - e-grave
	183: uChar(236),    // 0ec - i-grave
	184: uChar(242),    // 0f2 - o-grave
	185: uChar(249),    // 0f9 - u-grave
	186: uChar(192),    // 0c0 - A-grave
	187: uChar(200),    // 0c8 - E-grave
	188: uChar(204),    // 0cc - I-grave
	189: uChar(210),    // 0d2 - O-grave
	190: uChar(217),    // 0d9 - U-grave
	191: uChar(226),    // 0e2 - a-circumflex
	192: uChar(234),    // 0ea - e-circumflex
	193: uChar(238),    // 0ee - i-circumflex
	194: uChar(244),    // 0f4 - o-circumflex
	195: uChar(251),    // 0fb - u-circumflex
	196: uChar(194),    // 0c2 - A-circumflex
	197: uChar(202),    // 0ca - E-circumflex
	198: uChar(206),    // 0ce - I-circumflex
	199: uChar(212),    // 0d4 - O-circumflex
	200: uChar(219),    // 0db - U-circumflex
	201: uChar(229),    // 0e5 - a-ring
	202: uChar(197),    // 0c5 - A-ring
	203: uChar(248),    // 0f8 - o-slash
	204: uChar(216),    // 0d8 - O-slash
	205: uChar(227),    // 0e3 - a-tilde
	206: uChar(241),    // 0f1 - n-tilde
	207: uChar(245),    // 0f5 - o-tilde
	208: uChar(195),    // 0c3 - A-tilde
	209: uChar(209),    // 0d1 - N-tilde
	210: uChar(213),    // 0d5 - O-tilde
	211: uChar(230),    // 0e6 - ae-ligature
	212: uChar(198),    // 0c6 - AE-ligature
	213: uChar(231),    // 0e7 - c-cedilla
	214: uChar(199),    // 0c7 - C-cedilla
	215: uChar(254),    // 0fe - Icelandic thorn
	216: uChar(240),    // 0f0 - Icelandic eth
	217: uChar(222),    // 0de - Icelandic Thorn
	218: uChar(208),    // 0d0 - Icelandic Eth
	219: uChar(163),    // 0a3 - pound symbol
	220: uChar(339),    // 153 - oe-ligature
	221: uChar(338),    // 152 - OE-ligature
	222: uChar(161),    // 0a1 - inverted !
	223: uChar(191),    // 0bf - inverted ?
}
