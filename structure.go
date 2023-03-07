// structure.go contains the struct and methods of a PDB file.
package pdbhandler

type PDB struct {
	Name  string
	Model map[int]Model
}

type Model struct {
	Number int
	Chain  map[string]Chain
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
	AtomNumber int
	AtomName   string
	AltLoc     string
	ResName    string
	Chain      string
	ResNumber  int
	ICode      string
	X          float64
	Y          float64
	Z          float64
	Occup      float64
	Temp       float64
	Element    string
	Charge     string
}
