package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EMRInstanceFleetConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-elasticmapreduce-instancefleetconfig.html
type EMRInstanceFleetConfig struct {
	Type       string                           `yaml:"Type"`
	Properties EMRInstanceFleetConfigProperties `yaml:"Properties"`
	Condition  interface{}                      `yaml:"Condition,omitempty"`
	Metadata   interface{}                      `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                      `yaml:"DependsOn,omitempty"`
}

// EMRInstanceFleetConfig Properties
type EMRInstanceFleetConfigProperties struct {
	ClusterId              interface{} `yaml:"ClusterId"`
	InstanceFleetType      interface{} `yaml:"InstanceFleetType"`
	Name                   interface{} `yaml:"Name,omitempty"`
	TargetOnDemandCapacity interface{} `yaml:"TargetOnDemandCapacity,omitempty"`
	TargetSpotCapacity     interface{} `yaml:"TargetSpotCapacity,omitempty"`
	InstanceTypeConfigs    interface{} `yaml:"InstanceTypeConfigs,omitempty"`
	LaunchSpecifications   interface{} `yaml:"LaunchSpecifications,omitempty"`
}

// NewEMRInstanceFleetConfig constructor creates a new EMRInstanceFleetConfig
func NewEMRInstanceFleetConfig(properties EMRInstanceFleetConfigProperties, deps ...interface{}) EMRInstanceFleetConfig {
	return EMRInstanceFleetConfig{
		Type:       "AWS::EMR::InstanceFleetConfig",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEMRInstanceFleetConfig parses EMRInstanceFleetConfig
func ParseEMRInstanceFleetConfig(
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
	var resource EMRInstanceFleetConfig
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
					"Fn::Sub": "${AWS::StackName}-EMRInstanceFleetConfig-" + name,
				},
			},
		},
	}

	return
}

// ParseEMRInstanceFleetConfig validator
func (resource EMRInstanceFleetConfig) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEMRInstanceFleetConfigProperties validator
func (resource EMRInstanceFleetConfigProperties) Validate() []error {
	errors := []error{}
	return errors
}
