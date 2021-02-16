package main

type Key struct {
	Type  byte
	Value byte
}

const (
	// KeyLetter is a single logical key. Value is the ASCII code of one of these:
	KeyLetter = (1 << iota)
	// KeyCtrl is or'd to the modifier list if the ctrl key was pressed.
	// For technical reasons, C-(A-Z) is reported as C-(a-z).
	KeyCtrl
	KeyAlt
	// °: 0xB0
	// §: 0xA7
	KeyUml
	// KeySpecial indicates that this key should not be printed, but instead be interpreted
	KeySpecial
)

const (
	// a-z 0x61 - 0x71
	// A-Z 0x41 - 0x5A
	// 0-9 0x30 - 0x39
	AsciiSpace        = 0x20
	AsciiExcl         = 0x21
	AsciiDquote       = 0x22
	AsciiHash         = 0x23
	AsciiDollar       = 0x24
	AsciiPerc         = 0x25
	AsciiAnd          = 0x26
	AsciiSquote       = 0x27
	AsciiPLeft        = 0x28
	AsciiPRight       = 0x29
	AsciiAsterisk     = 0x2A
	AsciiStar         = 0x2B
	AsciiComma        = 0x2C
	AsciiMinus        = 0x2D
	AsciiDot          = 0x2E
	AsciiSlash        = 0x2F
	AsciiColon        = 0x3A
	AsciiScolon       = 0x3B
	AsciiLt           = 0x3C
	AsciiEq           = 0x3D
	AsciiGt           = 0x3E
	AsciiQuest        = 0x3F
	AsciiAt           = 0x40
	AsciiBracketLeft  = 0x5B
	AsciiBSlash       = 0x5C
	AsciiBracketRight = 0x5D
	AsciiCflex        = 0x5E
	AsciiUscore       = 0x5F
	AsciiBacktick     = 0x60
	AsciiBraceLeft    = 0x7B
	AsciiBar          = 0x7C
	AsciiBraceRight   = 0x7D
	AsciiTilde        = 0x7E
)

const (
	// UmlCAGrv is a capital A with an accent grave: À
	UmlCAGrv = 0x80

	// UmlCAAgu is a capital A with an accent aigu: Á
	UmlCAAgu = 0x81

	// UmlCACrc is a capital A with a circumflex: Â
	UmlCACrc = 0x82

	// UmlCATld is a capital A with a tilde: Ã
	UmlCATld = 0x83

	// UmlCADia is a capital A with a diaresis: Ä
	UmlCADia = 0x84

	// UmlCA is a capital A with a ring: Å
	UmlCA = 0x85

	// UmlC is a capital AE: Æ
	UmlCAE = 0x86

	// UmlCCCed is a capital C with cedilla: Ç
	UmlCCCed = 0x87

	// UmlCEGrv is a capital E with an accent grave: È
	UmlCEGrv = 0x88

	// UmlCEAgu is a capital E with an accent aigu: É
	UmlCEAgu = 0x89

	// UmlCECrc is a capital E with a circumflex: Ê
	UmlCECrc = 0x8a

	// UmlCEDia is a capital E with a diaresis: Ë
	UmlCEDia = 0x8b

	// UmlCIGrv is a capital I with an accent grave: Ì
	UmlCIGrv = 0x8c

	// UmlCIAgu is a capital I with an accent aigu: Í
	UmlCIAgu = 0x8d

	// UmlCICrc is a capital I with a circumflex: Î
	UmlCICrc = 0x8e

	// UmlCIDia is a capital I with a diaresis: Ï
	UmlCIDia = 0x8f

	// UmlC is a capital Eth: Ð
	UmlCEth = 0x90

	// UmlCNTld is a capital N with a tilde: Ñ
	UmlCNTld = 0x91

	// UmlCOGrv is a capital O with an accent grave: Ò
	UmlCOGrv = 0x92

	// UmlCOAgu is a capital O with an accent aigu: Ó
	UmlCOAgu = 0x93

	// UmlCOCrc is a capital O with a circumflex: Ô
	UmlCOCrc = 0x94

	// UmlCOTld is a capital O with a tilde: Õ
	UmlCOTld = 0x95

	// UmlCODia is a capital O with a diaresis: Ö
	UmlCODia = 0x96

	// UmlCOStrk is a capital O with strikethrough: Ø
	UmlCOStrk = 0x98

	// UmlCUGrv is a capital U with an accent grave: Ù
	UmlCUGrv = 0x99

	// UmlCUAgu is a capital U with an accent aigu: Ú
	UmlCUAgu = 0x9a

	// UmlCUCrc is a capital U with a circumflex: Û
	UmlCUCrc = 0x9b

	// UmlCUDia is a capital U with a diaresis: Ü
	UmlCUDia = 0x9c

	// UmlCYAgu is a capital Y with an accent aigu: Ý
	UmlCYAgu = 0x9d

	// UmlC is a capital Thorn: Þ
	UmlCThorn = 0x9e

	// UmlS is a small eszett ß:
	UmlSEszett = 0x9f

	// UmlSAGrv is a small A with an accent grave: à
	UmlSAGrv = 0xa0

	// UmlSAAgu is a small A with an accent aigu: á
	UmlSAAgu = 0xa1

	// UmlSACrc is a small A with a circumflex: â
	UmlSACrc = 0xa2

	// UmlSATld is a small A with a tilde: ã
	UmlSATld = 0xa3

	// UmlSADia is a small A with a diaresis: ä
	UmlSADia = 0xa4

	// UmlSA is a small A with a ring: å
	UmlSARing = 0xa5

	// UmlS is a small AE: æ
	UmlSAE = 0xa6

	// UmlSCCed is a small C with cedilla: ç
	UmlSCCed = 0xa7

	// UmlSEGrv is a small E with an accent grave: è
	UmlSEGrv = 0xa8

	// UmlSEAgu is a small E with an accent aigu: é
	UmlSEAgu = 0xa9

	// UmlSECrc is a small E with a circumflex: ê
	UmlSECrc = 0xaa

	// UmlSEDia is a small E with a diaresis: ë
	UmlSEDia = 0xab

	// UmlSIGrv is a small I with an accent grave: ì
	UmlSIGrv = 0xac

	// UmlSIAgu is a small I with an accent aigu: í
	UmlSIAgu = 0xad

	// UmlSICrc is a small I with a circumflex: î
	UmlSICrc = 0xae

	// UmlSIDia is a small I with a diaresis: ï
	UmlSIDia = 0xaf

	// UmlS is a small Eth: ð
	UmlSEth = 0xb0

	// UmlSNTld is a small N with a tilde: ñ
	UmlSNTld = 0xb1

	// UmlSOGrv is a small O with an accent grave: ò
	UmlSOGrv = 0xb2

	// UmlSOAgu is a small O with an accent aigu: ó
	UmlSOAgu = 0xb3

	// UmlSOCrc is a small O with a circumflex: ô
	UmlSOCrc = 0xb4

	// UmlSOTld is a small O with a tilde: õ
	UmlSOTld = 0xb5

	// UmlSODia is a small O with a diaresis: ö
	UmlSODia = 0xb6

	// UmlSOStrk is a small O with strikethrough: ø
	UmlSOStrk = 0xb8

	// UmlSUGrv is a small U with an accent grave: ù
	UmlSUGrv = 0xb9

	// UmlSUAgu is a small U with an accent aigu: ú
	UmlSUAgu = 0xba

	// UmlSUCrc is a small U with a circumflex: û
	UmlSUCrc = 0xbb

	// UmlSUDia is a small U with a diaresis: ü
	UmlSUDia = 0xbc

	// UmlSYAgu is a small Y with an accent aigu: ý
	UmlSYAgu = 0xbd

	// UmlS is a small Thorn: þ
	UmlSThorn = 0xbe

	// UmlSYDia is a small Y with a diaresis: ÿ
	UmlSYDia = 0xbf
)

const (
	Backspace = iota
)

// parseKeys splits the inputted bytes into logical characters
func parseKeys(in []byte) []Key {
	var i int
	var keys []Key

	for i = 0; i < len(in); i++ {
		if in[i] == 0x08 || in[i] == 0x7F {
			keys = append(keys, Key{KeySpecial, Backspace})
		} else if in[i] >= 0x01 && in[i] <= 0x1A {
			keys = append(keys, Key{KeyCtrl, in[i] - 0x01 + 0x61})
		} else if in[i] == 0xC3 {
			// parse umlaut:
			i++
			keys = append(keys, Key{KeyUml, in[i]})
		} else if in[i] == 0x1B {
			// parse escape sequence
		} else {
			keys = append(keys, Key{KeyLetter, in[i]})
		}
	}

	return keys
}