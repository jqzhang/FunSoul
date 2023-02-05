package config

import (
	"github.com/jinzhu/configor"
	"sync"
)

var once sync.Once

var Config Content

func init() {
	configor.Load(&Config, "config.yml")
}
