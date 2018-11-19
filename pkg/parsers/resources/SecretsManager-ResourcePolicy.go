package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// SecretsManagerResourcePolicy Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-resourcepolicy.html
type SecretsManagerResourcePolicy struct {
	Type       string                                 `yaml:"Type"`
	Properties SecretsManagerResourcePolicyProperties `yaml:"Properties"`
	Condition  interface{}                            `yaml:"Condition,omitempty"`
	Metadata   interface{}                            `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                            `yaml:"DependsOn,omitempty"`
}

// SecretsManagerResourcePolicy Properties
type SecretsManagerResourcePolicyProperties struct {
	ResourcePolicy interface{} `yaml:"ResourcePolicy"`
	SecretId       interface{} `yaml:"SecretId"`
}

// NewSecretsManagerResourcePolicy constructor creates a new SecretsManagerResourcePolicy
func NewSecretsManagerResourcePolicy(properties SecretsManagerResourcePolicyProperties, deps ...interface{}) SecretsManagerResourcePolicy {
	return SecretsManagerResourcePolicy{
		Type:       "AWS::SecretsManager::ResourcePolicy",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSecretsManagerResourcePolicy parses SecretsManagerResourcePolicy
func ParseSecretsManagerResourcePolicy(
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
	var resource SecretsManagerResourcePolicy
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
					"Fn::Sub": "${AWS::StackName}-SecretsManagerResourcePolicy-" + name,
				},
			},
		},
	}

	return
}

// ParseSecretsManagerResourcePolicy validator
func (resource SecretsManagerResourcePolicy) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSecretsManagerResourcePolicyProperties validator
func (resource SecretsManagerResourcePolicyProperties) Validate() []error {
	errors := []error{}
	return errors
}
