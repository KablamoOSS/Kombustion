package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ServiceAwsVpcConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-service-awsvpcconfiguration.html
type ServiceAwsVpcConfiguration struct {
	AssignPublicIp interface{} `yaml:"AssignPublicIp,omitempty"`
	SecurityGroups interface{} `yaml:"SecurityGroups,omitempty"`
	Subnets        interface{} `yaml:"Subnets"`
}

// ServiceAwsVpcConfiguration validation
func (resource ServiceAwsVpcConfiguration) Validate() []error {
	errors := []error{}

	return errors
}
