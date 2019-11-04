package main

import "syscall"

var user32 				*syscall.LazyDLL
var gdi32 				*syscall.LazyDLL
var getDC 				*syscall.LazyProc
var getDeviceGammaRamp 	*syscall.LazyProc
var setDeviceGammaRamp 	*syscall.LazyProc

func getGammaFunc()  {
	user32 				= 	syscall.NewLazyDLL("user32.dll")
	gdi32 				= 	syscall.NewLazyDLL("gdi32.dll")
	getDC 				= 	user32.NewProc("GetDC")
	getDeviceGammaRamp 	= 	gdi32.NewProc("GetDeviceGammaRamp")
	setDeviceGammaRamp 	= 	gdi32.NewProc("SetDeviceGammaRamp")
}
