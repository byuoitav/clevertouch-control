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
		log.Println("unblanking TV")
		//Blank ON 	= 	AA BB CC 07 4E 00 55 DD EE FF
		payload := []byte{0xAA, 0xBB, 0xCC, 0x07, 0x4E, 0x00, 0x55, 0xDD, 0xEE, 0xFF}
		_, err := PostHTTPWithContext(ctx, address, "blank", payload)
		if err != nil {
			return err
		}
		status = false
	} else {
		log.Println("blanking TV")
		//Blank OFF 	= 	AA BB CC 07 4E 00 55 DD EE FF
		payload := []byte{0xAA, 0xBB, 0xCC, 0x07, 0x4E, 0x00, 0x55, 0xDD, 0xEE, 0xFF}
		_, err := PostHTTPWithContext(ctx, address, "blank", payload)
		if err != nil {
			return err
		}
		status = true
	}

	return nil
}

func GetBlanked(address string) (string, error) {
	return "", nil
}
