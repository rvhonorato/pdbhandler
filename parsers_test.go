package pdbhandler

import (
	"reflect"
	"testing"

	"github.com/go-test/deep"
)

func TestParseAtomLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    Atom
		wantErr bool
	}{
		{
			name: "TestParseAtomLine",
			args: args{
				line: "ATOM     61  CA  ARG A  10       8.496   4.609   8.837  1.00  3.38           C  ",
			},
			want: Atom{
				AtomNumber: 61,
				AtomName:   "CA",
				AltLoc:     " ",
				ResName:    "ARG",
				Chain:      "A",
				ResNumber:  10,
				ICode:      " ",
				X:          8.496,
				Y:          4.609,
				Z:          8.837,
				Occup:      1.00,
				Temp:       3.38,
				Element:    " C",
				Charge:     "  ",
			},
		},
		{
			name: "TestParseAtomLineError",
			args: args{
				line: "ATOM 61  CA  ARG A  10       8.496   4.609   8.837  1.00  3.38           C  ",
			},
			want:    Atom{},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ParseAtomLine(tt.args.line)
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseAtomLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseAtomLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseModelLine(t *testing.T) {
	type args struct {
		line string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "Test_parseModelLine",
			args: args{
				line: "MODEL        1",
			},
			want:    1,
			wantErr: false,
		},
		{
			name: "Test_parseModelLineError",
			args: args{
				line: "MODEL 1",
			},
			want:    0,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseModelLine(tt.args.line)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseModelLine() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("parseModelLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_formatAtomLine(t *testing.T) {
	type args struct {
		atom Atom
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Test_formatAtomLine",
			args: args{
				atom: Atom{
					AtomNumber: 61,
					AtomName:   "CA",
					AltLoc:     " ",
					ResName:    "ARG",
					Chain:      "A",
					ResNumber:  10,
					ICode:      " ",
					X:          8.496,
					Y:          4.609,
					Z:          8.837,
					Occup:      1.00,
					Temp:       3.38,
					Element:    " C",
					Charge:     "  ",
				},
			},
			want: "ATOM     61   CA ARG A  10       8.496   4.609   8.837  1.00  3.38       C  ",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := formatAtomLine(tt.args.atom); got != tt.want {
				t.Errorf("formatAtomLine() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFormatModelLine(t *testing.T) {
	type args struct {
		modelNumber int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestFormatModelLine",
			args: args{
				modelNumber: 1,
			},
			want: "MODEL        1",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FormatModelLine(tt.args.modelNumber); got != tt.want {
				t.Errorf("FormatModelLine() = %v, want %v", got, tt.want)
			}
		})
	}
}
