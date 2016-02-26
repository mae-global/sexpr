/* mae.global/sexpr/outputer.go */
package sexpr

import (
	"fmt"		
	"github.com/mae-global/sexpr/cell"
)
const (
	Whitespace = " "
)

var (
	ErrRootIsNil = fmt.Errorf("Root is nil")
)

func OutputString(root *cell.Cell) (string,error) {
	
	if root == nil {
		return "",ErrRootIsNil
	}

	var out string

	c := root
	prev := false
	ws := false
	b := false
	for {

		if c == nil  {
			break
		}

		if c.IsList() {
			if !ws && !b && len(out) > 0 {
				out += Whitespace
				ws = true
			}
			out += "("
			b = true
			o,err := OutputString(c.List())
			if err != nil {
				return "",err
			}
			out += o + ")" 
			b = true
			prev = false

		} else {
			if val := c.Value(); val != nil {
				if prev {
					out += Whitespace
				}
				out += val.String() 
				prev = true
				b = false
			}
		}

		c = c.Next()	
	
	}

	return out,nil
}

