package dict

import (
	"github.com/UritMedical/qf2/define"
)

type baseBll struct {
	define.BaseBll
}

func (b *baseBll) GetConfig() string {
	return ""
}

func (b *baseBll) Apis() map[string]define.ApiHandler {
	return map[string]define.ApiHandler{}
}

func (b *baseBll) Subscribes() map[string]define.NoticeHandler {
	return map[string]define.NoticeHandler{}
}
