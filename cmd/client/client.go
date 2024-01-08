package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	hessian "github.com/apache/dubbo-go-hessian2"

	"dubbo-demo/api"
)

// export DUBBO_GO_CONFIG_PATH= PATH_TO_SAMPLES/direct/go-client/conf/dubbogo.yml
func main() {
	config.SetConsumerService(dubboDemoImpl)
	hessian.RegisterPOJO(&api.DubboRequest{})
	hessian.RegisterPOJO(&api.DubboResponse{})

	if err := config.Load(); err != nil {
		panic(err)
	}

	for {
		cost := fmt.Sprintf("%ss", strconv.Itoa(Random(3, 10)))
		req := &api.DubboRequest{
			Request: map[string]interface{}{
				"cost": cost,
			},
		}
		reply, err := dubboDemoImpl.SayHello(context.Background(), req)
		if err != nil {
			panic(err)
		}
		log.Printf("client response result: %s\n", reply.Reponse)
	}
}

var dubboDemoImpl = new(DubboDemoProvider)

type DubboDemoProvider struct {
	SayHello func(ctx context.Context, req *api.DubboRequest) (resp *api.DubboResponse, err error)
}

func Random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
