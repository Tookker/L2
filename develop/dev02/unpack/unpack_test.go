package unpack_test

import (
	"testing"

	"L2/develop/dev02/unpack"
)

func TestString(t *testing.T) {
	type args struct {
		str string
	}

	type want struct {
		res string
		err error
	}

	tests := []struct {
		name string
		args
		want    want
		wantErr bool
	}{
		{
			name: "String with digit and letters",
			args: args{"a4bc2d5e"},
			want: want{
				res: "aaaabccddddde",
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "String with letters",
			args: args{"abcd"},
			want: want{
				res: "abcd",
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "String with digit",
			args: args{"45"},
			want: want{
				res: "",
				err: unpack.ErrIncorectStr,
			},
			wantErr: true,
		},
		{
			name: "String with escape",
			args: args{"qwe\\4\\5"},
			want: want{
				res: "qwe45",
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "String with escape, digit, letters #1",
			args: args{"qwe\\45"},
			want: want{
				res: "qwe44444",
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "String with escape, digit, letters #2",
			args: args{"qwe\\\\5"},
			want: want{
				res: "qwe\\\\\\\\\\",
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "String with escape, digit, letters #3",
			args: args{"\\5qwe2"},
			want: want{
				res: "5qwee",
				err: nil,
			},
			wantErr: false,
		},
		{
			name: "String with escape, digit, letters #4",
			args: args{"\\53qwe2"},
			want: want{
				res: "555qwee",
				err: nil,
			},
			wantErr: false,
		},

		{
			name: "String with escape, digit, letters #5",
			args: args{"\\53qwe2\\"},
			want: want{
				res: "",
				err: unpack.ErrIncorectStr,
			},
			wantErr: true,
		},
	}

	for indx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := unpack.String(tt.args.str)
			if res != tt.want.res {
				t.Errorf("Test # %v. String res = %v error = %v.\nWant res = %v error = %v", indx+1, res, err, tt.want.res, tt.wantErr)
				return
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Test # %v. String res = %v error = %v.\nWant res = %v error = %v", indx+1, res, err, tt.want.res, tt.wantErr)
				return
			}
		})
	}
}
