package glispext

import (
	"github.com/KernelDeimos/emptyglisp/interpreter"
	"math/rand"
	"time"
)

var defaultRand = rand.New(rand.NewSource(time.Now().Unix()))

func RandomFunction(env *glisp.Glisp, name string,
	args []glisp.Sexp) (glisp.Sexp, error) {
	return glisp.SexpFloat(defaultRand.Float64()), nil
}

func ImportRandom(env *glisp.Glisp) {
	env.AddFunction("random", RandomFunction)
}
