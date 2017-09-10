package device

import (
	"io/ioutil"
)

type DeviceState string
var (
	StateOn 	DeviceState = "on"
	StateOff 	DeviceState = "off"
)

func ConvertToState(state string) (DeviceState, bool) {
	switch state {
	case "on":
		return StateOn, true
	case "off":
		return StateOff, true
	}

	return "", false
}

type Device struct {
	Name string `json:"name"`
	Pin uint 	`json:"pin"`
}

type DeviceStatus struct {
	Pin uint
	State DeviceState
}

var deviceInfo map[string]DeviceStatus

func InitDevice(filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return  err
	}

	devices, err := parseDevices(string(bytes))
	if err != nil {
		return err
	}

	deviceInfo = make(map[string]DeviceStatus)
	for _, device := range devices {
		deviceInfo[device.Name] = DeviceStatus{
			Pin: device.Pin,
			State: StateOff,
		}
	}

	err = initVisor()
	if err != nil {
		return err
	}

	return nil
}

func GetDeviceStatus(name string) (DeviceStatus, bool) {
	status, ok := deviceInfo[name]
	return status, ok
}