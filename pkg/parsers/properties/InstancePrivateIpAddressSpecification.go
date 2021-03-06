package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// InstancePrivateIpAddressSpecification Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-network-interface-privateipspec.html
type InstancePrivateIpAddressSpecification struct {
	Primary          interface{} `yaml:"Primary"`
	PrivateIpAddress interface{} `yaml:"PrivateIpAddress"`
}

// InstancePrivateIpAddressSpecification validation
func (resource InstancePrivateIpAddressSpecification) Validate() []error {
	errors := []error{}

	return errors
}
