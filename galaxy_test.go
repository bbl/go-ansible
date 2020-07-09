package ansible

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TestGalaxyBaseCmd    = "ansible-galaxy"
	TestRequirementsPath = "/requirements.yaml"
	TestRolesPath        = "/roles"
)

var TestGalaxyCmd = fmt.Sprintf("%s install -r %s -p %s", TestGalaxyBaseCmd, TestRequirementsPath, TestRolesPath)

func TestGalaxy(t *testing.T) {
	g := Galaxy()
	assert.Empty(t, g.requirementsPath)
	assert.Empty(t, g.rolesPath)
}

func TestGalaxy_Requirements(t *testing.T) {
	g := Galaxy().Requirements(TestRequirementsPath)
	assert.Equal(t, TestRequirementsPath, g.requirementsPath)
}

func TestGalaxy_RolesPath(t *testing.T) {
	g := Galaxy().RolesPath(TestRolesPath)
	assert.Equal(t, TestRolesPath, g.rolesPath)
}

func TestGalaxy_Command(t *testing.T) {
	g := Galaxy().
		RolesPath(TestRolesPath).
		Requirements(TestRequirementsPath)

	assert.Equal(t, g.Command(), TestGalaxyCmd)
}
