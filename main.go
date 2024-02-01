package main

import (
	"fmt"
	"time"

	log "github.com/schollz/logger"
	"golang.org/x/exp/io/i2c"
)

// ADS7830Read reads from the ADS7830 ADC on the specified channel.
func ADS7830Read(ch int) (val int, err error) {
	// Define the base command byte.
	var commandByte byte

	// Select the channel.
	switch ch {
	case 0:
		commandByte = 0b10000100
	case 1:
		commandByte = 0b11000100
	case 2:
		commandByte = 0b10010100
	case 3:
		commandByte = 0b11010100
	case 4:
		commandByte = 0b10100100
	case 5:
		commandByte = 0b11100100
	case 6:
		commandByte = 0b11000100
	case 7:
		commandByte = 0b11110100
	default:
		return 0, fmt.Errorf("invalid channel")
	}

	d, err := i2c.Open(&i2c.Devfs{Dev: "/dev/i2c-1"}, 0x48)
	if err != nil {
		log.Error(err)
		return
	}
	err = d.Write([]byte{commandByte})
	if err != nil {
		log.Error(err)
		return
	}

	// read the byte
	buf := make([]byte, 1)
	err = d.Read(buf)
	if err != nil {
		log.Error(err)
		return
	}

	// print out binary value
	log.Debugf("Read: %08b", buf[0])

	return
}

func main() {
	log.SetLevel("debug")
	for {
		_, err := ADS7830Read(0)
		if err != nil {
			log.Error(err)
		}
		time.Sleep(100 * time.Millisecond)
	}

}
