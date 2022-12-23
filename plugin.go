package oracle

import (
	"database/sql"
	"fmt"

	plugin "github.com/jsmzr/boot-plugin"
	_ "github.com/sijms/go-ora/v2"
	"github.com/spf13/viper"
)

const configPrefix = "boot.database"

var defaultConfig map[string]interface{} = map[string]interface{}{
	"enabled": true,
	"order":   5,
}

var globalConn *sql.DB

type OraclePlugin struct{}

func PrePare(statement string) (*sql.Stmt, error) {
	return globalConn.Prepare(statement)
}

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
	globalConn = conn
	return nil
}

func (o *OraclePlugin) Enabled() bool {
	return viper.GetBool(configPrefix + "enabled")
}

func (o *OraclePlugin) Order() int {
	return viper.GetInt(configPrefix + "order")
}

func init() {
	for key := range defaultConfig {
		viper.SetDefault(configPrefix+key, defaultConfig[key])
	}
	plugin.Register("database", &OraclePlugin{})
}
