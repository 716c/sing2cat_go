package initial

import (
	"os"
	"path/filepath"

	"github.com/bitly/go-simplejson"
)
func loadTemplate(){
	file,err := os.ReadFile(filepath.Join(GetValue("base_dir").(string),"/config/static/template.json"))
	ErrorLog(err,true,"读取模板文件失败")
	file_content,err := simplejson.NewJson(file)
	ErrorLog(err,true,"json序列化模板文件失败")
	SetValue("template",file_content)
}
func loadConfig(){
	file,err := os.ReadFile(filepath.Join(GetValue("base_dir").(string),"/config/config.json"))
	ErrorLog(err,true,"读取配置文件失败")
	file_content,err := simplejson.NewJson(file)
	ErrorLog(err,true,"json序列化配置文件失败")
	SetValue("config",file_content)
}
func InitializeLoad(){
	loadConfig()
	loadTemplate()
}