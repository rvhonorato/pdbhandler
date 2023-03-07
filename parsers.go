package pdbhandler

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

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

// formatAtomLine formats an Atom struct into a PDB ATOM line.
func formatAtomLine(atom Atom) string {

	// Format the atom struct into a PDB ATOM line
	atomLine := fmt.Sprintf("ATOM  %5d %4s%1s%3s %1s%4d%1s   %8.3f%8.3f%8.3f%6.2f%6.2f      %2s%2s",
		atom.AtomNumber,
		atom.AtomName,
		atom.AltLoc,
		atom.ResName,
		atom.Chain,
		atom.ResNumber,
		atom.ICode,
		atom.X,
		atom.Y,
		atom.Z,
		atom.Occup,
		atom.Temp,
		atom.Element,
		atom.Charge,
	)

	return atomLine
}

// FormatModelLine formats a model number into a PDB MODEL line.
func FormatModelLine(modelNumber int) string {

	// Format the model number into a PDB MODEL line
	modelLine := fmt.Sprintf("MODEL     %4d", modelNumber)

	return modelLine
}
