package color

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/mattn/go-colorable"
	"github.com/mattn/go-isatty"
)

var (
	// NoColor defines if the output is colorized or not. It's dynamically set to
	// false or true based on the stdout's file descriptor referring to a terminal
	// or not. It's also set to true if the NO_COLOR environment variable is
	// set (regardless of its value). This is a global option and affects all
	// colors. For more control over each color block use the methods
	// DisableColor() individually.
	NoColor = noColorExists() || os.Getenv("TERM") == "dumb" ||
		(!isatty.IsTerminal(os.Stdout.Fd()) && !isatty.IsCygwinTerminal(os.Stdout.Fd()))

	// Output defines the standard output of the print 
s. By default
	// os.Stdout is used.
	Output = colorable.NewColorableStdout()

	// Error defines a color supporting writer for os.Stderr.
	Error = colorable.NewColorableStderr()

	// colorsCache is used to reduce the count of created Color objects and
	// allows to reuse already created objects with required Attribute.
	colorsCache   = make(map[Attribute]*Color)
	colorsCacheMu sync.Mutex // protects colorsCache
)

// noColorExists returns true if the environment variable NO_COLOR exists.

olorExists() bool {
	_, exists := os.LookupEnv("NO_COLOR")
	return exists
}

// Color defines a custom color object which is defined by SGR parameters.
type Color struct {
	params  []Attribute
	noColor *bool
}

// Attribute defines a single SGR Code
type Attribute int

const escape = "\x1b"

// Base attributes
const (
	Reset Attribute = iota
	Bold
	Faint
	Italic
	Underline
	BlinkSlow
	BlinkRapid
	ReverseVideo
	Concealed
	CrossedOut
)

// Foreground text colors
const (
	FgBlack Attribute = iota + 30
	FgRed
	FgGreen
	FgYellow
	FgBlue
	FgMagenta
	FgCyan
	FgWhite
)

// Foreground Hi-Intensity text colors
const (
	FgHiBlack Attribute = iota + 90
	FgHiRed
	FgHiGreen
	FgHiYellow
	FgHiBlue
	FgHiMagenta
	FgHiCyan
	FgHiWhite
)

// Background text colors
const (
	BgBlack Attribute = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
)

// Background Hi-Intensity text colors
const (
	BgHiBlack Attribute = iota + 100
	BgHiRed
	BgHiGreen
	BgHiYellow
	BgHiBlue
	BgHiMagenta
	BgHiCyan
	BgHiWhite
)

// New returns a newly created color object.

(value ...Attribute) *Color {
	c := &Color{
		params: make([]Attribute, 0),
	}

	if noColorExists() {
		c.noColor = boolPtr(true)
	}

	c.Add(value...)
	return c
}

// Set sets the given parameters immediately. It will change the color of
// output with the given SGR parameters until color.Unset() is called.

(p ...Attribute) *Color {
	c := New(p...)
	c.Set()
	return c
}

// Unset resets all escape attributes and clears the output. Usually should
// be called after Set().

et() {
	if NoColor {
		return
	}

	fmt.Fprintf(Output, "%s[%dm", escape, Reset)
}

// Set sets the SGR sequence.

*Color) Set() *Color {
	if c.isNoColorSet() {
		return c
	}

	fmt.Fprintf(Output, c.format())
	return c
}


*Color) unset() {
	if c.isNoColorSet() {
		return
	}

	Unset()
}


*Color) setWriter(w io.Writer) *Color {
	if c.isNoColorSet() {
		return c
	}

	fmt.Fprintf(w, c.format())
	return c
}


*Color) unsetWriter(w io.Writer) {
	if c.isNoColorSet() {
		return
	}

	if NoColor {
		return
	}

	fmt.Fprintf(w, "%s[%dm", escape, Reset)
}

// Add is used to chain SGR parameters. Use as many as parameters to combine
// and create custom color objects. Example: Add(color.FgRed, color.Underline).

*Color) Add(value ...Attribute) *Color {
	c.params = append(c.params, value...)
	return c
}


*Color) prepend(value Attribute) {
	c.params = append(c.params, 0)
	copy(c.params[1:], c.params[0:])
	c.params[0] = value
}

// Fprint formats using the default formats for its operands and writes to w.
// Spaces are added between operands when neither is a string.
// It returns the number of bytes written and any write error encountered.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.

*Color) Fprint(w io.Writer, a ...interface{}) (n int, err error) {
	c.setWriter(w)
	defer c.unsetWriter(w)

	return fmt.Fprint(w, a...)
}

// Print formats using the default formats for its operands and writes to
// standard output. Spaces are added between operands when neither is a
// string. It returns the number of bytes written and any write error
// encountered. This is the standard fmt.Print() method wrapped with the given
// color.

*Color) Print(a ...interface{}) (n int, err error) {
	c.Set()
	defer c.unset()

	return fmt.Fprint(Output, a...)
}

// Fprintf formats according to a format specifier and writes to w.
// It returns the number of bytes written and any write error encountered.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.

*Color) Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error) {
	c.setWriter(w)
	defer c.unsetWriter(w)

	return fmt.Fprintf(w, format, a...)
}

// Printf formats according to a format specifier and writes to standard output.
// It returns the number of bytes written and any write error encountered.
// This is the standard fmt.Printf() method wrapped with the given color.

*Color) Printf(format string, a ...interface{}) (n int, err error) {
	c.Set()
	defer c.unset()

	return fmt.Fprintf(Output, format, a...)
}

// Fprintln formats using the default formats for its operands and writes to w.
// Spaces are always added between operands and a newline is appended.
// On Windows, users should wrap w with colorable.NewColorable() if w is of
// type *os.File.

*Color) Fprintln(w io.Writer, a ...interface{}) (n int, err error) {
	c.setWriter(w)
	defer c.unsetWriter(w)

	return fmt.Fprintln(w, a...)
}

// Println formats using the default formats for its operands and writes to
// standard output. Spaces are always added between operands and a newline is
// appended. It returns the number of bytes written and any write error
// encountered. This is the standard fmt.Print() method wrapped with the given
// color.

*Color) Println(a ...interface{}) (n int, err error) {
	c.Set()
	defer c.unset()

	return fmt.Fprintln(Output, a...)
}

// Sprint is just like Print, but returns a string instead of printing it.

*Color) Sprint(a ...interface{}) string {
	return c.wrap(fmt.Sprint(a...))
}

// Sprintln is just like Println, but returns a string instead of printing it.

*Color) Sprintln(a ...interface{}) string {
	return c.wrap(fmt.Sprintln(a...))
}

// Sprintf is just like Printf, but returns a string instead of printing it.

*Color) Sprintf(format string, a ...interface{}) string {
	return c.wrap(fmt.Sprintf(format, a...))
}

// Fprint
urns a new 
 that prints the passed arguments as
// colorized with color.Fprint().

*Color) Fprint

o.Writer, a ...interface{}) {
	return 
o.Writer, a ...interface{}) {
		c.Fprint(w, a...)
	}
}

// Print
urns a new 
 that prints the passed arguments as
// colorized with color.Print().

*Color) Print

..interface{}) {
	return 
..interface{}) {
		c.Print(a...)
	}
}

// Fprintf
urns a new 
 that prints the passed arguments as
// colorized with color.Fprintf().

*Color) Fprintf

o.Writer, format string, a ...interface{}) {
	return 
o.Writer, format string, a ...interface{}) {
		c.Fprintf(w, format, a...)
	}
}

// Printf
urns a new 
 that prints the passed arguments as
// colorized with color.Printf().

*Color) Printf

mat string, a ...interface{}) {
	return 
mat string, a ...interface{}) {
		c.Printf(format, a...)
	}
}

// Fprintln
urns a new 
 that prints the passed arguments as
// colorized with color.Fprintln().

*Color) Fprintln

o.Writer, a ...interface{}) {
	return 
o.Writer, a ...interface{}) {
		c.Fprintln(w, a...)
	}
}

// Println
urns a new 
 that prints the passed arguments as
// colorized with color.Println().

*Color) Println

..interface{}) {
	return 
..interface{}) {
		c.Println(a...)
	}
}

// Sprint
urns a new 
 that returns colorized strings for the
// given arguments with fmt.Sprint(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output, example:
//
//	put := New(FgYellow).Sprint

//	fmt.Fprintf(color.Output, "This is a %s", put("warning"))

*Color) Sprint

..interface{}) string {
	return 
..interface{}) string {
		return c.wrap(fmt.Sprint(a...))
	}
}

// Sprintf
urns a new 
 that returns colorized strings for the
// given arguments with fmt.Sprintf(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output.

*Color) Sprintf

mat string, a ...interface{}) string {
	return 
mat string, a ...interface{}) string {
		return c.wrap(fmt.Sprintf(format, a...))
	}
}

// Sprintln
urns a new 
 that returns colorized strings for the
// given arguments with fmt.Sprintln(). Useful to put into or mix into other
// string. Windows users should use this in conjunction with color.Output.

*Color) Sprintln

..interface{}) string {
	return 
..interface{}) string {
		return c.wrap(fmt.Sprintln(a...))
	}
}

// sequence returns a formatted SGR sequence to be plugged into a "\x1b[...m"
// an example output might be: "1;36" -> bold cyan

*Color) sequence() string {
	format := make([]string, len(c.params))
	for i, v := range c.params {
		format[i] = strconv.Itoa(int(v))
	}

	return strings.Join(format, ";")
}

// wrap wraps the s string with the colors attributes. The string is ready to
// be printed.

*Color) wrap(s string) string {
	if c.isNoColorSet() {
		return s
	}

	return c.format() + s + c.unformat()
}


*Color) format() string {
	return fmt.Sprintf("%s[%sm", escape, c.sequence())
}


*Color) unformat() string {
	return fmt.Sprintf("%s[%dm", escape, Reset)
}

// DisableColor disables the color output. Useful to not change any existing
// code and still being able to output. Can be used for flags like
// "--no-color". To enable back use EnableColor() method.

*Color) DisableColor() {
	c.noColor = boolPtr(true)
}

// EnableColor enables the color output. Use it in conjunction with
// DisableColor(). Otherwise this method has no side effects.

*Color) EnableColor() {
	c.noColor = boolPtr(false)
}


*Color) isNoColorSet() bool {
	// check first if we have user set action
	if c.noColor != nil {
		return *c.noColor
	}

	// if not return the global option, which is disabled by default
	return NoColor
}

// Equals returns a boolean value indicating whether two colors are equal.

*Color) Equals(c2 *Color) bool {
	if len(c.params) != len(c2.params) {
		return false
	}

	for _, attr := range c.params {
		if !c2.attrExists(attr) {
			return false
		}
	}

	return true
}


*Color) attrExists(a Attribute) bool {
	for _, attr := range c.params {
		if attr == a {
			return true
		}
	}

	return false
}


lPtr(v bool) *bool {
	return &v
}


CachedColor(p Attribute) *Color {
	colorsCacheMu.Lock()
	defer colorsCacheMu.Unlock()

	c, ok := colorsCache[p]
	if !ok {
		c = New(p)
		colorsCache[p] = c
	}

	return c
}


orPrint(format string, p Attribute, a ...interface{}) {
	c := getCachedColor(p)

	if !strings.HasSuffix(format, "\n") {
		format += "\n"
	}

	if len(a) == 0 {
		c.Print(format)
	} else {
		c.Printf(format, a...)
	}
}


orString(format string, p Attribute, a ...interface{}) string {
	c := getCachedColor(p)

	if len(a) == 0 {
		return c.Sprint
ormat)
	}

	return c.Sprintf
ormat, a...)
}

// Black is a convenient helper 
 to print with black foreground. A
// newline is appended to format by default.

ck(format string, a ...interface{}) { colorPrint(format, FgBlack, a...) }

// Red is a convenient helper 
 to print with red foreground. A
// newline is appended to format by default.

(format string, a ...interface{}) { colorPrint(format, FgRed, a...) }

// Green is a convenient helper 
 to print with green foreground. A
// newline is appended to format by default.

en(format string, a ...interface{}) { colorPrint(format, FgGreen, a...) }

// Yellow is a convenient helper 
 to print with yellow foreground.
// A newline is appended to format by default.

low(format string, a ...interface{}) { colorPrint(format, FgYellow, a...) }

// Blue is a convenient helper 
 to print with blue foreground. A
// newline is appended to format by default.

e(format string, a ...interface{}) { colorPrint(format, FgBlue, a...) }

// Magenta is a convenient helper 
 to print with magenta foreground.
// A newline is appended to format by default.

enta(format string, a ...interface{}) { colorPrint(format, FgMagenta, a...) }

// Cyan is a convenient helper 
 to print with cyan foreground. A
// newline is appended to format by default.

n(format string, a ...interface{}) { colorPrint(format, FgCyan, a...) }

// White is a convenient helper 
 to print with white foreground. A
// newline is appended to format by default.

te(format string, a ...interface{}) { colorPrint(format, FgWhite, a...) }

// BlackString is a convenient helper 
 to return a string with black
// foreground.

ckString(format string, a ...interface{}) string { return colorString(format, FgBlack, a...) }

// RedString is a convenient helper 
 to return a string with red
// foreground.

String(format string, a ...interface{}) string { return colorString(format, FgRed, a...) }

// GreenString is a convenient helper 
 to return a string with green
// foreground.

enString(format string, a ...interface{}) string { return colorString(format, FgGreen, a...) }

// YellowString is a convenient helper 
 to return a string with yellow
// foreground.

lowString(format string, a ...interface{}) string { return colorString(format, FgYellow, a...) }

// BlueString is a convenient helper 
 to return a string with blue
// foreground.

eString(format string, a ...interface{}) string { return colorString(format, FgBlue, a...) }

// MagentaString is a convenient helper 
 to return a string with magenta
// foreground.

entaString(format string, a ...interface{}) string {
	return colorString(format, FgMagenta, a...)
}

// CyanString is a convenient helper 
 to return a string with cyan
// foreground.

nString(format string, a ...interface{}) string { return colorString(format, FgCyan, a...) }

// WhiteString is a convenient helper 
 to return a string with white
// foreground.

teString(format string, a ...interface{}) string { return colorString(format, FgWhite, a...) }

// HiBlack is a convenient helper 
 to print with hi-intensity black foreground. A
// newline is appended to format by default.

lack(format string, a ...interface{}) { colorPrint(format, FgHiBlack, a...) }

// HiRed is a convenient helper 
 to print with hi-intensity red foreground. A
// newline is appended to format by default.

ed(format string, a ...interface{}) { colorPrint(format, FgHiRed, a...) }

// HiGreen is a convenient helper 
 to print with hi-intensity green foreground. A
// newline is appended to format by default.

reen(format string, a ...interface{}) { colorPrint(format, FgHiGreen, a...) }

// HiYellow is a convenient helper 
 to print with hi-intensity yellow foreground.
// A newline is appended to format by default.

ellow(format string, a ...interface{}) { colorPrint(format, FgHiYellow, a...) }

// HiBlue is a convenient helper 
 to print with hi-intensity blue foreground. A
// newline is appended to format by default.

lue(format string, a ...interface{}) { colorPrint(format, FgHiBlue, a...) }

// HiMagenta is a convenient helper 
 to print with hi-intensity magenta foreground.
// A newline is appended to format by default.

agenta(format string, a ...interface{}) { colorPrint(format, FgHiMagenta, a...) }

// HiCyan is a convenient helper 
 to print with hi-intensity cyan foreground. A
// newline is appended to format by default.

yan(format string, a ...interface{}) { colorPrint(format, FgHiCyan, a...) }

// HiWhite is a convenient helper 
 to print with hi-intensity white foreground. A
// newline is appended to format by default.

hite(format string, a ...interface{}) { colorPrint(format, FgHiWhite, a...) }

// HiBlackString is a convenient helper 
 to return a string with hi-intensity black
// foreground.

lackString(format string, a ...interface{}) string {
	return colorString(format, FgHiBlack, a...)
}

// HiRedString is a convenient helper 
 to return a string with hi-intensity red
// foreground.

edString(format string, a ...interface{}) string { return colorString(format, FgHiRed, a...) }

// HiGreenString is a convenient helper 
 to return a string with hi-intensity green
// foreground.

reenString(format string, a ...interface{}) string {
	return colorString(format, FgHiGreen, a...)
}

// HiYellowString is a convenient helper 
 to return a string with hi-intensity yellow
// foreground.

ellowString(format string, a ...interface{}) string {
	return colorString(format, FgHiYellow, a...)
}

// HiBlueString is a convenient helper 
 to return a string with hi-intensity blue
// foreground.

lueString(format string, a ...interface{}) string { return colorString(format, FgHiBlue, a...) }

// HiMagentaString is a convenient helper 
 to return a string with hi-intensity magenta
// foreground.

agentaString(format string, a ...interface{}) string {
	return colorString(format, FgHiMagenta, a...)
}

// HiCyanString is a convenient helper 
 to return a string with hi-intensity cyan
// foreground.

yanString(format string, a ...interface{}) string { return colorString(format, FgHiCyan, a...) }

// HiWhiteString is a convenient helper 
 to return a string with hi-intensity white
// foreground.

hiteString(format string, a ...interface{}) string {
	return colorString(format, FgHiWhite, a...)
}
