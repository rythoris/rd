package config

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/adrg/xdg"
)

type Config struct {
	Token string `json:"api_token"`
}

func ParseConfig() (Config, error) {
	confFile, err := xdg.SearchConfigFile("rd/config.json")
	if err != nil {
		return Config{}, err
	}

	fileContent, err := os.ReadFile(confFile)
	if err != nil {
		return Config{}, fmt.Errorf("config read error: %w", err)
	}

	var c Config
	err = json.Unmarshal(fileContent, &c)
	if err != nil {
		return Config{}, fmt.Errorf("config parse error: %w", err)
	}

	return c, nil
}
