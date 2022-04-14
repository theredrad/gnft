package geth

import (
	"testing"
)

func TestNewGETH(t *testing.T) {
	type args struct {
		rawurl string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "ok",
			args: args{
				rawurl: "https://cloudflare-eth.com/",
			},
			wantErr: false,
		},
		{
			name: "invalid url",
			args: args{
				rawurl: "ht:\\/asd.com",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.rawurl)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr && got == nil {
				t.Errorf("New() got = %v", got)
			}
		})
	}
}
