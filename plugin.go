package oracle

import (
	"database/sql"
	"fmt"

	plugin "github.com/jsmzr/boot-plugin"
	"github.com/jsmzr/boot-plugin-oracle/db"
	_ "github.com/sijms/go-ora/v2"
	"github.com/spf13/viper"
)

const configPrefix = "boot.database."

var defaultConfig map[string]interface{} = map[string]interface{}{
	"enabled": true,
	"order":   5,
}

type OraclePlugin struct{}

func (o *OraclePlugin) Load() error {
	username := viper.GetString(configPrefix + "username")
	password := viper.GetString(configPrefix + "password")
	url := viper.GetString(configPrefix + "url")
	conn, err := sql.Open("oracle", fmt.Sprintf("oracle://%s:%s@%s", username, password, url))
	if err != nil {
		return err
	}
	if err := conn.Ping(); err != nil {
		return err
	}
	db.DB = conn
	return nil
}

func (o *OraclePlugin) Enabled() bool {
	return viper.GetBool(configPrefix + "enabled")
}

func (o *OraclePlugin) Order() int {
	return viper.GetInt(configPrefix + "order")
}

func init() {
	plugin.InitDefaultConfig(configPrefix, defaultConfig)
	plugin.Register("database", &OraclePlugin{})
}
