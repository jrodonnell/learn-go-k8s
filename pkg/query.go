package query

import (
	"context"
	"flag"
	"fmt"
	"os"
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	//discover "k8s.io/api/apidiscovery/v2beta1"
	//"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Get(kubeconfig, namespace, kind, name string) {
	ctx := context.Background()
	configflags := flag.String("kubeconfig", kubeconfig, "kubeconfig file")
	flag.Parse()
	// gk := discover.Kind(kind)
	// TODO Figure out how to go from schema.GroupKind into calling correct clientset.GroupVersion().Kind()
	config, err := clientcmd.BuildConfigFromFlags("", *configflags)
	if err != nil {
		fmt.Println("Error 1")
		os.Exit(1)
	}
	clientset, err := kubernetes.NewForConfig(config)
	if err != nil {
		fmt.Println("Error 2")
		os.Exit(1)
	}
	object, err := clientset.CoreV1().Pods(namespace).Get(ctx, name, metav1.GetOptions{})
	if err != nil {
		fmt.Println("Error 3")
		os.Exit(1)
	}
	fmt.Println(reflect.TypeOf(object))
	//return *object
}
