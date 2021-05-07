package main
import(
	"fmt"
	js"github.com/dop251/goja"
)
func main(){
	vm:=js.New()
	{
		vm.Set("foo",4)
		vm.Set("bar",2)
		r,_:=vm.RunString(`
			foo+bar
		`)
		fmt.Println(r)
	}
	{
		vm.Set("baz",func()int{return 42})
		r,_:=vm.RunString(`
			baz()
		`)
		fmt.Println(r)
	}
	{
		log:=func(val string){
			fmt.Println(val)
		}
		console:=vm.NewObject()
		console.Set("log",log)
		vm.Set("console",console)
		vm.RunString("console.log('HELLO0')")
		vm.RunString("console.log(42)")
	}
	{
		log:=func(call js.FunctionCall)js.Value{
			str:=call.Argument(0)
			fmt.Println(str.String())
			return str
		}
		console:=vm.NewObject()
		console.Set("log",log)
		vm.Set("console",console)
		vm.RunString("console.log('HELLO1')")
		vm.RunString("console.log(24)")
	}
}
