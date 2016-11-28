/* pinhole/internal/outputer_test.go */
package sexpr

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"fmt"
)

func Test_Outputer(t *testing.T) {

	Convey("Outputer Simple",t,func() {
		
		root,err := ParseString("(+ 1 2)",config)
		So(err,ShouldBeNil)
		So(root,ShouldNotBeNil)

		out,err := OutputString(root)
		So(err,ShouldBeNil)
		So(out,ShouldEqual,"(+ 1 2)")


	})

	Convey("Outputer Complex",t,func() {

		root,err := ParseString("(- (+ 1 2)(+ 1 2))",config)
		So(err,ShouldBeNil)
		So(root,ShouldNotBeNil)
	
		out,err := OutputString(root)
		So(err,ShouldBeNil)
		So(out,ShouldEqual,"(- (+ 1 2)(+ 1 2))")	
	})

	Convey("Outputer Comments",t,func() {

		root,err := ParseString("(concat \"hello\" \"Alice\") ; comment",config)
		So(err,ShouldBeNil)
		So(root,ShouldNotBeNil)
		
		out,err := OutputString(root)
		So(err,ShouldBeNil)
		So(out,ShouldEqual,"(concat \"hello\" \"Alice\")")

	})

	Convey("Outputer Example 001",t,func() {
			
		l := "(card b0a0639e5ev565cb (version 1)(attributes (attribute \"Alice\" (string \"in wonderland\"))))"
		root,err := ParseString(l,config)
		So(err,ShouldBeNil)
		So(root,ShouldNotBeNil)

		fmt.Printf("root = %s\n",root)

		out,err := OutputString(root)
		So(err,ShouldBeNil)
		So(out,ShouldEqual,l)
	})
		

}
