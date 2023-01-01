package configurations

import (
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/xuliangTang/athena/athena/plugins"
	"log"
)

// LocalizeCfg @Configuration
type LocalizeCfg struct{}

func NewLocalizeCfg() *LocalizeCfg {
	return &LocalizeCfg{}
}

func (*LocalizeCfg) InitDefaultLocalize() *i18n.Localizer {
	localize, err := plugins.GetDefaultLocalize()
	if err != nil {
		log.Fatalln(err)
	}

	return localize
}
