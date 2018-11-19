package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// TopicRuleS3Action Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-iot-topicrule-s3action.html
type TopicRuleS3Action struct {
	BucketName interface{} `yaml:"BucketName"`
	Key        interface{} `yaml:"Key"`
	RoleArn    interface{} `yaml:"RoleArn"`
}

// TopicRuleS3Action validation
func (resource TopicRuleS3Action) Validate() []error {
	errors := []error{}

	return errors
}
