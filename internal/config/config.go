package config

import (
	"os"
	"encoding/json"
	"path/filepath"
)

const configFileName = ".gatorconfig.json"
const configPathFromHome = "bootdev/Gator/"

type Config struct {
	DB_Url			 string `json:"db_url"`
	CurrentUserName  string `json:"current_user_name"`
}
//REMEMBER !!!
//WHEN DOING STRUCTS THE START MUST ALWAYS BE IN CAPITAL LETTERS
//This is for packages

func (cfg *Config) SetUser (userName string) error {
	cfg.CurrentUserName = userName
	err := Write(cfg)
	if err != nil {
        return err
    }
	return nil
}

func Read() (Config, error) {	
	fullPath, err := GetConfigFilePath()
	if err != nil {
		return Config{}, err
	}

	file, err := os.Open(fullPath)
	if err != nil {
		return Config{}, err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	config := Config{}
	err = decoder.Decode(&config)
	if err != nil {
		return Config{}, err
	}

	return config, nil
}

func GetConfigFilePath() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}

	fullFilePath := filepath.Join(home, configFileName)
	return fullFilePath, nil
}

func Write(cfg *Config) error {
	fullPath, err := GetConfigFilePath()
	if err != nil {
		return err
	}

	file, err := os.Create(fullPath)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(cfg)
	if err != nil {
		return err
	}

	return nil
} 