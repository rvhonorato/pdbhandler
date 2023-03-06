// utils.go contains the functions that are used by multiple files.
package pdbhandler

import (
	"os"
	"path/filepath"
)

// FileExists checks if a file exists.
func FileExists(f string) bool {
	_, error := os.Stat(f)
	return !os.IsNotExist(error)
}

// IsPDBFile checks if a file is a PDB file.
func IsPDBFile(f string) bool {
	ext := filepath.Ext(f)
	return ext == ".pdb"
}
