package actions

import (
	"context"
	"log"
)

/*
Energy AA BB CC 07 4E 00 55 DD EE FF
*/

func SetBlank(ctx context.Context, address string, status bool) error {
	if status {
		log.Println("blanking TV")
		//Blank ON = set to android input
		payload := []byte{0x3A, 0x30, 0x31, 0x53, 0x3A, 0x31, 0x30, 0x31, 0x0d}
		_, err := sendCommand(address, payload)
		if err != nil {
			return err
		}
		//status = false
	} else {
		log.Println("un-blanking TV, restoring previous input")
		//Blank OFF = set to previous input
		SetInput(ctx, address, CurrentInput)
		//status = true
	}
	return nil
}

func GetBlanked(address string) (string, error) {
	return "", nil
}
