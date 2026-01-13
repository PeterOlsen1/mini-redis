package authtypes

import "fmt"

type Rule struct {
	// The regex matching which keys the user can access
	Regex string `yaml:"regex"`

	// Either ALLOW or DENY
	Mode bool `yaml:"mode"`

	// The operation this rule is set on. Either READ or WRITE, not admin
	Operation UserPermission `yaml:"operation"`
}

const ALLOW = true
const DENY = false

type Ruleset []Rule

func (rset Ruleset) Contains(rule Rule) bool {
	for _, r := range rset {
		if r.Regex == rule.Regex && r.Mode == rule.Mode && r.Operation == rule.Operation {
			return true
		}
	}

	return false
}

func (rset *Ruleset) Remove(rule Rule) {
	deref := *rset
	for i, r := range deref {
		if r.Regex == rule.Regex && r.Mode == rule.Mode && r.Operation == rule.Operation {
			deref = append(deref[:i], deref[i+1:]...)
			rset = &deref
			fmt.Println("removing rule")
			fmt.Println(rset)
			return
		}
	}
}

// Combines two rule sets. No identical items will be included
func (rset *Ruleset) Add(other Ruleset) {
	for _, rule := range other {
		if rset.Contains(rule) {
			continue
		}

		deref := append(*rset, rule)
		rset = &deref
	}
}

func (rset *Ruleset) Subtract(other Ruleset) {
	for _, rule := range other {
		rset.Remove(rule)
	}
}
