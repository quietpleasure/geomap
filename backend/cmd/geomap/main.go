package main

import (
	// "context"

	l "log"
	"net/http"
	_ "net/http/pprof"

	"geomap/internal/app"
)

func main() {
	// net/http/pprof
	go func() {
		l.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	app.Create().Run()

}
