package patterns

import (
	"github.com/go-playground/validator/v10"
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/util/intstr"
)

// ServiceConfig -
type ServiceConfig struct {
	Name        string
	Labels      map[string]string
	Annotations map[string]string
	Type        apiv1.ServiceType
	Port        int
	TargetPort  int
	Selector    map[string]string
	Replicas int
}

// Service -
type Service struct {
	Config     ServiceConfig
	Deployment *Deployment
}

// NewService -
func NewService(config ServiceConfig) *Service {
	return &Service{Config: config}
}

// SetDeployment -
func (s *Service) SetDeployment(deployment *Deployment) *Service {
	s.Deployment = deployment
	return s
}

// Generate -
func (s *Service) Generate() (interface{}, error) {
	validate = validator.New()
	if err := validate.Struct(s.Config); err != nil {
		return nil, err
	}

	service := &apiv1.Service{
		ObjectMeta: metav1.ObjectMeta{
			Name:        s.Config.Name,
			Labels:      s.Config.Labels,
			Annotations: s.Config.Annotations,
		},
		Spec: apiv1.ServiceSpec{
			Selector: s.Config.Selector,
			Type:     s.Config.Type,
			Ports: []apiv1.ServicePort{
				{
					Port:       int32(s.Config.Port),
					TargetPort: intstr.FromInt(s.Config.TargetPort),
					Protocol:   "TCP",
				},
			},
		},
	}

	if s.Deployment != nil {
		service.Spec.Ports[0].TargetPort = intstr.FromInt(s.Deployment.GetPort())
	}

	return service, nil
}
