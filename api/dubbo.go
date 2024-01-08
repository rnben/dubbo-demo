package api

type DubboRequest struct {
	Request map[string]interface{}
}

func (u *DubboRequest) JavaClassName() string {
	return "org.apache.dubbo.DubboRequest"
}

type DubboResponse struct {
	Reponse []byte
}

func (u *DubboResponse) JavaClassName() string {
	return "org.apache.dubbo.DubboResponse"
}
