package tcp2http

import (
	"bytes"
	"io"
	"net"
	"net/http"

	"go.uber.org/zap"
)

// Start запускает TCP-сервер.
func (t *TCPForwarder) Start() error {
	ln, err := net.Listen("tcp", t.Config.Listen)
	if err != nil {
		t.log.Error("Ошибка запуска TCP сервера", zap.Error(err))
		return err
	}
	t.ln = ln
	t.log.Info("TCP сервер запущен", zap.String("адрес", t.Config.Listen))
	t.wg.Add(1)
	go t.acceptConnections()
	return nil
}

// acceptConnections обрабатывает входящие соединения.
func (t *TCPForwarder) acceptConnections() {
	defer t.wg.Done()
	for {
		conn, err := t.ln.Accept()
		if err != nil {
			t.log.Error("Ошибка принятия соединения", zap.Error(err))
			continue
		}
		t.wg.Add(1)
		go t.handleConnection(conn)
	}
}

// handleConnection обрабатывает одно TCP-соединение.
func (t *TCPForwarder) handleConnection(conn net.Conn) {
	defer conn.Close()
	defer t.wg.Done()

	t.log.Info("Новое соединение", zap.String("адрес", conn.RemoteAddr().String()))

	// Читаем входящие данные
	data, err := io.ReadAll(conn)
	if err != nil {
		t.log.Error("Ошибка чтения TCP-сообщения", zap.Error(err))
		return
	}

	respBody, err := t.forwardToHTTP(data)
	if err != nil {
		return
	}

	// Отправляем ответ обратно в TCP
	if _, err := conn.Write(respBody); err != nil {
		t.log.Error("Ошибка отправки ответа клиенту", zap.Error(err))
	}
}

// forwardToHTTP отправляет данные на HTTP-сервер и возвращает ответ.
func (t *TCPForwarder) forwardToHTTP(data []byte) ([]byte, error) {
	t.log.Info("Отправляем данные через HTTP", zap.String("url", t.Config.TargetURL))

	req, err := http.NewRequest(http.MethodPost, t.Config.TargetURL, bytes.NewReader(data))
	if err != nil {
		t.log.Error("Ошибка создания HTTP-запроса", zap.Error(err))
		return nil, err
	}

	for key, value := range t.Config.Headers {
		req.Header.Set(key, value)
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		t.log.Error("Ошибка выполнения HTTP-запроса", zap.Error(err))
		return nil, err
	}
	defer resp.Body.Close()

	t.log.Info("HTTP-запрос выполнен", zap.Int("статус", resp.StatusCode))

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.log.Error("Ошибка чтения тела ответа", zap.Error(err))
		return nil, err
	}

	return respBody, nil
}
