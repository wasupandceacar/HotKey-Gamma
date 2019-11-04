package main

import (
	"math"
	"unsafe"
)

var HDC uintptr

type Gamma struct {
	Red		[256]uint16
	Green	[256]uint16
	Blue	[256]uint16
}

func getHDC() {
	HDC, _, _ = getDC.Call(0)
}

func getGamma() *Gamma {
	g := &Gamma{}
	success, _, _ := getDeviceGammaRamp.Call(HDC, uintptr(unsafe.Pointer(g)))
	if success == 0 {
		panic("读取 gamma 值失败")
	}
	return g
}

func initGamma() (*Gamma, float64) {
	g := getGamma()
	// 初始gamma设为0
	gamma := 0.0
	adjustGamma(g, gamma)
	return g, gamma
}

func adjustGamma(g *Gamma, gamma float64) {
	for i := 0; i < 256; i++ {
		val := uint16(math.Min(65535, math.Max(0, math.Pow(float64(i) / 256.0, math.Pow(4, gamma)) * 65535 + 0.5)))
		g.Red[i] = val
		g.Blue[i] = val
		g.Green[i] = val
	}
	setDeviceGammaRamp.Call(HDC, uintptr(unsafe.Pointer(g)))
}
