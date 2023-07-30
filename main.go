package main

import (
	"flag"
	"learn/pkg/k8s"
	"os"
)

func main() {
	kubeconfig := flag.String("k", os.ExpandEnv("$HOME/.kube/config"), "kubeconfig file")
	namespace := flag.String("ns", "default", "Namespace to Get from")
	name := flag.String("n", "", "Pod to Get")
	flag.Parse()
	k8s.Get(kubeconfig, namespace, name)
}
