package config

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"gopkg.in/yaml.v3"
)

type SerialConfig struct {
	BaudRate int     `yaml:"baudrate"`
	Parity   string  `yaml:"parity"`
	DataBits int     `yaml:"databits"`
	StopBits float32 `yaml:"stopbits"`
	Format   string  `yaml:"format"`
}

func GetConfig(portname string, args []string) (SerialConfig, error) {

	// Check in configuration file
	config, err := getConfigFromFile(portname)
	if err != nil {
		return config, err
	}

	// Apply the arguments modifications
	err = applyArguments(args, &config)
	if err != nil {
		return config, err
	}

	return config, nil
}

func (s *SerialConfig) UnmarshalYAML(unmarshal func(interface{}) error) error {
	s.BaudRate = 9600
	s.Parity = "none"
	s.DataBits = 8
	s.StopBits = 1
	s.Format = "ascii"

	type plain SerialConfig
	if err := unmarshal((*plain)(s)); err != nil {
		return err
	}

	return nil
}

func getConfigFromFile(portname string) (SerialConfig, error) {
	// Get the home directory
	base, err := os.UserHomeDir()
	if err != nil {
		return SerialConfig{}, err
	}

	// Search for the config file
	filepath := filepath.Join(base, ".tty", "config.yml")
	_, err = os.Stat(filepath)
	if err != nil {
		fmt.Println("No config file found at location ", filepath)
	}

	// Read the YAML file
	file, err := os.ReadFile(filepath)
	if err != nil {
		return SerialConfig{}, err
	}

	// Unmarshal the YAML data into the map
	configs := make(map[string]SerialConfig)
	err = yaml.Unmarshal(file, &configs)
	if err != nil {
		return SerialConfig{}, err
	}

	if config, ok := configs[portname]; ok {
		return config, nil
	} else {
		fmt.Println("No configuration found for this port in the config file")
	}

	return SerialConfig{}, nil
}

func applyArguments(args []string, config *SerialConfig) error {
	for i, arg := range args {
		if len(args)-i < 2 {
			break
		}
		switch arg {

		case "--baudrate":
			baudrate, err := strconv.Atoi(args[i+1])
			if err != nil {
				return fmt.Errorf("the --baudrate value \"%v\" is invalid", args[i+1])
			}
			config.BaudRate = baudrate

		case "--databits":
			databits, err := strconv.Atoi(args[i+1])
			if err != nil {
				return fmt.Errorf("the --databits value \"%v\" is invalid", args[i+1])
			}
			config.DataBits = databits

		case "--parity":
			switch args[i+1] {
			case "none", "even", "odd", "mark", "space":
				config.Parity = args[i+1]
			default:
				return fmt.Errorf("the --parity value \"%v\" is invalid", args[i+1])
			}

		case "--stopbits":
			stopbits, err := strconv.ParseFloat(args[i+1], 32)
			if err != nil {
				return fmt.Errorf("the --parity value \"%v\" is invalid", args[i+1])
			}
			switch stopbits {
			case 1, 1.5, 2:
				config.StopBits = float32(stopbits)
			default:
				return fmt.Errorf("the --parity value \"%v\" is invalid", args[i+1])
			}

		case "--format":
			switch args[i+1] {
			case "ascii":
			case "hex":
				config.Format = args[i+1]
			default:
				return fmt.Errorf("the --format value \"%v\" is invalid", args[i+1])
			}
		}
	}
	return nil
}
