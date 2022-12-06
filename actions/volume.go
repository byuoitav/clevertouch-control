package actions

import (
	"bytes"
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

Volume	=	AA BB CC 03 02 00 05 DD EE FF
			AA BB CC 82 00 xx ** DD EE FF Volume is xx
			xx = the amount of volume. EX

Audio Status	= 	AA BB CC 03 03 00 06 DD EE FF
					AA BB CC 82 01 00 83 DD EE FF (Mute)
					AA BB CC 82 01 01 84 DD EE FF (Unmute)

*/

type Volume struct {
	Volume int `json:"volume"`
}

type Mute struct {
	Muted bool `json:"muted"`
}

func GetVolume(address string) (Volume, error) {
	log.Printf("Getting volume for %v", address)
	var output Volume
	payload := []byte{0xAA, 0xBB, 0xCC, 0x03, 0x02, 0x00, 0x05, 0xDD, 0xEE, 0xFF}
	log.Println("getting volume status")
	resp, err := PostHTTP(address, payload, "audio")
	if err != nil {
		return Volume{}, err
	}

	output.Volume = int(binary.BigEndian.Uint64(resp[5:6]))

	return output, nil
}

func getAudioInfo(address string) (int, error) {
	//AA BB CC 03 02 00 05 DD EE FF
	payload := []byte{0xAA, 0xBB, 0xCC, 0x03, 0x02, 0x00, 0x05, 0xDD, 0xEE, 0xFF}
	log.Println("getting volume status")

	resp, err := PostHTTP(address, payload, "audio")
	if err != nil {
		return 0, err
	}

	vol := int(binary.BigEndian.Uint64(resp[5:6]))

	return vol, nil
}

func GetMute(address string) (Mute, error) {
	var output Mute

	// AA BB CC 82 01 00 83 DD EE FF (Mute)
	mute := []byte{0xAA, 0xBB, 0xCC, 0x82, 0x01, 0x00, 0x83, 0xDD, 0xEE, 0xFF}
	// AA BB CC 82 01 01 84 DD EE FF (Unmute)
	unmute := []byte{0xAA, 0xBB, 0xCC, 0x82, 0x01, 0x01, 0x84, 0xDD, 0xEE, 0xFF}

	//AA BB CC 03 03 00 06 DD EE FF
	payload := []byte{0xAA, 0xBB, 0xCC, 0x03, 0x03, 0x00, 0x06, 0xDD, 0xEE, 0xFF}
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
