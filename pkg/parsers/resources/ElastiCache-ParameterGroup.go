package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// ElastiCacheParameterGroup Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticache-parameter-group.html
type ElastiCacheParameterGroup struct {
	Type       string                              `yaml:"Type"`
	Properties ElastiCacheParameterGroupProperties `yaml:"Properties"`
	Condition  interface{}                         `yaml:"Condition,omitempty"`
	Metadata   interface{}                         `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                         `yaml:"DependsOn,omitempty"`
}

// ElastiCacheParameterGroup Properties
type ElastiCacheParameterGroupProperties struct {
	CacheParameterGroupFamily interface{} `yaml:"CacheParameterGroupFamily"`
	Description               interface{} `yaml:"Description"`
	Properties                interface{} `yaml:"Properties,omitempty"`
}

// NewElastiCacheParameterGroup constructor creates a new ElastiCacheParameterGroup
func NewElastiCacheParameterGroup(properties ElastiCacheParameterGroupProperties, deps ...interface{}) ElastiCacheParameterGroup {
	return ElastiCacheParameterGroup{
		Type:       "AWS::ElastiCache::ParameterGroup",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseElastiCacheParameterGroup parses ElastiCacheParameterGroup
func ParseElastiCacheParameterGroup(
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
	var resource ElastiCacheParameterGroup
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
					"Fn::Sub": "${AWS::StackName}-ElastiCacheParameterGroup-" + name,
				},
			},
		},
	}

	return
}

// ParseElastiCacheParameterGroup validator
func (resource ElastiCacheParameterGroup) Validate() []error {
	return resource.Properties.Validate()
}

// ParseElastiCacheParameterGroupProperties validator
func (resource ElastiCacheParameterGroupProperties) Validate() []error {
	errors := []error{}
	return errors
}
