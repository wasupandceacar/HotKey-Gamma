package main

import (
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

	getGammaFunc()
	getHDC()
	g, gamma := initGamma()

	hkey.Register(hotkey.None, 0xDD, func() {
		log.Println("gamma up")
		gamma = math.Max(-1, gamma - 0.05)
		adjustGamma(g, gamma)
	})
	hkey.Register(hotkey.None, 0xDB, func() {
		log.Println("gamma down")
		gamma = math.Min(1, gamma + 0.05)
		adjustGamma(g, gamma)
	})

	<-c
	adjustGamma(g, -0.0)
}
