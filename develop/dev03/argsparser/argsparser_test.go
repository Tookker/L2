package argsparser

import (
	"errors"
	"testing"
)

func TestSetBoolFlag(t *testing.T) {
	type args struct {
		flag  rune
		flags Flags
	}

	type want struct {
		flags Flags
		err   error
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name:    "set M",
			args:    args{flag: 'M', flags: Flags{}},
			want:    want{flags: Flags{M: true}, err: nil},
			wantErr: false,
		},

		{
			name:    "set u",
			args:    args{flag: 'u', flags: Flags{}},
			want:    want{flags: Flags{U: true}, err: nil},
			wantErr: false,
		},

		{
			name:    "set b",
			args:    args{flag: 'b', flags: Flags{}},
			want:    want{flags: Flags{B: true}, err: nil},
			wantErr: false,
		},

		{
			name:    "set c",
			args:    args{flag: 'c', flags: Flags{}},
			want:    want{flags: Flags{C: true}, err: nil},
			wantErr: false,
		},

		{
			name:    "set h",
			args:    args{flag: 'h', flags: Flags{}},
			want:    want{flags: Flags{H: true}, err: nil},
			wantErr: false,
		},

		{
			name:    "set r",
			args:    args{flag: 'r', flags: Flags{}},
			want:    want{flags: Flags{R: true}, err: nil},
			wantErr: false,
		},

		{
			name:    "set n",
			args:    args{flag: 'n', flags: Flags{}},
			want:    want{flags: Flags{N: true}, err: nil},
			wantErr: false,
		},

		{
			name:    "set undefind flag",
			args:    args{flag: 'q', flags: Flags{}},
			want:    want{flags: Flags{}, err: ErrUnknowFlag},
			wantErr: true,
		},
	}

	for indx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := setBoolFlag(tt.args.flag, tt.args.flags)
			if res != tt.want.flags {
				t.Errorf("Test # %v. SetBoolFlag res = %v. Want res %v", indx, res, tt.want.flags)
				return
			}

			if (err != nil) != tt.wantErr {
				if !errors.Is(err, ErrUnknowFlag) {
					t.Errorf("Test # %v. SetBoolFlag err = %v. Want err %v", indx, err, tt.want.err)
					return
				}
			}
		})
	}
}

func TestSetUintFlag(t *testing.T) {
	type args struct {
		flag  rune
		val   rune
		flags Flags
	}

	type want struct {
		flags Flags
		err   error
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name:    "set 5 to k flag",
			args:    args{flag: 'k', val: '4', flags: Flags{}},
			want:    want{flags: Flags{K: 4}, err: nil},
			wantErr: false,
		},

		{
			name:    "set letter to a flag",
			args:    args{flag: 'k', val: 't', flags: Flags{}},
			want:    want{flags: Flags{}, err: ErrAtioErr},
			wantErr: true,
		},

		{
			name:    "set simbol to a flag",
			args:    args{flag: 'k', val: '.', flags: Flags{}},
			want:    want{flags: Flags{}, err: ErrAtioErr},
			wantErr: true,
		},
	}

	for indx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := setUintFlag(tt.args.flag, tt.args.val, tt.args.flags)
			if res != tt.want.flags {
				t.Errorf("Test # %v. SetUintFlag res = %v. Want res %v", indx, res, tt.want.flags)
				return
			}

			if (err != nil) != tt.wantErr {
				if errors.Is(err, ErrAtioErr) {
					t.Errorf("Test # %v. SetUintFlag err = %v. Want err %v", indx, err, tt.want.err)
				}
				return
			}
		})
	}
}
