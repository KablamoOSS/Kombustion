package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// BucketEncryptionConfiguration Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-encryptionconfiguration.html
type BucketEncryptionConfiguration struct {
	ReplicaKmsKeyID interface{} `yaml:"ReplicaKmsKeyID"`
}

// BucketEncryptionConfiguration validation
func (resource BucketEncryptionConfiguration) Validate() []error {
	errors := []error{}

	return errors
}
