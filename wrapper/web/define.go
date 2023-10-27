package web

import "github.com/UritMedical/qf2/define"

type WebParams struct {
	AppInitFunc []func(adapter define.QAppAdapter)
	Bll         define.QBll
}

func LoadParams(bll define.QBll, initFunc ...func(adapter define.QAppAdapter)) WebParams {
	setting := WebParams{
		AppInitFunc: initFunc,
		Bll:         bll,
	}
	return setting
}

type setting struct {
	Port     int
	DefGroup string
}

func loadSetting() *setting {
	setting := &setting{
		Port:     10001,
		DefGroup: "",
	}
	return setting
}
