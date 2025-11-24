package utils

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Config struct {
	General struct {
		UseSpotify bool `toml:"use_spotify"`
	} `toml:"general"`

	Spotify struct {
		ClientID     string `toml:"client_id"`
		ClientSecret string `toml:"client_secret"`
	} `toml:"spotify"`
}

var Current Config

func LoadConfig(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	_, err = toml.Decode(string(data), &Current)
	return err
}

func SaveConfig(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return toml.NewEncoder(file).Encode(Current)
}