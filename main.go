package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"

	//"github.com/go-vgo/robotgo"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

//	type Config struct {
//		Commands []struct {
//			Type string `yaml:"type"`
//		} `yaml:"commands"`
//	}
type Config struct {
	Commands []struct {
		Type     string
		X        int
		Y        int
		MinX     int
		MinY     int
		MaxX     string
		MaxY     string
		Button   string
		Duration int64
	}
}

//type Command struct {
//	Type     string
//	X        int
//	Y        int
//	MinX     int
//	MinY     int
//	MaxX     string
//	MaxY     string
//	Button   string
//	Duration int
//}

func main() {

	file, err := os.Open("script.yaml")
	if err != nil {
		panic("unable to open configuration file script.yaml")
	}
	x, err := io.ReadAll(file)
	if err != nil {
		panic("failed reading script.yaml")
	}

	var config Config
	yaml.Unmarshal(x, &config)
	for _, command := range config.Commands {
		fmt.Println(command)
		switch command.Type {
		case "move":
			robotgo.Move(command.X, command.Y)
		case "click":
			robotgo.Click(command.Button)
		case "sleep":
			fmt.Printf("Sleeping for %d milliseconds.", command.Duration)
			time.Sleep(time.Duration(command.Duration) * time.Millisecond)
		}
	}

	//if ok {
	//
	//	for _, command := range commands {
	//		if a, ok := command.(map[string]any); ok {
	//			fmt.Println(a["type"])
	//		}
	//
	//	}
	//}
}
