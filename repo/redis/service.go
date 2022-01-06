package redis

import (
	"context"
	"net/url"

	"github.com/aceaura/libra/core/cast"
	"github.com/aceaura/libra/core/device"
)

type ServiceRequest struct {
	URL string
	Cmd []string
}

type ServiceResponse struct {
	Err     error
	Request *ServiceRequest
	Result  []string
}

type Service struct{}

func init() {
	device.Bus().WithService(&Service{})
}

func (s *Service) Redis(ctx context.Context, req *ServiceRequest) (resp *ServiceResponse, err error) {
	u, err := url.Parse(req.URL)
	if err != nil {
		return nil, err
	}

	addr := u.Host
	db := cast.ToInt(u.Path[1:])
	password, _ := u.User.Password()

	resp = new(ServiceResponse)
	c := NewClient().WithAddr(addr).WithDB(db).WithContext(ctx).WithPassword(password)
	result, err := c.Command(cast.ToSlice(req.Cmd)...)
	if err != nil {
		return nil, err
	}
	resp.Result = result
	return resp, nil
}

func (s *Service) RedisPipeline(ctx context.Context, req *ServiceRequest) (resp *ServiceResponse, err error) {
	return nil, nil
}
