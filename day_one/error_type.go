package main

import (
	"errors"
	"fmt"
)

func ErrorTypes() {
	err := errors.New("Data not complete")
	if err != nil {
		fmt.Println(err)
	}
}
