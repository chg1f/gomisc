package weak

import (
	"fmt"
	"io"
)

func ExampleCompressError_1() {
	var ce CompressError
	ce = append(ce, io.EOF)
	ce = ce.Shrink()
}
func ExampleCompressError_2() {
	f := func() error {
		var ce CompressError
		return ce.Shrink()
	}
	if err := f(); err != nil {
		fmt.Print(err)
	}
}
