package merge

import (
	"goproxy/initial"
	"os"
	"sync"

	"github.com/bitly/go-simplejson"
)

func GenerateConfigJson() {
	outbound_channel := make(chan *simplejson.Json,50)
	outbounds := []*simplejson.Json{}
	var jobs sync.WaitGroup
	jobs.Add(1)
	// 获取固定信息
	log := initial.GetValue("template").(*simplejson.Json).Get("log")
	dns := initial.GetValue("template").(*simplejson.Json).Get("dns")
	inbounds := initial.GetValue("template").(*simplejson.Json).Get("inbounds")
	experimental := initial.GetValue("template").(*simplejson.Json).Get("experimental")
	// 获取会变化的信息,出站和路由
	route := MergeRoute()
	go func(){
		outbounds := MergeOutbounds()
		for _,outbound := range(outbounds){
			outbound_channel <- outbound
		}
		
		defer jobs.Done()
		defer close(outbound_channel)
	}()
	// 设置json
	config := simplejson.New()
	config.Set("log",log)
	config.Set("dns",dns)
	config.Set("inbounds",inbounds)
	config.Set("route",route)
	config.Set("experimental",experimental)
	
	for outbound:= range(outbound_channel){
		outbounds = append(outbounds, outbound)
	}
	jobs.Wait()

	config.Set("outbounds",outbounds)

	config_file,err := config.EncodePretty()
	initial.ErrorLog(err,true,"解析json文件失败")
	// 写入文件
	_,err = os.Stat("./temp/config.json")
	if err!=nil{
		if os.IsExist(err){
			os.Remove("./temp/config.json")
		}
	}else{
		os.Remove("./temp/config.json")
	}
	file ,err := os.OpenFile("./temp/config.json",os.O_CREATE,0)
	initial.ErrorLog(err,true,"写入json文件失败")
	file.WriteString(string(config_file))
	defer file.Close()
}
