package encoding

import (
	"encoding/json"
	"fmt"
	"github.com/Yandex-Practicum/final-project-encoding-go/models"
	"gopkg.in/yaml.v3"
	"io"
	"os"
)

const (
	errorOpenFile   = "Произошла ошибка при открытии файла: "
	errorReadFile   = "Произошла ошибка при чтении файла: "
	errorWriteFile  = "Произошла ошибка при записи файла: "
	errorMarshal    = "Ошибка при сериализации: "
	errorUnmarshal  = "Ошибка при десериализации: "
	errorCreateFile = "Произошла ошибка при создании файла: "
	errorCloseFile  = "Произошла ошибка при закрытии файла: "
)

// JSONData тип для перекодирования из JSON в YAML
type JSONData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// YAMLData тип для перекодирования из YAML в JSON
type YAMLData struct {
	DockerCompose *models.DockerCompose
	FileInput     string
	FileOutput    string
}

// MyEncoder интерфейс для структур YAMLData и JSONData
type MyEncoder interface {
	Encoding() error
}

// Encoding перекодирует файл из JSON в YAML
func (j *JSONData) Encoding() error {
	var dockerCompose models.DockerCompose

	jsonFile, err := os.Open(j.FileInput)

	if err != nil {
		return fmt.Errorf(errorOpenFile, err)
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Printf(errorCloseFile, err.Error())
		}
	}(jsonFile)

	jsonData, err := io.ReadAll(jsonFile)

	if err != nil {
		return fmt.Errorf(errorReadFile, err)
	}

	if err = json.Unmarshal(jsonData, &dockerCompose); err != nil {
		return fmt.Errorf(errorUnmarshal, err)
	}

	fmt.Println(string(jsonData))
	fmt.Println()
	fmt.Println(dockerCompose)

	yamlData, err := yaml.Marshal(&dockerCompose)
	if err != nil {
		return fmt.Errorf(errorMarshal, err)
	}
	fmt.Println(string(yamlData))

	yamlFile, err := os.Create(j.FileOutput)
	if err != nil {
		return fmt.Errorf(errorCreateFile, err)
	}

	defer func(yamlFile *os.File) {
		err := yamlFile.Close()
		if err != nil {
			fmt.Printf(errorCloseFile, err.Error())
		}
	}(yamlFile)

	_, err = yamlFile.Write(yamlData)

	if err != nil {
		return fmt.Errorf(errorWriteFile, err)
	}

	return nil
}

// Encoding перекодирует файл из YAML в JSON
func (y *YAMLData) Encoding() error {
	var dockerCompose models.DockerCompose

	yamlFile, err := os.Open(y.FileInput)

	if err != nil {
		return fmt.Errorf(errorOpenFile, err)
	}

	defer func(yamlFile *os.File) {
		err := yamlFile.Close()
		if err != nil {
			fmt.Printf(errorCloseFile, err.Error())
		}
	}(yamlFile)

	yamlData, err := io.ReadAll(yamlFile)

	if err != nil {
		return fmt.Errorf(errorReadFile, err)
	}

	if err = yaml.Unmarshal(yamlData, &dockerCompose); err != nil {
		return fmt.Errorf(errorUnmarshal, err)
	}

	jsonData, err := json.Marshal(&dockerCompose)
	if err != nil {
		return fmt.Errorf(errorMarshal, err)
	}

	jsonFile, err := os.Create(y.FileOutput)
	if err != nil {
		return fmt.Errorf(errorCreateFile, err)
	}

	defer func(jsonFile *os.File) {
		err := jsonFile.Close()
		if err != nil {
			fmt.Printf(errorCloseFile, err.Error())
		}
	}(jsonFile)

	_, err = jsonFile.Write(jsonData)

	if err != nil {
		return fmt.Errorf(errorWriteFile, err)
	}

	return nil
}
