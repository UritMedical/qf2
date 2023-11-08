/**
 * @Author: Joey
 * @Description:
 * @Create Date: 2023/11/8 9:09
 */

package qdefine

type QBaseModule struct {
	Stop        func()
	StartInvoke func(route string, ctx QContext) *QFail
	EndInvoke   func(route string, ctx QContext)
}

func (m *QBaseModule) OnStartInvoke(route string, ctx QContext) *QFail {
	if m.StartInvoke != nil {
		return m.StartInvoke(route, ctx)
	}
	return nil
}
func (m *QBaseModule) OnStop() {
	if m.Stop != nil {
		m.Stop()
	}

}
func (m *QBaseModule) OnEndInvoke(route string, ctx QContext) {
	if m.EndInvoke != nil {
		m.EndInvoke(route, ctx)
	}
}
