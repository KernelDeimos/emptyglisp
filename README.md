# EmptyGLisp

This is a modified version of Glisp with all the built-in functions removed,
making it easy to use this interpreter to create new Glisp-like languages.

This package has two goals:

## Goal A
Completely de-couple the interpreter from the rest of Glisp, which means
eventually removing everything under `extensions/`,
and any functions that the interpreter doesn't depend on.

## Goal B
Make the interpreter as extensible as possible, so a user of this package can
add new syntax as limitlessly as possible, but while keeping the interpreter
itself as simple as possible.

For instance, the user might want to enable a special syntax marker for regular
expressions. EmptyGLisp will not provide a parseRegex function, but it will
provide a means of calling a custom function to parse regex when this marker
is found.

## Goal C
I'm gonna add a way to do everything without brackets. Fight me.

# GLisp

Below is the original readme text from Glisp; you can find more at the
[source repository](https://github.com/zhemao/glisp).

This is a LISP dialect designed as an embedded extension language for the Go
programming language. It is implemented in pure Go, so it can be easily ported
to all systems and architectures that Go targets.

Here is a list of what features are implemented and not implemented so far.

 * [x] Float, Int, Char, String, Symbol, List, Array, and Hash datatypes
 * [x] Arithmetic (`+`, `-`, `*`, `/`, `mod`)
 * [x] Shift Operators (`sll`, `srl`, `sra`)
 * [x] Bitwise operations (`bit-and`, `bit-or`, `bit-xor`)
 * [x] Comparison operations (`<`, `>`, `<=`, `>=`, `=`, and `not=`)
 * [x] Short-circuit boolean operators (`and` and `or`)
 * [x] Conditionals (`cond`)
 * [x] Lambdas (`fn`)
 * [x] Bindings (`def`, `defn`, and `let`)
 * [x] A Basic Repl
 * [x] Tail-call optimization
 * [x] Go API
 * [x] Macro System
 * [x] Syntax quoting (backticks)
 * [x] Channel and goroutine support
 * [x] Pre- and Post- function call hooks

The full documentation can be found in the [Wiki](https://github.com/zhemao/glisp/wiki).
