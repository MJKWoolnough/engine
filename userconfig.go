package main

import (
	"io"

	"github.com/BurntSushi/toml"
)

type UserConfig struct {
	Width  int    `toml:",omitempty"`
	Height int    `toml:",omitempty"`
	Title  string `toml:",omitempty"`
}

func (c *UserConfig) Load(r io.Reader) error {
	_, err := toml.DecodeReader(r, c)
	return err
}

func (c *UserConfig) Save(w io.Writer) error {
	return toml.NewEncoder(w).Encode(c)
}
