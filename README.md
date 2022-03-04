## 秒级定时任务


### 增加任务

    修改config/cron.yaml, 重启服务

### 部署

1. 创建应用配置文件 /etc/systemd/system/ycron.service
2. 使用 systemctl daemon-reload 重新加载服务;
3. 执行 systemctl start ycron 来启动服务;
4. 最后执行 systemctl status ycron 来查看服务运行的状态信息;
5. 执行 systemctl enable ycron 将服务添加到开机启动项;
6. 注意：执行的 ycron 是使用文件名作为服务名;
7. 常见的命令有: start(启动), stop(停止), restart(重启), status(查看运行状态), enable(添加到开机启动项), disable(将程序从开机启动中移除)
