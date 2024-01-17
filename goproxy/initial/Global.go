package initial

import (
	"fmt"
	"os"
	"path/filepath"

	"go.uber.org/zap"
)
var global_vars = make(map[string]interface{})
func GetWorkingDir() {
	root_dir,err:= os.Executable()
	if err!=nil{
		ErrorLog(err,true,"获取文件目录失败")
	}
	global_vars["base_dir"] = (filepath.Dir(root_dir))
}
func GetValue(key string) interface{}{
	result := global_vars[key]
	if result==nil {fmt.Println("没这个值")}
	return result
}
func SetValue(key string,value interface{}){
	global_vars[key] = value
}
func ErrorLog(err error,fatal bool,msg string){
	if err!=nil{
		if fatal {
			Logger.Panic(msg,zap.Error(err))
		}else {
			Logger.Error(msg,zap.Error(err))
		}		
	}
}