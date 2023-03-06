// io.go contains the functions for reading and writing PDB files.
package pdbhandler

import (
	"reflect"
	"testing"
)

func TestReadPDBFile(t *testing.T) {
	type args struct {
		f string
	}
	tests := []struct {
		name string
		args args
		want PDB
	}{
		{
			name: "TestReadPDBFile",
			args: args{
				f: "1crn.pdb",
			},
			want: PDB{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ReadPDBFile(tt.args.f); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ReadPDBFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
