package database

import (
	"testing"

	"github.com/chnmk/sample-authorization-backend/database/defaultDB"
	"github.com/stretchr/testify/require"
)

func TestUseDBsqlite(t *testing.T) {
	require.Panics(t, func() {
		Database := UseDB("sqlite")
		_ = Database
	})
}

func TestUseDBdefault(t *testing.T) {
	Database := UseDB("default")
	require.Implements(t, new(DB), Database)
	require.IsType(t, new(defaultDB.DB), Database)
}

func TestUseDBinvalid(t *testing.T) {
	Database := UseDB("sample_invalid_input")
	require.Implements(t, new(DB), Database)
	require.IsType(t, new(defaultDB.DB), Database)
}
