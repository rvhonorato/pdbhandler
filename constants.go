package pdbhandler

const (
	// MODEL: `https://www.wwpdb.org/documentation/file-format-content/format33/sect9.html#MODEL`
	MODEL_RECORD_REGEX = `MODEL     ((?:\s|\d){4})`

	// ATOM: `https://www.wwpdb.org/documentation/file-format-content/format33/sect9.html#ATOM`
	ATOM_RECORD_REGEX = `ATOM\s{2}((?:\s|\d){5})\s((?:\s|\w){3})\s((?:\s|\w){1})((?:\s|\w){3})\s((?:\s|\w){1})((?:\s|\w){4})((?:\s|\w){1})\s{3}((?:\s|.){8})((?:\s|.){8})((?:\s|.){8})((?:\s|.){6})((?:\s|.){6})\s{10}((?:\s|\w){2})((?:\s|\w){2})`

	// TER: `https://www.wwpdb.org/documentation/file-format-content/format33/sect9.html#TER`
	TERREGEX = `TER   ((?:\s|\d){5})      ((?:\s|\w){3}) (\s|\w)((?:\s|\w){4})(\s|\w)`

	// HETATM: `https://www.wwpdb.org/documentation/file-format-content/format33/sect9.html#HETATM`
	HETATMREGEX = `HETATM((?:\s|\d){5})\s((?:\s|\w){3})\s((?:\s|\w){1})((?:\s|\w){3})\s((?:\s|\w){1})((?:\s|\w){4})((?:\s|\w){1})\s{3}((?:\s|.){8})((?:\s|.){8})((?:\s|.){8})((?:\s|.){6})((?:\s|.){6})\s{10}((?:\s|\w){2})((?:\s|\w){2})`

	// ENDMDL: `https://www.wwpdb.org/documentation/file-format-content/format33/sect9.html#ENDMDL`
	ENDMDLREGEX = `ENDMDL`
)
