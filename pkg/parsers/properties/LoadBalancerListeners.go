package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// LoadBalancerListeners Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-listener.html
type LoadBalancerListeners struct {
	InstancePort     interface{} `yaml:"InstancePort"`
	InstanceProtocol interface{} `yaml:"InstanceProtocol,omitempty"`
	LoadBalancerPort interface{} `yaml:"LoadBalancerPort"`
	Protocol         interface{} `yaml:"Protocol"`
	SSLCertificateId interface{} `yaml:"SSLCertificateId,omitempty"`
	PolicyNames      interface{} `yaml:"PolicyNames,omitempty"`
}

// LoadBalancerListeners validation
func (resource LoadBalancerListeners) Validate() []error {
	errors := []error{}

	return errors
}
