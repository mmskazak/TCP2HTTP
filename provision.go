package tcp2http

import (
	"context"

	"github.com/caddyserver/caddy/v2"
	"go.uber.org/zap"
)

// Provision инициализирует модуль, создавая логгер и контекст.
// Этот метод вызывается системой Caddy во время инициализации модуля.
func (t *TCPForwarder) Provision(ctx caddy.Context) error {
	t.log = ctx.Logger(t)
	t.ctx, t.cancel = context.WithCancel(context.Background())

	t.log.Info("Загружена конфигурация",
		zap.String("listen", t.Config.Listen),
		zap.String("target_url", t.Config.TargetURL),
		zap.Any("headers", t.Config.Headers),
	)

	return nil
}
