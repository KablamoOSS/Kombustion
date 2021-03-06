package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ElastiCacheSecurityGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-security-group.html
type ElastiCacheSecurityGroup struct {
	Type       string                             `yaml:"Type"`
	Properties ElastiCacheSecurityGroupProperties `yaml:"Properties"`
	Condition  interface{}                        `yaml:"Condition,omitempty"`
	Metadata   interface{}                        `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                        `yaml:"DependsOn,omitempty"`
}

// ElastiCacheSecurityGroup Properties
type ElastiCacheSecurityGroupProperties struct {
	Description interface{} `yaml:"Description"`
}

// NewElastiCacheSecurityGroup constructor creates a new ElastiCacheSecurityGroup
func NewElastiCacheSecurityGroup(properties ElastiCacheSecurityGroupProperties, deps ...interface{}) ElastiCacheSecurityGroup {
	return ElastiCacheSecurityGroup{
		Type:       "AWS::ElastiCache::SecurityGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseElastiCacheSecurityGroup parses ElastiCacheSecurityGroup
func ParseElastiCacheSecurityGroup(
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
	var resource ElastiCacheSecurityGroup
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
					"Fn::Sub": "${AWS::StackName}-ElastiCacheSecurityGroup-" + name,
				},
			},
		},
	}

	return
}

// ParseElastiCacheSecurityGroup validator
func (resource ElastiCacheSecurityGroup) Validate() []error {
	return resource.Properties.Validate()
}

// ParseElastiCacheSecurityGroupProperties validator
func (resource ElastiCacheSecurityGroupProperties) Validate() []error {
	errors := []error{}
	return errors
}
