package patterns

import (
	apiv1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// SecretConfig -
type SecretConfig struct {
	Name   string
	Values map[string][]byte
}

// Secret -
type Secret struct {
	SecretConfig
}

// Generate -
func (c *Secret) Generate() *apiv1.Secret {
	secret := &apiv1.Secret{
		ObjectMeta: metav1.ObjectMeta{
			Name: c.SecretConfig.Name,
		},
		Data: c.SecretConfig.Values,
	}

	return secret
}
