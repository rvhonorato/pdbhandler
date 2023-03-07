// io.go contains the functions for reading and writing PDB files.
package pdbhandler

import (
	"bufio"
	"os"
	"testing"

	"github.com/go-test/deep"
)

// //go:embed test_data/test.pdb
// var testdata embed.FS

func TestReadPDBFile(t *testing.T) {
	// Write a test file
	pdbStr := []string{
		"MODEL        1",
		"ATOM      1  N   THR A   1      17.047  14.099   3.625  1.00 13.79           N  ",
		"MODEL        2",
		"ATOM     61  CA  ARG A  10       8.496   4.609   8.837  1.00  3.38           C  ",
	}

	file, err := os.OpenFile("test.pdb", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		t.Errorf("Error creating test file: %v", err)
	}

	w := bufio.NewWriter(file)
	for _, line := range pdbStr {
		_, _ = w.WriteString(line + "\n")
	}

	w.Flush()

	file.Close()

	defer os.Remove("test.pdb")

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
				f: "test.pdb",
			},
			want: PDB{
				Name: "test.pdb",
				Model: map[int]Model{
					1: {
						Number: 1,
						Chain: map[string]Chain{
							"A": {
								ID: "A",
								Residue: map[int]Residue{
									1: {
										ResNumber: 1,
										ResName:   "THR",
										Atom: map[int]Atom{
											1: {
												AtomNumber: 1,
												AtomName:   "N",
												AltLoc:     " ",
												ResName:    "THR",
												Chain:      "A",
												ResNumber:  1,
												ICode:      " ",
												X:          17.047,
												Y:          14.099,
												Z:          3.625,
												Occup:      1.00,
												Temp:       13.79,
												Element:    " N",
												Charge:     "  ",
											},
										},
									},
								},
							},
						},
					},
					2: {
						Number: 2,
						Chain: map[string]Chain{
							"A": {
								ID: "A",
								Residue: map[int]Residue{
									10: {
										ResNumber: 10,
										ResName:   "ARG",
										Atom: map[int]Atom{
											61: {
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
									},
								},
							},
						},
					},
				},
			},
		},
		{
			name: "TestReadPDBFile-non-existent",
			args: args{
				f: "non-existent.pdb",
			},
			want: PDB{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ReadPDBFile(tt.args.f)
			if diff := deep.Equal(got, tt.want); diff != nil {
				t.Error(diff)
			}

		})
	}
}
