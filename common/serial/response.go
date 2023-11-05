package serial

import (
	"fmt"
	"io"
)

func GetResponse(port io.ReadWriteCloser, header []byte, body []byte) error {
	err := FilterSerial(port, header)
	if err != nil {
		return fmt.Errorf("serial: failed to filter header")
	}

	err = FilterSerial(port, body)
	if err != nil {
		return fmt.Errorf("serial: failed to filter body")
	}

	return nil
}
