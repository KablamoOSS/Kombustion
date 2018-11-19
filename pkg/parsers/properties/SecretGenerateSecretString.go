package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// SecretGenerateSecretString Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-secretsmanager-secret-generatesecretstring.html
type SecretGenerateSecretString struct {
	ExcludeCharacters       interface{} `yaml:"ExcludeCharacters,omitempty"`
	ExcludeLowercase        interface{} `yaml:"ExcludeLowercase,omitempty"`
	ExcludeNumbers          interface{} `yaml:"ExcludeNumbers,omitempty"`
	ExcludePunctuation      interface{} `yaml:"ExcludePunctuation,omitempty"`
	ExcludeUppercase        interface{} `yaml:"ExcludeUppercase,omitempty"`
	GenerateStringKey       interface{} `yaml:"GenerateStringKey,omitempty"`
	IncludeSpace            interface{} `yaml:"IncludeSpace,omitempty"`
	PasswordLength          interface{} `yaml:"PasswordLength,omitempty"`
	RequireEachIncludedType interface{} `yaml:"RequireEachIncludedType,omitempty"`
	SecretStringTemplate    interface{} `yaml:"SecretStringTemplate,omitempty"`
}

// SecretGenerateSecretString validation
func (resource SecretGenerateSecretString) Validate() []error {
	errors := []error{}

	return errors
}
