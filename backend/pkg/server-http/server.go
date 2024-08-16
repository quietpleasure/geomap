package serverhttp

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"time"
)

type Option func(option *options) error

type options struct {
	host           *string
	port           *string
	maxheaderbytes *int
	writetimeout   *time.Duration
	readtimeout    *time.Duration
	idletimeout    *time.Duration
}

const (
	default_write_timeout = time.Duration(15 * time.Second)
	default_read_timeout  = time.Duration(15 * time.Second)
	default_idle_timeout  = time.Duration(60 * time.Second)
)

func New(ctx context.Context, handler http.Handler, opts ...Option) (*http.Server, error) {
	var opt options
	for _, option := range opts {
		if err := option(&opt); err != nil {
			return nil, err
		}
	}

	host := ""
	if opt.host != nil {
		host = *opt.host
	}
	port := ""
	if opt.port != nil {
		port = *opt.port
	}
	_, err := net.ResolveTCPAddr("tcp", fmt.Sprintf("%s:%s", host, port))
	if err != nil {
		return nil, err
	}
	var writetimeout, readtimeout, idletimeout time.Duration
	if opt.writetimeout == nil {
		writetimeout = default_write_timeout
	} else {
		writetimeout = *opt.writetimeout
	}
	if opt.readtimeout == nil {
		readtimeout = default_read_timeout
	} else {
		readtimeout = *opt.readtimeout
	}
	if opt.idletimeout == nil {
		idletimeout = default_idle_timeout
	} else {
		idletimeout = *opt.idletimeout
	}
	var maxheaderbytes int
	if opt.maxheaderbytes == nil {
		maxheaderbytes = http.DefaultMaxHeaderBytes
	} else {
		maxheaderbytes = *opt.maxheaderbytes
	}
	return &http.Server{
		Addr:           fmt.Sprintf("%s:%s", host, port),
		Handler:        handler,
		WriteTimeout:   writetimeout,
		ReadTimeout:    readtimeout,
		IdleTimeout:    idletimeout,
		MaxHeaderBytes: maxheaderbytes,
		BaseContext:    func(_ net.Listener) context.Context { return ctx },
	}, nil

}

// func Stop(ctx context.Context) error {
// 	// gracefullCtx, cancelShutdown := context.WithTimeout(s.BaseContext(nil), 5*time.Second)
// 	gracefullCtx, cancelShutdown := context.WithTimeout(ctx, 5*time.Second)
// 	defer cancelShutdown()
// 	s.SetKeepAlivesEnabled(false)
// 	return s.Shutdown(gracefullCtx)
// }

// func (s *Server) RegisterHandler(handler http.Handler) {
// 	if handler != nil {
// 		s.Handler = handler
// 	}
// }

func WithMaxHeaderBytes(bts int) Option {
	return func(options *options) error {
		options.maxheaderbytes = &bts
		return nil
	}
}

func WithWriteTimeout(timeout time.Duration) Option {
	return func(options *options) error {
		options.writetimeout = &timeout
		return nil
	}
}

func WithReadTimeout(timeout time.Duration) Option {
	return func(options *options) error {
		options.readtimeout = &timeout
		return nil
	}
}

func WithIdleTimeout(timeout time.Duration) Option {
	return func(options *options) error {
		options.idletimeout = &timeout
		return nil
	}
}

func WithHost(host string) Option {
	return func(options *options) error {
		options.host = &host
		return nil
	}
}

// if port=0 listening to random available port
func WithPort(port int) Option {
	return func(options *options) error {
		if port < 0 {
			return fmt.Errorf("port cannot be less than zero")
		}
		p := fmt.Sprintf("%d", port)
		options.port = &p
		return nil
	}
}

// func (s *Server) AwaitStop() {
// 	sig := make(chan os.Signal, 1)
// 	signal.Notify(sig,
// 		os.Interrupt,
// 		syscall.SIGINT,
// 		syscall.SIGABRT,
// 		syscall.SIGQUIT,
// 		syscall.SIGTERM,
// 		syscall.SIGHUP,
// 	)
// 	<-sig
// 	gracefullCtx, cancelShutdown := context.WithTimeout(s.BaseContext(nil), 5*time.Second)
// 	defer cancelShutdown()
// 	s.SetKeepAlivesEnabled(false)

// 	return s.Shutdown(gracefullCtx)
// }
