package actions

import (
	"bytes"
	"context"
	"log"
)

/*
We are setting the Backlight to 100% (0x64) to turn the TV on
and 0% (0x00) to turn the TV off. This is a workaround because the on/off command
takes a long time to respond and the NIC loses power when powered off.
*/

// CleverTouch Command Codes
/*

SET:

Backlight:
30 3A 30 31 53 30 30 30 30 0d = (OFF) Backlight 0%
30 3A 30 31 53 30 30 30 31 0d = (ON) Backlight 100%


GET:
30 3A 30 31 47 30 30 30 30 0D = Get Backlight (OFF/ON)
30 3A 30 31 72 30 30 30 30 0D = (OFF) Backlight 0%
30 3A 30 31 72 30 30 30 31 0D = (ON) Backlight 100%

*/

type Power struct {
	Power string `json:"power"`
}

type Booted struct {
	Booted string `json:"booted"`
}

func SetPower(ctx context.Context, address string, status bool) error {
	if status {
		log.Println("Turning on TV")
		//Power ON = 30 3A 30 31 53 30 30 30 31 0d
		payload := []byte{0x3A, 0x30, 0x31, 0x53, 0x30, 0x30, 0x30, 0x31, 0x0d}
		_, err := sendCommand(address, payload)
		if err != nil {
			return err
		}
	} else {
		log.Println("Turning off TV")
		//Power OFF = 30 3A 30 31 53 30 30 30 30 0d
		payload := []byte{0x3A, 0x30, 0x31, 0x53, 0x30, 0x30, 0x30, 0x30, 0x0d}
		_, err := sendCommand(address, payload)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetPower(ctx context.Context, address string) (Power, error) {
	//30 3A 30 31 72 30 30 30 31 0D (ON)
	on := []byte{0x3A, 0x30, 0x31, 0x72, 0x30, 0x30, 0x30, 0x31, 0x0D}
	//30 3A 30 31 72 30 30 30 30 0D (OFF)
	off := []byte{0x3A, 0x30, 0x31, 0x72, 0x30, 0x30, 0x30, 0x30, 0x0D}

	var output Power
	//GetPower = 30 3A 30 31 47 30 30 30 30 0D
	payload := []byte{0x3A, 0x30, 0x31, 0x47, 0x30, 0x30, 0x30, 0x30, 0x0D}
	log.Println("getting power status")
	response, err := sendCommand(address, payload)
	if err != nil {
		return Power{}, err
	}
	log.Println("power status: ", response)

	if bytes.Equal(response, on) {
		output.Power = "ON"
	} else if bytes.Equal(response, off) {
		output.Power = "OFF"
	} else {
		output.Power = "UNKNOWN"
	}

	return output, nil
}

func GetBooted(ctx context.Context, address string) (Booted, error) {
	var boot Booted
	boot.Booted = "System Booted"
	payload := []byte{0x3A, 0x30, 0x31, 0x47, 0x30, 0x30, 0x30, 0x30, 0x0D}
	log.Println("getting boot power status")
	response, err := sendCommand(address, payload)
	if err != nil {
		return Booted{}, err
	}
	log.Println("power booted status: ", response)

	if bytes.Contains(response, []byte{0x3A, 0x30, 0x31, 0x72, 0x30, 0x30, 0x30, 0x32, 0x0D}) {
		log.Println("Booting TV")
		//Power ON = 30 3A 30 31 53 30 30 30 33 0d
		payload := []byte{0x3A, 0x30, 0x31, 0x53, 0x30, 0x30, 0x30, 0x33, 0x0d}
		_, err2 := sendCommand(address, payload)
		if err != nil {
			return Booted{}, err2
		}
		log.Println("boot power was off, turning display on")
		boot.Booted = "Powering On"
	}

	return boot, nil
}
