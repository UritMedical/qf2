package define

type BaseBll struct {
	plugin BasePlugin
}

func (b *BaseBll) SendNotice(topic string, params ...interface{}) error {
	return b.plugin.SendNotice(topic, params...)
}
