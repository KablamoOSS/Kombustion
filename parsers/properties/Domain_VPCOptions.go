package properties

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

type Domain_VPCOptions struct {
	SecurityGroupIds interface{} `yaml:"SecurityGroupIds,omitempty"`
	SubnetIds        interface{} `yaml:"SubnetIds,omitempty"`
}

func (resource Domain_VPCOptions) Validate() []error {
	errs := []error{}

	return errs
}
