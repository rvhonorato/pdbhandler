package pdbhandler

import "sort"

// ListChains returns a list of chains in the PDB
//
// The list is sorted alphabetically.
func (p *PDB) ListChains() []string {
	var chains []string
	for _, m := range p.Model {
		for _, c := range m.Chain {
			chains = append(chains, c.ID)
		}
	}
	sort.Strings(chains)
	return chains
}

// ListResidues returns a list of residues in the PDB
//
// The list is sorted alphabetically.
func (p *PDB) ListResidues() []string {
	var residues []string
	for _, m := range p.Model {
		for _, c := range m.Chain {
			for _, r := range c.Residue {
				residues = append(residues, r.ResName)
			}
		}
	}
	sort.Strings(residues)
	return residues
}

// CountAtoms returns the number of atoms in the PDB
func (p *PDB) CountAtoms() int {
	var count int
	for _, m := range p.Model {
		for _, c := range m.Chain {
			for _, r := range c.Residue {
				count += len(r.Atom)
			}
		}
	}
	return count
}

// CountResidues returns the number of residues in the PDB
func (p *PDB) CountResidues() int {
	var count int
	for _, m := range p.Model {
		for _, c := range m.Chain {
			count += len(c.Residue)
		}
	}
	return count
}
