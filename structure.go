// structure.go contains the struct and methods of a PDB file.
package pdbhandler

type PDB struct {
	ID    string
	Model map[string]Model
}

type Model struct {
	ID    string
	Chain map[string]Chain
}

type Chain struct {
	ID      string
	Residue map[int]Residue
}

type Residue struct {
	ResNumber int
	ResName   string
	Atom      map[int]Atom
}

type Atom struct {
	AtomName string
	AltLoc   string
	ResName  string
	ICode    string
	X        float64
	Y        float64
	Z        float64
}
