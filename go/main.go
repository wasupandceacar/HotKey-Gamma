package main

import (
	"fmt"
	"github.com/MakeNowJust/hotkey"
	"log"
	"math"
	"os"
	"os/signal"
	"syscall"
)

var hkey = hotkey.New()

var hkeyset = hotkey.New()

var downkey = 0
var upkey = 0
var setkey = 0
var configDic = ""

var iddown hotkey.Id = -1
var idup hotkey.Id = -1

var g *Gamma
var gamma float64

func registerFunction() {
	iddown, _ = hkey.Register(hotkey.None, uint32(downkey), func() {
		gamma = math.Min(1, gamma + 0.05)
		log.Println("gamma down, current gamma:", fmt.Sprintf("%.2f", -gamma))
		adjustGamma(g, gamma)
	})

	idup, _ = hkey.Register(hotkey.None, uint32(upkey), func() {
		gamma = math.Max(-1, gamma - 0.05)
		log.Println("gamma up, current gamma:", fmt.Sprintf("%.2f", -gamma))
		adjustGamma(g, gamma)
	})
}

func registerSet() {
	_, _ = hkeyset.Register(hotkey.None, uint32(setkey), func() {
		log.Println("Unregister functions. down:", downkey, "up:", upkey)
		unregisterFunction()

		log.Println("Press your gamma down key:")
		downkey = readSingleKey()

		log.Println("Press your gamma up key:")
		upkey = readSingleKey()

		writeConfig(configDic, downkey, upkey, setkey)

		log.Println("Register functions. down:", downkey, "up:", upkey)
		registerFunction()
	})
}

func unregisterFunction() {
	hkey.Unregister(iddown)
	hkey.Unregister(idup)
}

func readSingleKey() int {
	ids := make([]hotkey.Id, 222)
	press := 0
	for k := 1; k <= 222; k++ {
		key := k
		id, _ := hkey.Register(hotkey.None, uint32(k), func() {
			press = key
		})
		ids[k-1] = id
	}
	for press == 0 {

	}
	for _, unid := range ids {
		hkey.Unregister(unid)
	}
	log.Println("detect:", press)
	return press
}

func main(){
	downkey, upkey, setkey, configDic = checkConfig()

	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	fmt.Println("Made by wasupandceacar")
	fmt.Println("=======================================================")
	fmt.Println("Guide:")
	fmt.Println("Press [(init) to decrease gamma (min value: -1)")
	fmt.Println("Press ](init) to increase gamma (max value: 1)")
	fmt.Println("Press \\(init) to reset gamma down and up key")
	fmt.Println("Press Ctrl+C to exit, and gamma will return to standard (value: 0)")
	fmt.Println("Don't click close button to exit, gamma will not return back")
	fmt.Println("If you accidentally do so, just restart and it will set gamma to standard initially")
	fmt.Println("=======================================================")

	getGammaFunc()
	getHDC()
	g, gamma = initGamma()

	registerSet()
	registerFunction()

	<-c
	adjustGamma(g, -0.0)
}
