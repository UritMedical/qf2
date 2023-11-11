package qweb

import (
	"github.com/UritMedical/qf2/qdefine"
)

type Module struct {
	QBll qdefine.QBll
	QDao qdefine.QDao
}

type setting struct {
	Port       int
	DefGroup   string
	StartDelay int
}

func defaultSetting() setting {
	return setting{
		Port:       10001,
		DefGroup:   "api",
		StartDelay: 5,
	}
}
