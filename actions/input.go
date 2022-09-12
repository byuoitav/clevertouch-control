package actions

// CleverTouch Command Codes
/*

Action Codes:

Input Tv 	= 	AA BB CC 02 01 00 03 DD EE FF
Input HDMI1 = 	AA BB CC 02 06 00 08 DD EE FF
Input HDMI2 = 	AA BB CC 02 07 00 09 DD EE FF
Input HDMI3 = 	AA BB CC 02 05 00 07 DD EE FF
Input PC 	= 	AA BB CC 02 08 00 0A DD EE FF

Response Codes:

Input Satus		=	AA BB CC 02 00 00 02 DD EE FF
					AA BB CC 81 01 00 82 DD EE FF (TV)
					AA BB CC 81 06 00 87 DD EE FF (HDMI1)
					AA BB CC 81 07 00 88 DD EE FF (HDMI2)
					AA BB CC 81 05 00 86 DD EE FF (HDMI3)
*/

type Input struct {
	Input string `json:"input,omitempty"`
}

// GetInput returns the input being shown on the display
func GetInput(address string) (Input, error) {
	var output Input

	return output, nil
}

// GetActiveSignal determines if the current input on the display is active or not
func GetActiveSignal(address string) (string, error) {
	return "", nil
}
