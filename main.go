package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
)

type SampleError struct {
	message string
	err     error
}

func (se *SampleError) Error() string {
	return se.message
}

func (se *SampleError) Unwrap() error {
	return se.err
}

func main() {
	if _, err := os.Open("non-existing"); err != nil {
		var pathError *fs.PathError

		wrapedErr := &SampleError{message: "this is wraped err", err: err}
		if errors.As(wrapedErr, &pathError) {
			fmt.Println("errors.As():Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
		if errors.Is(wrapedErr, pathError) {
			fmt.Println("errors.Is():Failed at path:", pathError.Path)
		} else {
			fmt.Println(err)
		}
	}
}

// Output:
// errors.As():Failed at path: non-existing
// errors.Is():Failed at path: non-existing
