// Creates a standardized input mechanism.
package machine

// A single user input action.
//
// Standard user input keys should use the Key value to reflect the text typed, while
// special control characters should use the constant translations below.
type UserInput struct {
	Ctrl ZsciiChar
	Key  rune
}

var DelInput = UserInput{Ctrl: DelChar}
var EscInput = UserInput{Ctrl: EscChar}
var UpInput = UserInput{Ctrl: UpChar}
var DownInput = UserInput{Ctrl: DownChar}
var LeftInput = UserInput{Ctrl: LeftChar}
var RightInput = UserInput{Ctrl: RightChar}
var F1Input = UserInput{Ctrl: F1Char}
var F2Input = UserInput{Ctrl: F2Char}
var F3Input = UserInput{Ctrl: F3Char}
var F4Input = UserInput{Ctrl: F4Char}
var F5Input = UserInput{Ctrl: F5Char}
var F6Input = UserInput{Ctrl: F6Char}
var F7Input = UserInput{Ctrl: F7Char}
var F8Input = UserInput{Ctrl: F8Char}
var F9Input = UserInput{Ctrl: F9Char}
var F10Input = UserInput{Ctrl: F10Char}
var F11Input = UserInput{Ctrl: F11Char}
var F12Input = UserInput{Ctrl: F12Char}
var Keypad0Input = UserInput{Ctrl: Keypad0Char}
var Keypad1Input = UserInput{Ctrl: Keypad1Char}
var Keypad2Input = UserInput{Ctrl: Keypad2Char}
var Keypad3Input = UserInput{Ctrl: Keypad3Char}
var Keypad4Input = UserInput{Ctrl: Keypad4Char}
var Keypad5Input = UserInput{Ctrl: Keypad5Char}
var Keypad6Input = UserInput{Ctrl: Keypad6Char}
var Keypad7Input = UserInput{Ctrl: Keypad7Char}
var Keypad8Input = UserInput{Ctrl: Keypad8Char}
var Keypad9Input = UserInput{Ctrl: Keypad9Char}
var MenuClickInput = UserInput{Ctrl: MenuClickChar}
var DoubleClickInput = UserInput{Ctrl: DoubleClickChar}
var SingleClickInput = UserInput{Ctrl: SingleClickChar}

type InputAction struct {
	UserInput
	Delay float32 // seconds since last input
}
