package main

import (
	"fmt"
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

func createInitConfig(configDic string) {
	file, _ := os.OpenFile(configDic, os.O_CREATE|os.O_WRONLY, 0666)
	defer file.Close()
	_, err := file.WriteString("[Key]\ndown: 219\nup: 221\nset: 32\n")
	if err != nil {
		panic(err)
	}
}

func writeConfig(configDic string, down int, up int, set int) {
	file, _ := os.OpenFile(configDic, os.O_WRONLY|os.O_TRUNC, 0)
	defer file.Close()
	_, err := file.WriteString(fmt.Sprintf("[Key]\ndown: %d\nup: %d\nset: %d\n", down, up, set))
	if err != nil {
		panic(err)
	}
}

func checkConfig() (up int, down int, set int, configDic string) {
	home, _ := home()
	configDic = home + "\\.hotkeygammaconfig"
	if !isExist(configDic) {
		log.Println("Config not found. Will create init config file. Path:", configDic)
		createInitConfig(configDic)
	}
	down, up, set = readConfig(configDic)
	log.Println("Read keys: down:", down, "up:", up, "set:", set)
	return down, up, set, configDic
}

func readConfig(configDic string) (up int, down int, set int) {
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
	set, err = keyConfig.Int("Key.set")
	if err != nil {
		panic(err)
	}
	return down, up, set
}
