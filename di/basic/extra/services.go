package extra

import "github.com/google/wire"

// DefaultSet, e.g. for local development.
var DefaultSet = wire.NewSet(NewLogger, NewFoo, NewBar)

type Logger struct{}

func NewLogger() *Logger {
	return &Logger{}
}

type Foo struct {
	l *Logger
}

func NewFoo(l *Logger) *Foo {
	return &Foo{
		l: l,
	}
}

type Bar struct {
	foo *Foo
}

func NewBar(foo *Foo) *Bar {
	return &Bar{
		foo: foo,
	}
}
