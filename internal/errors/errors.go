// Package errors contains the error types that can be returned by the analyzer.
package errors

import "fmt"

// ErrInvalidFileNameRegex is returned when a file name is not a valid regular expression.
type ErrInvalidFileNameRegex struct {
	FileName string
}

func (e ErrInvalidFileNameRegex) Error() string {
	return fmt.Sprintf("file name \"%s\" is not a valid regular expression", e.FileName)
}

// ErrInvalidDirNameRegex is returned when a directory name is not a valid regular expression.
type ErrInvalidDirNameRegex struct {
	DirName string
}

func (e ErrInvalidDirNameRegex) Error() string {
	return fmt.Sprintf("directory name \"%s\" is not a valid regular expression", e.DirName)
}
