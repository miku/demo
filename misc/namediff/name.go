// This package anyname lives under a different directory. Directories and
// package name do not need to match, although often they will.
package anyname

import "fmt"

func Render() string {
	return fmt.Sprintln("anyname")
}
