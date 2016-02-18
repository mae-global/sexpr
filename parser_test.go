/* mae.global/sexpr/parser_test.go */
package sexpr

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	
	"github.com/mae.global/sexpr/cell"
)

var (
	config *Configuration
)

func init() {

	config = &Configuration{}
	config.Verbose = true
	config.Annotate = true
	config.BufferSize = DefaultBufferSize
}

func Test_Simple(t *testing.T) {

	Convey("Simple",t,func() {
		
		root,err := ParseString("(+ 1 2)",config)
		So(err,ShouldBeNil)
		So(root,ShouldNotBeNil)
		So(cell.Compact(root),ShouldEqual,"cell{++cell{1+cell{2+empty}}}")
	})
}

func Test_Complex(t *testing.T) {

	Convey("Complex",t,func() {

		root,err := ParseString("(- (+ 1 2) (+ 1 2))",config)
		So(err,ShouldBeNil)
		So(root,ShouldNotBeNil)

		So(cell.Compact(root),ShouldEqual,"cell{-+cell{cell{++cell{1+cell{2+empty}}}+cell{cell{++cell{1+cell{2+empty}}}+empty}}}")
	})
}

func Test_Comments(t *testing.T) {

	Convey("Comments",t,func() {

		root,err := ParseString("(concat \"hello\" \"Alice\") ; comment",config)
		So(err,ShouldBeNil)
		So(root,ShouldNotBeNil)

		So(cell.Compact(root),ShouldEqual,"cell{concat+cell{\"hello\"+cell{\"Alice\"+empty}}}")
	})
}
