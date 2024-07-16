package qweb

import (
	"github.com/UritMedical/qf2/qdefine"
)

type Module struct {
	QBll qdefine.QBll
	QDao qdefine.QDao
}

type setting struct {
	Port        int
	StartDelay  int
	DefGroup    string
	StaticDir   string
	WebDir      string
	HistoryMode int
}

type webSetting struct {
}

func defaultSetting() setting {
	return setting{
		Port:        10001,
		StartDelay:  5,
		DefGroup:    "api",
		StaticDir:   "./assets",
		WebDir:      "./dist",
		HistoryMode: 1,
	}
}
