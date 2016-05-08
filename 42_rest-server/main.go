package main

import (
	"bytes"
	"encoding/binary"
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/huin/goserial"
)

// SensorData ...
type SensorData struct {
	Temp float32
	Hum  float32
}

var port io.ReadWriteCloser

// findArduino looks for the file that represents the Arduino
// serial connection. Returns the fully qualified path to the
// device if we are able to find a likely candidate for an
// Arduino, otherwise an empty string if unable to find
// something that 'looks' like an Arduino device.
func findArduino() string {
	contents, _ := ioutil.ReadDir("/dev")

	// Look for what is mostly likely the Arduino device
	for _, f := range contents {
		if strings.Contains(f.Name(), "tty.usbserial") ||
			strings.Contains(f.Name(), "ttyUSB") {
			return "/dev/" + f.Name()
		}
	}

	// Have not been able to find a USB device that 'looks'
	// like an Arduino.
	return ""
}

func init() {
	c := &goserial.Config{Name: findArduino(), Baud: 9600}
	port, _ = goserial.OpenPort(c)

}

// sendArduinoCommand transmits a new command over the nominated serial
// port to the arduino. Returns an error on failure. Each command is
// identified by a single byte and may take one argument (a float).
func sendArduinoCommand(
	command byte,
	argument float32) error {

	// Package argument for transmission
	bufOut := new(bytes.Buffer)
	err := binary.Write(bufOut, binary.LittleEndian, argument)
	if err != nil {
		return err
	}

	// Transmit command and argument down the pipe.
	for _, v := range [][]byte{[]byte{command}, bufOut.Bytes()} {
		_, err = port.Write(v)
		if err != nil {
			return err
		}
	}

	return nil
}

func getSensorData() SensorData {
	var data SensorData

	buf := make([]byte, 64)

	n, _ := port.Read(buf)

	if n > 0 {
		xml.Unmarshal(buf, &data)
	} else {
		log.Printf("%q", buf)
	}

	return data
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

		sendArduinoCommand('h', 0)
		time.Sleep(250 * time.Millisecond)
		sensorData := getSensorData()
		fmt.Fprintf(w, "Temperature: %fC, Humidity: %f%%", sensorData.Temp, sensorData.Hum)

	})

	log.Fatal(http.ListenAndServe(":8080", nil))

}
