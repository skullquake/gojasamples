package main
import(
	"log"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"github.com/dop251/goja"
)
func main(){
	if len(os.Args[1:])<1{
		log.Fatalf("invalid arguments")
		os.Exit(1)
	}
	vm:=goja.New()
	Register_jsext_consoleutils(vm)
	Register_jsext_fsutils(vm)
	for _,a:=range os.Args[1:]{
		buf,err:=ioutil.ReadFile(a)
		if err!=nil{
			log.Fatalf("failed to open file")
			panic(err)
		}
		vm.RunString(string(buf))
	}
}
func Register_jsext_fsutils(vm*goja.Runtime) {
	vm.SetFieldNameMapper(goja.TagFieldNameMapper("json",true))
	type Version struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Bump  int `json:"bump"`
	}
	type EntryInfo struct {
		Name string `json:"name"`
		Path string `json:"path"`
		Type string `json:"type"`
		Size int64  `json:"size"`
	}
	vm.Set("fsutils", struct {
		Version         Version                  `json:"version"`
		File2String     func(string) string      `json:"file2string"`
		String2File     func(string, string)     `json:"string2file"`
		Ls              func(string) []EntryInfo `json:"ls"`
		Glob            func(string) []EntryInfo `json:"glob"`
		Walk            func(string) []EntryInfo `json:"walk"`
		MkDir           func(string) bool        `json:"mkdir"`
		Rm              func(string) bool        `json:"rm"`
		Rename          func(string,string)bool  `json:"rename"`
		Mv              func(string,string)bool  `json:"mv"`
	}{
		Version: Version{
			Major: 0,
			Minor: 0,
			Bump:  4,
		},
		File2String: func(path string) string {
			content, err := ioutil.ReadFile(path)
			if err != nil {
				panic(vm.ToValue("Failed to open file"))
			}
			text := string(content)
			return text
		},
		String2File: func(path string, contents string) {
			file, err := os.Create(path)
			if err != nil {
				panic(vm.ToValue("Failed to create file"))
			}
			defer file.Close()
			_, err = io.WriteString(file, contents)
			if err != nil {
				panic(vm.ToValue("Failed to write to file"))
			}
			if file.Sync() != nil {
				panic(vm.ToValue("Failed to sync file"))
			}
		},
		Ls: func(path string) []EntryInfo {
			var ret []EntryInfo
			cwd, err := os.Getwd()
			if err != nil {
				panic(vm.ToValue("Failed to get cwd"))
			}
			entries, err := ioutil.ReadDir(path)
			if err != nil {
				panic(vm.ToValue("Failed to open directory"))
			}
			for _, info := range entries {
				var abspath string
				if filepath.IsAbs(path) {
					abspath = filepath.Join(path, info.Name())
				} else {
					abspath = filepath.Join(cwd, filepath.Join(path, info.Name()))
				}
				if info.Mode().IsRegular() {
					ret = append(ret, EntryInfo{
						Name: info.Name(),
						Path: abspath,
						Type: "File",
						Size: info.Size(),
					})
				} else if info.Mode().IsDir() {
					ret = append(ret, EntryInfo{
						Name: info.Name(),
						Path: abspath,
						Type: "Dir",
						Size: info.Size(),
					})
				}
			}
			return ret
		},
		Glob: func(path string) []EntryInfo {
			var ret []EntryInfo
			matches, err := filepath.Glob(path)
			if err != nil {
				panic(vm.ToValue("Failed to open directory"))
			}
			for _, match := range matches {
				ret = append(ret, EntryInfo{
					Path: match,
				})

			}
			return ret
		},
		Walk: func(path string) []EntryInfo {
			var ret []EntryInfo
			err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
				if err != nil {
					return err
				}
				if info.Mode().IsRegular() {
					ret = append(ret, EntryInfo{
						Name: info.Name(),
						Path: path,
						Type: "File",
						Size: info.Size(),
					})
				} else if info.Mode().IsDir() {
					ret = append(ret, EntryInfo{
						Name: info.Name(),
						Path: path,
						Type: "Dir",
						Size: info.Size(),
					})
				}
				return nil
			})
			if err != nil {
				panic(vm.ToValue("Failed to open directory"))
			}
			return ret
		},
		MkDir: func(path string)(bool){
			err:=os.MkdirAll(path,0777)
			if err!=nil{
				panic(vm.ToValue("Failed to create directory"))
			}
			return true
		},
		Rm: func(path string)(bool){
			err:=os.RemoveAll(path)
			if err!=nil{
				panic(vm.ToValue("Failed to remove"))
			}
			return true
		},
        Rename: func(src string,tgt string)(bool){
			//stub
			return true
		},
        Mv: func(src string,tgt string)(bool){
			//stub
			return true
		},
	})
}
func Register_jsext_consoleutils(vm *goja.Runtime) {
	vm.SetFieldNameMapper(goja.TagFieldNameMapper("json",true))
	type Version struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Bump  int `json:"bump"`
	}
	//log.SetFlags(log.LstdFlags|log.Lmicroseconds)
	l:=log.New(os.Stdout,"",0)//no timestamp
	vm.Set("console", struct {
		Version Version              `json:"version"`
		Log     func(...interface{}) `json:"log"`
		Warn    func(...interface{}) `json:"warn"`
		Error   func(...interface{}) `json:"error"`
		Debug   func(...interface{}) `json:"debug"`
		Trace   func(...interface{}) `json:"trace"`
	}{
		Version: Version{
			Major: 0,
			Minor: 0,
			Bump:  1,
		},
		Log: func(msg ...interface{}) {
			l.Println(msg...)
		},
		Warn: func(msg ...interface{}) {
			l.Println(msg...)
		},
		Error: func(msg ...interface{}) {
			l.Println(msg...)
		},
		Debug: func(msg ...interface{}) {
			l.Println(msg...)
		},
		Trace: func(msg ...interface{}) {
			l.Println(msg...)
		},
	})
}
