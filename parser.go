/* mae.global/sexpr/parser.go */
package sexpr

import (
	"fmt"
	"io"
	"strings"

	"github.com/mae.global/sexpr/cell"
)


const (
	DefaultBufferSize = 128
)

var (
	ErrZeroStack = fmt.Errorf("Zero Stack")
	ErrNoSource  = fmt.Errorf("No Source")
)

type Configuration struct {
	Verbose bool
	Annotate bool
	BufferSize int
}

func Parse(reader io.Reader,config *Configuration) (*cell.Cell,error) {

	verbose := false
	annotate := false
	buffersize := DefaultBufferSize

	if config != nil {
		verbose = config.Verbose
		annotate = config.Annotate
		if config.BufferSize > 0 {
			buffersize = config.BufferSize
		}
	}

	buf := make([]byte,buffersize)

	/* create helper */
	h := &helper{verbose:verbose,annotate:annotate}
	h.stack = make([]*cell.Cell,0)


	word := ""
	instr := false /* currently in string */
	incom := false /* currently in comment */

	s := -1
	f := -1
	lit := false
	idx := 0
	for {
		n,err := reader.Read(buf)
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil,err
		}

		for _,c := range buf[:n] {
			sc := string(c)
			if incom {
				if sc == "\n" {
					incom = false
				}
				idx ++
				continue
			}

			if sc == ";" && !instr {
				if verbose {
					fmt.Printf("\tcomment,s=%d\n",s)
				}
				incom = true
			} else if sc == "(" && !instr {
				if err := h.Push(idx); err != nil {
					return nil,err
				}
			} else if sc == ")" && !instr {
				if len(word) > 0 {
					if err := h.Append(s,f,lit,word); err != nil {
						return nil,err
					}
					if verbose {
						fmt.Printf("\ts=%d,f=%d,literal=%v [%s]\n",s,f,lit,word)
					}
					word = cell.Empty
					s = -1; f = -1; lit = false
				}
				if err := h.Pop(idx); err != nil {
					return nil,err
				}
				
			} else if (sc == " " || sc == "\n" || sc == "\t") && !instr {
				if len(word) > 0 {
					if err := h.Append(s,f,lit,word); err != nil {
						return nil,err
					}
					if verbose {
						fmt.Printf("\ts=%d,f=%d,literal=%v [%s]\n",s,f,lit,word)
					}
					word = cell.Empty
					s = -1; f = -1; lit = false
				}
			} else if sc == "\"" {
				if instr {
					lit = true
				}
				instr = !instr
				
			} else {
				if s == -1 {
					s = idx
				}
				f = idx
				word = word + string(c)
			}
			
			idx ++		
		}
	}
	/* tail */
	if len(word) > 0 {
		if err := h.Append(s,idx,lit,word); err != nil {
			return nil,err
		}
	}

	root := h.Root()
	if root == nil {
		return nil,ErrNoSource
	}

	return root,nil
}

func ParseString(s string,config *Configuration) (*cell.Cell,error) {
	return Parse(strings.NewReader(s),config)
}

/* see https://github.com/dvyukov/go-fuzz */
func Fuzz(data []byte) int {
	if _,err := ParseString(string(data),&Configuration{true,true,50}); err != nil {
		return 0
	}
	return 1
}	

	
