package config

type Config interface {
	LibraryRootPath() string
}

type ConfigImpl struct {
	LibraryRoot string
}

var _ Config = &ConfigImpl{}

func NewConfigImpl(libraryRootPath string) *ConfigImpl {
	return &ConfigImpl{
		LibraryRoot: libraryRootPath,
	}
}

func (c *ConfigImpl) LibraryRootPath() string {
	return c.LibraryRoot
}
