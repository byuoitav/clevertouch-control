package actions

import (
	"bytes"
	"context"
	"errors"
)

// CleverTouch Command Codes
/*

SET:

Video Source:

3A 3A 30 31 53 3A 30 30 31 0d = HDMI1
3A 3A 30 31 53 3A 30 30 32 0d = HDMI2
3A 3A 30 31 53 3A 30 32 31 0d = HDMI3
3A 3A 30 31 53 3A 30 32 32 0d = HDMI4
3A 3A 30 31 53 3A 30 35 31 0d = TV
3A 3A 30 31 53 3A 30 30 37 0d = DP
3A 3A 30 31 53 3A 31 30 34 0d = Type-C1
3A 3A 30 31 53 3A 31 30 35 0d = Type-C2

GET:

3A 3A 30 31 47 3A 30 30 30 0D = Get Input
3A 3A 30 31 72 3A 30 30 31 0D = HDMI1
3A 3A 30 31 72 3A 30 30 32 0D = HDMI2
3A 3A 30 31 72 3A 30 32 31 0D = HDMI3
3A 3A 30 31 72 3A 30 32 32 0D = HDMI4
3A 3A 30 31 72 3A 30 35 31 0D = TV
3A 3A 30 31 72 3A 30 30 37 0D = DP
3A 3A 30 31 72 3A 31 30 34 0D = Type-C1
3A 3A 30 31 72 3A 31 30 35 0D = Type-C2

*/

type Input struct {
	Input string `json:"input,omitempty"`
}

// GetInput returns the input being shown on the display
func GetInput(address string) (Input, error) {
	var output Input

	// AA BB CC 81 01 00 82 DD EE FF (TV)
	tv := []byte{0xAA, 0xBB, 0xCC, 0x81, 0x01, 0x00, 0x82, 0xDD, 0xEE, 0xFF}
	// AA BB CC 81 06 00 87 DD EE FF (HDMI1)
	hdmi1 := []byte{0xAA, 0xBB, 0xCC, 0x81, 0x06, 0x00, 0x87, 0xDD, 0xEE, 0xFF}
	// AA BB CC 81 07 00 88 DD EE FF (HDMI2)
	hdmi2 := []byte{0xAA, 0xBB, 0xCC, 0x81, 0x07, 0x00, 0x88, 0xDD, 0xEE, 0xFF}
	// AA BB CC 81 05 00 86 DD EE FF (HDMI3)
	hdmi3 := []byte{0xAA, 0xBB, 0xCC, 0x81, 0x05, 0x00, 0x86, 0xDD, 0xEE, 0xFF}

	pwrState, err := GetPower(context.TODO(), address)
	if err != nil {
		return output, err
	}
	if pwrState.Power != "on" {
		return output, errors.New("display is off")
	}

	// AA BB CC 02 00 00 02 DD EE FF
	payload := []byte{0xAA, 0xBB, 0xCC, 0x02, 0x00, 0x00, 0x02, 0xDD, 0xEE, 0xFF}

	response, err := PostHTTP(address, payload, "avContent")
	if err != nil {
		return output, err
	}

	if bytes.Equal(response, tv) {
		output.Input = "tv"
	} else if bytes.Equal(response, hdmi1) {
		output.Input = "hdmi1"
	} else if bytes.Equal(response, hdmi2) {
		output.Input = "hdmi2"
	} else if bytes.Equal(response, hdmi3) {
		output.Input = "hdmi3"
	} else {
		return output, errors.New("unknown input")
	}

	return output, nil
}

// GetActiveSignal determines if the current input on the display is active or not
func GetActiveSignal(address string) (string, error) {
	return "", nil
}
