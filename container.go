package patterns

import (
	apiv1 "k8s.io/api/core/v1"
	"strings"
)

type Env apiv1.EnvVar

// ToContainer
func ToContainer(c *Container) apiv1.Container {
	return apiv1.Container{
		Name:            strings.ToLower(c.Config.Name),
		Image:           strings.ToLower(c.Config.Image),
		Command:         c.Config.Command,
		Args:            c.Config.Args,
		Env:             c.Config.Env,
		ImagePullPolicy: apiv1.PullAlways,
		VolumeMounts:    ToVolumeMounts(c.Config.VolumeMounts),
		Ports: []apiv1.ContainerPort{
			{
				ContainerPort: int32(c.Config.Port),
			},
		},
	}
}

// ToContainers -
func ToContainers(containers []*Container) []apiv1.Container {
	updated := make([]apiv1.Container, 0)
	for _, val := range containers {
		updated = append(updated, ToContainer(val))
	}
	return updated
}

// Config -
type ContainerConfig struct {
	Name         string
	Image        string
	Port         int
	Command      []string
	Env          []apiv1.EnvVar
	Args         []string
	Volumes      []Volume
	VolumeMounts []VolumeMount
}

// Container -
type Container struct {
	Config ContainerConfig
}

// AddVolumeMount -
func (c *Container) AddVolumeMount(volume Volume) *Container {
	c.Config.Volumes = append(c.Config.Volumes, volume)
	return c
}

// SetEnv -
func (c *Container) SetEnv(env []apiv1.EnvVar) *Container {
	c.Config.Env = env
	return c
}

// SetVolumeMounts -
func (c *Container) SetVolumeMounts(vms []VolumeMount) *Container {
	c.Config.VolumeMounts = vms
	return c
}

// AddEnv -
func (c *Container) AddEnv(env apiv1.EnvVar) *Container {
	c.Config.Env = append(c.Config.Env, env)
	return c
}

// AddArgs -
func (c *Container) AddArgs(args []string) *Container {
	c.Config.Args = args
	return c
}

// NewContainer -
func NewContainer(config ContainerConfig) *Container {
	return &Container{Config: config}
}

// Generate -
func (c *Container) Generate() apiv1.Container {
	val := apiv1.Container{
		Name:            strings.ToLower(c.Config.Name),
		Image:           strings.ToLower(c.Config.Image),
		Command:         c.Config.Command,
		Args:            c.Config.Args,
		Env:             c.Config.Env,
		ImagePullPolicy: apiv1.PullAlways,
		VolumeMounts:    ToVolumeMounts(c.Config.VolumeMounts),
		Ports: []apiv1.ContainerPort{
			{
				ContainerPort: int32(c.Config.Port),
			},
		},
	}

	return val
}
