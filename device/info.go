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
	Pin uint			`json:"pin"`
	State DeviceState	`json:"state"`
}

var deviceInfo map[string]*DeviceStatus

func InitDevice(filename string) error {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return  err
	}

	devices, err := parseDevices(string(bytes))
	if err != nil {
		return err
	}

	err = initVisor()
	if err != nil {
		return err
	}

	deviceInfo = make(map[string]*DeviceStatus)
	for _, device := range devices {
		setOutput(device.Pin)
		setState(device.Pin, StateOff)
		if device.Name != "disable" {
			deviceInfo[device.Name] = &DeviceStatus{
				Pin:   device.Pin,
				State: StateOff,
			}
		}
	}

	return nil
}

func GetDeviceStatus(name string) (DeviceStatus, bool) {
	status, ok := deviceInfo[name]
	return *status, ok
}

func SetDeviceStatus(name string, state DeviceState) {
	if _, ok := deviceInfo[name]; !ok {
		return
	}

	setState(deviceInfo[name].Pin, state)
	deviceInfo[name].State = state
}

func GetDevices() map[string]*DeviceStatus {
	return deviceInfo
}