package main

import (
	"context"
	"dubbo-demo/api"
	"fmt"
	"log"
	"time"

	"dubbo.apache.org/dubbo-go/v3/config"
	_ "dubbo.apache.org/dubbo-go/v3/imports"
	hessian "github.com/apache/dubbo-go-hessian2"
)

// export DUBBO_GO_CONFIG_PATH= PATH_TO_SAMPLES/direct/go-server/conf/dubbogo.yml
func main() {
	config.SetProviderService(&DubboDemoProvider{})
	hessian.RegisterPOJO(&api.DubboRequest{})
	hessian.RegisterPOJO(&api.DubboResponse{})

	if err := config.Load(); err != nil {
		panic(err)
	}
	select {}
}

type DubboDemoProvider struct{}

func (d *DubboDemoProvider) SayHello(ctx context.Context, req *api.DubboRequest) (resp *api.DubboResponse, err error) {
	st := time.Now()

	defer func() {
		log.Printf("SayHello cost:%dms\n", time.Since(st).Milliseconds())
	}()

	cost, _ := req.Request["cost"].(string)

	t, err := time.ParseDuration(cost)
	if err != nil {
		return nil, err
	}

	time.Sleep(t)

	msg := fmt.Sprintf("Hello, this request cost %v", t)

	return &api.DubboResponse{Reponse: []byte(msg)}, nil
}
