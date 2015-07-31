package metrics_matcher

import (
	"github.com/onsi/gomega/types"
)

func ContainsMetric(expected interface{}) types.GomegaMatcher {
	return &containsMetric{
		expected: expected,
	}
}

type containsMetric struct {
	expected interface{}
}

func (matcher *containsMetric) Match(actual interface{}) (success bool, err error) {

}

func (matcher *representJSONMatcher) FailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nto contain the metric \n\t%#v", actual, matcher.expected)
}

func (matcher *representJSONMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return fmt.Sprintf("Expected\n\t%#v\nnot to contain the metric\n\t%#v", actual, matcher.expected)
}
