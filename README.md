# K8s Patterns

Library for creating Kubernetes resources programmatically.

## Example

```golang
package main

import (
	patterns "github.com/EwanValentine/k8s-patterns"
	"log"
)

func main() {
	app := patterns.NewApp("ais-service")

	container := patterns.NewContainer(patterns.ContainerConfig{
		Name: "test",
		Image: "nginx:latest",
		Port: 8080,
	})
	deployment := patterns.NewDeployment(patterns.DeploymentConfig{
		Name: "test",
	})
	deployment.SetReplicas(1)
	deployment.AddContainer(container)

	service := patterns.NewService(patterns.ServiceConfig{})
	service.SetDeployment(deployment)

	//ingress := patterns.NewIngress(patterns.IngressConfig{})
	//ingress.SetService(service)

	if err := app.
		AddDeployment(deployment).
		AddService(service).
		Deploy(); err != nil {
		log.Panic(err)
	}
}
```
