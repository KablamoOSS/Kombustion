package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ProjectRegistryCredential Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-registrycredential.html
type ProjectRegistryCredential struct {
	Credential         interface{} `yaml:"Credential"`
	CredentialProvider interface{} `yaml:"CredentialProvider"`
}

// ProjectRegistryCredential validation
func (resource ProjectRegistryCredential) Validate() []error {
	errors := []error{}

	return errors
}
