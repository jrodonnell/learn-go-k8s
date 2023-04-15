package query

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"reflect"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/tools/clientcmd"
)

func Get(kubeconfig, namespace, name *string) {
	ctx := context.Background()
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig)

	if err != nil {
		fmt.Println("1: ", err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config)

	if err != nil {
		fmt.Println("2: ", err)
		os.Exit(1)
	}

	object, err := clientset.CoreV1().Pods(*namespace).Get(ctx, *name, metav1.GetOptions{})

	if err != nil {
		fmt.Println("3: ", err)
		os.Exit(1)
	}

	objectmeta := object.ObjectMeta
	rt := reflect.TypeOf(objectmeta)
	output := make(map[string]any, rt.NumField())
	marshal, err := json.Marshal(objectmeta)

	if err != nil {
		fmt.Println("4: ", err)
		os.Exit(1)
	}

	err = json.Unmarshal(marshal, &output)

	if err != nil {
		fmt.Println("5: ", err)
		os.Exit(1)
	}

	delete(output, "managedFields")
	fmt.Println(output)
}
