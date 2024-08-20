package defaultDB

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAddTwoUsers(t *testing.T) {
	Database := new(DB)
	Database.Users = make(map[string][2]string)

	err := Database.Add("NewUser", "samplepasswordjwt", "admin")
	require.NoError(t, err)

	err = Database.Add("NewUser2", "samplepasswordjwt", "admin")
	assert.NoError(t, err)
}

func TestAddExistingUser(t *testing.T) {
	Database := new(DB)
	Database.Users = make(map[string][2]string)

	err := Database.Add("NewUser", "samplepasswordjwt", "admin")
	require.NoError(t, err)

	err = Database.Add("NewUser", "samplepasswordjwt", "admin")
	assert.Error(t, err)
}

func TestAddEmptyUser(t *testing.T) {
	Database := new(DB)
	Database.Users = make(map[string][2]string)

	err := Database.Add("", "", "")
	require.NoError(t, err)

	err = Database.Add("", "", "")
	assert.Error(t, err)
}

func TestFindExistingUser(t *testing.T) {
	Database := new(DB)
	Database.Users = make(map[string][2]string)

	err := Database.Add("NewUser", "samplepasswordjwt", "admin")
	require.NoError(t, err)

	group, err := Database.Find("NewUser", "samplepasswordjwt")
	assert.NoError(t, err)
	assert.Equal(t, "admin", group)
}

func TestFindInvalidUsername(t *testing.T) {
	Database := new(DB)
	Database.Users = make(map[string][2]string)

	err := Database.Add("NewUser", "samplepasswordjwt", "admin")
	require.NoError(t, err)

	group, err := Database.Find("NewUser2", "samplepasswordjwt")
	assert.Error(t, err)
	assert.Empty(t, group)
}

func TestFindInvalidPassword(t *testing.T) {
	Database := new(DB)
	Database.Users = make(map[string][2]string)

	err := Database.Add("NewUser", "samplepasswordjwt", "admin")
	require.NoError(t, err)

	group, err := Database.Find("NewUser", "invalidpassword")
	assert.Error(t, err)
	assert.Empty(t, group)
}

func TestFindEmptyInput(t *testing.T) {
	Database := new(DB)
	Database.Users = make(map[string][2]string)

	err := Database.Add("NewUser", "samplepasswordjwt", "admin")
	require.NoError(t, err)

	group, err := Database.Find("", "")
	assert.Error(t, err)
	assert.Empty(t, group)
}
