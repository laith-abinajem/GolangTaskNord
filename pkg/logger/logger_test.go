// Developer: Saif Hamdan
// Date: 18/7/2023

package logger

import (
	"task/pkg/config"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLogger(t *testing.T) {
	logger, err := NewLogger(&config.Config{
		Logger: config.Logger{},
	})
	require.NoError(t, err)
	require.NotNil(t, logger)
}
