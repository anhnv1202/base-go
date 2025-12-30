package main

import "github.com/anhnv1202/base-go/internal/routers"

func main() {
  router := routers.NewRouter()
  router.Run() // listens on 0.0.0.0:8080 by default
}


