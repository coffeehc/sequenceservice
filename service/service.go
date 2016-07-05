package main

import (
	"baseservices/sequenceservice"
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

func (this *SequenceService) Init(configPath string, server *web.Server) error {
	serviceConfig := new(Config)
	err := base.LoadConfig(base.GetDefaultConfigPath(configPath), serviceConfig)
	if err != nil {
		return err
	}
	this.config = serviceConfig
	this.serviceDiscoveryRegister, err = consultool.NewConsulServiceRegister(consultool.WarpConsulConfig(serviceConfig.ConsulConfig))
	if err != nil {
		return err
	}
	this.sequenceService = cfsequence.NewSequenceService(*nodeId)
	return nil
}

func (this *SequenceService) Run() error {
	return nil

}
func (this *SequenceService) Stop() error {
	return nil
}
func (this *SequenceService) GetServiceInfo() base.ServiceInfo {
	return &sequenceservice.SequenceServiceInfo{}
}
func (this *SequenceService) GetEndPoints() []base.EndPoint {
	return []base.EndPoint{
		base.EndPoint{
			Metadata:    sequenceservice.POST_SEQUENCE,
			HandlerFunc: this.post_sequence,
		},
		base.EndPoint{
			Metadata:    sequenceservice.GET_SEQUENCE,
			HandlerFunc: this.get_sequence,
		},
	}
}

func (this *SequenceService) GetServiceDiscoveryRegister() base.ServiceDiscoveryRegister {
	return this.serviceDiscoveryRegister
}
