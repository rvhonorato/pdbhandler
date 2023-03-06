// structure.go contains the struct of a PDB file.
package pdbhandler

type PDB struct {
	Model []Model
}

type Model struct {
	Chain []Chain
}

type Chain struct {
	Residue []Residue
}

type Residue struct {
	Atom []Atom
}

type Atom struct {
	AtomName string
	AtomType string
	AtomNum  int
	ResName  string
	ResNum   int
	X        float64
	Y        float64
	Z        float64
}
