package toml

import "github.com/BurntSushi/toml"

type Config struct {
	metaData toml.MetaData `json:"-",toml:"-"`
}

func (c *Config) DecodeFile(fpath string, v interface{}) (err error) {
	c.metaData, err = toml.DecodeFile(fpath, v)
	return
}

func (c Config) IsDefined(key ...string) bool {
	return c.metaData.IsDefined(key...)
}
