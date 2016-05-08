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

	defer f.Close()

	read5BytesFromBeginning(f)
	seekToLocation(f)
	rewind(f)
	bufferedRead(f)

	fw, err := os.Create("./hello3.txt")
	check(err)

	defer fw.Close()

	dumpStringIntoFile()
	writeByteSlice(fw)
	writeString(fw)

	// flush to storage
	fw.Sync()

	bufferedWriting(fw)

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

func dumpStringIntoFile() {
	buffer := []byte("hello")
	err := ioutil.WriteFile("./hello2.txt", buffer, 0644)
	check(err)

}

func writeByteSlice(fw *os.File) {
	byteSlice := []byte{115, 111, 109, 101, 10}
	wroteBytesCount, err := fw.Write(byteSlice)
	check(err)
	fmt.Printf("wrote %d bytes\n", wroteBytesCount)

}

func writeString(fw *os.File) {
	wroteBytesCount, err := fw.WriteString("hello world")
	check(err)
	fmt.Printf("wrote %d bytes\n", wroteBytesCount)
}

func bufferedWriting(fw *os.File) {
	bufferedWriter := bufio.NewWriter(fw)
	wroteBytesCount, err := bufferedWriter.WriteString("buffered hello world\n")
	check(err)
	fmt.Printf("wrote %d bytes\n", wroteBytesCount)
	bufferedWriter.Flush()
}
