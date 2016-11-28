/* pinhole/internal/sexpr/cell/functions_test.go */
package cell

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
)

type value string

func (v value) String() string {
	return fmt.Sprintf("\"%s\"",string(v))
}

func Test_Functions(t *testing.T) {

	Convey("Cons",t,func() {
		c := Cons(New(value("Hello"),nil),New(value("Alice"),nil))
		So(c,ShouldNotBeNil)
		So(compact(c.String()),ShouldEqual,"cell{cell{\"Hello\"+empty}+cell{cell{\"Alice\"+empty}+empty}}")

		h := c.List()
		So(h,ShouldNotBeNil)
		So(compact(h.String()),ShouldEqual,"cell{\"Hello\"+empty}")
		So(h.IsValue(),ShouldBeTrue)
		v := h.Value()
		av,ok := v.(value)
		So(ok,ShouldBeTrue)
		So(string(av),ShouldEqual,"Hello")

		t := c.Next()
		So(t,ShouldNotBeNil)
		So(compact(t.String()),ShouldEqual,"cell{cell{\"Alice\"+empty}+empty}")
		So(t.IsList(),ShouldBeTrue)
		t = t.List()
		So(t,ShouldNotBeNil)
		So(compact(t.String()),ShouldEqual,"cell{\"Alice\"+empty}")

		So(t.IsValue(),ShouldBeTrue)
		v = t.Value()
		av,ok = v.(value)
		So(ok,ShouldBeTrue)
		So(string(av),ShouldEqual,"Alice")
			
	})

	Convey("Append",t,func() {

		c := New(value("Alice"),nil)
		So(c,ShouldNotBeNil)

		c = Append(c,New(value("in"),nil))
		So(c,ShouldNotBeNil)
		So(compact(c.String()),ShouldEqual,"cell{\"Alice\"+cell{\"in\"+empty}}")

		c = Append(c,New(value("Wonderland"),nil))
		So(c,ShouldNotBeNil)
		So(compact(c.String()),ShouldEqual,"cell{\"Alice\"+cell{\"in\"+cell{\"Wonderland\"+empty}}}")
	})

	Convey("List and Count",t,func() {

		c := List(value("Alice"),value("in"),value("Wonderland"))
		So(c,ShouldNotBeNil)
		So(compact(c.String()),ShouldEqual,"cell{\"Alice\"+cell{\"in\"+cell{\"Wonderland\"+empty}}}")

		So(Count(nil),ShouldEqual,0)
		So(Count(c),ShouldEqual,3)		
	})

	Convey("First and Rest",t,func() {

		So(First(nil),ShouldBeNil)
		So(Last(nil),ShouldBeNil)
		
		c := List(value("Madhatters"),value("tea"),value("party"))
		So(c,ShouldNotBeNil)

		f := First(c)
		So(f,ShouldNotBeNil)
		So(compact(f.String()),ShouldEqual,"cell{\"Madhatters\"+empty}")
		
		l := Rest(c)
		So(l,ShouldNotBeNil)
		So(compact(l.String()),ShouldEqual,"cell{\"tea\"+cell{\"party\"+empty}}")
	})

}

