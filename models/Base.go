package models

import (
	"ginReact/conf"
	"github.com/xormplus/xorm"
)

type Base struct {
}

func (b *Base) GetDb() *xorm.Engine {
	return conf.Engine
}