package main

import (
	"github.com/coffeehc/baseservices/sequenceservice"
	"net/http"
	"strconv"

	"github.com/coffeehc/microserviceboot/base"
	"github.com/coffeehc/microserviceboot/serviceboot"
	"github.com/coffeehc/web"
)

func (this *SequenceService) post_sequence(request *http.Request, pathFragments map[string]string, reply web.Reply) {
	defer serviceboot.ErrorRecover(reply)
	reply.With(strconv.FormatInt(this.sequenceService.NextId(), 10)).As(web.Transport_Text)
}

func (this *SequenceService) get_sequence(request *http.Request, pathFragments map[string]string, reply web.Reply) {
	defer serviceboot.ErrorRecover(reply)
	sequence, err := strconv.ParseInt(pathFragments[sequenceservice.PATHPARAM_SQUENCE], 10, 64)
	if err != nil {
		panic(base.NewBizErr(base.ERROR_CODE_BASE_INVALID_PARAMTER, "非法的 sequence"))
	}
	sequenceId := this.sequenceService.ParseSequence(sequence)
	reply.With(strconv.FormatInt(sequenceId.CreateTime.UnixNano(), 10)).As(web.Transport_Text)
}

func (this *SequenceService) get_minSequence(request *http.Request, pathFragments map[string]string, reply web.Reply) {
	defer serviceboot.ErrorRecover(reply)
	sequence, err := strconv.ParseInt(pathFragments[sequenceservice.PATHPARAM_TIMESTEMP], 10, 64)
	if err != nil {
		panic(base.NewBizErr(base.ERROR_CODE_BASE_INVALID_PARAMTER, "非法的 sequence"))
	}
	reply.With(strconv.FormatInt(this.sequenceService.MinId(sequence), 10)).As(web.Transport_Text)
}
