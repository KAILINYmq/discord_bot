package settings

import (
	"github.com/hedzr/assert"
	"testing"
)

func TestInitConfig(t *testing.T) {
	path := "../../conf/config_dev.toml"

	config := InitConfig(path)

	assert.Equal(t, ":8091", config.Base.Port)
	assert.Equal(t, "", config.Base.Feishu)

	assert.Equal(t, 20, config.Log.MaxSize)
	assert.Equal(t, 10, config.Log.Backups)
}
