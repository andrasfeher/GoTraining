package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	readEntireFileIntoMemory()

	f, err := os.Open("./hello.txt")
	check(err)

	read5BytesFromBeginning(f)
	seekToLocation(f)
	rewind(f)
	bufferedRead(f)

	f.Close()
}

func readEntireFileIntoMemory() {
	dat, err := ioutil.ReadFile("./hello.txt")
	check(err)
	fmt.Print(string(dat))
}

func read5BytesFromBeginning(f *os.File) {
	buffer := make([]byte, 5)
	bytesRead, err := f.Read(buffer)
	check(err)
	fmt.Printf("%d bytes read from beginning: %s\n", bytesRead, string(buffer))
}

func seekToLocation(f *os.File) {
	// 0 means: from the beginning of the file
	// 1 would mean: from current position
	// 2 would mean: from the end of the file
	seekPosition, err := f.Seek(8, 1)
	check(err)
	buffer := make([]byte, 2)
	bytesRead, err := f.Read(buffer)
	check(err)
	fmt.Printf("%d bytes read @ %d position: %s\n", bytesRead, seekPosition, string(buffer))
}

func rewind(f *os.File) {
	_, err := f.Seek(0, 0)
	check(err)
}

func bufferedRead(f *os.File) {
	bufReader := bufio.NewReader(f)

	// peek: returns n bytes without advancing the reader
	buffer, err := bufReader.Peek(5)
	check(err)

	fmt.Printf("5 bytes: %s\n", string(buffer))
}
