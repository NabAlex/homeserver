package main

import (
	"net/http"
	"log"
	"homeserver/device"
)

func sendEasyOk(w http.ResponseWriter) {
	w.WriteHeader(http.StatusOK)
}

func sendBadParams(w http.ResponseWriter) {
	w.WriteHeader(http.StatusBadRequest)
}

func simpleParam(r *http.Request, key string) (string, bool) {
	value, ok := r.Form[key]
	if !ok || len(value) != 1 {
		return "", false
	}

	return value[0], true
}

func EnableDevice(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	log.Println("parse:", r.Method, r.URL)
	devicename, ok := simpleParam(r, "devicename")
	if !ok {
		sendBadParams(w)
		return
	}

	stateParam, ok := simpleParam(r, "state")
	if !ok {
		sendBadParams(w)
		return
	}

	var state device.DeviceState
	if state, ok = device.ConvertToState(stateParam); !ok {
		sendBadParams(w)
		return
	}

	status, ok := device.GetDeviceStatus(devicename)
	if  !ok {
		log.Println("cannot find", devicename)
		sendBadParams(w)
		return
	}

	device.SetPin(status.Pin, state)
	sendEasyOk(w)
}

func main() {
	err := device.InitDevice("conf/device.json")
	if err != nil {
		log.Println("Cannot init deviceinfo: ", err)
		return
	}

	http.HandleFunc("/enable", EnableDevice)

	err = http.ListenAndServe(":9090", nil)
	if err != nil {
		log.Fatal("Fail: ", err)
		return
	}

	log.Println("Start server", ":9090")
}