package lint

// Global is what one would intuitive think of as being the global context of the configuration file.
// That is, given the following configuration...
//
// some_flag = true
// some_string = "the greatest song in the world"
//
// [e_some_lint]
// some_other_flag = false
//
// The fields `some_flag` and `some_string` will be targeted to land into this struct.
type Global struct{}

func (g Global) namespace() string {
	return "Global"
}

// RFC5280Config is the higher scoped configuration which services as the deserialization target for...
//
// [RFC5280Config]
// ...
// ...
type RFC5280Config struct{}

func (r RFC5280Config) namespace() string {
	return "RFC5280Config"
}

// RFC5480Config is the higher scoped configuration which services as the deserialization target for...
//
// [RFC5480Config]
// ...
// ...
type RFC5480Config struct{}

func (r RFC5480Config) namespace() string {
	return "RFC5480Config"
}

// RFC5891Config is the higher scoped configuration which services as the deserialization target for...
//
// [RFC5891Config]
// ...
// ...
type RFC5891Config struct{}

func (r RFC5891Config) namespace() string {
	return "RFC5891Config"
}

// CABFBaselineRequirementsConfig is the higher scoped configuration which services as the deserialization target for...
//
// [CABFBaselineRequirementsConfig]
// ...
// ...
type CABFBaselineRequirementsConfig struct{}

func (c CABFBaselineRequirementsConfig) namespace() string {
	return "CABFBaselineRequirementsConfig"
}

// CABFEVGuidelinesConfig is the higher scoped configuration which services as the deserialization target for...
//
// [CABFEVGuidelinesConfig]
// ...
// ...
type CABFEVGuidelinesConfig struct{}

func (c CABFEVGuidelinesConfig) namespace() string {
	return "CABFEVGuidelinesConfig"
}

// MozillaRootStorePolicyConfig is the higher scoped configuration which services as the deserialization target for...
//
// [MozillaRootStorePolicyConfig]
// ...
// ...
type MozillaRootStorePolicyConfig struct{}

func (m MozillaRootStorePolicyConfig) namespace() string {
	return "MozillaRootStorePolicyConfig"
}

// AppleRootStorePolicyConfig is the higher scoped configuration which services as the deserialization target for...
//
// [AppleRootStorePolicyConfig]
// ...
// ...
type AppleRootStorePolicyConfig struct{}

func (a AppleRootStorePolicyConfig) namespace() string {
	return "AppleRootStorePolicyConfig"
}

// CommunityConfig is the higher scoped configuration which services as the deserialization target for...
//
// [CommunityConfig]
// ...
// ...
type CommunityConfig struct{}

func (c CommunityConfig) namespace() string {
	return "CommunityConfig"
}

// EtsiEsiConfig is the higher scoped configuration which services as the deserialization target for...
//
// [EtsiEsiConfig]
// ...
// ...
type EtsiEsiConfig struct{}

func (e EtsiEsiConfig) namespace() string {
	return "EtsiEsiConfig"
}

type GlobalConfiguration interface {
	namespace() string
}
