package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// BucketLifecycleConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-lifecycleconfig.html
type BucketLifecycleConfiguration struct {
	Rules interface{} `yaml:"Rules"`
}

// BucketLifecycleConfiguration validation
func (resource BucketLifecycleConfiguration) Validate() []error {
	errors := []error{}

	return errors
}
