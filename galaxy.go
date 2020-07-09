package ansible

import (
	"fmt"
)

type galaxy struct {
	requirementsPath string
	rolesPath        string
}

func Galaxy() *galaxy {
	return &galaxy{}
}

func (g *galaxy) Requirements(path string) *galaxy {
	g.requirementsPath = path
	return g
}

func (g *galaxy) RolesPath(path string) *galaxy {
	g.rolesPath = path
	return g
}

func (g *galaxy) baseCmd() string {
	return "ansible-galaxy"
}

func (g *galaxy) Install() error {
	return Exec(g.Command())
}

func (g *galaxy) Command() string {
	return fmt.Sprintf("%s install -r %s -p %s", g.baseCmd(), g.requirementsPath, g.rolesPath)
}
