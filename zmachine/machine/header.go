// The story file header management.
package machine

type StatusLineType int

const (
	ScoreTurnsStatusLine StatusLineType = iota
	HourMinutesStatusLine
)

// Header has version-specific header value parsing.
//
// Setting a value should only return an error if the operation itself causes an error.
// It must not return an error if that feature is unavailable.
type Header interface {
	VersionNumber() int
	RevisionNumber() uint16
	FileChecksum() uint16
	SetInterpreter(number uint8, version uint8) error
	SetScreenHeight(lines uint8) error // 255 means "infinite"
	SetScreenWidth(chars uint8) error
	SetScreenHeightUnits(units uint16) error
	SetScreenWidthUnits(units uint16) error
	SetFontWidthUnits(units uint8) error  // As per the width of a '0'
	SetFontHeightUnits(units uint8) error // As per the height of a '0'

	MarkInterpreterStart() // Marks the interpreter-set header values for reset or restart operations.
	OnReset()              // Resets the header data as per a 'reset' or 'undo' operation.
	OnRestart()            // Restart resets the header data as per a game restart operation.

	SetOutputStream3TextSentPixelWidth(pixels uint16) error // Total width in pixels of text sent to output stream 3

	StatusLineType() StatusLineType
	TwoDiscStory() bool

	StatusLineAvailable() bool
	SetStatusLineAvailable(val bool) error

	ScreenSplitAvailable() bool
	SetScreenSplitAvailable(val bool) error

	DefaultVariablePitchFont() bool
	SetDefaultVariablePitchFont(val bool) error

	ColorsAvailable() bool
	SetColorsAvailable(val bool) error

	PictureDisplayAvailable() bool
	SetPictureDisplayAvailable(val bool) error

	BoldfaceAvailable() bool
	SetBoldfaceAvailable(val bool) error

	ItalicAvailable() bool
	SetItalicAvailable(val bool) error

	FixedSpaceFontAvailable() bool
	SetFixedSpaceFontAvailable(val bool) error

	SoundEffectsAvailable() bool
	SetSoundEffectsAvailable(val bool) error

	TimedKeyboardInputAvailable() bool
	SetTimedKeyboardInputAvailable(val bool) error

	DefaultColors() (uint8, uint8)
	SetDefaultColors(foreground uint8, background uint8) error

	HighMemoryBaseAddress() AbsAddr
	StaticMemoryBaseAddress() AbsAddr
	DictionaryAddress() AbsAddr
	ObjectTableAddress() AbsAddr
	GlobalVariableTableAddress() AbsAddr
	AbbreviationsTableAddress() AbsAddr
	TerminatingCharactersTableAddress() AbsAddr
	UnicodeTranslationTableAddress() AbsAddr
	FileLength() int // As marked by the story file and version-interpreted.

	TranscriptEnabled() bool
	SetTranscriptEnabled(val bool) error // Also settable by the game.
	ForcedFixedPitchFontEnabled() bool   // Game may change it, interpreter may not.
	RequestsStatusLineRedraw() bool
	SetRequestStatusLineRedraw() error // Interpreter requests it, game clears it.
	GameRequestsPictures() bool        // Game may change it, interpreter may clear it if not available.
	SetPicturesNotAvailable() error
	GameRequestsUndoOpcodes() bool // Game may change it, interpreter may not.
	SetUndoNotAvailable() error
	GameRequestsMouseSupport() bool // Game may change it, interpreter may not.
	SetMouseNotAvailable() error
	RequestsColors() bool     // Game hard-coded in story file.
	GameRequestsSounds() bool // Game may change it, interpreter may not.
	SetSoundsNotAvailable() error
	GameRequestsMenus() bool // Game may change it, interpreter may not.
	SetMenusNotAvailable() error

	SetMouseClickPos(x uint16, y uint16) error

	// ValueSettable checks if the given header position and value are legal for the game to change.
	//
	// The interpreter allows the game to set only some values.
	ValueSettable(pos AbsAddr, val uint8) bool
}
