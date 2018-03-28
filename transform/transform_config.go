package transform

type TransformConfig struct {
	Transform_config []struct{
		Name string
		Probe_name string
		Rule string
	}
}
