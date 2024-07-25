package configHelper

import (
	tomlHandler "github.com/BurntSushi/toml"
	"os"
	"path/filepath"
)

type Toml struct {
	Env  string
	Path string
}

func (t *Toml) Get(config any) (any, error) {
	var err error
	var dirEntries []os.DirEntry
	//获取配置文件
	if dirEntries, err = os.ReadDir(t.Path); err != nil {
		return nil, err
	} else {
		for _, dirEntry := range dirEntries {
			filePath := filepath.Join(t.Path, dirEntry.Name())
			fileName := filepath.Base(filePath)
			if filepath.Ext(fileName) == ".toml" {
				if _, err = tomlHandler.DecodeFile(filePath, &config); err != nil {
					return nil, err
				}
			}
		}
	}
	return config, nil
}
