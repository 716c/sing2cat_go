#!/bin/bash
cp /opt/singbox/config.json /opt/sing2cat/temp/config.json.bak
/opt/sing2cat/sing2cat
config=true
if [ $config == true ]; then
	if [ -e /opt/sing2cat/temp/config.json ]; then
		if [ -s /opt/sing2cat/temp/config.json ]; then
			echo "下载成功"
		else
			echo "下载失败"
			config=false
		fi
	else
		echo "不存在"
		config=false
	fi
fi


if [ $config == true ]; then
	systemctl stop sing-box.service
  rm -rf /opt/singbox/config.json
	cp /opt/sing2cat/temp/config.json /opt/singbox/config.json
	systemctl start sing-box.service
else
	echo "不更新"
fi

sleep 3
if systemctl status sing-box.service |grep -q "running"; then
	echo "更新完成"
	rm -rf /opt/sing2cat/temp/config.json.bak
else
	echo "似乎哪里暴毙了,开始恢复"
	rm -rf /opt/singbox/config.json
	cp /opt/sing2cat/temp/config.json.bak /opt/singbox/config.json
	rm -rf /opt/sing2cat/temp/config.json.bak
	systemctl start sing-box.service
fi
