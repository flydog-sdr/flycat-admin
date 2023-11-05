package config

import (
	"encoding/json"
	"os"
)

const (
	DEFAULT_GAIN   int = 20
	DEFAULT_FILTER int = 1
)

func (pref *Preference) Init(path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&Preference{
		Gain:   DEFAULT_GAIN,
		Filter: DEFAULT_FILTER,
	})
	if err != nil {
		return err
	}

	return nil
}

func (pref *Preference) Update(path string) error {
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_TRUNC, 0666)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(&pref)
	if err != nil {
		return err
	}

	return nil
}

func (pref *Preference) Read(path string) error {
	file, err := os.Open(path)
	if err != nil {
		return err
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(&pref)
	if err != nil {
		return err
	}

	return nil
}
