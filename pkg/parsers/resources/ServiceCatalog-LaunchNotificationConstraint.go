package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ServiceCatalogLaunchNotificationConstraint Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-servicecatalog-launchnotificationconstraint.html
type ServiceCatalogLaunchNotificationConstraint struct {
	Type       string                                               `yaml:"Type"`
	Properties ServiceCatalogLaunchNotificationConstraintProperties `yaml:"Properties"`
	Condition  interface{}                                          `yaml:"Condition,omitempty"`
	Metadata   interface{}                                          `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                          `yaml:"DependsOn,omitempty"`
}

// ServiceCatalogLaunchNotificationConstraint Properties
type ServiceCatalogLaunchNotificationConstraintProperties struct {
	AcceptLanguage   interface{} `yaml:"AcceptLanguage,omitempty"`
	Description      interface{} `yaml:"Description,omitempty"`
	PortfolioId      interface{} `yaml:"PortfolioId"`
	ProductId        interface{} `yaml:"ProductId"`
	NotificationArns interface{} `yaml:"NotificationArns"`
}

// NewServiceCatalogLaunchNotificationConstraint constructor creates a new ServiceCatalogLaunchNotificationConstraint
func NewServiceCatalogLaunchNotificationConstraint(properties ServiceCatalogLaunchNotificationConstraintProperties, deps ...interface{}) ServiceCatalogLaunchNotificationConstraint {
	return ServiceCatalogLaunchNotificationConstraint{
		Type:       "AWS::ServiceCatalog::LaunchNotificationConstraint",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseServiceCatalogLaunchNotificationConstraint parses ServiceCatalogLaunchNotificationConstraint
func ParseServiceCatalogLaunchNotificationConstraint(
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
	var resource ServiceCatalogLaunchNotificationConstraint
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
					"Fn::Sub": "${AWS::StackName}-ServiceCatalogLaunchNotificationConstraint-" + name,
				},
			},
		},
	}

	return
}

// ParseServiceCatalogLaunchNotificationConstraint validator
func (resource ServiceCatalogLaunchNotificationConstraint) Validate() []error {
	return resource.Properties.Validate()
}

// ParseServiceCatalogLaunchNotificationConstraintProperties validator
func (resource ServiceCatalogLaunchNotificationConstraintProperties) Validate() []error {
	errors := []error{}
	return errors
}
