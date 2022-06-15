// Package million implements functions to demonstrate how to work with panic and recovery functions
//
// The MillionFiles creates million files in the current directory
//
// MillionFiles()
//
// You are allowed to call this function and get one million files
package million

import (
	"fmt"
	"os"
)

func MillionFiles() {
	for i := 0; i < 10; i++ {
		create(i)
	}
}

func create(i int) {
	defer func() {
		if v := recover(); v != nil {
			fmt.Printf("Recovered, last index %v\n", i)
		}
	}()
	file, err := os.Create(fmt.Sprintf("trash/%d", i))
	if err != nil {
		panic("Cannot create the file")
	}
	if i == 1000 {
		panic("cannot create the file")
	}
	defer file.Close()
}
