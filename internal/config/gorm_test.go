package config

import (
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

// TestInitDB tests if the database initializes without errors
func TestInitDB(t *testing.T) {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "user")
	os.Setenv("DB_PASSWORD", "password")
	os.Setenv("DB_NAME", "mydatabase")

	viper.AutomaticEnv()
	InitDB()

	assert.NotNil(t, DB, "Database instance should not be nil")
	sqlDB, err := DB.DB()
	assert.NoError(t, err, "Should retrieve database instance without errors")
	assert.NoError(t, sqlDB.Ping(), "Database should be reachable")
}
