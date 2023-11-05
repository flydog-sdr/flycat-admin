package serial

import (
	"fmt"
	"io"

	"github.com/albenik/go-serial/v2"
)

func PairDevice() (io.ReadWriteCloser, error) {
	list, err := serial.GetPortsList()
	if err != nil {
		return nil, err
	}

	for _, port := range list {
		p := OpenSerial(port, 9600)

		_, err := p.Write([]byte{0xFF})
		if err != nil {
			CloseSerial(p)
			continue
		}

		err = GetResponse(p, []byte{0x2B}, []byte{
			0x53, 0x54, 0x41, 0x54, 0x55, 0x53,
			0x3A, 0x52, 0x45, 0x41, 0x44, 0x59,
		})
		if err != nil {
			CloseSerial(p)
			continue
		}

		return p, nil
	}

	return nil, fmt.Errorf("serial: failed to pair device")
}
