package pdbhandler

import (
	"reflect"
	"testing"
)

func TestPDB_ListChains(t *testing.T) {
	type fields struct {
		Name  string
		Model map[int]Model
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "TestPDB_ListChains",
			fields: fields{
				Name: "1abc",
				Model: map[int]Model{
					1: {
						Chain: map[string]Chain{
							"A": {
								ID: "A",
							},
							"B": {
								ID: "B",
							},
						},
					},
				},
			},
			want: []string{"A", "B"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PDB{
				Name:  tt.fields.Name,
				Model: tt.fields.Model,
			}
			if got := p.ListChains(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PDB.ListChains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPDB_ListResidues(t *testing.T) {
	type fields struct {
		Name  string
		Model map[int]Model
	}
	tests := []struct {
		name   string
		fields fields
		want   []string
	}{
		{
			name: "TestPDB_ListResidues",
			fields: fields{
				Name: "1abc",
				Model: map[int]Model{
					1: {
						Chain: map[string]Chain{
							"A": {
								ID: "A",
								Residue: map[int]Residue{
									1: {
										ResName: "ALA",
									},
									2: {
										ResName: "GLY",
									},
								},
							},
							"B": {
								ID: "B",
								Residue: map[int]Residue{
									1: {
										ResName: "ALA",
									},
									2: {
										ResName: "ARG",
									},
								},
							},
						},
					},
				},
			},
			want: []string{"ALA", "ALA", "ARG", "GLY"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PDB{
				Name:  tt.fields.Name,
				Model: tt.fields.Model,
			}
			if got := p.ListResidues(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PDB.ListResidues() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPDB_CountAtoms(t *testing.T) {
	type fields struct {
		Name  string
		Model map[int]Model
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "TestPDB_CountAtoms",
			fields: fields{
				Name: "1abc",
				Model: map[int]Model{
					1: {
						Chain: map[string]Chain{
							"A": {
								ID: "A",
								Residue: map[int]Residue{
									1: {
										ResName: "ALA",
										Atom: map[int]Atom{
											1: {},
											2: {},
										},
									},
									2: {
										ResName: "GLY",
										Atom: map[int]Atom{
											1: {},
											2: {},
										},
									},
								},
							},
							"B": {
								ID: "B",
								Residue: map[int]Residue{
									1: {
										ResName: "ALA",
										Atom: map[int]Atom{
											1: {},
										},
									},
									2: {
										ResName: "ARG",
										Atom: map[int]Atom{
											1: {},
											2: {},
											3: {},
										},
									},
								},
							},
						},
					},
				},
			},
			want: 8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PDB{
				Name:  tt.fields.Name,
				Model: tt.fields.Model,
			}
			if got := p.CountAtoms(); got != tt.want {
				t.Errorf("PDB.CountAtoms() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPDB_CountResidues(t *testing.T) {
	type fields struct {
		Name  string
		Model map[int]Model
	}
	tests := []struct {
		name   string
		fields fields
		want   int
	}{
		{
			name: "TestPDB_CountResidues",
			fields: fields{
				Name: "1abc",
				Model: map[int]Model{
					1: {
						Chain: map[string]Chain{
							"A": {
								ID: "A",
								Residue: map[int]Residue{
									1: {
										ResName: "ALA",
									},
									2: {
										ResName: "GLY",
									},
								},
							},
							"B": {
								ID: "B",
								Residue: map[int]Residue{
									1: {
										ResName: "ALA",
									},
									2: {
										ResName: "ARG",
									},
								},
							},
						},
					},
				},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &PDB{
				Name:  tt.fields.Name,
				Model: tt.fields.Model,
			}
			if got := p.CountResidues(); got != tt.want {
				t.Errorf("PDB.CountResidues() = %v, want %v", got, tt.want)
			}
		})
	}
}
