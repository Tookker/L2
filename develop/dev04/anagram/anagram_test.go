package anagram

import (
	"reflect"
	"testing"
)

func TestIsRussianWord(t *testing.T) {
	type args struct {
		word string
	}

	type want struct {
		err error
	}

	tests := []struct {
		name string
		args
		want
		wantErr bool
	}{
		{
			name:    "enter Russian word",
			args:    args{"Привет"},
			want:    want{nil},
			wantErr: false,
		},
		{
			name:    "enter not Russian word",
			args:    args{"Hello"},
			want:    want{ErrLetterIsNotRussian},
			wantErr: true,
		},

		{
			name:    "enter word with digit",
			args:    args{"Привет1"},
			want:    want{ErrDigitInWord},
			wantErr: true,
		},
	}

	for indx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := isRussianWord(tt.word)
			if (err != nil) != tt.wantErr {
				t.Errorf("Test # %v. Error = %v.\nWant Error = %v", indx+1, err, tt.wantErr)
				return
			}
		})
	}
}

func TestGetAnagramm(t *testing.T) {
	type args struct {
		word string
	}

	type want struct {
		res string
	}

	tests := []struct {
		name string
		args
		want
	}{
		{
			name: "get anagram",
			args: args{"гдабвежиз"},
			want: want{"абвгдежзи"},
		},
	}

	for indx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := getAnagram(tt.word)
			if tt.res != getAnagram(tt.word) {
				t.Errorf("Test # %v. Res = %v.\nWant res = %v", indx+1, res, tt.res)
				return
			}
		})
	}
}

func TestFindAnagramKey(t *testing.T) {
	type args struct {
		word       string
		anagramMap map[string][]string
	}

	type want struct {
		val string
		ok  bool
	}
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "find anagram",
			args: args{word: "пятка", anagramMap: map[string][]string{"пятка": {"тяпка"}}},
			want: want{val: "пятка", ok: true},
		},

		{
			name: "not find anagram",
			args: args{word: "ватка", anagramMap: map[string][]string{"пятка": {"тяпка"}}},
			want: want{val: "", ok: false},
		},
	}

	for indx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, ok := findAnagramKey(tt.args.word, tt.args.anagramMap)
			if ok != tt.want.ok || tt.want.val != res {
				t.Errorf("Test # %v. Res = %v Ok = %v.\nWant Res = %v Ok = %v", indx+1, res, ok, tt.want.val, tt.want.ok)
				return
			}
		})
	}
}

func TestMakeRes(t *testing.T) {
	type args struct {
		arg map[string][]string
	}

	type want struct {
		res map[string][]string
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "make res without only key anagramm",
			args: args{map[string][]string{"пятка": {"тяпка", "пятак"}, "листок": {"слиток", "столик"}}},
			want: want{map[string][]string{"пятка": {"тяпка", "пятак"}, "листок": {"слиток", "столик"}}},
		},

		{
			name: "make res with only key anagramm",
			args: args{map[string][]string{"пятка": {"тяпка", "пятак"}, "листок": {"слиток", "столик"}, "мышь": {}}},
			want: want{map[string][]string{"пятка": {"тяпка", "пятак"}, "листок": {"слиток", "столик"}}},
		},
	}

	for indx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := makeRes(tt.args.arg)
			if !reflect.DeepEqual(res, tt.want.res) {
				t.Errorf("Test # %v. Res = %v .\nWant Res = %v ", indx+1, res, tt.want.res)
				return
			}
		})
	}
}

func TestContainsWord(t *testing.T) {
	type args struct {
		word  string
		words []string
	}

	type want struct {
		ok bool
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "word contains",
			args: args{word: "привет", words: []string{"привет", "пока"}},
			want: want{ok: true},
		},

		{
			name: "word contains",
			args: args{word: "досвидания", words: []string{"привет", "пока"}},
			want: want{ok: false},
		},
	}

	for indx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res := containsWord(tt.args.word, tt.args.words)
			if res != tt.want.ok {
				t.Errorf("Test # %v. Res = %v .\nWant Res = %v ", indx+1, res, tt.want.ok)
				return
			}
		})
	}
}

func TestMake(t *testing.T) {
	type args struct {
		words []string
	}

	type want struct {
		res map[string][]string
		err error
	}

	tests := []struct {
		name    string
		args    args
		want    want
		wantErr bool
	}{
		{
			name:    "correct make anagramm",
			args:    args{[]string{"пятка", "пятак", "тяпка", "тяпка", "листок", "слиток", "столик"}},
			want:    want{map[string][]string{"листок": {"слиток", "столик"}, "пятка": {"пятак", "тяпка"}}, nil},
			wantErr: false,
		},

		{
			name:    "input with digit word",
			args:    args{[]string{"пятак1", "тяпка", "пятка", "тяпка", "листок", "слиток", "столик"}},
			want:    want{nil, ErrDigitInWord},
			wantErr: true,
		},

		{
			name:    "input with no russian word",
			args:    args{[]string{"пятак", "asdf", "пятка", "тяпка", "листок", "слиток", "столик"}},
			want:    want{nil, ErrLetterIsNotRussian},
			wantErr: true,
		},

		{
			name:    "correct make anagramm with signle anagramm key",
			args:    args{[]string{"пятка", "пятак", "тяпка", "тяпка", "листок", "слиток", "столик", "привет"}},
			want:    want{map[string][]string{"листок": {"слиток", "столик"}, "пятка": {"пятак", "тяпка"}}, nil},
			wantErr: false,
		},
	}

	for indx, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, err := Make(tt.args.words)
			if (err != nil) != tt.wantErr {
				t.Errorf("Test # %v. Err = %v .\nWant Err = %v ", indx+1, err, tt.want.err)
				return
			}

			if !reflect.DeepEqual(res, tt.want.res) {
				t.Errorf("Test # %v. Res = %v .\nWant Res = %v ", indx+1, res, tt.want.res)
				return
			}
		})
	}
}
