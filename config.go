package microfiber

import (
	"errors"
	"os"

	"github.com/spf13/viper"
)

type ReadConfigOptions struct {
	Filename string
	FileType string
}

type Config struct {
	NAME string
}

func ReadConfig(out interface{}, option *ReadConfigOptions) error {
	file, err := os.Open(option.Filename)
	if os.IsNotExist(err) {
		return errors.New("file is not exist")
	} else if err != nil {
		return err
	}
	v := viper.New()
	v.SetConfigType(option.FileType)
	err = v.ReadConfig(file)
	if err != nil {
		return err
	}
	err = v.Unmarshal(out)
	if err != nil {
		return err
	}
	return nil
}
