package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// SecretsManagerSecretTargetAttachment Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-secretsmanager-secrettargetattachment.html
type SecretsManagerSecretTargetAttachment struct {
	Type       string                                         `yaml:"Type"`
	Properties SecretsManagerSecretTargetAttachmentProperties `yaml:"Properties"`
	Condition  interface{}                                    `yaml:"Condition,omitempty"`
	Metadata   interface{}                                    `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                    `yaml:"DependsOn,omitempty"`
}

// SecretsManagerSecretTargetAttachment Properties
type SecretsManagerSecretTargetAttachmentProperties struct {
	SecretId   interface{} `yaml:"SecretId"`
	TargetId   interface{} `yaml:"TargetId"`
	TargetType interface{} `yaml:"TargetType"`
}

// NewSecretsManagerSecretTargetAttachment constructor creates a new SecretsManagerSecretTargetAttachment
func NewSecretsManagerSecretTargetAttachment(properties SecretsManagerSecretTargetAttachmentProperties, deps ...interface{}) SecretsManagerSecretTargetAttachment {
	return SecretsManagerSecretTargetAttachment{
		Type:       "AWS::SecretsManager::SecretTargetAttachment",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSecretsManagerSecretTargetAttachment parses SecretsManagerSecretTargetAttachment
func ParseSecretsManagerSecretTargetAttachment(
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
	var resource SecretsManagerSecretTargetAttachment
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
					"Fn::Sub": "${AWS::StackName}-SecretsManagerSecretTargetAttachment-" + name,
				},
			},
		},
	}

	return
}

// ParseSecretsManagerSecretTargetAttachment validator
func (resource SecretsManagerSecretTargetAttachment) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSecretsManagerSecretTargetAttachmentProperties validator
func (resource SecretsManagerSecretTargetAttachmentProperties) Validate() []error {
	errors := []error{}
	return errors
}
