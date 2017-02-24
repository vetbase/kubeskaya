package client

import "k8s.io/client-go/kubernetes"

// Client for interacting with Kubernetes
type Client struct {
	client *kubernetes.Clientset
}

func New() *Client {
	c := new(Client)
	return c
}
