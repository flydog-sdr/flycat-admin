package controller

import (
	"fmt"
	"io"
	"math"

	"github.com/flydog-sdr/flycat-admin/common/serial"
)

func SetRegister(port io.ReadWriteCloser, gain, antenna, filter int) error {
	if gain > 34 || gain < -11 {
		return fmt.Errorf("gain value is out of range")
	}

	if antenna > 2 || antenna < 0 {
		return fmt.Errorf("antenna value is out of range")
	}

	if filter != 0 && filter != 1 {
		return fmt.Errorf("notch filter value is out of range")
	}

	var symbol int
	if gain < 0 {
		symbol = NEGATIVE_GAIN
	} else {
		symbol = POSITIVE_GAIN
	}

	_, err := port.Write([]byte{SET_FILTER, byte(filter)})
	if err != nil {
		return fmt.Errorf("serial: failed to write notch filter value")
	}

	err = serial.GetResponse(port, []byte{RESPONSE_HEADER}, []byte{
		0x52, 0x45, 0x53, 0x50, 0x4F, 0x4E, 0x53,
		0x45, 0x3A, 0x44, 0x4F, 0x4E, 0x45,
	})
	if err != nil {
		return err
	}

	_, err = port.Write([]byte{SET_REGISTER, byte(int(math.Abs(float64(gain)))), byte(symbol), byte(antenna)})
	if err != nil {
		return fmt.Errorf("serial: failed to write register values")
	}

	err = serial.GetResponse(port, []byte{RESPONSE_HEADER}, []byte{
		0x52, 0x45, 0x53, 0x50, 0x4F, 0x4E, 0x53,
		0x45, 0x3A, 0x44, 0x4F, 0x4E, 0x45,
	})
	if err != nil {
		return err
	}

	return nil
}
