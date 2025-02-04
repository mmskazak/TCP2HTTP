package tcp2http

// Config представляет конфигурацию модуля, которую можно задавать через config.json.
// Пример конфигурации:
//	  "tcp2http": {
//		"config": {
//		  "listen": ":9004",
//		  "target_url": "http://localhost:8081",
//		  "headers": {
//			"X-Test-Header": "test-value"
//		  }
//		}
//	  }

type Config struct {
	// Listen задаёт адрес и порт, на которых будет слушать TCP-соединения.
	Listen string `json:"listen"`

	// TargetURL указывает адрес, на который будет отправляться HTTP-запрос.
	TargetURL string `json:"target_url"`

	// Headers содержит набор HTTP-заголовков, которые добавляются к запросу.
	Headers map[string]string `json:"headers"`
}
