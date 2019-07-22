package api

import (
	"../api/hello"
)

type (
	HelloServlet struct{ hello.Servlet }

	Servlet struct {
		HelloServlet
	}
)
