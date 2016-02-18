/* mae.global/sexpr/cell/cell_test.go */
package cell

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
	"strings"
)

func compact(s string) string {
	return strings.Replace(strings.Replace(strings.Replace(s,"\n","",-1),"\t","",-1)," ","",-1)
}

func Test_Basic(t *testing.T) {

	Convey("Basics",t,func() {

		c := New(nil,nil)
		So(c,ShouldNotBeNil)
		So(compact(c.String()),ShouldEqual,"cell{empty+empty}")

		c0 := New(c,nil)
		So(c0,ShouldNotBeNil)
		So(compact(c0.String()),ShouldEqual,"cell{cell{empty+empty}+empty}")

		c1 := New(nil,c)
		So(c1,ShouldNotBeNil)
		So(compact(c1.String()),ShouldEqual,"cell{empty+cell{empty+empty}}")

		c01 := New(c0,c1)
		So(c01,ShouldNotBeNil)
		So(compact(c01.String()),ShouldEqual,"cell{cell{cell{empty+empty}+empty}+cell{empty+cell{empty+empty}}}")

		fmt.Printf("%s\n",c01)
	})

	Convey("User Information",t,func() {
		
		c := New(nil,nil)
		So(c,ShouldNotBeNil)
		So(c.Set("hello"),ShouldNotBeNil)
		
		v := c.Get()
		So(v,ShouldNotBeNil)
		s,ok := v.(string)
		So(ok,ShouldBeTrue)
		So(s,ShouldEqual,"hello")
	})

	Convey("Cell",t,func() {

		c := New(New(nil,nil),New(nil,nil))
		So(c,ShouldNotBeNil)
		So(compact(c.String()),ShouldEqual,"cell{cell{empty+empty}+cell{empty+empty}}")

		So(c.IsList(),ShouldBeTrue)
		So(c.IsEmpty(),ShouldBeFalse)

		c0 := c.List()
		So(c0,ShouldNotBeNil)
		So(compact(c0.String()),ShouldEqual,"cell{empty+empty}")

		So(c0.List(),ShouldBeNil)
		So(c0.Next(),ShouldBeNil)

		c1 := c.Next()
		So(c1,ShouldNotBeNil)
		So(compact(c1.String()),ShouldEqual,"cell{empty+empty}")	

		So(c1.List(),ShouldBeNil)
		So(c1.Next(),ShouldBeNil)
	})
				
}
