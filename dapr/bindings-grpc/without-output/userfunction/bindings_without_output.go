package userfunction

import (
	"github.com/dapr/go-sdk/service/common"
	ofctx "github.com/tpiperatgod/offf-go/openfunction-context"
	"log"
)

func BindingsNoOutput(ctx *ofctx.OpenFunctionContext, in interface{}) int {
	input := in.(*common.BindingEvent)
	log.Printf("binding - Data:%s, Meta:%v", input.Data, input.Metadata)
	return 200
}
