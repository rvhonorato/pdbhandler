// io.go contains the functions for reading and writing PDB files.
package pdbhandler

import (
	"bufio"
	"fmt"
	"os"

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
	var fileLines []string

	for fileScanner.Scan() {
		line := fileScanner.Text()
		fmt.Println(line)
		// fileLines = append(fileLines, fileScanner.Text())
	}

	readFile.Close()

	for _, line := range fileLines {
		fmt.Println(line)
	}

	return PDB{}
}
