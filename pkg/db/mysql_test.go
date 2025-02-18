package db

import (
	"task/pkg/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewMySQLDB(t *testing.T) {

	testCases := []struct {
		Name          string
		MySQLConfig   config.MySQL
		CheckResponse func(t *testing.T, d *MySQLDB, err error)
	}{
		{
			Name: "OK",
			MySQLConfig: config.MySQL{
				MysqlHost:     "127.0.0.1",
				MysqlPort:     "3306",
				MysqlUser:     "root",
				MysqlPassword: "root",
			},
			CheckResponse: func(t *testing.T, d *MySQLDB, err error) {
				require.NoError(t, err)
				require.NotNil(t, d)
			},
		},
		{
			Name: "Invalid Credentials",
			MySQLConfig: config.MySQL{
				MysqlHost:     "127.0.0.1",
				MysqlPort:     "3306",
				MysqlUser:     "invalid_user",
				MysqlPassword: "wrong_pass",
			},
			CheckResponse: func(t *testing.T, d *MySQLDB, err error) {
				require.Error(t, err)
				require.Nil(t, d)
			},
		},
		{
			Name: "Invalid Host",
			MySQLConfig: config.MySQL{
				MysqlHost:     "invalid_host",
				MysqlPort:     "3306",
				MysqlUser:     "root",
				MysqlPassword: "root",
			},
			CheckResponse: func(t *testing.T, d *MySQLDB, err error) {
				require.Error(t, err)
				require.Nil(t, d)
			},
		},
		{
			Name: "Invalid Port",
			MySQLConfig: config.MySQL{
				MysqlHost:     "127.0.0.1",
				MysqlPort:     "9999", // Invalid port
				MysqlUser:     "root",
				MysqlPassword: "root",
			},
			CheckResponse: func(t *testing.T, d *MySQLDB, err error) {
				require.Error(t, err)
				require.Nil(t, d)
			},
		},
	}

	for i := range testCases {
		t.Run(testCases[i].Name, func(t *testing.T) {
			db, err := NewMySQLDB(&config.Config{
				MySQL: testCases[i].MySQLConfig,
			})
			testCases[i].CheckResponse(t, db, err)
		})
	}
}
