package redis

import (
	"context"
	"net/url"

	"github.com/aceaura/libra/core/cast"
	"github.com/aceaura/libra/core/device"
)

type CommandRequest struct {
	URL string
	Cmd []string
}

type CommandResponse struct {
	Result []string
}

type PipelineRequest struct {
}

type PipelineResponse struct {
}

type Service struct{}

func init() {
	device.Bus().WithDevice(device.NewRouter().WithName("Redis").WithService(&Service{}))
}

func (s *Service) Command(ctx context.Context, req *CommandRequest) (resp *CommandResponse, err error) {
	u, err := url.Parse(req.URL)
	if err != nil {
		return nil, err
	}

	addr := u.Host
	db := cast.ToInt(u.Path[1:])
	password, _ := u.User.Password()

	resp = new(CommandResponse)
	c := NewClient().WithAddr(addr).WithDB(db).WithContext(ctx).WithPassword(password)
	result, err := c.Command(cast.ToSlice(req.Cmd)...)
	if err != nil {
		return nil, err
	}
	resp.Result = result
	return resp, nil
}

// TODO ADD POOL

func (s *Service) Pipeline(ctx context.Context, req *PipelineRequest) (resp *PipelineResponse, err error) {
	return nil, nil
}
