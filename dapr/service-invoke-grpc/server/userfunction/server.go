package userfunction

import (
	"github.com/dapr/go-sdk/service/common"
	ofctx "github.com/tpiperatgod/offf-go/openfunction-context"
	"log"
)

func Server(ctx *ofctx.OpenFunctionContext, in interface{}) int {
	input := in.(*common.InvocationEvent)
	if input == nil {
		log.Printf("nil invocation parameter")
		return 500
	}
	log.Printf(
		"echo - ContentType:%s, Verb:%s, QueryString:%s, %s",
		input.ContentType, input.Verb, input.QueryString, input.Data,
	)
	return 200
}
