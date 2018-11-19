package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ProjectLogsConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-codebuild-project-logsconfig.html
type ProjectLogsConfig struct {
	S3Logs         interface{} `yaml:"S3Logs,omitempty"`
	CloudWatchLogs interface{} `yaml:"CloudWatchLogs,omitempty"`
}

// ProjectLogsConfig validation
func (resource ProjectLogsConfig) Validate() []error {
	errors := []error{}

	return errors
}
