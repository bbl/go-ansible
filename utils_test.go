package ansible

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExec(t *testing.T) {
	// todo it's actually an integration test
	err := Exec("echo test")
	assert.NoError(t, err)
}
