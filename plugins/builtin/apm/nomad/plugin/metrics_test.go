package plugin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_parseQuery(t *testing.T) {
	testCases := []struct {
		name        string
		input       string
		expected    *Query
		expectError bool
	}{
		{
			name:  "avg_cpu",
			input: "avg_cpu/group/job",
			expected: &Query{
				Metric:    "nomad.client.allocs.cpu.total_percent",
				Job:       "job",
				Group:     "group",
				Operation: "avg",
			},
			expectError: false,
		},
		{
			name:  "avg_memory",
			input: "avg_memory/group/job",
			expected: &Query{
				Metric:    "nomad.client.allocs.memory.usage",
				Job:       "job",
				Group:     "group",
				Operation: "avg",
			},
			expectError: false,
		},
		{
			name:  "arbritary metric",
			input: "avg_nomad.client.allocs.cpu.total_percent/group/job",
			expected: &Query{
				Metric:    "nomad.client.allocs.cpu.total_percent",
				Job:       "job",
				Group:     "group",
				Operation: "avg",
			},
			expectError: false,
		},
		{
			name:  "job with fwd slashes",
			input: "avg_cpu/group/my/super/job//",
			expected: &Query{
				Metric:    "nomad.client.allocs.cpu.total_percent",
				Job:       "my/super/job//",
				Group:     "group",
				Operation: "avg",
			},
			expectError: false,
		},
		{
			name:        "empty query",
			input:       "",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "invalid query format",
			input:       "invalid",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "missing job",
			input:       "avg_cpu/group",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "invalid op_metric format",
			input:       "invalid/groups/job",
			expected:    nil,
			expectError: true,
		},
		{
			name:        "invalid operation",
			input:       "op_invalid/group/job",
			expected:    nil,
			expectError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := parseQuery(tc.input)

			assert.Equal(t, tc.expected, actual)

			if tc.expectError {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}