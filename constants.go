// constants.go contains the constants that are used by multiple files.
package pdbhandler

const (
	// https://regex101.com/r/LGV5z8/2
	ATOMREGEX = `ATOM  ((?:\s|\d){5})  ((?:\s|\w){3})(\s|\w)((?:\s|\w){3}) (\s|\w)((?:\s|\w){4})(\s|\w)   ((?:\s|.){8})((?:\s|.){8})((?:\s|.){8})((?:\s|.){6})((?:\s|.){6})\s{11}(\w)(\w)?`
)
