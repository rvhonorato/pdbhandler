// io.go contains the functions for reading and writing PDB files.
package pdbhandler

import (
	"bufio"
	"os"
	"regexp"

	"github.com/golang/glog"
)

var (
	ATOMREGEXP  = regexp.MustCompile(ATOM_RECORD_REGEX)
	MODELREGEXP = regexp.MustCompile(MODEL_RECORD_REGEX)
)

// ReadPDBFile reads a PDB file and returns a PDB object.
func ReadPDBFile(f string) PDB {

	// read the pdb file
	readFile, err := os.Open(f)
	if err != nil {
		glog.Error(err)
		return PDB{}
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// Initialize the PDB struct
	p := PDB{
		Name:  f,
		Model: make(map[int]Model),
	}

	// Compile the regex
	// atomR := regexp.MustCompile(ATOMREGEX)
	// modelR := regexp.MustCompile(MODELREGEX)

	// Loop through the file
	currentModel := 1
	for fileScanner.Scan() {
		line := fileScanner.Text()

		modelNb, err := parseModelLine(line)
		if err == nil {
			currentModel = modelNb
		}

		// ------------------------------
		atom, err := ParseAtomLine(line)
		if err != nil {
			continue
		}

		// Populate the model struct
		_, ok := p.Model[currentModel]
		if !ok {
			p.Model[currentModel] = Model{
				Number: currentModel,
				Chain:  make(map[string]Chain),
			}
		}

		// Populate the chain struct
		_, ok = p.Model[currentModel].Chain[atom.Chain]
		if !ok {
			p.Model[currentModel].Chain[atom.Chain] = Chain{
				ID:      atom.Chain,
				Residue: make(map[int]Residue),
			}
		}

		// Populate the residue struct
		_, ok = p.Model[currentModel].Chain[atom.Chain].Residue[atom.ResNumber]
		if !ok {
			p.Model[currentModel].Chain[atom.Chain].Residue[atom.ResNumber] = Residue{
				ResNumber: atom.ResNumber,
				ResName:   atom.ResName,
				Atom:      make(map[int]Atom),
			}
		}

		// Populate the atom struct
		p.Model[currentModel].Chain[atom.Chain].Residue[atom.ResNumber].Atom[atom.AtomNumber] = atom

	}
	_ = readFile.Close()

	return p

}

// // WritePDBFile writes a PDB object to a file.
// func WritePDBFile(p PDB) error {

// 	return nil

// }
