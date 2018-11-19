package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// WorkspaceWorkspaceProperties Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-workspaces-workspace-workspaceproperties.html
type WorkspaceWorkspaceProperties struct {
	ComputeTypeName                     interface{} `yaml:"ComputeTypeName,omitempty"`
	RootVolumeSizeGib                   interface{} `yaml:"RootVolumeSizeGib,omitempty"`
	RunningMode                         interface{} `yaml:"RunningMode,omitempty"`
	RunningModeAutoStopTimeoutInMinutes interface{} `yaml:"RunningModeAutoStopTimeoutInMinutes,omitempty"`
	UserVolumeSizeGib                   interface{} `yaml:"UserVolumeSizeGib,omitempty"`
}

// WorkspaceWorkspaceProperties validation
func (resource WorkspaceWorkspaceProperties) Validate() []error {
	errors := []error{}

	return errors
}
