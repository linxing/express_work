package main

import (
	"./routers"
	"./api"
)

func main() {
	servlet := new(api.Servlet)
	r := routers.SetupRouter(servlet)
	r.Run()
}
