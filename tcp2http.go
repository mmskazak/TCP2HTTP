package tcp2http

import (
	"context"
	"log"
	"net"
	"sync"

	"go.uber.org/zap"

	"github.com/caddyserver/caddy/v2"
)

// TCPForwarder является основным типом модуля для Caddy, реализующим логику получения
// данных по TCP и пересылки их в виде HTTP-запроса.
type TCPForwarder struct {
	// Конфигурация модуля, загруженная из config.json.
	Config Config `json:"config,omitempty"`

	// Логгер для записи логов.
	log *zap.Logger

	// TCP-слушатель для принятия входящих соединений.
	ln net.Listener

	// Контекст для управления жизненным циклом сервера.
	ctx    context.Context
	cancel context.CancelFunc

	// WaitGroup для ожидания завершения всех горутин.
	wg sync.WaitGroup
}

// Интерфейсная защита: убеждаемся, что TCPForwarder реализует необходимые интерфейсы Caddy.
var (
	_ caddy.Module      = (*TCPForwarder)(nil)
	_ caddy.Provisioner = (*TCPForwarder)(nil)
	_ caddy.App         = (*TCPForwarder)(nil)
)

// init регистрирует модуль в системе Caddy.
func init() {
	log.Println("Модуль tcp2http загружен!") // Проверим, вызывается ли init()

	caddy.RegisterModule(new(TCPForwarder))
}

// CaddyModule возвращает информацию о модуле для Caddy.
func (t *TCPForwarder) CaddyModule() caddy.ModuleInfo {
	return caddy.ModuleInfo{
		ID:  "tcp2http",
		New: func() caddy.Module { return new(TCPForwarder) },
	}
}
