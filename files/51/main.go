package main

import (
	"errors"
	"fmt"
	"log"
	"os"

	"gopkg.in/yaml.v3"
)

// Config хранит настройки программы
type Config struct {
	Host   string `yaml:"host"`
	Port   int    `yaml:"port"`
	DBName string `yaml:"dbname"`
	DBUser string `yaml:"dbuser"`
}

// LoadConfig загружает конфигурацию из YAML файла в структуру Config
func LoadConfig(cfgFile string, cfg *Config) error {
	data, err := os.ReadFile(cfgFile)
	if err != nil {
		// если файла нет, используем значения по умолчанию
		if errors.Is(err, os.ErrNotExist) {
			return nil
		}
		return err
	}
	return yaml.Unmarshal(data, cfg)
}

// SaveConfig сохраняет структуру Config в YAML файл
func SaveConfig(cfgFile string, cfg Config) error {
	// конвертируем структуру в YAML
	data, err := yaml.Marshal(cfg)
	if err != nil {
		return err
	}
	// записываем данные в файл
	return os.WriteFile(cfgFile, data, 0755)
}

// переменная конфигурации со значениями по умолчанию
var cfg = Config{
	Host:   "localhost",
	Port:   8080,
	DBName: "app",
	DBUser: "admin",
}

func main() {
	// путь к файлу конфигурации
	cfgFile := "settings.yaml"

	// Загружаем конфигурацию из файла
	err := LoadConfig(cfgFile, &cfg)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%#v\n", cfg)
	// изменяем настройки
	cfg.Host = "127.0.0.1"
	cfg.Port = 6001
	// сохраняем обновленную конфигурацию в файл
	err = SaveConfig(cfgFile, cfg)
	if err != nil {
		log.Fatal(err)
	}
}
