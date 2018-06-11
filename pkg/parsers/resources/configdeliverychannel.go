package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// ConfigDeliveryChannel Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-config-deliverychannel.html
type ConfigDeliveryChannel struct {
	Type       string                          `yaml:"Type"`
	Properties ConfigDeliveryChannelProperties `yaml:"Properties"`
	Condition  interface{}                     `yaml:"Condition,omitempty"`
	Metadata   interface{}                     `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                     `yaml:"DependsOn,omitempty"`
}

// ConfigDeliveryChannel Properties
type ConfigDeliveryChannelProperties struct {
	Name                             interface{}                                                 `yaml:"Name,omitempty"`
	S3BucketName                     interface{}                                                 `yaml:"S3BucketName"`
	S3KeyPrefix                      interface{}                                                 `yaml:"S3KeyPrefix,omitempty"`
	SnsTopicARN                      interface{}                                                 `yaml:"SnsTopicARN,omitempty"`
	ConfigSnapshotDeliveryProperties *properties.DeliveryChannelConfigSnapshotDeliveryProperties `yaml:"ConfigSnapshotDeliveryProperties,omitempty"`
}

// NewConfigDeliveryChannel constructor creates a new ConfigDeliveryChannel
func NewConfigDeliveryChannel(properties ConfigDeliveryChannelProperties, deps ...interface{}) ConfigDeliveryChannel {
	return ConfigDeliveryChannel{
		Type:       "AWS::Config::DeliveryChannel",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseConfigDeliveryChannel parses ConfigDeliveryChannel
func ParseConfigDeliveryChannel(ctx map[string]interface{}, name string, data string) (cf types.TemplateObject, err error) {
	var resource ConfigDeliveryChannel
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: ConfigDeliveryChannel - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseConfigDeliveryChannel validator
func (resource ConfigDeliveryChannel) Validate() []error {
	return resource.Properties.Validate()
}

// ParseConfigDeliveryChannelProperties validator
func (resource ConfigDeliveryChannelProperties) Validate() []error {
	errs := []error{}
	if resource.S3BucketName == nil {
		errs = append(errs, fmt.Errorf("Missing required field 'S3BucketName'"))
	}
	return errs
}
