package auth

import (
	"net/http"
	"strings"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	tests := map[string]struct {
		key       string
		value     string
		expect    string
		expectErr error
	}{
		"no header": {
			expectErr: ErrNoAuthHeaderIncluded,
		},
		"empty header": {
			key:       "Authorization",
			expectErr: ErrNoAuthHeaderIncluded,
		},
		"wrong value": {
			key:       "Authorization",
			value:     "-",
			expectErr: ErrMalformedAuthHeader,
		},
		"wrong prefix": {
			key:       "Authorization",
			value:     "Bearer xxxxx",
			expectErr: ErrMalformedAuthHeader,
		},
		"correct": {
			key:    "Authorization",
			value:  "ApiKey xxxxx",
			expect: "xxxxx",
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			header := http.Header{}
			header.Add(test.key, test.value)

			output, err := GetAPIKey(header)
			if err != nil {
				if test.expectErr != nil && strings.Contains(err.Error(), test.expectErr.Error()) {
					return
				}
				t.Errorf("Unexpected: TestGetAPIKey:%v\n", err)
				return
			}

			if output != test.expect {
				t.Errorf("Unexpected: TestGetAPIKey:%s", output)
				return
			}
		})
	}
}
