package toml

import (
	"testing"

	"github.com/WindomZ/testify/assert"
)

var c TestConfig

type TestConfig struct {
	Config
	Debug bool
	Test  bool
	Log   TestLogConfig
	Http  TestHttpConfig
	Https TestHttpsConfig
}

type TestLogConfig struct {
	Debug    bool
	FileName string `toml:"file"`
}

type TestHttpConfig struct {
	Port int
}

type TestHttpsConfig struct {
	Enable bool
	Port   int
}

func TestConfig_DecodeFile(t *testing.T) {
	if err := c.DecodeFile("./test.toml", &c); err != nil {
		t.Fatal(err)
	}
	// root
	assert.Equal(t, c.Debug, true)
	assert.Equal(t, c.Test, false)
	// root.log
	assert.Equal(t, c.Log.Debug, true)
	assert.NotEmpty(t, c.Log.FileName)
	// root.http
	assert.Equal(t, c.Http.Port, 80)
	// root.https
	assert.Equal(t, c.Https.Enable, true)
	assert.Equal(t, c.Https.Port, 443)
}

func TestConfig_IsDefined(t *testing.T) {
	// root
	assert.Equal(t, c.IsDefined("debug"), true)
	assert.Equal(t, c.IsDefined("test"), true)
	assert.Equal(t, c.IsDefined("test1"), false)
	// root.log
	assert.Equal(t, c.IsDefined("log", "debug"), true)
	assert.Equal(t, c.IsDefined("log", "debug1"), false)
	// root.http
	assert.Equal(t, c.IsDefined("http", "port"), true)
	assert.Equal(t, c.IsDefined("http", "debug"), false)
	// root.https
	assert.Equal(t, c.IsDefined("https", "enable"), true)
	assert.Equal(t, c.IsDefined("https", "debug"), false)
	assert.Equal(t, c.IsDefined("https", "port"), true)
}
