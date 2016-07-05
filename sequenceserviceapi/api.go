package sequenceserviceapi

import (
	"baseservices/sequenceservice"
	"fmt"
	"strconv"

	"github.com/coffeehc/logger"
	"github.com/coffeehc/microserviceboot/base"
	"github.com/coffeehc/microserviceboot/serviceclient"
)

type SequenceServiceApi interface {
	NextId() (int64, *base.Error)
}

const (
	POST_SEQUENCE = "POST_SEQUENCE"
)

func NewSqeuenceServiceApi(discoveryConfig *serviceclient.ServiceClientConsulConfig) (SequenceServiceApi, error) {
	serviceClient, err := serviceclient.NewServiceClient(&sequenceservice.SequenceServiceInfo{}, discoveryConfig)
	if err != nil {
		return nil, err
	}
	serviceClient.ApiRegister(POST_SEQUENCE, sequenceservice.POST_SEQUENCE)
	sequenceServiceApi := &_SequenceServiceApi{
		serviceClient: serviceClient,
	}
	logger.Info("创建 sequenceServiceApi")
	return sequenceServiceApi, nil
}

type _SequenceServiceApi struct {
	serviceClient *serviceclient.ServiceClient
}

func (this *_SequenceServiceApi) NextId() (int64, *base.Error) {
	response, err := this.serviceClient.SyncCallApi(POST_SEQUENCE, nil, nil, nil)
	if err != nil {
		return 0, base.NewSimpleError(-1, fmt.Sprintf("%s", err.Error()))
	}
	id, err := strconv.ParseInt(fmt.Sprintf("%s", response.Body()), 10, 64)
	if err != nil {
		return 0, base.NewSimpleError(-1, fmt.Sprintf("%s", err.Error()))
	}
	return id, nil
}
