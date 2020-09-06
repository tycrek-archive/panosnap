package main

import (
	//"fmt"
	//"time"
	"strings"
	//"strconv"

	"github.com/go-vgo/robotgo"
)

var wait = 50

func keyTap(key string) {
	robotgo.KeyTap(key)
	robotgo.MilliSleep(wait)
}

func shiftTap(key string) {
	robotgo.KeyTap(key, "shift")
	robotgo.MilliSleep(wait)
}

func sendCommand(yaw string, pitch string) {
	// Open chat
	keyTap("t")

	// /tp
	keyTap("/")
	keyTap("t")
	keyTap("p")
	keyTap("space")

	// @p

	shiftTap("2")
	keyTap("p")
	keyTap("space")

	// x y z
	shiftTap("~")
	keyTap("space")
	shiftTap("~")
	keyTap("space")
	shiftTap("~")
	keyTap("space")

	// yaw
	shiftTap("~")
	for i := 0; i < len(strings.Split(yaw, "")); i++ { 
		keyTap(strings.Split(yaw, "")[i]) 
	}

	keyTap("space")

	// pitch
	for i := 0; i < len(strings.Split(pitch, "")); i++ { 
		keyTap(strings.Split(pitch, "")[i]) 
	}
	
	// send command
	keyTap("enter")

	// take screenshot
	robotgo.MilliSleep(wait)
	keyTap("f2")
	robotgo.MilliSleep(wait)
}

/* Main */
func main() {
	robotgo.MilliSleep(3000)

	sendCommand("0", "0")
	sendCommand("90", "0")
	sendCommand("90", "0")
	sendCommand("90", "0")
	sendCommand("90", "-90")
	sendCommand("0", "90")
}