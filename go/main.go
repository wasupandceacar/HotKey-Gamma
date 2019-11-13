package main

import (
	"fmt"
	"log"
	"github.com/MakeNowJust/hotkey"
	"math"
	"os"
	"os/signal"
	"syscall"
)

var hkey = hotkey.New()

func main(){
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	fmt.Println("Made by wasupandceacar")
	fmt.Println("=======================================================")
	fmt.Println("Guide:")
	fmt.Println("Press ] to increase gamma (max value: 1)")
	fmt.Println("Press [ to decrease gamma (min value: -1)")
	fmt.Println("Press Ctrl+C to exit, and gamma will return to standard (value: 0)")
	fmt.Println("Don't click close button to exit, gamma will not return back")
	fmt.Println("If you accidentally do so, just restart and it will set gamma to standard initially")
	fmt.Println("=======================================================")

	getGammaFunc()
	getHDC()
	g, gamma := initGamma()

	hkey.Register(hotkey.None, 0xDD, func() {
		gamma = math.Max(-1, gamma - 0.05)
		log.Println("gamma up, current gamma:", fmt.Sprintf("%.2f", -gamma))
		adjustGamma(g, gamma)
	})
	hkey.Register(hotkey.None, 0xDB, func() {
		gamma = math.Min(1, gamma + 0.05)
		log.Println("gamma down, current gamma:", fmt.Sprintf("%.2f", -gamma))
		adjustGamma(g, gamma)
	})

	<-c
	adjustGamma(g, -0.0)
}
