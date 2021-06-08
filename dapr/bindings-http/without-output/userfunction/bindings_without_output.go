package userfunction

import (
	ofctx "github.com/tpiperatgod/offf-go/openfunction-context"
	"io/ioutil"
	"log"
	"net/http"
)

func BindingsNoOutput(ctx *ofctx.OpenFunctionContext, in interface{}) int {
	input := in.(*http.Request)
	content, err := ioutil.ReadAll(input.Body)
	if err != nil {
		log.Printf("Error: %v\n", err)
		return 500
	}
	log.Printf("binding - Data:%s, Header:%v", string(content), input.Header)
	return 200
}
