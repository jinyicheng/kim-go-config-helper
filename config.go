package configHelper

import (
	"os"
	"path"
	"path/filepath"
)

type Config struct {
	Env  string
	Ext  string
	Path string
}

func (c *Config) Get(decodeFileCallback func(env string, filePath string)) error {
	var err error
	var files []os.DirEntry
	//获取配置文件
	if files, err = os.ReadDir(c.Path); err != nil {
		return err
	} else {
		for _, file := range files {
			filePath := c.Path + "/" + file.Name()
			fileName := path.Base(filePath)
			if path.Ext(fileName) == c.Ext {
				if filePath, err = filepath.Abs(filePath); err != nil {
					return err
				} else {
					decodeFileCallback(c.Env, filePath)
				}
			}
		}
	}
	return nil
}
