package main
import(
	"fmt"
	js"github.com/dop251/goja"
)
func main(){
	vm:=js.New()
	r,_:=vm.RunString(`
		1+1
	`)
	fmt.Println(r)
}
