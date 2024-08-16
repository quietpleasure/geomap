package server

type API interface {
	// http.Handler

	Version() string
}
