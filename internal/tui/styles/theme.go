package styles

import (
	"os"

	"github.com/BurntSushi/toml"
)

type Theme struct {
	Colors struct {
		Primary    string `toml:"primary"`
		Background string `toml:"background"`
		Foreground string `toml:"foreground"`
		Border     string `toml:"border"`
		Accent     string `toml:"accent"`
	} `toml:"colors"`

	Styles struct {
		Padding     int    `toml:"padding"`
		Margin      int    `toml:"margin"`
		BorderStyle string `toml:"border_style"`
	} `toml:"styles"`
}

var Current Theme

func LoadTheme(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	_, err = toml.Decode(string(data), &Current)
	return err
}

func SaveTheme(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	return toml.NewEncoder(file).Encode(Current)
}