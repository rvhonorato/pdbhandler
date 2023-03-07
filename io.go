// io.go contains the functions for reading and writing PDB files.
package pdbhandler

import (
	"bufio"
	"errors"
	"os"
	"regexp"
	"strconv"
	"strings"

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

// ParseAtomLine parses an ATOM line and returns an Atom struct.
func ParseAtomLine(line string) (Atom, error) {

	// Check if the line matches the ATOM regex
	atomMatch := ATOMREGEXP.FindStringSubmatch(line)
	if len(atomMatch) == 0 {
		// If the line does not match the ATOM regex,
		//  skip to the next line
		err := errors.New("line does not match ATOM regex")
		return Atom{}, err
	}

	// Extract the matching groups
	atomNumber, _ := strconv.Atoi(strings.Trim(atomMatch[1], " "))
	atomName := strings.Trim(atomMatch[2], " ")
	altLoc := atomMatch[3]
	resName := atomMatch[4]
	chainID := atomMatch[5]
	resNumber, _ := strconv.Atoi(strings.Trim(atomMatch[6], " "))
	iCode := atomMatch[7]
	x, _ := strconv.ParseFloat(strings.Trim(atomMatch[8], " "), 64)
	y, _ := strconv.ParseFloat(strings.Trim(atomMatch[9], " "), 64)
	z, _ := strconv.ParseFloat(strings.Trim(atomMatch[10], " "), 64)
	occup, _ := strconv.ParseFloat(strings.Trim(atomMatch[11], " "), 64)
	temp, _ := strconv.ParseFloat(strings.Trim(atomMatch[12], " "), 64)
	element := atomMatch[13]
	charge := atomMatch[14]

	// Populate the atom struct
	atom := Atom{
		AtomNumber: atomNumber,
		AtomName:   atomName,
		AltLoc:     altLoc,
		ResName:    resName,
		Chain:      chainID,
		ResNumber:  resNumber,
		ICode:      iCode,
		X:          x,
		Y:          y,
		Z:          z,
		Occup:      occup,
		Temp:       temp,
		Element:    element,
		Charge:     charge,
	}

	return atom, nil

}

// parseModelLine parses a MODEL line and returns the model number.
func parseModelLine(line string) (int, error) {

	modelMatch := MODELREGEXP.FindStringSubmatch(line)
	if len(modelMatch) == 0 {
		err := errors.New("line does not match MODEL regex")
		return 0, err
	}

	modelNumber, _ := strconv.Atoi(strings.Trim(modelMatch[1], " "))

	return modelNumber, nil
}
