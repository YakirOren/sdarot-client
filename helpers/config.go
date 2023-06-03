package helpers

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/YakirOren/sdarot"
)

const ConfigFile = "sdarot_config.json"

func CreateConfig() error {
	qs := []*survey.Question{
		{
			Name: "username",
			Prompt: &survey.Input{
				Message: "Enter your username",
			},
		},
		{
			Name: "password",
			Prompt: &survey.Password{
				Message: "Enter your password",
			},
		},
	}

	conf := sdarot.Config{}

	if err := survey.Ask(qs, &conf); err != nil {
		return err
	}

	bytes, err := json.Marshal(conf)
	if err != nil {
		return err
	}

	file, err := os.Create(ConfigFile)
	if err != nil {
		return err
	}

	_, err = file.Write(bytes)
	if err != nil {
		return err
	}

	return nil
}

func LoadConfig() (*sdarot.Config, error) {
	if _, err := os.Stat(ConfigFile); errors.Is(err, os.ErrNotExist) {
		if err2 := CreateConfig(); err2 != nil {
			return nil, err2
		}
	}

	open, err := os.Open(ConfigFile)
	if err != nil {
		return nil, err
	}

	conf := &sdarot.Config{}

	if decodeErr := json.NewDecoder(open).Decode(&conf); decodeErr != nil {
		return nil, decodeErr
	}
	return conf, nil
}
