package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	//
	// "fmt"
	//
	//
	//"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	//
)

// DirectoryServiceMicrosoftAD Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-directoryservice-microsoftad.html
type DirectoryServiceMicrosoftAD struct {
	Type       string                                `yaml:"Type"`
	Properties DirectoryServiceMicrosoftADProperties `yaml:"Properties"`
	Condition  interface{}                           `yaml:"Condition,omitempty"`
	Metadata   interface{}                           `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                           `yaml:"DependsOn,omitempty"`
}

// DirectoryServiceMicrosoftAD Properties
type DirectoryServiceMicrosoftADProperties struct {
	CreateAlias interface{} `yaml:"CreateAlias,omitempty"`
	Edition     interface{} `yaml:"Edition,omitempty"`
	EnableSso   interface{} `yaml:"EnableSso,omitempty"`
	Name        interface{} `yaml:"Name"`
	Password    interface{} `yaml:"Password"`
	ShortName   interface{} `yaml:"ShortName,omitempty"`
	VpcSettings interface{} `yaml:"VpcSettings"`
}

// NewDirectoryServiceMicrosoftAD constructor creates a new DirectoryServiceMicrosoftAD
func NewDirectoryServiceMicrosoftAD(properties DirectoryServiceMicrosoftADProperties, deps ...interface{}) DirectoryServiceMicrosoftAD {
	return DirectoryServiceMicrosoftAD{
		Type:       "AWS::DirectoryService::MicrosoftAD",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseDirectoryServiceMicrosoftAD parses DirectoryServiceMicrosoftAD
func ParseDirectoryServiceMicrosoftAD(
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
	var resource DirectoryServiceMicrosoftAD
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
					"Fn::Sub": "${AWS::StackName}-DirectoryServiceMicrosoftAD-" + name,
				},
			},
		},
	}

	return
}

// ParseDirectoryServiceMicrosoftAD validator
func (resource DirectoryServiceMicrosoftAD) Validate() []error {
	return resource.Properties.Validate()
}

// ParseDirectoryServiceMicrosoftADProperties validator
func (resource DirectoryServiceMicrosoftADProperties) Validate() []error {
	errors := []error{}
	return errors
}
