package main

import (
	"github.com/coffeehc/sequenceservice"
	"flag"

	"github.com/coffeehc/cfsequence"
	"github.com/coffeehc/microserviceboot/base"
	"github.com/coffeehc/microserviceboot/consultool"
	"github.com/coffeehc/microserviceboot/serviceboot"
	"github.com/coffeehc/web"
)

var nodeId = flag.Int64("nodeId", 0, "sequence 发生器的节点编号")

func main() {
	serviceboot.ServiceLauncher(&SequenceService{})
}

type SequenceService struct {
	config                   *Config
	serviceDiscoveryRegister base.ServiceDiscoveryRegister
	sequenceService          cfsequence.SequenceService
}

func (service *SequenceService) Init(configPath string, server *web.Server) base.Error {
	serviceConfig := new(Config)
	err := base.LoadConfig(base.GetDefaultConfigPath(configPath), serviceConfig)
	if err != nil {
		return err
	}
	service.config = serviceConfig
	service.serviceDiscoveryRegister, err = consultool.NewConsulServiceRegister(consultool.WarpConsulConfig(serviceConfig.ConsulConfig))
	if err != nil {
		return err
	}
	if *nodeId == 0{
		return base.NewError(base.ERROR_CODE_BASE_CONFIG_ERROR,"nodeId 不能为0,请设置非零值,如:-nodeid=1")
	}
	service.sequenceService = cfsequence.NewSequenceService(*nodeId)
	return nil
}

func (*SequenceService) Run() base.Error {
	return nil

}
func (*SequenceService) Stop() base.Error {
	return nil
}
func (*SequenceService) GetServiceInfo() base.ServiceInfo {
	return &sequenceservice.SequenceServiceInfo{}
}
func (service *SequenceService) GetEndPoints() []base.EndPoint {
	return []base.EndPoint{
		base.EndPoint{
			Metadata:    sequenceservice.POST_SEQUENCE,
			HandlerFunc: service.post_sequence,
		},
		base.EndPoint{
			Metadata:    sequenceservice.GET_SEQUENCE,
			HandlerFunc: service.get_sequence,
		},
		base.EndPoint{
			Metadata:    sequenceservice.GET_MINSEQUENCE,
			HandlerFunc: service.get_minSequence,
		},
	}
}

func (service *SequenceService) GetServiceDiscoveryRegister() base.ServiceDiscoveryRegister {
	return service.serviceDiscoveryRegister
}
