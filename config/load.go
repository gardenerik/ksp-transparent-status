package config

import (
	"encoding/json"
	"github.com/nsf/termbox-go"
	"io/ioutil"
	"os"
	"zahradnik.xyz/mirror-stats/logger"
)

type Person struct {
	CandleName  string   `json:"candle_name"`
	DisplayName string   `json:"display_name"`
	RGB         [3]uint8 `json:"rgb"`
}

func (p Person) Color() termbox.Attribute {
	return termbox.RGBToAttribute(p.RGB[0], p.RGB[1], p.RGB[2])
}

var People []Person

func ReadConfig() error {
	logger.Log.Println("Reloading config")

	file, err := os.Open("config.json")
	if err != nil {
		return err
	}
	defer file.Close()

	bytes, err := ioutil.ReadAll(file)
	if err != nil {
		return err
	}

	err = json.Unmarshal(bytes, &People)
	return err
}
