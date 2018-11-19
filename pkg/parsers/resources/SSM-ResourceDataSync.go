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
)

// SSMResourceDataSync Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-ssm-resourcedatasync.html
type SSMResourceDataSync struct {
	Type       string                        `yaml:"Type"`
	Properties SSMResourceDataSyncProperties `yaml:"Properties"`
	Condition  interface{}                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                   `yaml:"DependsOn,omitempty"`
}

// SSMResourceDataSync Properties
type SSMResourceDataSyncProperties struct {
	BucketName   interface{} `yaml:"BucketName"`
	BucketPrefix interface{} `yaml:"BucketPrefix,omitempty"`
	BucketRegion interface{} `yaml:"BucketRegion"`
	KMSKeyArn    interface{} `yaml:"KMSKeyArn,omitempty"`
	SyncFormat   interface{} `yaml:"SyncFormat"`
	SyncName     interface{} `yaml:"SyncName"`
}

// NewSSMResourceDataSync constructor creates a new SSMResourceDataSync
func NewSSMResourceDataSync(properties SSMResourceDataSyncProperties, deps ...interface{}) SSMResourceDataSync {
	return SSMResourceDataSync{
		Type:       "AWS::SSM::ResourceDataSync",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSSMResourceDataSync parses SSMResourceDataSync
func ParseSSMResourceDataSync(
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
	var resource SSMResourceDataSync
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
					"Fn::Sub": "${AWS::StackName}-SSMResourceDataSync-" + name,
				},
			},
		},
	}

	return
}

// ParseSSMResourceDataSync validator
func (resource SSMResourceDataSync) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSSMResourceDataSyncProperties validator
func (resource SSMResourceDataSyncProperties) Validate() []error {
	errors := []error{}
	return errors
}
