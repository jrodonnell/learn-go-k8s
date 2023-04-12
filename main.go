package main

import query "query/pkg"

func main() {
	query.Get("/home/jrodonnell/.kube/config.master", "cattle-system", "pod", "rancher-5f9b79d5fc-x89zh")
}
