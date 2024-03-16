package consul

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/google/uuid"
	"github.com/hashicorp/consul/api"
)

type ConsulClient struct {
	*api.Client
}

func ServiceRegistryWithConsul(ipAddress string, port int, myUUID uuid.UUID) {
	config := api.DefaultConfig()
	consul, err := api.NewClient(config)
	if err != nil {
		log.Println(err)
	}

	/* Mỗi trường hợp dịch vụ nên có một dịch vụ duy nhất */
	serviceID := fmt.Sprintf("hello-%v", myUUID)
	/* Tag nên tuân theo quy tắc của fabio: urlprefix- */
	tags := []string{"urlprefix-/"}

	// Dockerport: Điều này được tiêm trong lệnh `Docker Run`. Nó không tồn tại khi ứng dụng Go chạy bên ngoài container docker
	dockerContainerPort, _ := strconv.Atoi(os.Getenv("DOCKERPORT"))

	registration := &api.AgentServiceRegistration{
		ID:      serviceID,
		Name:    "hello-todo",
		Port:    dockerContainerPort,
		Address: ipAddress,
		Tags:    tags, /* Thêm thẻ để đăng ký */
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("http://%s:%v/health", ipAddress, dockerContainerPort),
			Interval: "10s",
			Timeout:  "30s",
		},
	}

	registrationErr := consul.Agent().ServiceRegister(registration)

	if registrationErr != nil {
		log.Printf("Failed to register service: %s:%v ", ipAddress, dockerContainerPort)
	} else {
		log.Printf("successfully register service: %s:%v", ipAddress, dockerContainerPort)
	}
}

func NewClient(addr string) (*ConsulClient, error) {
	conf := &api.Config{
		Address: addr,
	}
	client, err := api.NewClient(conf)
	if err != nil {
		log.Println("error initiating new consul client: ", err)
		return &ConsulClient{}, err
	}

	return &ConsulClient{
		client,
	}, nil
}
