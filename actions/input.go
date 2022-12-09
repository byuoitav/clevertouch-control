package actions

import (
	"bytes"
	"context"
	"errors"
	"log"
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

	// 3A 3A 30 31 47 3A 30 30 30 0D = Get Input
	payload := []byte{0x3A, 0x3A, 0x30, 0x31, 0x47, 0x3A, 0x30, 0x30, 0x30, 0x0D}
	log.Println("getting input")
	resp, err := PostHTTP(address, payload, "input")
	if err != nil {
		return output, err
	}
	log.Println("got input")
	switch {
	case bytes.Contains(resp, []byte{0x3A, 0x3A, 0x30, 0x31, 0x72, 0x3A, 0x30, 0x30, 0x31, 0x0D}):
		output.Input = "hdmi1"
	case bytes.Contains(resp, []byte{0x3A, 0x3A, 0x30, 0x31, 0x72, 0x3A, 0x30, 0x30, 0x32, 0x0D}):
		output.Input = "hdmi2"
	case bytes.Contains(resp, []byte{0x3A, 0x3A, 0x30, 0x31, 0x72, 0x3A, 0x30, 0x32, 0x31, 0x0D}):
		output.Input = "hdmi3"
	case bytes.Contains(resp, []byte{0x3A, 0x3A, 0x30, 0x31, 0x72, 0x3A, 0x30, 0x32, 0x32, 0x0D}):
		output.Input = "hdmi4"
	case bytes.Contains(resp, []byte{0x3A, 0x3A, 0x30, 0x31, 0x72, 0x3A, 0x30, 0x35, 0x31, 0x0D}):
		output.Input = "tv"
	case bytes.Contains(resp, []byte{0x3A, 0x3A, 0x30, 0x31, 0x72, 0x3A, 0x30, 0x30, 0x37, 0x0D}):
		output.Input = "dp"
	case bytes.Contains(resp, []byte{0x3A, 0x3A, 0x30, 0x31, 0x72, 0x3A, 0x31, 0x30, 0x34, 0x0D}):
		output.Input = "type-c1"
	case bytes.Contains(resp, []byte{0x3A, 0x3A, 0x30, 0x31, 0x72, 0x3A, 0x31, 0x30, 0x35, 0x0D}):
		output.Input = "type-c2"
	default:
		return output, errors.New("unknown input")
	}
	return output, nil
}

// SetInput sets the input on the display
func SetInput(ctx context.Context, address string, input Input) error {
	var payload []byte

	log.Println("input: ", input.Input)
	switch input.Input {
	case "tv":
		// 3A 3A 30 31 53 3A 30 35 31 0d = TV
		payload = []byte{0x3A, 0x3A, 0x30, 0x31, 0x53, 0x3A, 0x30, 0x35, 0x31, 0x0d}
	case "hdmi1":
		// 3A 3A 30 31 53 3A 30 30 31 0d = HDMI1
		payload = []byte{0x3A, 0x3A, 0x30, 0x31, 0x53, 0x3A, 0x30, 0x30, 0x31, 0x0d}
	case "hdmi2":
		// 3A 3A 30 31 53 3A 30 30 32 0d = HDMI2
		payload = []byte{0x3A, 0x3A, 0x30, 0x31, 0x53, 0x3A, 0x30, 0x30, 0x32, 0x0d}
	case "hdmi3":
		// 3A 3A 30 31 53 3A 30 32 31 0d = HDMI3
		payload = []byte{0x3A, 0x3A, 0x30, 0x31, 0x53, 0x3A, 0x30, 0x32, 0x31, 0x0d}
	case "hdmi4":
		// 3A 3A 30 31 53 3A 30 32 32 0d = HDMI4
		payload = []byte{0x3A, 0x3A, 0x30, 0x31, 0x53, 0x3A, 0x30, 0x32, 0x32, 0x0d}
	case "dp":
		// 3A 3A 30 31 53 3A 30 30 37 0d = DP
		payload = []byte{0x3A, 0x3A, 0x30, 0x31, 0x53, 0x3A, 0x30, 0x30, 0x37, 0x0d}
	case "type-c1":
		// 3A 3A 30 31 53 3A 31 30 34 0d = Type-C1
		payload = []byte{0x3A, 0x3A, 0x30, 0x31, 0x53, 0x3A, 0x31, 0x30, 0x34, 0x0d}
	case "type-c2":
		// 3A 3A 30 31 53 3A 31 30 35 0d = Type-C2
		payload = []byte{0x3A, 0x3A, 0x30, 0x31, 0x53, 0x3A, 0x31, 0x30, 0x35, 0x0d}
	default:
		return errors.New("unknown input")
	}
	_, err := PostHTTPWithContext(ctx, address, "input", payload)
	if err != nil {
		return err
	}
	return nil
}

// GetActiveSignal determines if the current input on the display is active or not
func GetActiveSignal(address string) (string, error) {
	return "", nil
}
