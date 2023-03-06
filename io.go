// io.go contains the functions for reading and writing PDB files.
package pdbhandler

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang/glog"
)

// ReadPDBFile reads a PDB file and returns a PDB object.
func ReadPDBFile(f string) PDB {

	// Check if the file exists
	if !FileExists(f) {
		glog.Error("file %s does not exist", f)
		return PDB{}
	}

	// Check if the file is a PDB file
	if !IsPDBFile(f) {
		glog.Error("file %s is not a PDB file", f)
		return PDB{}
	}

	// TODO: Check if the PDB is valid
	// if !IsValidPDB(f) {
	// 	glog.Infof("file %s is not a valid PDB file", f)
	// 	return PDB{}
	// }

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
		ID:    f,
		Model: make(map[int]Model),
	}

	// Compile the regex
	r := regexp.MustCompile(ATOMREGEX)

	// Loop through the file
	for fileScanner.Scan() {
		line := fileScanner.Text()

		// TODO: Implement something here to find out the model number
		modelNumber := 1
		// ------------------------------

		// Check if the line matches the ATOM regex
		match := r.FindStringSubmatch(line)
		if len(match) == 0 {
			// If the line does not match the ATOM regex,
			//  skip to the next line
			continue
		}

		// Extract the matching groups
		atomNumber, _ := strconv.Atoi(strings.Trim(match[1], " "))
		atomName := strings.Trim(match[2], " ")
		altLoc := match[3]
		resName := match[4]
		chainID := match[5]
		resNumber, _ := strconv.Atoi(strings.Trim(match[6], " "))
		iCode := match[7]
		x, _ := strconv.ParseFloat(strings.Trim(match[8], " "), 64)
		y, _ := strconv.ParseFloat(strings.Trim(match[9], " "), 64)
		z, _ := strconv.ParseFloat(strings.Trim(match[10], " "), 64)

		// Populate the atom struct
		atom := Atom{
			AtomName: atomName,
			AltLoc:   altLoc,
			ICode:    iCode,
			X:        x,
			Y:        y,
			Z:        z,
		}

		// Populate the model struct
		_, ok := p.Model[modelNumber]
		if !ok {
			p.Model[modelNumber] = Model{
				ID:    modelNumber,
				Chain: make(map[string]Chain),
			}
		}

		// Populate the chain struct
		_, ok = p.Model[modelNumber].Chain[chainID]
		if !ok {
			p.Model[modelNumber].Chain[chainID] = Chain{
				ID:      chainID,
				Residue: make(map[int]Residue),
			}
		}

		// Populate the residue struct
		_, ok = p.Model[modelNumber].Chain[chainID].Residue[resNumber]
		if !ok {
			p.Model[modelNumber].Chain[chainID].Residue[resNumber] = Residue{
				ResNumber: resNumber,
				ResName:   resName,
				Atom:      make(map[int]Atom),
			}
		}

		// Populate the atom struct
		p.Model[modelNumber].Chain[chainID].Residue[resNumber].Atom[atomNumber] = atom

	}
	_ = readFile.Close()

	fmt.Println(p.Model[1].Chain["A"].Residue[42])
	// return p

	return PDB{}
}
