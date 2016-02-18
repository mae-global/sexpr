/* mae.global/sexpr/cell/functions.go */
package cell

func Cons(a,b Stringer) *Cell {
	cons := New(a,nil)
	if b != nil {
		tail := New(b,nil)
		tail.count++
		cons.t = tail
	}
	return cons
}

func Append(a,b *Cell) *Cell {
	
	c := a
	l := c
	count := 0
	for {
		if c == nil {
			break
		}
		l = c
		count = c.count
		c = c.Next()
	}

	count++
	l.t = b
	c = b
	for {
		if c == nil {
			break
		}
		c.count = count
		count++
		c = c.Next()
	}
	return a
}

func List(atoms... Stringer) *Cell {

	h := New(nil,nil)
	c := h
	next := false
	count := 0
	for _,v := range atoms {
		if next {
			count++
			n := New(nil,nil)
			n.count = count
			c.t = n
			c = n
			next = false
		}
		c.h = v
		next = true
	}

	return h
}

func Count(h *Cell) int {
	if h == nil {
		return 0
	}
	c := h
	count := h.count
	for {
		if c == nil {
			break
		}
		count = c.count
		c = c.Next()
	}
	return (count + 1) 
}

func First(c *Cell) *Cell {
	if c == nil {
		return nil
	}
	return c.Copy()
}

func Last(h *Cell) *Cell {
	if h == nil {
		return nil
	}
	c := h
	for {
		if c == nil {
			break
		}
		c = c.Next()
	}
	return c
}
 
func Rest(h *Cell) *Cell {
	var r,n *Cell
	c := h
	i := 0
	for {
		if c == nil {
			break
		}
		if i > 0 {
			if n == nil {
				n = c.Copy()
				r = n
			} else {
				n.t = c.Copy()
				n = n.Next()
				n.count = (i - 1)
			}
		}

		c = c.Next()
		i++
	}
	return r
}


