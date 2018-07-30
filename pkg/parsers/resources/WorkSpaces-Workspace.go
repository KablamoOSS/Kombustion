package resources

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

import (
	"fmt"
	"github.com/KablamoOSS/kombustion/types"
	yaml "github.com/KablamoOSS/yaml"
)

// WorkSpacesWorkspace Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-resource-workspaces-workspace.html
type WorkSpacesWorkspace struct {
	Type       string                        `yaml:"Type"`
	Properties WorkSpacesWorkspaceProperties `yaml:"Properties"`
	Condition  interface{}                   `yaml:"Condition,omitempty"`
	Metadata   interface{}                   `yaml:"Metadata,omitempty"`
	DependsOn  interface{}                   `yaml:"DependsOn,omitempty"`
}

// WorkSpacesWorkspace Properties
type WorkSpacesWorkspaceProperties struct {
	BundleId                    interface{} `yaml:"BundleId"`
	DirectoryId                 interface{} `yaml:"DirectoryId"`
	RootVolumeEncryptionEnabled interface{} `yaml:"RootVolumeEncryptionEnabled,omitempty"`
	UserName                    interface{} `yaml:"UserName"`
	UserVolumeEncryptionEnabled interface{} `yaml:"UserVolumeEncryptionEnabled,omitempty"`
	VolumeEncryptionKey         interface{} `yaml:"VolumeEncryptionKey,omitempty"`
}

// NewWorkSpacesWorkspace constructor creates a new WorkSpacesWorkspace
func NewWorkSpacesWorkspace(properties WorkSpacesWorkspaceProperties, deps ...interface{}) WorkSpacesWorkspace {
	return WorkSpacesWorkspace{
		Type:       "AWS::WorkSpaces::Workspace",
		Properties: properties,
		DependsOn:  deps,
	}
}

// ParseWorkSpacesWorkspace parses WorkSpacesWorkspace
func ParseWorkSpacesWorkspace(
	name string,
	data string,
) (
	source string,
	conditions types.TemplateObject,
	metadata types.TemplateObject,
	mappings types.TemplateObject,
	outputs types.TemplateObject,
	parameters types.TemplateObject,
	resources types.TemplateObject,
	transform types.TemplateObject,
	errors []error,
) {
	source = "kombustion-core-resources"
	var resource WorkSpacesWorkspace
	err := yaml.Unmarshal([]byte(data), &resource)

	if err != nil {
		errors = append(errors, err)
		return
	}

	if validateErrs := resource.Properties.Validate(); len(errors) > 0 {
		errors = append(errors, validateErrs...)
		return
	}

	resources = types.TemplateObject{name: resource}

	return
}

// ParseWorkSpacesWorkspace validator
func (resource WorkSpacesWorkspace) Validate() []error {
	return resource.Properties.Validate()
}

// ParseWorkSpacesWorkspaceProperties validator
func (resource WorkSpacesWorkspaceProperties) Validate() []error {
	errors := []error{}
	if resource.BundleId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'BundleId'"))
	}
	if resource.DirectoryId == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'DirectoryId'"))
	}
	if resource.UserName == nil {
		errors = append(errors, fmt.Errorf("Missing required field 'UserName'"))
	}
	return errors
}