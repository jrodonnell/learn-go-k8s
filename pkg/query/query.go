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
	config, err := clientcmd.BuildConfigFromFlags("", *kubeconfig) // read and parse kuebconfig file

	if err != nil {
		fmt.Println("1: ", err)
		os.Exit(1)
	}

	clientset, err := kubernetes.NewForConfig(config) // create new client for querying, clientset is dynamic

	if err != nil {
		fmt.Println("2: ", err)
		os.Exit(1)
	}

	object, err := clientset.CoreV1().Pods(*namespace).Get(ctx, *name, metav1.GetOptions{}) // get specific object by walking clientset inheritance tree

	if err != nil {
		fmt.Println("3: ", err)
		os.Exit(1)
	}

	objectmeta := object.ObjectMeta               // get metav1.ObjectMeta, everything in the .metadata field from kubectl get
	rt := reflect.TypeOf(objectmeta)              // get Type of objectmeta so we can make a map of the right size
	output := make(map[string]any, rt.NumField()) // create map to store json of objectmeta so we can delete managedFields later
	marshal, err := json.Marshal(objectmeta)      // capture values of objectmeta in json so we can write it to the output map

	if err != nil {
		fmt.Println("4: ", err)
		os.Exit(1)
	}

	err = json.Unmarshal(marshal, &output) // write output json values back to output map object

	if err != nil {
		fmt.Println("5: ", err)
		os.Exit(1)
	}

	delete(output, "managedFields") // delete managedFields from output map object
	fmt.Println(output)
}
