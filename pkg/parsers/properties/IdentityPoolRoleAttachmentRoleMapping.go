package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// IdentityPoolRoleAttachmentRoleMapping Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cognito-identitypoolroleattachment-rolemapping.html
type IdentityPoolRoleAttachmentRoleMapping struct {
	AmbiguousRoleResolution interface{} `yaml:"AmbiguousRoleResolution,omitempty"`
	Type                    interface{} `yaml:"Type"`
	RulesConfiguration      interface{} `yaml:"RulesConfiguration,omitempty"`
}

// IdentityPoolRoleAttachmentRoleMapping validation
func (resource IdentityPoolRoleAttachmentRoleMapping) Validate() []error {
	errors := []error{}

	return errors
}
