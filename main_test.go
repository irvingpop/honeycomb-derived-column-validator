package main

import (
	"testing"
)

func TestParseDerived(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		wantErr bool
		errMsg  string
	}{
		{
			name:    "empty expression",
			input:   "",
			wantErr: true,
			errMsg:  "empty expression",
		},
		{
			name:    "invalid expression (mismatched parenthesis)",
			input:   "IF(1,2", // missing closing parenthesis
			wantErr: true,
		},
		{
			name:    "valid expression",
			input:   "IF(1,2)",
			wantErr: false,
		},
		// {
		// 	name:    "valid JSON input",
		// 	input:   `{"expression": "IF(EQUALS($http.response.status_code, 200), 1)"}`,
		// 	wantErr: false,
		// },
		// {
		// 	name:    "invalid JSON (missing expression key)",
		// 	input:   `{"wrong_key": "IF(1,2)"}`,
		// 	wantErr: true,
		// },
		// {
		// 	name:    "invalid JSON syntax",
		// 	input:   `{"expression": "IF(1,2)"`, // missing closing brace
		// 	wantErr: true,
		// },
		// {
		// 	name:    "complex valid expression",
		// 	input:   `IF(AND(NOT(EXISTS($trace.parent_id)),EXISTS($duration_ms)),LTE($duration_ms,300))`,
		// 	wantErr: false,
		// },
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			err := ParseDerived(tc.input, false)
			if tc.wantErr {
				if err == nil {
					t.Errorf("expected error for input %q; got nil", tc.input)
				} else if tc.errMsg != "" && err.Error() != tc.errMsg {
					t.Errorf("expected error message %q, got %q", tc.errMsg, err.Error())
				}
			} else {
				if err != nil {
					t.Errorf("expected no error for input %q, got %v", tc.input, err)
				}
			}
		})
	}
}
