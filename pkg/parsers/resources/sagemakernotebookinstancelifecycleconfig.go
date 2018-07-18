package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// SageMakerNotebookInstanceLifecycleConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-sagemaker-notebookinstancelifecycleconfig.html
type SageMakerNotebookInstanceLifecycleConfig struct {
	Type       string                                             `yaml:"Type"`
	Properties SageMakerNotebookInstanceLifecycleConfigProperties `yaml:"Properties"`
	Condition  interface{}                                        `yaml:"Condition,omitempty"`
	Metadata   interface{}                                        `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                                        `yaml:"DependsOn,omitempty"`
}

// SageMakerNotebookInstanceLifecycleConfig Properties
type SageMakerNotebookInstanceLifecycleConfigProperties struct {
	NotebookInstanceLifecycleConfigName interface{} `yaml:"NotebookInstanceLifecycleConfigName,omitempty"`
	OnCreate                            interface{} `yaml:"OnCreate,omitempty"`
	OnStart                             interface{} `yaml:"OnStart,omitempty"`
}

// NewSageMakerNotebookInstanceLifecycleConfig constructor creates a new SageMakerNotebookInstanceLifecycleConfig
func NewSageMakerNotebookInstanceLifecycleConfig(properties SageMakerNotebookInstanceLifecycleConfigProperties, deps ...interface{}) SageMakerNotebookInstanceLifecycleConfig {
	return SageMakerNotebookInstanceLifecycleConfig{
		Type:       "AWS::SageMaker::NotebookInstanceLifecycleConfig",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseSageMakerNotebookInstanceLifecycleConfig parses SageMakerNotebookInstanceLifecycleConfig
func ParseSageMakerNotebookInstanceLifecycleConfig(name string, data string) (cf types.TemplateObject, err error) {
	var resource SageMakerNotebookInstanceLifecycleConfig
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: SageMakerNotebookInstanceLifecycleConfig - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseSageMakerNotebookInstanceLifecycleConfig validator
func (resource SageMakerNotebookInstanceLifecycleConfig) Validate() []error {
	return resource.Properties.Validate()
}

// ParseSageMakerNotebookInstanceLifecycleConfigProperties validator
func (resource SageMakerNotebookInstanceLifecycleConfigProperties) Validate() []error {
	errs := []error{}
	return errs
}
