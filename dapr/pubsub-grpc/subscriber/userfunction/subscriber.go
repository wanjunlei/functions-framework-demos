package userfunction

import (
	ofctx "github.com/OpenFunction/functions-framework-go/openfunction-context"
	"github.com/dapr/go-sdk/service/common"
	"log"
)

func Subscriber(ctx *ofctx.OpenFunctionContext, in interface{}) int {
	input := in.(*common.TopicEvent)
	log.Printf("event - PubsubName:%s, Topic:%s, ID:%s, Data: %s", input.PubsubName, input.Topic, input.ID, input.Data)
	return 200
}
