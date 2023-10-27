package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"math/rand"
	"time"

	//"github.com/go-vgo/robotgo"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

type Config struct {
	Commands []struct {
		Type     *string
		X        *int
		Y        *int
		MinX     *int `yaml:"minX"`
		MinY     *int `yaml:"minY"`
		MaxX     *int `yaml:"maxX"`
		MaxY     *int `yaml:"maxY"`
		Button   *string
		Duration *int64
	}
}

func main() {
	file, err := os.Open("script.yaml")
	if err != nil {
		panic("unable to open configuration file script.yaml")
	}
	rawConfig, err := io.ReadAll(file)
	if err != nil {
		panic(fmt.Errorf("failed reading script.yaml: %w", err))
	}

	var config Config
	err = yaml.Unmarshal(rawConfig, &config)
	if err != nil {
		panic(fmt.Errorf("failed to parse script.yaml: %w", err))
	}
	for {
		for _, command := range config.Commands {
			switch *command.Type {
			case "move":
				if command.X != nil && command.Y != nil {
					robotgo.MoveSmooth(*command.X, *command.Y)
				} else if command.MinX != nil && command.MinY != nil && command.MaxX != nil && command.MaxY != nil {
					x := rand.Intn(*command.MaxX-*command.MinX) + *command.MinX
					y := rand.Intn(*command.MaxY-*command.MinY) + *command.MinY
					robotgo.MoveSmooth(x, y)
				} else {
					panic(fmt.Sprintf("Invalid move configuration: %+v", command))
				}
			case "click":
				robotgo.Click(command.Button)
			case "sleep":
				fmt.Printf("Sleeping for %d milliseconds.", command.Duration)
				time.Sleep(time.Duration(*command.Duration) * time.Millisecond)
			}
		}
	}
}
