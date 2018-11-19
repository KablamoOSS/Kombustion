package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// CognitoUserPoolUser Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-userpooluser.html
type CognitoUserPoolUser struct {
	Type       string                        `yaml:"Type"`
	Properties CognitoUserPoolUserProperties `yaml:"Properties"`
	Condition  interface{}                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                   `yaml:"DependsOn,omitempty"`
}

// CognitoUserPoolUser Properties
type CognitoUserPoolUserProperties struct {
	ForceAliasCreation     interface{} `yaml:"ForceAliasCreation,omitempty"`
	MessageAction          interface{} `yaml:"MessageAction,omitempty"`
	UserPoolId             interface{} `yaml:"UserPoolId"`
	Username               interface{} `yaml:"Username,omitempty"`
	DesiredDeliveryMediums interface{} `yaml:"DesiredDeliveryMediums,omitempty"`
	UserAttributes         interface{} `yaml:"UserAttributes,omitempty"`
	ValidationData         interface{} `yaml:"ValidationData,omitempty"`
}

// NewCognitoUserPoolUser constructor creates a new CognitoUserPoolUser
func NewCognitoUserPoolUser(properties CognitoUserPoolUserProperties, deps ...interface{}) CognitoUserPoolUser {
	return CognitoUserPoolUser{
		Type:       "AWS::Cognito::UserPoolUser",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCognitoUserPoolUser parses CognitoUserPoolUser
func ParseCognitoUserPoolUser(
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
	var resource CognitoUserPoolUser
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
					"Fn::Sub": "${AWS::StackName}-CognitoUserPoolUser-" + name,
				},
			},
		},
	}

	return
}

// ParseCognitoUserPoolUser validator
func (resource CognitoUserPoolUser) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCognitoUserPoolUserProperties validator
func (resource CognitoUserPoolUserProperties) Validate() []error {
	errors := []error{}
	return errors
}
