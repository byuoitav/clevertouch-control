package actions

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
)

// CleverTouch Command Codes
/*

Action Codes:

Power ON 	= 	AA BB CC 01 00 00 01 DD EE FF
Power OFF 	= 	AA BB CC 01 01 00 02 DD EE FF

Input Tv 	= 	AA BB CC 02 01 00 03 DD EE FF
Input HDMI1 = 	AA BB CC 02 06 00 08 DD EE FF
Input HDMI2 = 	AA BB CC 02 07 00 09 DD EE FF
Input HDMI3 = 	AA BB CC 02 05 00 07 DD EE FF
Input PC 	= 	AA BB CC 02 08 00 0A DD EE FF

Mute TV 	= 	AA BB CC 03 01 00 04 DD EE FF
Unmute TV 	= 	AA BB CC 03 01 01 05 DD EE FF
Volume		=	AA BB CC 03 00 xx ** DD EE FF (000-100)


Response Codes:

Power	=	AA BB CC 01 02 00 03 DD EE FF
			AA BB CC 80 00 00 80 DD EE FF (ON)
			AA BB CC 80 01 00 81 DD EE FF (OFF)

Volume	=	AA BB CC 03 02 00 05 DD EE FF
			AA BB CC 82 00 xx ** DD EE FF Volume is xx
			xx = the amount of volume. EX

Audio Status	= 	AA BB CC 03 03 00 06 DD EE FF
					AA BB CC 82 01 00 83 DD EE FF (Mute)
					AA BB CC 82 01 01 84 DD EE FF (Unmute)

Input Satus		=	AA BB CC 02 00 00 02 DD EE FF
					AA BB CC 81 01 00 82 DD EE FF (TV)
					AA BB CC 81 06 00 87 DD EE FF (HDMI1)
					AA BB CC 81 07 00 88 DD EE FF (HDMI2)
					AA BB CC 81 05 00 86 DD EE FF (HDMI3)

*/

// CTAudioResponse is the struct that is returned when we query the adio state
type CTAudioResponse struct {
	Result [][]CTAudioSettings `json:"result"`
	ID     int                 `json:"id"`
}

// CTAudioSettings is the struct that holds the audio settings data
type CTAudioSettings struct {
	Target    string `json:"target"`
	Volume    int    `json:"volume"`
	Mute      bool   `json:"mute"`
	MaxVolume int    `json:"maxVolume"`
	MinVolume int    `json:"minVolume"`
}

type CTAVContentSettings struct {
	URI        string `json:"uri"`
	Source     string `json:"source"`
	Title      string `json:"title"`
	Status     string `json:"status"`
	Connection string `json:"connection"`
}

type CTAVContentResponse struct {
	Result [][]CTAVContentSettings `json:"result"`
	ID     int                     `json:"id"`
}

type CTMultiAvContentResponse struct {
	Result []CTAVContentSettings `json:"result"`
	ID     int                   `json:"id"`
}

type CTTvRequest struct {
	Method  string                   `json:"method"`
	Version string                   `json:"version"`
	ID      int                      `json:"id"`
	Params  []map[string]interface{} `json:"params"`
}

type CTTvSystemResponse struct {
	Result []CTTvSystemInformation `json:"result"`
	ID     int                     `json:"id"`
}

type CTTvSystemInformation struct {
	Product    string `json:"product"`
	Region     string `json:"region,omitempty"`
	Language   string `json:"language,omitempty"`
	Model      string `json:"model"`
	Serial     string `json:"serial,omitempty"`
	MAC        string `json:"macAddr,omitempty"`
	Name       string `json:"name"`
	Generation string `json:"generation,omitempty"`
	Area       string `json:"area,omitempty"`
	CID        string `json:"cid,omitempty"`
}

type CTNetworkResponse struct {
	ID     int                        `json:"id"`
	Result [][]CTTvNetworkInformation `json:"result"`
}

type CTTvNetworkInformation struct {
	NetworkInterface string `json:"networkInterface"`
	HardwareAddress  string `json:"hardwareAddress"`
	IPv4             string `json:"ipv4"`
	IPv6             string `json:"ipv6"`
	Netmask          string `json:"netmask"`
	Gateway          string `json:"gateway"`
	DNS              string `json:"dns"`
}

func PostHTTPWithContext(ctx context.Context, address string, service string, payload []byte) ([]byte, error) {

	reqBody, err := json.Marshal(payload)
	if err != nil {
		return []byte{}, err
	}

	addr := fmt.Sprintf("http://%s/CleverTouch/%s", address, service)

	req, err := http.NewRequestWithContext(ctx, "POST", addr, bytes.NewBuffer(reqBody))
	if err != nil {
		return []byte{}, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-Auth-PSK", os.Getenv(""))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return []byte{}, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	switch {
	case err != nil:
		return []byte{}, err
	case resp.StatusCode != http.StatusOK:
		return []byte{}, errors.New(string(body))
	case body == nil:
		return []byte{}, errors.New("Response from device was blank")
	}

	return body, nil
}

func PostHTTP(address string, payload []byte, service string) ([]byte, error) {

	return PostHTTPWithContext(context.TODO(), address, service, payload)
}

func BuildAndSendPayload(address string, service string, method string, params map[string]interface{}) error {

	payload := CTTvRequest{
		Params:  []map[string]interface{}{params},
		Method:  method,
		Version: "1.0",
		ID:      1,
	}

	_, err := PostHTTP(address, payload, service)

	return err
}
