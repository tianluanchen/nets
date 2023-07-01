package pkg

import (
	"fmt"
	"os"
)

func CheckErr(err error) {
	if err != nil {
		ExitWithError(err)
	}
}

func ExitWithError(err interface{}) {
	fmt.Fprintln(os.Stderr, err)
	os.Exit(1)
}
