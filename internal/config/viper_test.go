package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Test LoadConfig in production mode (reading from OS environment variables)
func TestLoadConfig_Production(t *testing.T) {
	defer os.Clearenv()
	os.Setenv("HTTP_PORT", "9090")
	os.Setenv("DB_HOST", "prod-db")
	os.Setenv("DB_PORT", "5433")
	os.Setenv("DB_USER", "prod_user")
	os.Setenv("DB_PASSWORD", "prod_password")
	os.Setenv("DB_NAME", "prod_db")
	os.Setenv("LOG_LEVEL", "warn")

	viper.Reset()
	LoadConfig("")

	assert.Equal(t, "9090", viper.GetString("HTTP_PORT"))
	assert.Equal(t, "prod-db", viper.GetString("DB_HOST"))
	assert.Equal(t, 5433, viper.GetInt("DB_PORT"))
	assert.Equal(t, "prod_user", viper.GetString("DB_USER"))
	assert.Equal(t, "prod_password", viper.GetString("DB_PASSWORD"))
	assert.Equal(t, "prod_db", viper.GetString("DB_NAME"))
	assert.Equal(t, "warn", viper.GetString("LOG_LEVEL"))
}

// Test LoadConfig in development mode (reading from .env)
func TestLoadConfig_Development(t *testing.T) {
	fileName := ".env.test"
	mockEnv := `HTTP_PORT=8080
				
				DB_HOST=localhost
				DB_PORT=5432
				DB_USER=postgres
				DB_PASSWORD=password
				DB_NAME=userdb
				
				LOG_LEVEL=info
				`
	err := os.WriteFile(fileName, []byte(mockEnv), 0644)
	require.NoError(t, err)
	defer os.Remove(fileName)

	viper.Reset()
	LoadConfig(fileName)

	assert.Equal(t, "8080", viper.GetString("HTTP_PORT"))
	assert.Equal(t, "localhost", viper.GetString("DB_HOST"))
	assert.Equal(t, 5432, viper.GetInt("DB_PORT"))
	assert.Equal(t, "postgres", viper.GetString("DB_USER"))
	assert.Equal(t, "password", viper.GetString("DB_PASSWORD"))
	assert.Equal(t, "userdb", viper.GetString("DB_NAME"))
	assert.Equal(t, "info", viper.GetString("LOG_LEVEL"))
}
