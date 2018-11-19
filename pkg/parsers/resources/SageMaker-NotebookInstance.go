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

// SageMakerNotebookInstance Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstance.html
type SageMakerNotebookInstance struct {
	Type       string                              `yaml:"Type"`
	Properties SageMakerNotebookInstanceProperties `yaml:"Properties"`
	Condition  interface{}                         `yaml:"Condition,omitempty"`
	Metadata   interface{}                         `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                         `yaml:"DependsOn,omitempty"`
}

// SageMakerNotebookInstance Properties
type SageMakerNotebookInstanceProperties struct {
	DirectInternetAccess interface{} `yaml:"DirectInternetAccess,omitempty"`
	InstanceType         interface{} `yaml:"InstanceType"`
	KmsKeyId             interface{} `yaml:"KmsKeyId,omitempty"`
	LifecycleConfigName  interface{} `yaml:"LifecycleConfigName,omitempty"`
	NotebookInstanceName interface{} `yaml:"NotebookInstanceName,omitempty"`
	RoleArn              interface{} `yaml:"RoleArn"`
	SubnetId             interface{} `yaml:"SubnetId,omitempty"`
	VolumeSizeInGB       interface{} `yaml:"VolumeSizeInGB,omitempty"`
	SecurityGroupIds     interface{} `yaml:"SecurityGroupIds,omitempty"`
	Tags                 interface{} `yaml:"Tags,omitempty"`
}

// NewSageMakerNotebookInstance constructor creates a new SageMakerNotebookInstance
func NewSageMakerNotebookInstance(properties SageMakerNotebookInstanceProperties, deps ...interface{}) SageMakerNotebookInstance {
	return SageMakerNotebookInstance{
		Type:       "AWS::SageMaker::NotebookInstance",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSageMakerNotebookInstance parses SageMakerNotebookInstance
func ParseSageMakerNotebookInstance(
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
	var resource SageMakerNotebookInstance
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
					"Fn::Sub": "${AWS::StackName}-SageMakerNotebookInstance-" + name,
				},
			},
		},
	}

	return
}

// ParseSageMakerNotebookInstance validator
func (resource SageMakerNotebookInstance) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSageMakerNotebookInstanceProperties validator
func (resource SageMakerNotebookInstanceProperties) Validate() []error {
	errors := []error{}
	return errors
}
