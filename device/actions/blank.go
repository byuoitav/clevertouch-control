package actions

import (
	"context"
)

/*
Energy AA BB CC 07 4E 00 55 DD EE FF
*/

func SetBlank(ctx context.Context, address string, status bool) error {
	// payload := []byte{0x3A, 0x30, 0x31, 0x53, 0x3A, 0x31, 0x30, 0x31, 0x0d} //set input to android
	var payload []byte
	if status {
		//Blank ON = set to android input
		payload = []byte{0x3A, 0x30, 0x31, 0x53, 0x30, 0x30, 0x30, 0x30, 0x0d} //set backlight to 0 for blank
	} else {
		//Blank OFF = set to previous input
		payload = []byte{0x3A, 0x30, 0x31, 0x53, 0x30, 0x30, 0x30, 0x31, 0x0d} //set backlight on for un-blank
		//SetInput(ctx, address, CurrentInput)
	}
	_, err := sendCommand(address, payload)
	if err != nil {
		return err
	}
	return nil
}

func GetBlanked(address string) (string, error) {
	return "", nil
}
