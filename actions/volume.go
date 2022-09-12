package actions

/ CleverTouch Command Codes
/*

Action Codes:

Mute TV 	= 	AA BB CC 03 01 00 04 DD EE FF
Unmute TV 	= 	AA BB CC 03 01 01 05 DD EE FF
Volume		=	AA BB CC 03 00 xx ** DD EE FF (000-100)


Response Codes:

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
	var output Volume
	
	return output, nil
}

func getAudioInfo(address string) (string, error) {
	return "", nil
}

func GetMute(address string) (Mute, error) {
	var output Mute

	return output, nil
}
