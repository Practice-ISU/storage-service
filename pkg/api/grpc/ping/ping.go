package ping

import (
	"context"
	"fmt"
	"time"

	// "net"
	// main_conf "storage-service/configs/grpc/main"

	"storage-service/pkg/grpc/discovery/ping"
	"storage-service/pkg/grpc/discovery/registration"

	"google.golang.org/grpc"
)

type service struct {
	Addr string
	Name string
}

type discoveryConfig interface {
	GetDiscoveryAddr() string
	GetPingPort() string
}

type serviceConfig interface {
	GetFilePort() string
	GetFolderPort() string

	GetIp() string

	GetFolderServiceName() string
	GetFileServiceName() string
}

type discoveryPingServer struct {
	ping.UnimplementedDiscoveryPingServer
	pingPort     string
	ip           string
	services     []service
	register     registration.ServiceRegistrationClient
	lastCallTime time.Time
}

func NewDiscoveryPingServer(discoveryCnf discoveryConfig, cnfService serviceConfig) (*discoveryPingServer, error) {
	fmt.Println("Discovery chanel = " + discoveryCnf.GetDiscoveryAddr())
	options := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	conn, err := grpc.Dial(discoveryCnf.GetDiscoveryAddr(), options...)
	if err != nil {
		return nil, err
	}
	register := registration.NewServiceRegistrationClient(conn)

	// folderPort := cnfService.GetFolderPort()

	return &discoveryPingServer{
		register: register,
		pingPort: discoveryCnf.GetPingPort(),
		services: []service{
			{Addr: cnfService.GetIp() + ":" + cnfService.GetFilePort(), Name: cnfService.GetFileServiceName()},
			{Addr: cnfService.GetIp() + ":" + cnfService.GetFolderPort(), Name: cnfService.GetFolderServiceName()},
		},
		ip: cnfService.GetIp(),
	}, nil
}

func (s *discoveryPingServer) Ping(context.Context, *ping.PingRequest) (*ping.PingResponse, error) {
	s.lastCallTime = time.Now()
	// fmt.Println("Pinged in", s.lastCallTime.String())
	return &ping.PingResponse{
		Timestamp: time.Now().Format("2006-01-02 15:04:05.000"),
		Success:   true,
	}, nil
}

func (s *discoveryPingServer) SendRegistrationRequest() {
	for {
		datas := make([]registration.ServiceRequest, len(s.services))
		flag := true

		for i, service := range s.services {
			datas[i] = registration.ServiceRequest{
				Timestamp:   time.Now().Format("2006-01-02 15:04:05.000"),
				ServiceName: service.Name,
				Channel:     "http://" + service.Addr,
				ChannelPing: "http://" + s.ip + ":" + s.pingPort,
			}

			result, err := s.register.Registration(context.TODO(), &datas[i])
			if err != nil {
				fmt.Println(err.Error())
				flag = false
				break
			}
			if result.Success {
				fmt.Println(service.Name+" registered in", result.Timestamp, "success -", result.Success)
			} else {
				fmt.Println("Error for", datas[i].ServiceName)
				flag = false
				break
			}
		}

		if flag {
			break
		}

		time.Sleep(time.Minute)
	}
}

func (s *discoveryPingServer) StartTimeout(f func()) {
	time.AfterFunc(6*time.Minute, func() {
		if time.Since(s.lastCallTime) > 6*time.Minute {
			fmt.Println("Lost ping from discovery-service!")
			f()
		}
	})

}
