package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// EC2VolumeAttachment Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ec2-ebs-volumeattachment.html
type EC2VolumeAttachment struct {
	Type       string                        `yaml:"Type"`
	Properties EC2VolumeAttachmentProperties `yaml:"Properties"`
	Condition  interface{}                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                   `yaml:"DependsOn,omitempty"`
}

// EC2VolumeAttachment Properties
type EC2VolumeAttachmentProperties struct {
	Device     interface{} `yaml:"Device"`
	InstanceId interface{} `yaml:"InstanceId"`
	VolumeId   interface{} `yaml:"VolumeId"`
}

// NewEC2VolumeAttachment constructor creates a new EC2VolumeAttachment
func NewEC2VolumeAttachment(properties EC2VolumeAttachmentProperties, deps ...interface{}) EC2VolumeAttachment {
	return EC2VolumeAttachment{
		Type:       "AWS::EC2::VolumeAttachment",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseEC2VolumeAttachment parses EC2VolumeAttachment
func ParseEC2VolumeAttachment(
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
	var resource EC2VolumeAttachment
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
					"Fn::Sub": "${AWS::StackName}-EC2VolumeAttachment-" + name,
				},
			},
		},
	}

	return
}

// ParseEC2VolumeAttachment validator
func (resource EC2VolumeAttachment) Validate() []error {
	return resource.Properties.Validate()
}

// ParseEC2VolumeAttachmentProperties validator
func (resource EC2VolumeAttachmentProperties) Validate() []error {
	errors := []error{}
	return errors
}
