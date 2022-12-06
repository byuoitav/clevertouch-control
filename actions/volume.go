package actions

import (
	"bytes"
	"context"
	"encoding/binary"
	"log"
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

func SetMute(ctx context.Context, address string, status bool) error {
	log.Printf("Setting mute to %v", status)
	//Mute ON = 39 3A 30 31 53 39 30 30 31 0d
	//Mute OFF = 39 3A 30 31 53 39 30 30 30 0d
	if status {
		payload := []byte{0x39, 0x3A, 0x30, 0x31, 0x53, 0x39, 0x30, 0x30, 0x31, 0x0d}
		_, err := PostHTTPWithContext(ctx, address, "audio", payload)
		if err != nil {
			return err
		}
		status = false
	} else {
		payload := []byte{0x39, 0x3A, 0x30, 0x31, 0x53, 0x39, 0x30, 0x30, 0x30, 0x0d}
		_, err := PostHTTPWithContext(ctx, address, "audio", payload)
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
	mute := []byte{0x39, 0x3A, 0x30, 0x31, 0x72, 0x39, 0x30, 0x30, 0x31, 0x0D}
	//39 3A 30 31 72 39 30 30 30 0D = Mute OFF
	unmute := []byte{0x39, 0x3A, 0x30, 0x31, 0x72, 0x39, 0x30, 0x30, 0x30, 0x0D}

	//39 3A 30 31 47 39 30 30 30 0D = Get Mute
	payload := []byte{0x39, 0x3A, 0x30, 0x31, 0x47, 0x39, 0x30, 0x30, 0x30, 0x0D}
	log.Println("getting mute status")
	resp, err := PostHTTP(address, payload, "audio")
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
	//38 3A 30 31 53 38 30 35 30 0d = 50%
	payload := []byte{0x38, 0x3A, 0x30, 0x31, 0x53, 0x38, 0x30, 0x35, 0x30, 0x0d}
	_, err := PostHTTPWithContext(ctx, address, "audio", payload)

	if err != nil {
		return err
	}

	return nil
}

func GetVolume(address string) (Volume, error) {
	log.Printf("Getting volume for %v", address)
	var output Volume

	//38 3A 30 31 47 38 30 30 30 0D = Get Volume
	payload := []byte{0x38, 0x3A, 0x30, 0x31, 0x47, 0x38, 0x30, 0x30, 0x30, 0x0D}
	log.Println("getting volume status")
	resp, err := PostHTTP(address, payload, "audio")
	if err != nil {
		return Volume{}, err
	}

	output.Volume = int(binary.BigEndian.Uint64(resp[5:6]))

	return output, nil
}
