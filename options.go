package tcpwrapper

import (
	isrequest "github.com/mmskazak/tcpwrapper/is_request"
	isresponse "github.com/mmskazak/tcpwrapper/is_response"
	"go.uber.org/zap"
)

// Option defines a functional option for configuring tcpWrapper
type Option func(*tcpWrapper)

// WithRequestDelimiter sets the request delimiter
func WithRequestDelimiter(delimiter []byte) Option {
	return func(w *tcpWrapper) {
		w.requestDelimiter = delimiter
	}
}

// WithResponseDelimiter sets the response delimiter
func WithResponseDelimiter(delimiter []byte) Option {
	return func(w *tcpWrapper) {
		w.responseDelimiter = delimiter
	}
}

// WithRequestChecker sets the request checker function
func WithRequestChecker(checker isrequest.IsRequestFunc) Option {
	return func(w *tcpWrapper) {
		w.isRequest = checker
	}
}

// WithResponseChecker sets the response checker function
func WithResponseChecker(checker isresponse.IsResponseFunc) Option {
	return func(w *tcpWrapper) {
		w.isResponse = checker
	}
}

// WithLogger sets the logger
func WithLogger(logger *zap.Logger) Option {
	return func(w *tcpWrapper) {
		w.logger = logger.Sugar()
	}
}
