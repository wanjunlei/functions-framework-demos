package userfunction

import (
	ofctx "github.com/tpiperatgod/offf-go/openfunction-context"
	"github.com/dapr/go-sdk/service/common"
	"log"
)

func BindingsOutput(ctx *ofctx.OpenFunctionContext, in interface{}) int {
	input := in.(*common.BindingEvent)
	log.Printf("binding - Data:%s, Meta:%v", input.Data, input.Metadata)

	greeting := []byte("Hello")
	err := ctx.SendTo(greeting, "echo")
	if err != nil {
		log.Printf("Error: %v\n", err)
		return 500
	}
	return 200
}
