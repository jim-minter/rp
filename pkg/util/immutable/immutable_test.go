package immutable

import (
	"testing"
)

type ts struct {
	Mutable     string `json:"mutable,omitempty" mutable:"true"` // should be able to change
	Case        string `json:"case,omitempty" mutable:"case"`    // should be case insensitive
	Empty       string `json:"empty,omitempty" mutable:""`       // default to immutable
	EmptyNoJSON string `mutable:"false"`                         // handle no json tag
	None        string // default to immutable
}

func TestValidate(t *testing.T) {
	before := ts{
		Mutable:     "before",
		Case:        "before",
		Empty:       "before",
		EmptyNoJSON: "before",
		None:        "before",
	}
	tests := []struct {
		name    string
		modify  func(*ts)
		wantErr string
	}{
		{
			name: "no change",
		},
		{
			name: "can change mutable",
			modify: func(s *ts) {
				s.Mutable = "after"
			},
		},
		{
			name: "can change case caps",
			modify: func(s *ts) {
				s.Case = "BEFORE"
			},
		},
		{
			name: "can NOT change case",
			modify: func(s *ts) {
				s.Case = "after"
			},
			wantErr: "400: PropertyChangeNotAllowed: case: Changing property 'case' is not allowed.",
		},
		{
			name: "can NOT change empty",
			modify: func(s *ts) {
				s.Empty = "after"
			},
			wantErr: "400: PropertyChangeNotAllowed: empty: Changing property 'empty' is not allowed.",
		},
		{
			name: "can NOT change EmptyNoJSON",
			modify: func(s *ts) {
				s.EmptyNoJSON = "after"
			},
			wantErr: "400: PropertyChangeNotAllowed: EmptyNoJSON: Changing property 'EmptyNoJSON' is not allowed.",
		},
		{
			name: "can NOT change None",
			modify: func(s *ts) {
				s.None = "after"
			},
			wantErr: "400: PropertyChangeNotAllowed: None: Changing property 'None' is not allowed.",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			after := before

			if tt.modify != nil {
				tt.modify(&after)
			}

			err := Validate("", &after, &before)
			if err == nil {
				if tt.wantErr != "" {
					t.Error(err)
				}
			} else {
				if err.Error() != tt.wantErr {
					t.Error(err)
				}
			}
		})
	}
}
