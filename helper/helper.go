package helper

// The `type HelperConfig struct` is defining a new struct type called `HelperConfig`.S
type HelperConfig struct {
	Password string
	FilePath string
}

// The NewHelperConfig function creates a new HelperConfig object
func NewHelperConfig(conf *HelperConfig) *HelperConfig {
	return &HelperConfig{
		Password: conf.Password,
		FilePath: conf.FilePath,
	}
}

func DefaultConfig() *HelperConfig {
	return NewHelperConfig(&HelperConfig{})
}
