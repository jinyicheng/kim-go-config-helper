package configHelper

import (
	jsoniter "github.com/json-iterator/go"
	"io"
	"log"
	"os"
	"path/filepath"
)

type Json struct {
	Env  string
	Path string
}

var jsonHandler = jsoniter.ConfigCompatibleWithStandardLibrary

func (j *Json) Get(config any) (any, error) {
	var err error
	var dirEntries []os.DirEntry
	var file *os.File
	var content []byte
	//获取配置文件
	if dirEntries, err = os.ReadDir(j.Path); err != nil {
		return nil, err
	}

	for _, dirEntry := range dirEntries {
		filePath := filepath.Join(j.Path, dirEntry.Name())
		fileName := filepath.Base(filePath)
		if filepath.Ext(fileName) == ".json" {
			// 打开 JSON 文件
			if file, err = os.Open(filePath); err != nil {
				return nil, err
			}
			if err = readAndProcessFile(file, &content, &config); err != nil {
				return nil, err
			}
		}
	}
	return config, nil
}

func readAndProcessFile(file *os.File, content *[]byte, config any) error {
	var err error
	defer func(file *os.File) {
		err = file.Close()
		if err != nil {
			log.Println("Failed to close file:", err)
		}
	}(file)
	// 读取文件内容
	if *content, err = io.ReadAll(file); err != nil {
		return err
	}
	if err = jsonHandler.Unmarshal(*content, &config); err != nil {
		return err
	}
	return nil
}
