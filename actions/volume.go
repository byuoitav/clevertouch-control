package actions

import (
	"bytes"
	"context"
	"log"
	"strconv"
)

// CleverTouch Command Codes

/*

SET:

Mute:
39 3A 30 31 53 39 30 30 30 0d = (OFF)
39 3A 30 31 53 39 30 30 31 0d = (ON)

Volume (Range 0 - 100):
To set volume change 7-9th byte digit to match the desired volume value.

38 3A 30 31 53 38 30 35 30 0d = 50%
38 3A 30 31 53 38 31 30 30 0d = 100%
38 3A 30 31 53 38 30 30 30 0d = 0%


GET:

38 3A 30 31 47 38 30 30 30 0D = Get Volume

The response is changes the  7-9th byte to the current volume value.
38 3A 30 31 72 38 30 35 30 0D = 50%

39 3A 30 31 47 39 30 30 30 0D = Get Mute
39 3A 30 31 72 39 30 30 30 0D = Mute OFF
39 3A 30 31 72 39 30 30 31 0D = Mute ON

*/

type Volume struct {
	Volume int `json:"volume"`
}

type Mute struct {
	Muted bool `json:"mute"`
}

var volCodes = map[string]byte{
	"0": 0x30,
	"1": 0x31,
	"2": 0x32,
	"3": 0x33,
	"4": 0x34,
	"5": 0x35,
	"6": 0x36,
	"7": 0x37,
	"8": 0x38,
	"9": 0x39,
}

var volCodesReverse = map[byte]string{
	0x30: "0",
	0x31: "1",
	0x32: "2",
	0x33: "3",
	0x34: "4",
	0x35: "5",
	0x36: "6",
	0x37: "7",
	0x38: "8",
	0x39: "9",
}

func SetMute(ctx context.Context, address string, status bool) error {
	log.Printf("Setting mute to %v", status)
	//Mute ON = 39 3A 30 31 53 39 30 30 31 0d
	//Mute OFF = 39 3A 30 31 53 39 30 30 30 0d
	if status {
		payload := []byte{0x3A, 0x30, 0x31, 0x53, 0x39, 0x30, 0x30, 0x31, 0x0d}
		_, err := sendCommand(address, payload)
		if err != nil {
			return err
		}
		status = false
	} else {
		payload := []byte{0x3A, 0x30, 0x31, 0x53, 0x39, 0x30, 0x30, 0x30, 0x0d}
		_, err := sendCommand(address, payload)
		if err != nil {
			return err
		}
		status = true
	}
	return nil
}

func GetMute(address string) (Mute, error) {
	var output Mute

	//39 3A 30 31 72 39 30 30 31 0D = Mute ON
	mute := []byte{0x3A, 0x30, 0x31, 0x72, 0x39, 0x30, 0x30, 0x31, 0x0D}
	//39 3A 30 31 72 39 30 30 30 0D = Mute OFF
	unmute := []byte{0x3A, 0x30, 0x31, 0x72, 0x39, 0x30, 0x30, 0x30, 0x0D}

	//39 3A 30 31 47 39 30 30 30 0D = Get Mute
	payload := []byte{0x3A, 0x30, 0x31, 0x47, 0x39, 0x30, 0x30, 0x30, 0x0D}
	log.Println("getting mute status")
	resp, err := sendCommand(address, payload)
	if err != nil {
		return Mute{}, err
	} else if bytes.Equal(resp, mute) {
		output.Muted = true
	} else if bytes.Equal(resp, unmute) {
		output.Muted = false
	} else {
		return Mute{}, err
	}

	return output, nil
}

func SetVolume(ctx context.Context, address string, volume int) error {
	log.Printf("Setting volume to %v", volume)
	//38 3A 30 31 53 38 30 30 30 0d = 000%
	vol := strconv.Itoa(volume)
	if len(vol) == 1 {
		vol = "00" + vol
	} else if len(vol) == 2 {
		vol = "0" + vol
	}
	v1 := string(vol[0])
	v2 := string(vol[1])
	v3 := string(vol[2])

	payload := []byte{0x3A, 0x30, 0x31, 0x53, 0x38, volCodes[v1], volCodes[v2], volCodes[v3], 0x0d}
	_, err := sendCommand(address, payload)

	if err != nil {
		return err
	}

	return nil
}

func GetVolume(address string) (Volume, error) {
	log.Printf("Getting volume for %v", address)
	var output Volume

	//38 3A 30 31 47 38 30 30 30 0D = Get Volume
	payload := []byte{0x3A, 0x30, 0x31, 0x47, 0x38, 0x30, 0x30, 0x30, 0x0D}
	log.Println("getting volume status")
	resp, err := sendCommand(address, payload)
	if err != nil {
		return Volume{}, err
	}
	vol := volCodesReverse[resp[5]] + volCodesReverse[resp[6]] + volCodesReverse[resp[7]]
	output.Volume, _ = strconv.Atoi(vol)

	return output, nil
}
