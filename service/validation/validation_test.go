package validation

import (
	"fmt"
	"github.com/seed95/shortening/build/messages"
	"github.com/seed95/shortening/pkg/derrors"
	"testing"
)

func TestAlias(t *testing.T) {

	setupService(t)

	tests := []struct {
		alias   string
		wantErr error
	}{
		{
			alias:   "lolo",
			wantErr: derrors.New(derrors.Invalid, messages.InvalidAliasLength),
		},
		{
			alias:   "lolo95",
			wantErr: nil,
		},
		{
			alias:   "alias12",
			wantErr: nil,
		},
		{
			alias:   "",
			wantErr: derrors.New(derrors.Invalid, messages.InvalidAliasLength),
		},
	}

	for _, tt := range tests {
		name := fmt.Sprintf("name")
		t.Run(name, func(t *testing.T) {
			if err := serviceTest.Alias(tt.alias); err != tt.wantErr {
				t.Fatalf("got error: %v, want error: %v", err, tt.wantErr)
			}
		})
	}

}
