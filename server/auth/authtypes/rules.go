package authtypes

import "iter"

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

func (rset *Ruleset) Contains(rule Rule) bool {
	for _, r := range *rset {
		if r.Regex == rule.Regex && r.Mode == rule.Mode && r.Operation == rule.Operation {
			return true
		}
	}

	return false
}

func (rset *Ruleset) Remove(rule Rule) {
	for i, r := range *rset {
		if r.Regex == rule.Regex && r.Mode == rule.Mode && r.Operation == rule.Operation {
			*rset = append((*rset)[:i], (*rset)[i+1:]...)
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

		*rset = append(*rset, rule)
	}
}

func (rset *Ruleset) Subtract(other Ruleset) {
	for _, rule := range other {
		rset.Remove(rule)
	}
}

func (rset Ruleset) Negatives() iter.Seq2[int, Rule] {
	return func(yield func(int, Rule) bool) {
		for i, rule := range rset {
			if rule.Mode == DENY {
				if !yield(i, rule) {
					return
				}
			}
		}
	}
}

func (rset Ruleset) Positives() iter.Seq2[int, Rule] {
	return func(yield func(int, Rule) bool) {
		for i, rule := range rset {
			if rule.Mode == ALLOW {
				if !yield(i, rule) {
					return
				}
			}
		}
	}
}
