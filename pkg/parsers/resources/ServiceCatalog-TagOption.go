package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ServiceCatalogTagOption Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-tagoption.html
type ServiceCatalogTagOption struct {
	Type       string                            `yaml:"Type"`
	Properties ServiceCatalogTagOptionProperties `yaml:"Properties"`
	Condition  interface{}                       `yaml:"Condition,omitempty"`
	Metadata   interface{}                       `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                       `yaml:"DependsOn,omitempty"`
}

// ServiceCatalogTagOption Properties
type ServiceCatalogTagOptionProperties struct {
	Active interface{} `yaml:"Active,omitempty"`
	Key    interface{} `yaml:"Key"`
	Value  interface{} `yaml:"Value"`
}

// NewServiceCatalogTagOption constructor creates a new ServiceCatalogTagOption
func NewServiceCatalogTagOption(properties ServiceCatalogTagOptionProperties, deps ...interface{}) ServiceCatalogTagOption {
	return ServiceCatalogTagOption{
		Type:       "AWS::ServiceCatalog::TagOption",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseServiceCatalogTagOption parses ServiceCatalogTagOption
func ParseServiceCatalogTagOption(
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
	errors []error,
) {
	source = "kombustion-core-resources"
	var resource ServiceCatalogTagOption
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

	return
}

// ParseServiceCatalogTagOption validator
func (resource ServiceCatalogTagOption) Validate() []error {
	return resource.Properties.Validate()
}

// ParseServiceCatalogTagOptionProperties validator
func (resource ServiceCatalogTagOptionProperties) Validate() []error {
	errors := []error{}
	if resource.Key == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Key'"))
	}
	if resource.Value == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'Value'"))
	}
	return errors
}
