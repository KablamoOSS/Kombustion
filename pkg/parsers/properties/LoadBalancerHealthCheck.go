package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// LoadBalancerHealthCheck Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-elb-health-check.html
type LoadBalancerHealthCheck struct {
	HealthyThreshold   interface{} `yaml:"HealthyThreshold"`
	Interval           interface{} `yaml:"Interval"`
	Target             interface{} `yaml:"Target"`
	Timeout            interface{} `yaml:"Timeout"`
	UnhealthyThreshold interface{} `yaml:"UnhealthyThreshold"`
}

// LoadBalancerHealthCheck validation
func (resource LoadBalancerHealthCheck) Validate() []error {
	errors := []error{}

	return errors
}
