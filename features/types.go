package features

import (
	"errors"
	. "mk3cli/s7cli/colors"
	. "mk3cli/s7cli/commands"
)

// Feature Structs

type Feature struct {
	Name         string
	Description  string
	Args         Args
	Dependencies []string
	ReturnsData  bool
	GenerateCode func(args FeatureSetArgsList) (string, error)
}

type Features []Feature

type FeatureSet struct {
	Feature Feature
	Enabled bool
	Args    FeatureSetArgsList
}

type FeatureSetArg struct {
	Arg   Arg
	Value interface{}
}

type FeatureSetArgsList []FeatureSetArg

func (this Feature) DisplayUsage() {
	usage := this.Name + " "

	for _, a := range this.Args {
		if a.Required {
			usage += "--" + a.Name + " [" + a.Datatype + "] "
		} else {
			usage += "--" + Gray + a.Name + " [" + Cyan + a.Datatype + Gray + "] "
		}
	}
	println(usage)
}

func (this FeatureSetArgsList) Find(name string) (FeatureSetArg, error) {
	for _, f_arg := range this {
		if f_arg.Arg.Name == name {
			return f_arg, nil
		}
	}

	return FeatureSetArg{}, errors.New("Failed to find argument '" + name + "' in FeatureSetArgsList")
}
