package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// CognitoIdentityPool Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-cognito-identitypool.html
type CognitoIdentityPool struct {
	Type       string                        `yaml:"Type"`
	Properties CognitoIdentityPoolProperties `yaml:"Properties"`
	Condition  interface{}                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                   `yaml:"DependsOn,omitempty"`
}

// CognitoIdentityPool Properties
type CognitoIdentityPoolProperties struct {
	AllowUnauthenticatedIdentities interface{} `yaml:"AllowUnauthenticatedIdentities"`
	CognitoEvents                  interface{} `yaml:"CognitoEvents,omitempty"`
	DeveloperProviderName          interface{} `yaml:"DeveloperProviderName,omitempty"`
	IdentityPoolName               interface{} `yaml:"IdentityPoolName,omitempty"`
	SupportedLoginProviders        interface{} `yaml:"SupportedLoginProviders,omitempty"`
	PushSync                       interface{} `yaml:"PushSync,omitempty"`
	CognitoIdentityProviders       interface{} `yaml:"CognitoIdentityProviders,omitempty"`
	OpenIdConnectProviderARNs      interface{} `yaml:"OpenIdConnectProviderARNs,omitempty"`
	SamlProviderARNs               interface{} `yaml:"SamlProviderARNs,omitempty"`
	CognitoStreams                 interface{} `yaml:"CognitoStreams,omitempty"`
}

// NewCognitoIdentityPool constructor creates a new CognitoIdentityPool
func NewCognitoIdentityPool(properties CognitoIdentityPoolProperties, deps ...interface{}) CognitoIdentityPool {
	return CognitoIdentityPool{
		Type:       "AWS::Cognito::IdentityPool",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseCognitoIdentityPool parses CognitoIdentityPool
func ParseCognitoIdentityPool(
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
	var resource CognitoIdentityPool
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
					"Fn::Sub": "${AWS::StackName}-CognitoIdentityPool-" + name,
				},
			},
		},
	}

	return
}

// ParseCognitoIdentityPool validator
func (resource CognitoIdentityPool) Validate() []error {
	return resource.Properties.Validate()
}

// ParseCognitoIdentityPoolProperties validator
func (resource CognitoIdentityPoolProperties) Validate() []error {
	errors := []error{}
	return errors
}
