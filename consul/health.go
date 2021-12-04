package consul

import (
	"context"
	"google.golang.org/grpc/health/grpc_health_v1"
	"log"
)

type HealthImpl struct {
	Status grpc_health_v1.HealthCheckResponse_ServingStatus
	Reason string
}

func (h *HealthImpl) Watch(*grpc_health_v1.HealthCheckRequest, grpc_health_v1.Health_WatchServer) error {
	return nil
}

func (h *HealthImpl) OffLine(reason string) {
	h.Status = grpc_health_v1.HealthCheckResponse_NOT_SERVING
	h.Reason = reason
	log.Println(reason)
}

func (h *HealthImpl) OnLine(reason string) {
	h.Status = grpc_health_v1.HealthCheckResponse_SERVING
	h.Reason = reason
	log.Println(reason)
}

//Check 实现健康检查接口，这里直接返回健康状态，这里也可以有更复杂的健康检查策略，比如根据服务器负载来返回
func (h *HealthImpl) Check(ctx context.Context, req *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{
		Status: h.Status,
	}, nil
}
