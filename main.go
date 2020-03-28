package main

import (
	"flag"
)
import "./cli"
import "./gui"
import "os"

func main() {
	retVal := 0
	//rotors := [5] rotor {rotor.New(0, 1), rotor.New(0, 2), rotor.New(0, 3), rotor.New(0, 4), rotor.New(0, 5)}

	noGui := flag.Bool("nogui", false, "Disable the GUI")
	//config := flag.String("config", "", "Engima Machine Configuation File")
	//message := flag.String("message", "", "Enigma Machine Message File")

	flag.Parse()

	if !*noGui {
		iface := gui.New()
		retVal = iface.Exec()
	} else {
		iface := cli.New()
		retVal = iface.Exec()
	}

	os.Exit(retVal)
}
