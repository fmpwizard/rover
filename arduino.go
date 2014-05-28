package main

//Most code on this page is from http://reprage.com/post/using-golang-to-connect-raspberrypi-and-arduino/

import (
	"bytes"
	"encoding/binary"
	"errors"
	"github.com/huin/goserial"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

var c = &goserial.Config{}
var s io.ReadWriteCloser

//We init the usb connection only once, at boot time.
func init() {
	// Find the device that represents the arduino serial
	// connection.
	c = &goserial.Config{Name: findArduino(), Baud: 9600}
	log.Printf("the USB port the Arduino service will use is %v\n", findArduino())
	s, _ = goserial.OpenPort(c)

}

//rduino converts the string command (on/off) to the one leter command
// the arduino board expects. And call sendArduinoCommand
func arduinoDo(command string, value int) error {
	return sendArduinoCommand(command, uint32(value), s) //u for up, d for down
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
			strings.Contains(f.Name(), "ttyUSB") ||
			strings.Contains(f.Name(), "ttyACM") {
			return "/dev/" + f.Name()
		}
	}

	// Have not been able to find a USB device that 'looks'
	// like an Arduino.
	return ""
}

// sendArduinoCommand transmits a new command over the nominated serial
// port to the arduino. Returns an error on failure. Each command is
// identified by a single byte and may take one argument (a float).
func sendArduinoCommand(command string, argument uint32, serialPort io.ReadWriteCloser) error {
	if serialPort == nil {
		return errors.New("No serial port to talk to the Arduino.")
	}

	// Package argument for transmission
	value := new(bytes.Buffer)
	err := binary.Write(value, binary.LittleEndian, argument)
	if err != nil {
		return err
	}
	cmd := new(bytes.Buffer)
	err = binary.Write(value, binary.LittleEndian, command)
	if err != nil {
		return err
	}

	// Transmit command and argument down the pipe.
	for _, v := range [][]byte{cmd.Bytes(), value.Bytes()} {
		_, err = serialPort.Write(v)
		if err != nil {
			return err
		}
	}

	return nil
}
