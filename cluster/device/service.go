package device

import (
	"context"
	"reflect"
	"unicode"
	"unicode/utf8"

	"github.com/aceaura/libra/codec"
	"github.com/aceaura/libra/scheduler"
)

type Service struct {
	Device
	codec         codec.Codec
	schedulerFunc func(context.Context) *scheduler.Scheduler
	handlers      map[string]*Handler
	gateway       Device
}

func (s *Service) String() string {
	return reflect.Indirect(reflect.ValueOf(s)).Type().Name()
}

func (s *Service) Gateway(device Device) {
	s.gateway = device
}

func (s *Service) Process(ctx context.Context, route Route, data []byte) error {
	deviceType := route.deviceType()
	if deviceType == DeviceTypeBus {
		return s.gateway.Process(ctx, route, data)
	} else if deviceType == DeviceTypeService {
		return s.localProcess(ctx, route, data)
	}

	return ErrRouteDeadEnd
}

func (s *Service) localProcess(ctx context.Context, route Route, data []byte) error {
	name := route.deviceName()
	handler, ok := s.handlers[name]
	if !ok {
		return ErrRouteMissingDevice
	}
	return handler.Process(ctx, route, data)
}

func (s *Service) scheduler(ctx context.Context) *scheduler.Scheduler {
	if s.schedulerFunc != nil {
		return s.schedulerFunc(ctx)
	}
	return scheduler.Default()
}

func (s *Service) ExtractHandlers() {
	if !s.isExported() {
		return
	}

	t := reflect.TypeOf(s)

	for index := 0; index < t.NumMethod(); index++ {
		method := t.Method(index)
		if !s.isMethodHandler(method) {
			continue
		}

		h := &Handler{
			method: method,
		}
		h.Gateway(s)
		s.handlers[h.String()] = h
	}
}

func (s *Service) isExported() bool {
	typeName := reflect.Indirect(reflect.ValueOf(s)).Type().Name()
	w, _ := utf8.DecodeRuneInString(typeName)
	return unicode.IsUpper(w)
}

func (*Service) isMethodHandler(method reflect.Method) bool {
	mt := method.Type
	// Check method is exported
	if mt.PkgPath() != "" {
		return false
	}

	// Check num in
	if mt.NumIn() != 3 {
		return false
	}

	// Check num out
	if mt.NumOut() != 2 {
		return false
	}

	// Check context.Context
	if t := mt.In(1); !t.Implements(typeOfContext) {
		return false
	}

	// Check error
	if t := mt.Out(1); !t.Implements(typeOfBytes) {
		return false
	}

	// Check request:  pointer or bytes
	if t := mt.In(2); t.Kind() != reflect.Ptr && t != typeOfBytes {
		return false
	}

	// Check response: pointer or bytes
	if t := mt.Out(0); t.Kind() != reflect.Ptr && t != typeOfBytes {
		return false
	}

	return true
}