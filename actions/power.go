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

func SetPower(ctx context.Context, address string, status bool) error {
	if status {
		log.Println("Turning on TV")
		//Power ON = 30 3A 30 31 53 30 30 30 31 0d
		payload := []byte{0x30, 0x3A, 0x30, 0x31, 0x53, 0x30, 0x30, 0x30, 0x31, 0x0d}
		_, err := PostHTTPWithContext(ctx, address, "power", payload)
		if err != nil {
			return err
		}
	} else {
		log.Println("Turning off TV")
		//Power OFF = 30 3A 30 31 53 30 30 30 30 0d
		payload := []byte{0x30, 0x3A, 0x30, 0x31, 0x53, 0x30, 0x30, 0x30, 0x30, 0x0d}
		_, err := PostHTTPWithContext(ctx, address, "power", payload)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetPower(ctx context.Context, address string) (Power, error) {
	//30 3A 30 31 72 30 30 30 31 0D (ON)
	on := []byte{0x30, 0x3A, 0x30, 0x31, 0x72, 0x30, 0x30, 0x30, 0x31, 0x0D}
	//30 3A 30 31 72 30 30 30 30 0D (OFF)
	off := []byte{0x30, 0x3A, 0x30, 0x31, 0x72, 0x30, 0x30, 0x30, 0x30, 0x0D}

	var output Power
	//GetPower = 30 3A 30 31 47 30 30 30 30 0D
	payload := []byte{0x30, 0x3A, 0x30, 0x31, 0x47, 0x30, 0x30, 0x30, 0x30, 0x0D}
	log.Println("getting power status")
	response, err := PostHTTPWithContext(ctx, address, "system", payload)
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
