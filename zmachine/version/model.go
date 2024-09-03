// General mechanics for the version specific information.
package version

type Version interface {
	// Number returns the version number for the interpreter.
	Number() int
}
