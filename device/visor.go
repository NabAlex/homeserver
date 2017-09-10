package device

import "github.com/stianeikeland/go-rpio"

func initVisor() error {
	return rpio.Open()
}

func setPin(pinNumber uint, state DeviceState) {
	pin := rpio.Pin(pinNumber)
	pin.Output()
	switch state {
	case StateOn:
		pin.High()
	case StateOff:
		pin.Low()
	}
}

