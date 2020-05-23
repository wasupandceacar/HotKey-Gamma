package main

import (
	"log"
	"os"
)

func isExist(fileAddr string)(bool){
	_, err := os.Stat(fileAddr)
	if err!=nil{
		if os.IsExist(err){
			return true
		}
		return false
	}
	return true
}

func writeInitConfig(configDic string) {
	file, _ := os.OpenFile(configDic, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	defer file.Close()
	_, err := file.WriteString("[Key]\ndown: 219\nup: 221\n")
	if err != nil {
		panic(err)
	}
}

func checkConfig() (up int, down int) {
	home, _ := home()
	configDic := home + "\\.hotkeygammaconfig"
	if !isExist(configDic) {
		log.Println("Config not found. Will create init config file. Path:", configDic)
		writeInitConfig(configDic)
	}
	down, up = readConfig(configDic)
	log.Println("Read keys: down:", down, "up:", up)
	return down, up
}

func readConfig(configDic string) (up int, down int) {
	keyConfig, err := NewFileConf(configDic)
	if err != nil {
		panic(err)
	}
	down, err = keyConfig.Int("Key.down")
	if err != nil {
		panic(err)
	}
	up, err = keyConfig.Int("Key.up")
	if err != nil {
		panic(err)
	}
	return down, up
}
