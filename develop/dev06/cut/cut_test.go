package cut

import (
	"errors"
	"reflect"
	"testing"
)

func TestReadArgF(t *testing.T) {
	type args struct {
		str string
	}

	type want struct {
		err error
		cut MyCut
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name:    "parse with , correct",
			args:    args{str: "1,2,3,4,5,6,7"},
			want:    want{cut: MyCut{showStr: []uint{1, 2, 3, 4, 5, 6, 7}}, err: nil},
			wantErr: false,
		},
		{
			name:    "parse with , uncorrect has negative val",
			args:    args{str: "1,2,3,4,5,6,-7"},
			want:    want{cut: MyCut{}, err: ErrParseArgFNegative},
			wantErr: true,
		},
		{
			name:    "parse with , uncorrect has letter",
			args:    args{str: "1,2,3,4,5,6,a"},
			want:    want{cut: MyCut{}, err: ErrParseArgF},
			wantErr: true,
		},
		{
			name:    "parse with - correct",
			args:    args{str: "-2"},
			want:    want{cut: MyCut{skipStr: 2}, err: nil},
			wantErr: false,
		},
		{
			name:    "parse with - negative num",
			args:    args{str: "--2"},
			want:    want{cut: MyCut{}, err: ErrParseArgFNegative},
			wantErr: true,
		},
		{
			name:    "parse with - letter",
			args:    args{str: "-b"},
			want:    want{cut: MyCut{}, err: ErrParseArgF},
			wantErr: true,
		},
		{
			name:    "parse with - letter",
			args:    args{str: "-b"},
			want:    want{cut: MyCut{}, err: ErrParseArgF},
			wantErr: true,
		},

		{
			name:    "parse with from - to correct",
			args:    args{str: "1-6"},
			want:    want{cut: MyCut{from: 1, to: 6}, err: nil},
			wantErr: false,
		},

		{
			name:    "parse with from - to negative val",
			args:    args{str: "-1-6"},
			want:    want{cut: MyCut{}, err: ErrParseArgF},
			wantErr: true,
		},

		{
			name:    "parse with from - to negative val",
			args:    args{str: "-1--6"},
			want:    want{cut: MyCut{}, err: ErrParseArgF},
			wantErr: true,
		},

		{
			name:    "parse with from - to letter",
			args:    args{str: "-a-6"},
			want:    want{cut: MyCut{}, err: ErrParseArgF},
			wantErr: true,
		},

		{
			name:    "parse with only - ",
			args:    args{str: "-"},
			want:    want{cut: MyCut{}, err: ErrParseArgF},
			wantErr: true,
		},

		{
			name:    "parse with only 1",
			args:    args{str: "1"},
			want:    want{cut: MyCut{showStr: []uint{1}}, err: nil},
			wantErr: false,
		},
	}

	for indx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var cut MyCut
			err := cut.readArgF(tt.args.str)
			if (err != nil) != tt.wantErr {
				t.Errorf("Test # %v, error = %v, want = %v", indx+1, err, tt.want.err)
				return
			}

			if tt.wantErr {
				if !errors.Is(err, tt.want.err) {
					t.Errorf("Test # %v, error = %v, want = %v", indx+1, err, tt.want.err)
					return
				}
			} else {
				if !reflect.DeepEqual(tt.want.cut, cut) {
					t.Errorf("Test # %v, cut = %v, want cut = %v", indx+1, cut, tt.want.cut)
					return
				}
			}
		})
	}
}

func TestGetWorkColumn(t *testing.T) {
	tests := []struct {
		cut  MyCut
		name string
		want [][]string
	}{
		{
			cut:  MyCut{str: "kek mek cheburek\nkek lol arbidol"},
			name: "cut default",
			want: [][]string{{"kek", "mek", "cheburek"}, {"kek", "lol", "arbidol"}},
		},

		{
			cut:  MyCut{str: "kek mek cheburek\nkek lol arbidol", skipStr: 2},
			name: "cut with escape second colunn",
			want: [][]string{{"kek", "cheburek"}, {"kek", "arbidol"}},
		},

		{
			cut:  MyCut{str: "kek mek cheburek\nkek lol arbidol", showStr: []uint{2, 3}},
			name: "cut diapason 2-3",
			want: [][]string{{"mek", "cheburek"}, {"lol", "arbidol"}},
		},

		{
			cut:  MyCut{str: "kek mek cheburek\nkek lol arbidol", from: 1, to: 4},
			name: "cut from, to",
			want: [][]string{{"kek", "mek", "cheburek"}, {"kek", "lol", "arbidol"}},
		},

		{
			cut:  MyCut{str: "kek mek cheburek\nkek lol arbidol", showStr: []uint{2}},
			name: "cut show only 1 column",
			want: [][]string{{"mek"}, {"lol"}},
		},
	}

	for index, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.cut.getWorkColumn()
			if !reflect.DeepEqual(res, tt.want) {
				t.Errorf("Test #%v res = %v, want %v", index+1, res, tt.want)
				return
			}
		})
	}
}

func TestDelimite(t *testing.T) {
	type args struct {
		cut MyCut
		msg [][]string
	}

	type want struct {
		res [][]string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "delemite with a",
			args: args{cut: MyCut{d: "a"}, msg: [][]string{{"arm", "var", "dar"}, {"mar", "dar", "csar"}}},
			want: want{res: [][]string{{"rm", "vr", "dr"}, {"mr", "dr", "csr"}}},
		},

		{
			name: "delemite with a and set flag s",
			args: args{cut: MyCut{d: "a", s: true}, msg: [][]string{{"arm", "vir", "dar"}, {"mar", "dar", "csar"}}},
			want: want{res: [][]string{{"rm", "dr"}, {"mr", "dr", "csr"}}},
		},
	}

	for index, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := tt.args.cut.delemite(tt.args.msg)
			if !reflect.DeepEqual(tt.want.res, res) {
				t.Errorf("Test #%v, res = %v, want = %v", index+1, res, tt.want.res)
				return
			}
		})
	}
}

func TestGetRes(t *testing.T) {
	type args struct {
		strs [][]string
	}

	type want struct {
		res string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "get res",
			args: args{[][]string{{"rm", "vr", "dr"}, {"mr", "dr", "csr"}}},
			want: want{"rm vr dr\nmr dr csr\n"},
		},
	}

	for index, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var cut MyCut
			res := cut.getRes(tt.args.strs)
			if res != tt.want.res {
				t.Errorf("Test #%v, res = %v, want = %v", index+1, res, tt.want.res)
				return
			}
		})
	}
}
