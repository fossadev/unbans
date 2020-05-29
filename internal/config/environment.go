package config

const (
	EnvDevelopment = "development"
	EnvProduction  = "production"
)

func getEnvironment(val string) string {
	switch val {
	case EnvDevelopment, EnvProduction:
		return val
	default:
		return EnvProduction
	}
}
