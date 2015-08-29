package deploy

import (
	"github.com/voxxit/drone/plugin/condition"
	"github.com/voxxit/drone/shared/build/buildfile"
)

type Bash struct {
	Script  []string `yaml:"script,omitempty"`
	Command string   `yaml:"command,omitempty"`

	Condition *condition.Condition `yaml:"when,omitempty"`
}

func (g *Bash) Write(f *buildfile.Buildfile) {
	g.Script = append(g.Script, g.Command)

	for _, cmd := range g.Script {
		f.WriteCmd(cmd)
	}
}

func (g *Bash) GetCondition() *condition.Condition {
	return g.Condition
}
