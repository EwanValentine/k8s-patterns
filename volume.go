package patterns

import apiv1 "k8s.io/api/core/v1"

// VolumeConfig -
type VolumeConfig struct {
	Name   string
	Source VolumeSource
}

// ToVolumes -
func ToVolumes(volumes []Volume) []apiv1.Volume {
	updated := make([]apiv1.Volume, 0)
	for _, v := range volumes {
		updated = append(updated, ToVolume(v))
	}
	return updated
}

// VolumeSource -
type VolumeSource struct{}

// ToVolumeSource -
func ToVolumeSource(vs VolumeSource) apiv1.VolumeSource {
	return apiv1.VolumeSource{}
}

func ToVolumeMount(vm VolumeMount) apiv1.VolumeMount {
	return apiv1.VolumeMount{}
}

func ToVolumeMounts(vms []VolumeMount) []apiv1.VolumeMount {
	updated := make([]apiv1.VolumeMount, 0)
	for _, v := range vms {
		updated = append(updated, ToVolumeMount(v))
	}
	return updated
}

// ToVolume -
func ToVolume(volume Volume) apiv1.Volume {
	return apiv1.Volume{
		Name:         volume.Config.Name,
		VolumeSource: ToVolumeSource(volume.Config.Source),
	}
}

// Volume -
type Volume struct {
	Config VolumeConfig
}

// VolumeMount -
type VolumeMount struct{}

// NewVolume -
func NewVolume(config VolumeConfig) *Volume {
	return &Volume{Config: config}
}
