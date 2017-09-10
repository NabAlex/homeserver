package device

import "github.com/stianeikeland/go-rpio"

var gpioPool *Pool

func initVisor() error {
	gpioPool = NewPool(5)
	return rpio.Open()
}

func setOutput(pinNumber uint) {
	pin := rpio.Pin(pinNumber)
	pin.Output()
}

/* with relay %)*/
func setStateOnImpl(obj interface{}) {
	pin := obj.(rpio.Pin)
	pin.Low()
}

func setStateOffImpl(obj interface{}) {
	pin := obj.(rpio.Pin)
	pin.High()
}

func setState(pinNumber uint, state DeviceState) {
	pin := rpio.Pin(pinNumber)

	switch state {
	case StateOn:
		gpioPool.ThrowTask(setStateOnImpl, pin)
	case StateOff:
		gpioPool.ThrowTask(setStateOffImpl, pin)
	}
}

