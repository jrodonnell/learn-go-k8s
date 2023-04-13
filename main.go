package main

import (
	"learn/pkg"
)

func main() {
	pkg.Get("/home/jrodonnell/.kube/config.master", "cattle-system", "pod", "rancher-5f9b79d5fc-x89zh")
}
