package actions

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

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
		return nil, err
	}

	req, err := http.NewRequestWithContext(ctx, "POST", address, bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	switch {
	case err != nil:
		return nil, err
	case resp.StatusCode != http.StatusOK:
		return nil, errors.New(string(body))
	case body == nil:
		return nil, errors.New("response from device was blank")
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
