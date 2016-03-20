/* mae.global/sexpr/parser_test.go */
package sexpr

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	
	"github.com/mae-global/sexpr/cell"
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
		So(cell.Compact(root),ShouldEqual,"cell{cell{++cell{1+cell{2+empty}}}+empty}")
	})
}

func Test_Complex(t *testing.T) {

	Convey("Complex",t,func() {

		root,err := ParseString("(- (+ 1 2) (+ 1 2))",config)
		So(err,ShouldBeNil)
		So(root,ShouldNotBeNil)

		So(cell.Compact(root),ShouldEqual,"cell{cell{-+cell{cell{++cell{1+cell{2+empty}}}+cell{cell{++cell{1+cell{2+empty}}}+empty}}}+empty}")
	})
}

func Test_Comments(t *testing.T) {

	Convey("Comments",t,func() {

		root,err := ParseString("(concat \"hello\" \"Alice\") ; comment",config)
		So(err,ShouldBeNil)
		So(root,ShouldNotBeNil)

		So(cell.Compact(root),ShouldEqual,"cell{cell{concat+cell{\"hello\"+cell{\"Alice\"+empty}}}+empty}")
	})
}

//go test -bench=.

func Benchmark_Simple(b *testing.B) {

	test := "(+ 1 2)"

	for i := 0; i < b.N; i++ {
		if _,err := ParseString(test,nil); err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_Medium(b *testing.B) {

	test := "((concat (\"Alice\" \"in Wonderland\")) (- (+ 9 2 3 4) (+ 2 3 4 (* 2 3))))"

	for i := 0; i < b.N; i++ {
		if _,err := ParseString(test,nil); err != nil {
			b.Error(err)
		}
	}
}

func Benchmark_Complex(b *testing.B) {

	test := "(card 88745ec816d24fc5 5 (attributes (attribute \"Alice\" (string \"in wonderland\"))(attribute \"List of Things\" (delimited-string \",\" \"Butterfly,Rabbit,Hamster\"))(attribute \"time\" (time \"11:39\"))))"

	for i := 0; i < b.N; i++ {
		if _,err := ParseString(test,nil); err != nil {
			b.Error(err)
		}
	}
}


		
