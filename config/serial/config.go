package serial

import (
	"fmt"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

const (
	configFolder = ".tty"
	configFile   = "config.yml"
)

type Config struct {
	// Serial Port Configuration
	BaudRate int    `yaml:"baudrate"`
	DataBits int    `yaml:"databits"`
	Parity   string `yaml:"parity"`
	StopBits string `yaml:"stopbits"`
	RTS      bool   `yaml:"rts"`
	DTR      bool   `yaml:"dtr"`

	// Terminal Configuration
	Format string `yaml:"format"`
}

func New() *Config {
	config := &Config{}
	config.setDefault()
	return config
}

func GetConfig(name string) (*Config, error) {
	file, err := readFile()
	if err != nil {
		return nil, err
	}

	if file != nil {
		return parseFile(file, name)
	}

	return New(), nil
}

func (config *Config) setDefault() {
	*config = Config{
		BaudRate: 9600,
		Parity:   "none",
		DataBits: 8,
		StopBits: "1",
		Format:   "ascii",
	}
}

func (config *Config) UnmarshalYAML(unmarshal func(interface{}) error) error {
	config.setDefault()

	type plain Config
	if err := unmarshal((*plain)(config)); err != nil {
		return err
	}

	return nil
}

func readFile() ([]byte, error) {
	// Get the home directory
	base, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	// Search for the config file
	filepath := filepath.Join(base, configFolder, configFile)
	_, err = os.Stat(filepath)
	if err != nil {
		fmt.Println("No config file found at location ", filepath)
		return nil, nil
	}

	// Read the YAML file
	file, err := os.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	return file, nil
}

func parseFile(file []byte, name string) (*Config, error) {
	// Unmarshal the YAML data into the map
	configs := make(map[string]Config)
	err := yaml.Unmarshal(file, &configs)
	if err != nil {
		return nil, err
	}

	// Check if the file contains information about the name
	if config, ok := configs[name]; ok {
		return &config, nil
	} else {
		return New(), nil
	}
}
