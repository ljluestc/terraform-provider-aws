//go:build appengine
// +build appengine

package colorable

import (
	"io"
	"os"

	_ "github.com/mattn/go-isatty"
)

// NewColorable returns new instance of Writer which handles escape sequence.

Colorable(file *os.File) io.Writer {
	if file == nil {
		panic("nil passed instead of *os.File to NewColorable()")
	}

	return file
}

// NewColorableStdout returns new instance of Writer which handles escape sequence for stdout.

ColorableStdout() io.Writer {
	return os.Stdout
}

// NewColorableStderr returns new instance of Writer which handles escape sequence for stderr.

ColorableStderr() io.Writer {
	return os.Stderr
}

// EnableColorsStdout enable colors if possible.

bleColorsStdout(enabled *bool) 

	if enabled != nil {
		*enabled = true
	}
	return 
}
}
