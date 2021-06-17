module github.com/miku/demo

go 1.16

require (
	github.com/fatih/color v1.10.0
	github.com/fatih/structs v1.1.0
	github.com/google/wire v0.5.0
)

// Cannot just invent a name an redirect it.
// replace vanity.xyz/learn => ./localpkg/learn
