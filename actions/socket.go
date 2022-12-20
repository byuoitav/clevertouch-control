package actions

import (
	"bufio"
	//"context"
	//"encoding/xml"
	//"errors"
	"fmt"
	"net"

	//"regexp"
	//"strconv"
	"time"
)

// Creating Connection
func createConnection(address string, port string) (*net.TCPConn, error) {
	radder, err := net.ResolveTCPAddr("tcp", address+":"+port)
	if err != nil {
		err = fmt.Errorf("error resolving address : %s", err.Error())
		return nil, err
	}

	conn, err := net.DialTCP("tcp", nil, radder)
	if err != nil {
		err = fmt.Errorf("error dialing address : %s", err.Error())
		return nil, err
	}

	return conn, nil
}

// SendCommand opens a connection with <addr> and sends the <command> to the via, returning the response from the via, or an error if one occured.
func sendCommand(address string, cmd []byte) ([]byte, error) {
	port := "4999"

	// get the connection
	fmt.Printf("Opening telnet connection with %s\n", address)
	conn, err := createConnection(address, port)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	bufClear := []byte{0x0D}

	timeoutDuration := 3 * time.Second

	// Set Read Connection Duration
	conn.SetReadDeadline(time.Now().Add(timeoutDuration))

	// write command
	if len(cmd) > 0 {
		firstCmd := append(bufClear, cmd...)
		fcmd := append(firstCmd, bufClear...)
		_, err = conn.Write(fcmd)
		if err != nil {
			return nil, err
		}
	}

	reader := bufio.NewReader(conn)
	resp, err := reader.ReadBytes('\r')
	if err != nil {
		err = fmt.Errorf("error reading from system: %s\n", err.Error())
		fmt.Printf(err.Error())
		return nil, err
	}

	fmt.Printf("Response from device: %x\n", resp)

	return resp, nil
}
