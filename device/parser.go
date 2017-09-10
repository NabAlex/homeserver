package device

import (
	"encoding/json"
)

type DevicesData struct {
	Devices []Device `json:"devices"`
}

func parseDevices(jsonString string) ([]Device, error) {
	var devices DevicesData
	err := json.Unmarshal([]byte(jsonString), &devices)
	if err != nil {
		return nil, err
	}

	return devices.Devices, nil
}