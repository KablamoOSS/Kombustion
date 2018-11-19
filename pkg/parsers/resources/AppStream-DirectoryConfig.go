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

// AppStreamDirectoryConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-appstream-directoryconfig.html
type AppStreamDirectoryConfig struct {
	Type       string                             `yaml:"Type"`
	Properties AppStreamDirectoryConfigProperties `yaml:"Properties"`
	Condition  interface{}                        `yaml:"Condition,omitempty"`
	Metadata   interface{}                        `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                        `yaml:"DependsOn,omitempty"`
}

// AppStreamDirectoryConfig Properties
type AppStreamDirectoryConfigProperties struct {
	DirectoryName                        interface{} `yaml:"DirectoryName"`
	ServiceAccountCredentials            interface{} `yaml:"ServiceAccountCredentials"`
	OrganizationalUnitDistinguishedNames interface{} `yaml:"OrganizationalUnitDistinguishedNames"`
}

// NewAppStreamDirectoryConfig constructor creates a new AppStreamDirectoryConfig
func NewAppStreamDirectoryConfig(properties AppStreamDirectoryConfigProperties, deps ...interface{}) AppStreamDirectoryConfig {
	return AppStreamDirectoryConfig{
		Type:       "AWS::AppStream::DirectoryConfig",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseAppStreamDirectoryConfig parses AppStreamDirectoryConfig
func ParseAppStreamDirectoryConfig(
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
	var resource AppStreamDirectoryConfig
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
					"Fn::Sub": "${AWS::StackName}-AppStreamDirectoryConfig-" + name,
				},
			},
		},
	}

	return
}

// ParseAppStreamDirectoryConfig validator
func (resource AppStreamDirectoryConfig) Validate() []error {
	return resource.Properties.Validate()
}

// ParseAppStreamDirectoryConfigProperties validator
func (resource AppStreamDirectoryConfigProperties) Validate() []error {
	errors := []error{}
	return errors
}
