package main

import (
	"flag"
	"learn/pkg/query"
	"os"
)

func main() {
	kubeconfig := flag.String("k", os.ExpandEnv("$HOME/.kube/config"), "kubeconfig file")
	namespace := flag.String("ns", "default", "Namespace to Get from")
	name := flag.String("n", "", "Pod to Get")
	flag.Parse()
	query.Get(kubeconfig, namespace, name)
}
