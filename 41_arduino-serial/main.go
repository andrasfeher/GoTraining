package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/huin/goserial"
)

// SensorData ...
type SensorData struct {
	Temp float32
	Hum  float32
}

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

func main() {
	// Find the device that represents the arduino serial
	// connection.
	c := &goserial.Config{Name: findArduino(), Baud: 9600}
	s, _ := goserial.OpenPort(c)
	var data SensorData

	buf := make([]byte, 128)
	/*
		testBuf := []byte(`{"Temp":10.20, "Hum":50.10}`)
		err := json.Unmarshal(testBuf, &data)
		fmt.Println("error:", err)
		fmt.Println(data)
	*/

	for {
		n, _ := s.Read(buf)

		if n > 0 {
			xml.Unmarshal(buf, &data)
			// fmt.Println("error:", err)

			// log.Printf("%q", buf[:n])
			fmt.Println(data.Temp, "C", " - ", data.Hum, "%")
		}

		time.Sleep(1 * time.Second)
	}
}
