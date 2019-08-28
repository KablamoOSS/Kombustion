package properties

// Code generated by go generate; DO NOT EDIT.
// It's generated by "github.com/KablamoOSS/kombustion/generate"

// ListenerAuthenticateOidcConfig Documentation: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticloadbalancingv2-listener-authenticateoidcconfig.html
type ListenerAuthenticateOidcConfig struct {
	AuthorizationEndpoint            interface{} `yaml:"AuthorizationEndpoint"`
	ClientId                         interface{} `yaml:"ClientId"`
	ClientSecret                     interface{} `yaml:"ClientSecret"`
	Issuer                           interface{} `yaml:"Issuer"`
	OnUnauthenticatedRequest         interface{} `yaml:"OnUnauthenticatedRequest,omitempty"`
	Scope                            interface{} `yaml:"Scope,omitempty"`
	SessionCookieName                interface{} `yaml:"SessionCookieName,omitempty"`
	SessionTimeout                   interface{} `yaml:"SessionTimeout,omitempty"`
	TokenEndpoint                    interface{} `yaml:"TokenEndpoint"`
	UserInfoEndpoint                 interface{} `yaml:"UserInfoEndpoint"`
	AuthenticationRequestExtraParams interface{} `yaml:"AuthenticationRequestExtraParams,omitempty"`
}

// ListenerAuthenticateOidcConfig validation
func (resource ListenerAuthenticateOidcConfig) Validate() []error {
	errors := []error{}

	return errors
}