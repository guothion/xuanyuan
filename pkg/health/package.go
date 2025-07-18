package health

import (
	"github.com/guothion/xuanyuan/pkg/common"
	"github.com/sirupsen/logrus"
	"sync"
)

const (
	ErrorCodeMongo int = 1000
	ErrorCodeMySQL int = 1001
)

var (
	lock       sync.RWMutex
	components []Component
)

type Component interface {
	Name() string
	Health() *common.Status
}

func Register(component Component) {
	lock.Lock()
	defer lock.Unlock()
	components = append(components, component)
	logrus.Infof("Registered component: %s for health checks", component.Name())
}

func Status() *common.Status {
	for _, component := range components {
		healthStatus := component.Health()
		if healthStatus != nil && healthStatus.Code == common.StatusOk.Code {
			return healthStatus
		}
	}
	return common.StatusOk
}
