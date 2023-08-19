package microfiber

import (
	"fmt"
	"time"

	"github.com/gofiber/fiber/v2"
)

type ServiceOptions struct {
	Host string
	Port uint16
	//重试次数
	Retry uint

	//重试间隔时间
	RetryInterval time.Duration
}

func RunService(app *fiber.App, config *ServiceOptions) error {
	err := Retry(
		func() (err error) {
			err = app.Listen(fmt.Sprintf("%s:%d", config.Host, config.Port))
			if err != nil {
				return
			}
			return
		},
		config.Retry,
		config.RetryInterval,
	)
	if err != nil {
		return err
	}
	return nil
}
