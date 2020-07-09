package ansible

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	TestPlaybookPath  = "/playbook.yaml"
	TestInventoryPath = "/inventory.ini"
)

var (
	TestExtraVars = map[string]string{
		"k1": "v1",
		"k2": "v2",
	}
	TestBaseCmd = fmt.Sprintf("ansible-playbook -i %s %s", TestInventoryPath, TestPlaybookPath)
)

func TestPlaybook(t *testing.T) {
	pb := Playbook()
	assert.NotNil(t, pb.extraVars)
	assert.Empty(t, pb.inventoryPath)
	assert.Empty(t, pb.playbookPath)
}

func TestPlaybook_ExtraVars(t *testing.T) {
	pb := Playbook()
	vars := map[string]string{
		"test": "test",
	}
	pb.ExtraVars(vars)
	assert.Equal(t, vars, pb.extraVars)

	pb.ExtraVars(map[string]string{"test-2": "test"})
	assert.NotEqual(t, vars, pb.extraVars)
}

func TestPlaybook_ExtraVarsString(t *testing.T) {
	pb := Playbook()
	pb.ExtraVars(TestExtraVars)
	assert.Contains(t, pb.extraVarsToString(), "-e k1=v1")
	assert.Contains(t, pb.extraVarsToString(), "-e k2=v2")
}

func TestPlaybook_Inventory(t *testing.T) {
	pb := Playbook().Inventory(TestInventoryPath)
	assert.Equal(t, TestInventoryPath, pb.inventoryPath)
}

func TestPlaybook_Path(t *testing.T) {
	pb := Playbook().Path(TestPlaybookPath)
	assert.Equal(t, TestPlaybookPath, pb.playbookPath)
}

func TestPlaybook_BaseCmd(t *testing.T) {
	pb := Playbook().
		Inventory(TestInventoryPath).
		Path(TestPlaybookPath)

	assert.Equal(t, TestBaseCmd, pb.baseCmd())
}

func TestPlaybook_Command(t *testing.T) {
	pb := Playbook().
		Inventory(TestInventoryPath).
		Path(TestPlaybookPath).
		ExtraVars(TestExtraVars)

	assert.Contains(t, pb.Command(), TestBaseCmd)
	assert.Contains(t, pb.Command(), "-e k1=v1")
	assert.Contains(t, pb.Command(), "-e k2=v2")
}
