package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import "fmt"

// BucketSseKmsEncryptedObjects Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-s3-bucket-ssekmsencryptedobjects.html
type BucketSseKmsEncryptedObjects struct {
	Status interface{} `yaml:"Status"`
}

// BucketSseKmsEncryptedObjects validation
func (resource BucketSseKmsEncryptedObjects) Validate() []error {
	errors := []error{}

	return errors
}
