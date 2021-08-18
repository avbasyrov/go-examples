package main

import (
	"encoding/json"
	"gopkg.in/yaml.v3"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type Employee struct {
	Name string `json:"name" yaml:"name"`
	Data struct {
		Age int `json:"age" yaml:"age"`
	} `json:"info" yaml:"info"`
}

func main() {
	// Читаем полностью файл
	data, err := ioutil.ReadFile("data.json")
	if err != nil {
		log.Fatalln(err)
	}

	// Создаем новый файл и открываем его для записи
	output, err := os.Create("data.yaml")
	if err != nil {
		log.Fatalln(err)
	}
	defer output.Close()

	err = ConvertToYaml(data, output)
	if err != nil {
		log.Fatalln(err)
	}
}

func ConvertToYaml(jsonData []byte, to io.Writer) error {
	// Декодируем из JSON
	var employees []Employee
	err := json.Unmarshal(jsonData, &employees)
	if err != nil {
		return err
	}

	// Конвертируем структуру в YAML и записываем в файл
	yamlEncoder := yaml.NewEncoder(to)
	err = yamlEncoder.Encode(employees)
	if err != nil {
		return err
	}

	return nil
}
