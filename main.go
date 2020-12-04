package main

import (
	"context"
	"errors"
	"fmt"
	//"encoding/json"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	"k8s.io/client-go/rest"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	done := make(chan os.Signal, 1)
	// catch SIGETRM or SIGINTERRUPT
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	go func() {
		for {
			// get an in-cluster client (will need sufficient RBAC access and a serviceaccount)
			config, err := rest.InClusterConfig()
			if err != nil {
				panic(err.Error())
			}
			clientset, err := kubernetes.NewForConfig(config)
			if err != nil {
				panic(err.Error())
			}

			// get current namespace using an env var from the downward API, which must be configured as such
			namespace := os.Getenv("POD_NAMESPACE")
			if namespace == "" {
				err := errors.New("environment variable POD_NAMESPACE not defined, define it with the downward API")
				panic(err.Error())
			}

			// now use the client we configured above to get the list of pods in the provided namespace
			pods, err := clientset.CoreV1().Pods(namespace).List(context.TODO(), metav1.ListOptions{})
			if err != nil {
				panic(err.Error())
			}

			// marshal it to something slightly more readable
			//podsData, err := json.Marshal(pods)
			//if err != nil {
			//	panic(err.Error())
			//}

			// just dump it out as json because we weren't asked for anything fancier
			//fmt.Printf("%s\n", podsData)

			// ok make it prettier than that!
			for _, m := range pods.Items {
				fmt.Println(m.ObjectMeta.Name, m.ObjectMeta.Namespace)
			}
		}
	}()
	<-done
	// any clean up in here, <- done will block until SIGTERM or SIGINTERRUPT
}
