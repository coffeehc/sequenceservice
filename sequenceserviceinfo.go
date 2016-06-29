package sequenceservice

type SequenceServiceInfo struct {
}

//获取 Api 定义的内容
func (this *SequenceServiceInfo) GetApiDefine() string {
	return ""
}

//获取 Service 名称
func (this *SequenceServiceInfo) GetServiceName() string {
	return "sequence-service"
}

//获取服务版本号
func (this *SequenceServiceInfo) GetVersion() string {
	return "1.0.0"
}

//获取服务描述
func (this *SequenceServiceInfo) GetDescriptor() string {
	return "Sequence service"
}

//获取 Service tags
func (this *SequenceServiceInfo) GetServiceTags() []string {
	return nil
}
