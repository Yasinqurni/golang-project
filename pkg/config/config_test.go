package config_test

import (
	"golang-project/pkg/config"
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestConfig_NewConfig(t *testing.T) {

	testCases := []struct {
		name           string
		req            map[string]string
		expected       error
		isNegativeTest bool
	}{
		{
			name: "success with env",
			req: map[string]string{
				"PORT_APP":    "8080",
				"DB_HOST":     "localhost",
				"DB_NAME":     "test",
				"DB_USER":     "test",
				"DB_PASSWORD": "test",
				"DB_PORT":     "3306",
				"DB_DRIVER":   "mysql",
			},
			expected:       nil,
			isNegativeTest: false,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {

			for k, v := range testCase.req {
				os.Setenv(k, v)
			}

			cfg, err := config.NewConfig(".")
			require.NoError(t, err)

			var cfgMap = map[string]any{
				"PORT_APP":    cfg.App.Port,
				"DB_HOST":     cfg.DB.Host,
				"DB_NAME":     cfg.DB.Name,
				"DB_USER":     cfg.DB.User,
				"DB_PASSWORD": cfg.DB.Password,
				"DB_PORT":     cfg.DB.Port,
				"DB_DRIVER":   cfg.DB.Driver,
			}

			if testCase.isNegativeTest {
				assert.Error(t, err)
				assert.Equal(t, testCase.expected, err)
			} else {
				for k, v := range testCase.req {
					if val, ok := cfgMap[k].(uint); ok {
						value, err := strconv.ParseUint(v, 10, 64)
						if err != nil {
							log.Fatalf("Error converting string to uint: %v", err)
						}
						assert.Equal(t, uint(value), val)
					} else {
						assert.Equal(t, v, cfgMap[k])
					}
				}
			}
		})
	}
}
