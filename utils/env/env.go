package env

import "os"

// This retrieve environment variables.
type EnvironmentGetter interface {
	Getenv(key string) string
}

// RealEnvGetter implements EnvGetter using os.Getenv.
type EnvironmentGetterImpl struct{}

func New() EnvironmentGetter {
	return &EnvironmentGetterImpl{}
}
// Getenv retrieves the value of the environment variable named by the key.
func (r *EnvironmentGetterImpl) Getenv(key string) string {
	return os.Getenv(key)
}
