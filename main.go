package main

import (
	"fmt"
	"os"
)

type Config struct {
	UserConfig
	Title string
}

func main() {
	var c Config
	c.Width = 640
	c.Height = 480
	c.Title = "Test"

	f, err := os.Open("config.toml")
	if err != nil {
		if !os.IsNotExist(err) {
			fmt.Fprintf(os.Stderr, "Error opening configuration file, using defaults. Err: %s", err)
		}
	} else {
		if err := c.Load(f); err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing configuration file, using defaults. Err: %s", err)
		}
		f.Close()
	}

	if err := loop(c); err != nil {
		fmt.Fprintf(os.Stderr, "Error occurred: %s", err)
		return
	}
}
