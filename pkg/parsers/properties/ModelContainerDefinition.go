package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// ModelContainerDefinition Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-sagemaker-model-containerdefinition.html
type ModelContainerDefinition struct {
	ContainerHostname interface{} `yaml:"ContainerHostname,omitempty"`
	Environment       interface{} `yaml:"Environment,omitempty"`
	Image             interface{} `yaml:"Image"`
	ModelDataUrl      interface{} `yaml:"ModelDataUrl,omitempty"`
}

// ModelContainerDefinition validation
func (resource ModelContainerDefinition) Validate() []error {
	errors := []error{}

	return errors
}
