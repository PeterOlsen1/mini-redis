package authtypes

import (
	"iter"
	"slices"
)

type Rule struct {
	// The regex matching which keys the user can access
	Regex string `yaml:"regex"`

	// Either ALLOW or DENY
	Mode bool `yaml:"mode"`

	// The operation this rule is set on. Either READ or WRITE. ADMIN for privileges
	Operation UserPermission `yaml:"operation"`
}

const ALLOW = true
const DENY = false

type Ruleset []Rule

var READALL = Rule{
	Regex:     "@",
	Mode:      ALLOW,
	Operation: READ,
}

var READNONE = Rule{
	Regex:     "@",
	Mode:      DENY,
	Operation: READ,
}

var WRITEALL = Rule{
	Regex:     "@",
	Mode:      ALLOW,
	Operation: WRITE,
}

var WRITENONE = Rule{
	Regex:     "@",
	Mode:      DENY,
	Operation: WRITE,
}

var ADMIN_RULE = Rule{
	Regex:     "",
	Mode:      ALLOW,
	Operation: ADMIN,
}

func (rset *Ruleset) ExtractPerms() int {
	out := 0
	for _, r := range *rset {
		if r == READALL {
			out |= READ
		}
		if r == WRITEALL {
			out |= WRITE
		}
		if r == ADMIN_RULE {
			out |= ADMIN
		}
	}

	return out
}

func (rset *Ruleset) Contains(rule Rule) bool {
	return slices.Contains(*rset, rule)
}

func (rset *Ruleset) Remove(rule Rule) {
	for i, r := range *rset {
		if r == rule {
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

// Playing with go iterators
//
// Not strictly necessary but cool
type RulesetIterator iter.Seq2[int, Rule]

func (rset Ruleset) Iter() RulesetIterator {
	return func(yield func(int, Rule) bool) {
		for i, rule := range rset {
			if !yield(i, rule) {
				return
			}
		}
	}
}

func (rset RulesetIterator) Negatives() RulesetIterator {
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

func (rset RulesetIterator) Positives() RulesetIterator {
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

func (rset RulesetIterator) Read() RulesetIterator {
	return func(yield func(int, Rule) bool) {
		for i, rule := range rset {
			if rule.Operation == READ {
				if !yield(i, rule) {
					return
				}
			}
		}
	}
}

func (rset RulesetIterator) Write() RulesetIterator {
	return func(yield func(int, Rule) bool) {
		for i, rule := range rset {
			if rule.Operation == WRITE {
				if !yield(i, rule) {
					return
				}
			}
		}
	}
}
