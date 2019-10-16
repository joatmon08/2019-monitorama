package mock

import (
	"io/ioutil"
	"sort"

	yaml "gopkg.in/yaml.v2"
)

const (
	TeamSecurity = "security"
)

type Alert struct {
	Name         string
	Acknowledged bool
	Tag          []string
}

type FalcoRule struct {
	Rule string   `yaml:"rule"`
	Tags []string `yaml:"tags"`
}

func GetAlertsWithinLastDay(team string) []*Alert {
	return []*Alert{
		&Alert{
			Name:         "Unauthorized process",
			Acknowledged: false,
		},
		&Alert{
			Name:         "Root login",
			Acknowledged: true,
		},
		&Alert{
			Name:         "User not authorized",
			Acknowledged: false,
		},
		&Alert{
			Name:         "Open security group",
			Acknowledged: false,
		},
	}
}

func GetRulesByTag(tag string) ([]*FalcoRule, error) {
	var rules []*FalcoRule
	var tagged []*FalcoRule
	data, err := ioutil.ReadFile("mock/falco_rules.yaml")
	if err != nil {
		return nil, err
	}
	err = yaml.Unmarshal([]byte(data), &rules)
	if err != nil {
		return nil, err
	}

	for _, rule := range rules {
		tags := rule.Tags
		sort.Strings(tags)
		i := sort.Search(len(tags), func(i int) bool { return tags[i] == tag })
		if i < len(tags) && tags[i] == tag {
			tagged = append(tagged, rule)
		}
	}

	return tagged, nil
}
