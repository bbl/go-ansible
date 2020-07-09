package ansible

import (
	"fmt"
	"strings"
)

type playbook struct {
	inventoryPath string
	playbookPath  string
	extraVars     map[string]string
}

func (p *playbook) Run() error {
	return Exec(p.Command())
}

func (p *playbook) baseCmd() string {
	return fmt.Sprintf("ansible-playbook -i %s %s", p.inventoryPath, p.playbookPath)
}

func Playbook() *playbook {
	a := &playbook{}
	a.extraVars = make(map[string]string)
	return a
}

func (p *playbook) Path(path string) *playbook {
	p.playbookPath = path
	return p
}

func (p *playbook) Inventory(path string) *playbook {
	p.inventoryPath = path
	return p
}

func (p *playbook) ExtraVars(vars map[string]string) *playbook {
	for k, v := range vars {
		p.extraVars[k] = v
	}
	return p
}

func (p *playbook) extraVarsToString() string {
	var params strings.Builder

	if p.extraVars != nil {
		for k, v := range p.extraVars {
			fmt.Fprintf(&params, "-e %s=%s ", k, v)
		}
	}

	return params.String()
}

func (p *playbook) Command() string {
	return fmt.Sprintf("%s %s", p.baseCmd(), p.extraVarsToString())
}
