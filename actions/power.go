package actions

import (
	"context"
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
	return nil
}

func GetPower(ctx context.Context, address string) (Power, error) {
	var output Power

	response, err := PostHTTPWithContext(ctx, address, []byte{0xAA, 0xBB, 0xCC, 0x01, 0x02, 0x00, 0x03, 0xDD, 0xEE, 0xFF})
	if err != nil {
		return Power, err
	}

	return output, nil
}
