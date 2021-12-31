package i18n

import (
	"github.com/seed95/shortening/pkg/translate"
	"github.com/seed95/shortening/pkg/translate/i18n/testdata"
	"testing"
)

func TestTranslatorBundle_Translate(t *testing.T) {

	type args struct {
		key      string
		language translate.Language
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "english database error",
			args: args{
				key:      testdata.DBError,
				language: translate.EN,
			},
			want: "error is exist",
		},
		{
			name: "persian invalid role error",
			args: args{
				key:      testdata.InvalidRole,
				language: translate.FA,
			},
			want: "نقش نامعتبر است",
		},
		{
			name: "key not exist",
			args: args{
				key:      "NotExist",
				language: translate.FA,
			},
			want: "NotExist",
		},
		{
			name: "empty key",
			args: args{
				key:      "",
				language: translate.EN,
			},
			want: "",
		},
	}

	translator, err := New("testdata")
	if err != nil {
		t.Fatalf("New error: %v", err)
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := translator.Translate(tt.args.key, tt.args.language)
			if got != tt.want {
				t.Fatalf("Translate() got: %v, want: %v", got, tt.want)
			}
		})
	}
}
