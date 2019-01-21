package di

import (
	"github.com/tampopos/go.api-template/src/2-interface-adapters/ctrl"
	"go.uber.org/dig"
)

// CreateContainer はDIContainerを生成します
func CreateContainer() (*dig.Container, error) {
	container := dig.New()
	if err := container.Provide(ctrl.NewHome); err != nil {
		return nil, err
	}
	return container, nil
}
