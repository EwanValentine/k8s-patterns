package patterns

import (
	"github.com/go-playground/validator/v10"
	v12 "k8s.io/api/apps/v1"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

var validate *validator.Validate

// DeploymentConfig -
type DeploymentConfig struct {
	Name           string `validate:"required"`
	Replicas       *int32 `validate:"numeric"`
	Labels         *map[string]string
	Containers     []*Container
	InitContainers []*Container
	Volumes        []Volume
	Annotations    *map[string]string
	NodeSelector   *map[string]string
}

// Deployment -
type Deployment struct {
	Config DeploymentConfig
}

// NewDeployment -
func NewDeployment(config DeploymentConfig) *Deployment {
	return &Deployment{Config: config}
}

// SetReplicas -
func (d *Deployment) SetReplicas(replicas int32) *Deployment {
	r := &replicas
	d.Config.Replicas = r
	return d
}

// SetAnnotations -
func (d *Deployment) SetAnnotations(annotations map[string]string) *Deployment {
	d.Config.Annotations = &annotations
	return d
}

// SetContainers -
func (d *Deployment) SetContainers(containers []*Container) *Deployment {
	d.Config.Containers = containers
	return d
}

// AddContainer -
func (d *Deployment) AddContainer(container *Container) *Deployment {
	d.Config.Containers = append(d.Config.Containers, container)
	return d
}

// SetInitContainers -
func (d *Deployment) SetInitContainers(initContainers []*Container) *Deployment {
	d.Config.InitContainers = initContainers
	return d
}

// SetLabels -
func (d *Deployment) SetLabels(labels map[string]string) *Deployment {
	d.Config.Labels = &labels
	return d
}

// GetPort -
func (d *Deployment) GetPort() int {
	return d.Config.Containers[0].Config.Port
}

// Generate -
func (d *Deployment) Generate() (interface{}, error) {
	validate = validator.New()
	if err := validate.Struct(d.Config); err != nil {
		return nil, err
	}

	deployment := &v12.Deployment{
		ObjectMeta: metav1.ObjectMeta{
			Name: d.Config.Name,
		},
		Spec: v12.DeploymentSpec{
			Replicas: d.Config.Replicas,
			Template: apiv1.PodTemplateSpec{
				Spec: apiv1.PodSpec{
					InitContainers: ToContainers(d.Config.InitContainers),
					Containers:     ToContainers(d.Config.Containers),
					RestartPolicy:  apiv1.RestartPolicyAlways,
					Volumes:        ToVolumes(d.Config.Volumes),
				},
			},
		},
	}

	if d.Config.Annotations != nil {
		deployment.Spec.Template.ObjectMeta.Annotations = *d.Config.Annotations
	}

	if d.Config.NodeSelector != nil {
		deployment.Spec.Template.Spec.NodeSelector = *d.Config.NodeSelector
	}

	if len(d.Config.InitContainers) > 0 {
		deployment.Spec.Template.Spec.InitContainers = ToContainers(d.Config.InitContainers)
	}

	if d.Config.Labels != nil {
		deployment.Spec.Template.ObjectMeta.Labels = *d.Config.Labels
		deployment.Spec.Template.Labels = *d.Config.Labels
	}

	return deployment, nil
}
