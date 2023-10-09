package base

func Dict_GetDict(id string) (interface{}, error) {
	rs, err := Plugin.Invoke("Dict_GetDict", id)
	return rs, err
}
