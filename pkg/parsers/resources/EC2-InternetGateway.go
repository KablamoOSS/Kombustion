package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EC2InternetGateway Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ec2-internetgateway.html
type EC2InternetGateway struct {
	Type       string                       `yaml:"Type"`
	Properties EC2InternetGatewayProperties `yaml:"Properties"`
	Condition  interface{}                  `yaml:"Condition,omitempty"`
	Metadata   interface{}                  `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                  `yaml:"DependsOn,omitempty"`
}

// EC2InternetGateway Properties
type EC2InternetGatewayProperties struct {
	Tags interface{} `yaml:"Tags,omitempty"`
}

// NewEC2InternetGateway constructor creates a new EC2InternetGateway
func NewEC2InternetGateway(properties EC2InternetGatewayProperties, deps ...interface{}) EC2InternetGateway {
	return EC2InternetGateway{
		Type:       "AWS::EC2::InternetGateway",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2InternetGateway parses EC2InternetGateway
func ParseEC2InternetGateway(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"

	// Resources
	var resource EC2InternetGateway
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	// Outputs

	outputs = types.TemplateObject{
		name: types.TemplateObject{
			"Description": name + " Object",
			"Value": map[string]interface{}{
				"Ref": name,
			},
			"Export": map[string]interface{}{
				"Name": map[string]interface{}{
					"Fn::Sub": "${AWS::StackName}-EC2InternetGateway-" + name,
				},
			},
		},
	}

	return
}

// ParseEC2InternetGateway validator
func (resource EC2InternetGateway) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2InternetGatewayProperties validator
func (resource EC2InternetGatewayProperties) Validate() []error {
	errors := []error{}
	return errors
}
