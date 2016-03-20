/* mae.global/sexpr/helper.go */
package sexpr

import (
	"github.com/mae-global/sexpr/cell"
	"fmt"
)

type Literal string

func (l Literal) String() string {
	return fmt.Sprintf("\"%s\"",string(l))
}

type String string

func (s String) String() string {
	return string(s)
}



type helper struct {
	verbose bool
	annotate bool

	stack []*cell.Cell
	root *cell.Cell
}

func (h *helper) Push(start int) error {

	if len(h.stack) == 0 {
		/*if h.verbose {
			fmt.Printf("Push\n") 
		}*/
		
		list := cell.New(nil,nil)
		h.stack = append(h.stack,list)
		return nil
	}

	/*c := h.stack[len(h.stack) - 1]

	if h.verbose {
		fmt.Printf("Push\tempty=%v\n",c.IsEmpty())
	}*/

	list := cell.New(nil,nil)
	h.stack = append(h.stack,list)
	return nil
}

func (h *helper) Pop(finish int) error {
	/*
	if h.verbose {
		fmt.Printf("Pop\n")
	}*/

	if len(h.stack) == 1 {
		h.root = h.stack[len(h.stack) - 1]
		h.stack = h.stack[:len(h.stack) - 1]
	} else if len(h.stack) > 1 {
	
		list := h.stack[len(h.stack) - 1]
		c := h.stack[len(h.stack) - 2]
		cons := cell.Cons(list,nil)
		c = cell.Append(c,cons)

		h.stack = h.stack[:len(h.stack) - 1]	
		h.stack[len(h.stack) - 1] = c
	}
	return nil
}

func (h *helper) Append(start,finish int,literal bool,word string) error {

	if len(h.stack) == 0 {
		return nil
	}

	c := h.stack[len(h.stack) - 1]

	if c.IsEmpty() {

		if literal {			
			c = cell.Cons(Literal(word),nil)
		} else {
			c = cell.Cons(String(word),nil)
		}
	} else {

		var cons *cell.Cell		
		if literal {		
			cons = cell.Cons(Literal(word),nil)
		} else {
			cons = cell.Cons(String(word),nil)
		}

		c = cell.Append(c,cons)
	}

	h.stack[len(h.stack) - 1] = c
	return nil
}

func (h *helper) Root() *cell.Cell {
	return h.root
}


	




