package utf8

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

// ParseUTF8 splits the inputted bytes into logical keypresses
func ParseUTF8(in []byte) []Key {
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
