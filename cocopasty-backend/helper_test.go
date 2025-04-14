package main

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetEnvVar(t *testing.T) {
	t.Setenv("TEST_VAR", "my-value")

	result := GetEnvVar("TEST_VAR")

	assert.Equal(t, "my-value", result)
}

func TestGetEnvVarWithDefault(t *testing.T) {

	result := GetEnvVarWithDefault("TEST_VAR_DEFAULT", "my-default")

	assert.Equal(t, "my-default", result)
}

func TestIsEnabled_NoEnvVarDefined(t *testing.T) {
	os.Unsetenv("NOT_DEFINED")

	result := IsEnabled("NOT_DEFINED")
	assert.False(t, result)
}

func TestIsEnabled_EnvVarDefinedAsTrue(t *testing.T) {
	for _, envValue := range []string{"true", "True", "Yes", "yes", "1", "YES", "TRUE", "YeS"} {
		t.Run("", func(tt *testing.T) {
			os.Setenv("TEST_IS_ENABLED", envValue)

			result := IsEnabled("TEST_IS_ENABLED")
			assert.True(t, result)
		})
	}
}

func TestIsEnabled_EnvVarDefinedAsFalse(t *testing.T) {
	for _, envValue := range []string{"false", "", "no", "0", "something"} {
		t.Run("", func(tt *testing.T) {
			os.Setenv("TEST_IS_ENABLED", envValue)

			result := IsEnabled("TEST_IS_ENABLED")
			assert.False(t, result)
		})
	}
}
