package main

import (
	"fmt"
	"net/http"

	"github.com/hussainmuzamil/students-api/internal/config"
)

func main() {
	cfg := config.MustLoad()
	http.ListenAndServe(cfg.Addr,cfg.)
}
 