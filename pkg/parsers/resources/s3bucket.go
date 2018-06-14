package resources

// DO NOT EDIT: This file is autogenerated by running 'go generate'
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"github.com/KablamoOSS/kombustion/pkg/parsers/properties"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
	"log"
)

// S3Bucket Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket.html
type S3Bucket struct {
	Type       string             `yaml:"Type"`
	Properties S3BucketProperties `yaml:"Properties"`
	Condition  interface{}        `yaml:"Condition,omitempty"`
	Metadata   interface{}        `yaml:"Metadata,omitempty"`
	DependsOn  interface{}        `yaml:"DependsOn,omitempty"`
}

// S3Bucket Properties
type S3BucketProperties struct {
	AccessControl             interface{}                                 `yaml:"AccessControl,omitempty"`
	BucketName                interface{}                                 `yaml:"BucketName,omitempty"`
	WebsiteConfiguration      *properties.BucketWebsiteConfiguration      `yaml:"WebsiteConfiguration,omitempty"`
	VersioningConfiguration   *properties.BucketVersioningConfiguration   `yaml:"VersioningConfiguration,omitempty"`
	ReplicationConfiguration  *properties.BucketReplicationConfiguration  `yaml:"ReplicationConfiguration,omitempty"`
	NotificationConfiguration *properties.BucketNotificationConfiguration `yaml:"NotificationConfiguration,omitempty"`
	LoggingConfiguration      *properties.BucketLoggingConfiguration      `yaml:"LoggingConfiguration,omitempty"`
	AnalyticsConfigurations   interface{}                                 `yaml:"AnalyticsConfigurations,omitempty"`
	InventoryConfigurations   interface{}                                 `yaml:"InventoryConfigurations,omitempty"`
	MetricsConfigurations     interface{}                                 `yaml:"MetricsConfigurations,omitempty"`
	Tags                      interface{}                                 `yaml:"Tags,omitempty"`
	LifecycleConfiguration    *properties.BucketLifecycleConfiguration    `yaml:"LifecycleConfiguration,omitempty"`
	CorsConfiguration         *properties.BucketCorsConfiguration         `yaml:"CorsConfiguration,omitempty"`
	BucketEncryption          *properties.BucketBucketEncryption          `yaml:"BucketEncryption,omitempty"`
	AccelerateConfiguration   *properties.BucketAccelerateConfiguration   `yaml:"AccelerateConfiguration,omitempty"`
}

// NewS3Bucket constructor creates a new S3Bucket
func NewS3Bucket(properties S3BucketProperties, deps ...interface{}) S3Bucket {
	return S3Bucket{
		Type:       "AWS::S3::Bucket",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseS3Bucket parses S3Bucket
func ParseS3Bucket(name string, data string) (cf types.TemplateObject, err error) {
	var resource S3Bucket
	if err = yaml.Unmarshal([]byte(data), &resource); err != nil {
		return
	}
	if errs := resource.Properties.Validate(); len(errs) > 0 {
		for _, err = range errs {
			log.Println("WARNING: S3Bucket - ", err)
		}
		return
	}
	cf = types.TemplateObject{name: resource}
	return
}

// ParseS3Bucket validator
func (resource S3Bucket) Validate() []error {
	return resource.Properties.Validate()
}

// ParseS3BucketProperties validator
func (resource S3BucketProperties) Validate() []error {
	errs := []error{}
	return errs
}
