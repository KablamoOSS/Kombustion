package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// TaskDefinitionContainerDefinition Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ecs-taskdefinition-containerdefinitions.html
type TaskDefinitionContainerDefinition struct {
	Cpu                    interface{} `yaml:"Cpu,omitempty"`
	DisableNetworking      interface{} `yaml:"DisableNetworking,omitempty"`
	Essential              interface{} `yaml:"Essential,omitempty"`
	Hostname               interface{} `yaml:"Hostname,omitempty"`
	Image                  interface{} `yaml:"Image,omitempty"`
	Memory                 interface{} `yaml:"Memory,omitempty"`
	MemoryReservation      interface{} `yaml:"MemoryReservation,omitempty"`
	Name                   interface{} `yaml:"Name,omitempty"`
	Privileged             interface{} `yaml:"Privileged,omitempty"`
	ReadonlyRootFilesystem interface{} `yaml:"ReadonlyRootFilesystem,omitempty"`
	User                   interface{} `yaml:"User,omitempty"`
	WorkingDirectory       interface{} `yaml:"WorkingDirectory,omitempty"`
	RepositoryCredentials  interface{} `yaml:"RepositoryCredentials,omitempty"`
	DockerLabels           interface{} `yaml:"DockerLabels,omitempty"`
	LogConfiguration       interface{} `yaml:"LogConfiguration,omitempty"`
	VolumesFrom            interface{} `yaml:"VolumesFrom,omitempty"`
	DockerSecurityOptions  interface{} `yaml:"DockerSecurityOptions,omitempty"`
	EntryPoint             interface{} `yaml:"EntryPoint,omitempty"`
	Environment            interface{} `yaml:"Environment,omitempty"`
	ExtraHosts             interface{} `yaml:"ExtraHosts,omitempty"`
	Command                interface{} `yaml:"Command,omitempty"`
	Links                  interface{} `yaml:"Links,omitempty"`
	Ulimits                interface{} `yaml:"Ulimits,omitempty"`
	DnsServers             interface{} `yaml:"DnsServers,omitempty"`
	MountPoints            interface{} `yaml:"MountPoints,omitempty"`
	PortMappings           interface{} `yaml:"PortMappings,omitempty"`
	DnsSearchDomains       interface{} `yaml:"DnsSearchDomains,omitempty"`
	LinuxParameters        interface{} `yaml:"LinuxParameters,omitempty"`
	HealthCheck            interface{} `yaml:"HealthCheck,omitempty"`
}

// TaskDefinitionContainerDefinition validation
func (resource TaskDefinitionContainerDefinition) Validate() []error {
	errors := []error{}

	return errors
}
