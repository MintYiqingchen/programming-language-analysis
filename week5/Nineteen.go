package main

import (
	"encoding/json"
	"fmt"
	"os"
	"plugin"
)

type PairInterface interface {
	GetStr() string
	GetCount() int
}

func main() {
	configfile, err := os.Open("NineteenConfig.json")
	if err != nil {
		return
	}
	jsonDecoder := json.NewDecoder(configfile)
	type LoadConfig struct {
		WordsPluginName, FrequencyPluginName string
	}
	var loadconf LoadConfig
	err = jsonDecoder.Decode(&loadconf)
	if err != nil {
		fmt.Println(err)
		return
	}
	wordsPluginName := loadconf.WordsPluginName         // "./plugin/ExtractWords1.so"
	frequencyPluginName := loadconf.FrequencyPluginName // "./plugin/Frequency1.so"
	fmt.Println(wordsPluginName, frequencyPluginName)
	wordsPlugin, err := plugin.Open(wordsPluginName)
	if err != nil {
		panic(err)
	}
	wordsFuncSym, _ := wordsPlugin.Lookup("ExtractWords")
	frequencyPlugin, err := plugin.Open(frequencyPluginName)
	if err != nil {
		panic(err)
	}
	frequencyFuncSym, _ := frequencyPlugin.Lookup("Top25")
	wordlist := wordsFuncSym.(func(string) []string)(os.Args[1])
	top25 := frequencyFuncSym.(func([]string) interface{})(wordlist)
	for _, intf := range top25.([]interface{}) {
		fmt.Println(intf.(PairInterface).GetStr(), " - ", intf.(PairInterface).GetCount())
	}
}
