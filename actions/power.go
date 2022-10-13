package actions

import (
	"bytes"
	"context"
	"log"
)

// CleverTouch Command Codes
/*

Action Codes:

Power ON 	= 	AA BB CC 01 00 00 01 DD EE FF
Power OFF 	= 	AA BB CC 01 01 00 02 DD EE FF

Response Codes:

Power	=	AA BB CC 01 02 00 03 DD EE FF
			AA BB CC 80 00 00 80 DD EE FF (ON)
			AA BB CC 80 01 00 81 DD EE FF (OFF)
*/

type Power struct {
	Power string `json:"power"`
}

func SetPower(ctx context.Context, address string, status bool) error {
	if status {
		log.Println("Turning on TV")
		//Power ON 	= 	AA BB CC 01 00 00 01 DD EE FF
		payload := []byte{0xAA, 0xBB, 0xCC, 0x01, 0x00, 0x00, 0x01, 0xDD, 0xEE, 0xFF}
		_, err := PostHTTPWithContext(ctx, address, "power", payload)
		if err != nil {
			return err
		}
	} else {
		log.Println("Turning off TV")
		//Power OFF 	= 	AA BB CC 01 01 00 02 DD EE FF
		payload := []byte{0xAA, 0xBB, 0xCC, 0x01, 0x01, 0x00, 0x02, 0xDD, 0xEE, 0xFF}
		_, err := PostHTTPWithContext(ctx, address, "power", payload)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetPower(ctx context.Context, address string) (Power, error) {
	//AA BB CC 80 00 00 80 DD EE FF (ON)
	on := []byte{0xAA, 0xBB, 0xCC, 0x80, 0x00, 0x00, 0x80, 0xDD, 0xEE, 0xFF}
	//AA BB CC 80 01 00 81 DD EE FF (OFF)
	off := []byte{0xAA, 0xBB, 0xCC, 0x80, 0x01, 0x00, 0x81, 0xDD, 0xEE, 0xFF}

	var output Power
	//Power	=	AA BB CC 01 02 00 03 DD EE FF
	payload := []byte{0xAA, 0xBB, 0xCC, 0x01, 0x02, 0x00, 0x03, 0xDD, 0xEE, 0xFF}
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
