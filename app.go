package patterns

import (
	"fmt"
	v12 "k8s.io/api/apps/v1"
	batchv1 "k8s.io/api/batch/v1beta1"
	apiv1 "k8s.io/api/core/v1"
	v1beta13 "k8s.io/api/networking/v1beta1"
	v1 "k8s.io/client-go/kubernetes/typed/apps/v1"
	v1b "k8s.io/client-go/kubernetes/typed/batch/v1beta1"
	corev1 "k8s.io/client-go/kubernetes/typed/core/v1"
	v1beta12 "k8s.io/client-go/kubernetes/typed/networking/v1beta1"
	"log"
)

// NewApp -
func NewApp(namespace string) *App {
	client := NewClientInstance()

	deployClient := client.AppsV1()
	serviceClient := client.CoreV1()
	jobsClient := client.BatchV1beta1()
	ingressClient := client.NetworkingV1beta1()

	items := make([]Item, 0)

	return &App{
		namespace,
		deployClient,
		serviceClient,
		jobsClient,
		ingressClient,
		items,
	}
}

// Generator -
type Generator interface {
	Generate() (interface{}, error)
}

// Item -
type Item struct {
	Type string
	Generator
}

// App -
type App struct {
	Namespace        string
	DeploymentClient v1.AppsV1Interface
	ServiceClient    corev1.CoreV1Interface
	JobsClient       v1b.BatchV1beta1Interface
	IngressClient    v1beta12.NetworkingV1beta1Interface
	Items            []Item
}

type Stringer interface {
	String() string
}

// String -
func (a *App) String() {
	for _, val := range a.Items {
		v, err := val.Generator.Generate()
		if err != nil {
			log.Panic(err)
		}
		s := v.(Stringer).String()
		fmt.Println(s)
	}
}

// AddDeployment -
func (a *App) AddDeployment(deployment Generator) *App {
	a.Items = append(a.Items, Item{
		Type:      "deployment",
		Generator: deployment,
	})
	return a
}

// AddService -
func (a *App) AddService(service Generator) *App {
	a.Items = append(a.Items, Item{
		Type:      "service",
		Generator: service,
	})
	return a
}

// AddIngress -
func (a *App) AddIngress(ingress Generator) *App {
	a.Items = append(a.Items, Item{
		Type:      "ingress",
		Generator: ingress,
	})
	return a
}

// Preview -
func (a *App) Preview() {
	a.String()
}

// Deploy -
func (a *App) Deploy() error {
	for _, val := range a.Items {
		switch val.Type {
		case "deployment":
			result, err := val.Generator.Generate()
			if err != nil {
				return err
			}

			r := result.(*v12.Deployment)
			if _, err := a.DeploymentClient.Deployments(a.Namespace).Create(r); err != nil {
				return err
			}
		case "service":
			result, err := val.Generator.Generate()
			if err != nil {
				return err
			}

			r := result.(*apiv1.Service)
			if _, err := a.ServiceClient.Services(a.Namespace).Create(r); err != nil {
				return err
			}
		case "secret":
			result, err := val.Generator.Generate()
			if err != nil {
				return err
			}

			r := result.(*apiv1.Secret)
			if _, err := a.ServiceClient.Secrets(a.Namespace).Create(r); err != nil {
				return err
			}
		case "ingress":
			result, err := val.Generator.Generate()
			if err != nil {
				return err
			}

			r := result.(*v1beta13.Ingress)
			if _, err := a.IngressClient.Ingresses(a.Namespace).Create(r); err != nil {
				return err
			}
		case "job":
			result, err := val.Generator.Generate()
			if err != nil {
				return err
			}

			r := result.(*batchv1.CronJob)
			if _, err := a.JobsClient.CronJobs(a.Namespace).Create(r); err != nil {
				return err
			}
		}
	}
	return nil
}
