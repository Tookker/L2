package main

import (
	"errors"
	"testing"
)

func TestGetTime(t *testing.T) {
	type Args struct {
		addres string
	}

	type Want struct {
		err error
	}

	tests := []struct {
		name    string
		args    Args
		want    Want
		wantErr bool
	}{
		{
			name: "Get time without err",
			args: Args{
				addres: "0.beevik-ntp.pool.ntp.org",
			},
			want: Want{
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "Get time with err",
			args: Args{
				addres: "0beevik-ntp.pool.ntp.org",
			},
			want: Want{
				err: errors.New("no such host"),
			},
			wantErr: true,
		},
	}

	for indx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := GetTime(tt.args.addres)
			if (err != nil) != tt.wantErr {
				t.Errorf("Test # %v. GetTime error = %v. Want error %v", indx, err, tt.wantErr)
				return
			}
		})
	}
}
