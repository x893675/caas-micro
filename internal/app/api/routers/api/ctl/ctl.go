package ctl

import "go.uber.org/dig"

func Inject(container *dig.Container) error {
	container.Provide(NewLogin)
	return nil
}
