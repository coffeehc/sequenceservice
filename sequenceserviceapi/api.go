package sequenceserviceapi

import (
	"baseservices/sequenceservice"
	"fmt"
	"strconv"

	"github.com/coffeehc/logger"
	"github.com/coffeehc/microserviceboot/base"
	"github.com/coffeehc/microserviceboot/serviceclient"
	"time"
)

type SequenceServiceApi interface {
	GetSequenceCreateTime(sequence int64) (time.Time, *base.Error)
	NextId() (int64, *base.Error)
	MinId(milTime int64) (int64, *base.Error)
}

const (
	POST_SEQUENCE   = "POST_SEQUENCE"
	GET_SEQUENCE    = "GET_SEQUENCE"
	GET_MINSEQUENCE = "GET_MINSEQUENCE"
)

func NewSqeuenceServiceApi(discoveryConfig *serviceclient.ServiceClientConsulConfig) (SequenceServiceApi, error) {
	serviceClient, err := serviceclient.NewServiceClient(&sequenceservice.SequenceServiceInfo{}, discoveryConfig)
	if err != nil {
		return nil, err
	}
	serviceClient.ApiRegister(POST_SEQUENCE, sequenceservice.POST_SEQUENCE)
	serviceClient.ApiRegister(GET_SEQUENCE, sequenceservice.GET_SEQUENCE)
	serviceClient.ApiRegister(GET_MINSEQUENCE, sequenceservice.GET_MINSEQUENCE)
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

func (this *_SequenceServiceApi) GetSequenceCreateTime(sequence int64) (time.Time, *base.Error) {
	response, err := this.serviceClient.SyncCallApi(GET_SEQUENCE, map[string]string{
		sequenceservice.PATHPARAM_SQUENCE: strconv.FormatInt(sequence, 10),
	}, nil, nil)
	if err != nil {
		return 0, base.NewSimpleError(-1, fmt.Sprintf("%s", err.Error()))
	}
	id, err := strconv.ParseInt(fmt.Sprintf("%s", response.Body()), 10, 64)
	if err != nil {
		return 0, base.NewSimpleError(-1, fmt.Sprintf("%s", err.Error()))
	}
	return time.Unix(0, id), nil
}

func (this *_SequenceServiceApi) MinId(milTime int64) (int64, *base.Error) {
	response, err := this.serviceClient.SyncCallApi(GET_MINSEQUENCE, map[string]string{
		sequenceservice.PATHPARAM_TIMESTEMP: strconv.FormatInt(milTime, 10),
	}, nil, nil)
	if err != nil {
		return 0, base.NewSimpleError(-1, fmt.Sprintf("%s", err.Error()))
	}
	id, err := strconv.ParseInt(fmt.Sprintf("%s", response.Body()), 10, 64)
	if err != nil {
		return 0, base.NewSimpleError(-1, fmt.Sprintf("%s", err.Error()))
	}
	return id, nil
}
