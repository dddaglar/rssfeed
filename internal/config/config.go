package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DBUrl           string `json:"db_url"`
	CurrentUserName string `json:"current_user_name"`
}

func getConfigFilePath() (string, error) {
	homePath, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	return homePath + "/" + configFileName, nil
}

func Read() (Config, error) {
	cfgPath, err := getConfigFilePath()
	if err != nil {
		return Config{}, err
	}
	jsonFile, err := os.ReadFile(cfgPath)
	if err != nil {
		return Config{}, err
	}
	var newConfig Config
	err = json.Unmarshal(jsonFile, &newConfig)
	if err != nil {
		return Config{}, err
	}
	return newConfig, nil
}

func write(cfg Config) error {
	jsonCfg, err := json.Marshal(cfg)
	if err != nil {
		return err
	}
	cfgPath, err := getConfigFilePath()
	if err != nil {
		return err
	}
	err = os.WriteFile(cfgPath, jsonCfg, 0666)
	if err != nil {
		return err
	}
	return nil
}

func (cfg *Config) SetUser(user string) error {
	cfg.CurrentUserName = user
	err := write(*cfg)
	return err
}
