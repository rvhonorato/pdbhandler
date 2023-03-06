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

	p := PDB{}

	// Check if the file exists
	if !FileExists(f) {
		glog.Infof("file %s does not exist", f)
		return p
	}

	// Check if the file is a PDB file
	if !IsPDBFile(f) {
		glog.Infof("file %s is not a PDB file", f)
		return p
	}

	// read the pdb file
	readFile, err := os.Open(f)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	c := Chain{
		Residue: make(map[int]Residue),
	}
	r := regexp.MustCompile(ATOMREGEX)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		m := r.FindStringSubmatch(line)
		if len(m) == 0 {
			continue
		}
		// Populate the Atom struct
		atomNumber, _ := strconv.Atoi(strings.Trim(m[1], " "))
		atomName := strings.Trim(m[2], " ")
		altLoc := m[3]
		resName := m[4]
		// chainID := m[5]
		resNumber, _ := strconv.Atoi(strings.Trim(m[6], " "))
		iCode := m[7]
		x, _ := strconv.ParseFloat(strings.Trim(m[8], " "), 64)
		y, _ := strconv.ParseFloat(strings.Trim(m[9], " "), 64)
		z, _ := strconv.ParseFloat(strings.Trim(m[10], " "), 64)
		atom := Atom{
			AtomName: atomName,
			AltLoc:   altLoc,
			ICode:    iCode,
			X:        x,
			Y:        y,
			Z:        z,
		}

		_, ok := c.Residue[resNumber]
		if !ok {
			c.Residue[resNumber] = Residue{
				ResNumber: resNumber,
				ResName:   resName,
				Atom:      make(map[int]Atom),
			}
		}

		c.Residue[resNumber].Atom[atomNumber] = atom

	}
	_ = readFile.Close()

	fmt.Println(c)

	// p = PDB{
	// 	Model: make(map[string]Model),
	// }
	// p.Model["1"] = Model{
	// 	Chain: make(map[string]Chain),
	// }
	// p.Model["1"].Chain["A"] = c

	return PDB{}
}
