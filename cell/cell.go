/* pinhole/internal/sexpr/cell/cell.go */
package cell

import (
	"fmt"
	"strings"
)

const (
	Empty = ""
)

type Stringer interface {
	String() string
}

func Compact(c *Cell) string {
	if c == nil {
		return "nil"
	}
	s := c.String()
	return strings.Replace(strings.Replace(strings.Replace(s,"\n","",-1),"\t","",-1)," ","",-1)
}


/* Cell struct for placing two atoms into */
type Cell struct {
	h,t Stringer    /* head & tail, car & cdr, first & rest ... */
  u interface{} /* user-data, useful for high-level parser annotation */
  count int
}

func (c *Cell) String() string {
	if c == nil {
		return "nil"
	}

	h := "empty"
	t := "empty" 
	
	if c.h != nil {
		h = c.h.String()
	}
	if c.t != nil {
		t = c.t.String()
	} 

	return fmt.Sprintf("cell{\n\t%s + \n\t%s\n}",h,t) /* TODO: formating needs work, maybe use tree? */
}

/* Set the user information */
func (c *Cell) Set(u interface{}) *Cell {
	c.u = u
	return c
}

/* Get the user information that is set */
func (c *Cell) Get() interface{} {
	return c.u
}

/* Copy not a deep copy of contents */
func (c *Cell) Copy() *Cell {
	n := New(c.h,nil)
  n.u = c.u
 	return n
}

func (c *Cell) IsEmpty() bool {
	return (c.h == nil && c.t == nil)
}

func (c *Cell) IsList() bool {
	if c.h == nil {
		return false
	}
	_,ok := c.h.(*Cell)
	return ok
}

func (c *Cell) IsValue() bool {
	return (c.IsList() == false)
}

func (c *Cell) Value() Stringer {
	return c.h
}

func (c *Cell) ToSValue() (string,bool) {
	ok := c.IsValue()
	if !ok {
		return "",false
	}
	return c.h.String(),ok
}

func (c *Cell) SValue() string {
	if c.IsValue() {
		return c.h.String()
	}
	return ""
}

func (c *Cell) List() *Cell {
	if c.h == nil {
		return nil
	}
	if a,ok := c.h.(*Cell); ok {
		return a
	}
	return nil
}

func (c *Cell) Next() *Cell {
	if c.t == nil {
		return nil
	}
	if a,ok := c.t.(*Cell); ok {
		return a
	}
	return nil
}

func Open(cell *Cell) *Cell {
	return &Cell{h:cell,t:nil}
}

func New(h,t Stringer) *Cell {
	return &Cell{h:h,t:t}
}





