# wrk 命令做压测示例
wrk -t4 -c100 -d30s http://x.x.x.x:3000/ > tmp.log

