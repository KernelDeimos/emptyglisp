package emptyglisp

import (
	"errors"
	"fmt"
	"os"
)

var WrongNargs error = errors.New("wrong number of arguments")

type GlispFunction []Instruction
type GlispUserFunction func(*Glisp, string, []Sexp) (Sexp, error)

var MissingFunction = SexpFunction{"__missing", true, 0, false, nil, nil, nil}

func MakeFunction(name string, nargs int, varargs bool,
	fun GlispFunction) SexpFunction {
	var sfun SexpFunction
	sfun.name = name
	sfun.user = false
	sfun.nargs = nargs
	sfun.varargs = varargs
	sfun.fun = fun
	return sfun
}

func MakeUserFunction(name string, ufun GlispUserFunction) SexpFunction {
	var sfun SexpFunction
	sfun.name = name
	sfun.user = true
	sfun.userfun = ufun
	return sfun
}

func SourceFileFunction(env *Glisp, name string, args []Sexp) (Sexp, error) {
	if len(args) < 1 {
		return SexpNull, WrongNargs
	}

	var sourceItem func(item Sexp) error

	sourceItem = func(item Sexp) error {
		switch t := item.(type) {
		case SexpArray:
			for _, v := range t {
				if err := sourceItem(v); err != nil {
					return err
				}
			}
		case SexpPair:
			expr := item
			for expr != SexpNull {
				list := expr.(SexpPair)
				if err := sourceItem(list.head); err != nil {
					return err
				}
				expr = list.tail
			}
		case SexpStr:
			var f *os.File
			var err error

			if f, err = os.Open(string(t)); err != nil {
				return err
			}

			if err = env.SourceFile(f); err != nil {
				return err
			}

			f.Close()
		default:
			return fmt.Errorf("%v: Expected `string`, `list`, `array` given type %T val %v", name, item, item)
		}

		return nil
	}

	for _, v := range args {
		if err := sourceItem(v); err != nil {
			return SexpNull, err
		}
	}

	return SexpNull, nil
}

func EvalFunction(env *Glisp, name string, args []Sexp) (Sexp, error) {
	if len(args) != 1 {
		return SexpNull, WrongNargs
	}
	newenv := env.Duplicate()
	err := newenv.LoadExpressions(args)
	if err != nil {
		return SexpNull, errors.New("failed to compile expression")
	}
	newenv.pc = 0
	return newenv.Run()
}
