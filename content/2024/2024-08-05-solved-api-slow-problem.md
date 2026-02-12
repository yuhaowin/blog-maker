## http 接口偶然响应慢问题排查

apt-get update && apt-get install -y tcpdump hping3 mtr traceroute iputils-ping --no-install-recommends tcpdump hping3 mtr

tcpdump -i eth0 host api.dify.ai -w capture.pcap

kubectl cp crm-qa-12579-7ff6459456-rzzz6:/app/bin/capture.pcap ~/Desktop/crm-qa.pcap


curl --location 'https://api.dify.ai/v1/chat-messages' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer fake-token' \
--data '{
"query": "拉卡拉校园食堂开通分账业务审批填写规范",
"response_mode": "streaming",
"user": "yuhao",
"inputs": {}
}'

https://weiborao.link/docker-traceroute-tcpdump.html

curl -o /dev/null -s -w time_namelookup:"\t"%{time_namelookup}"\n"time_connect:"\t\t"%{time_connect}"\n"time_appconnect:"\t"%{time_appconnect}"\n"time_pretransfer:"\t"%{time_pretransfer}"\n"time_starttransfer:"\t"%{time_starttransfer}"\n"time_total:"\t\t"%{time_total}"\n"time_redirect:"\t\t"%{time_redirect}"\n" --location 'https://api.dify.ai/v1/chat-messages' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer fake-token' \
--data '{
"query": "拉卡拉校园食堂开通分账业务审批填写规范",
"response_mode": "streaming",
"user": "yuhao",
"inputs": {}
}'

tcpdump   -i lo  port  6310 -w http.pcap

tcpdump -i eth0 host sg-dify.wosai-inc.com -w capture.pcap

tcpdump -i eth0 host api.dify.ai -w capture.pcap


apt-get update && apt-get install -y traceroute iputils-ping --no-install-recommends tcpdump

apt-get install hping3

hping3 -S -p 443 -c 100 api.dify.ai


kubectl cp crm-qa-base-5799dbd479-kjxcv:/app/bin/capture.pcap ~/Desktop/crm-qa-1.pcap



kubectl cp crm-qa-12579-7ff6459456-rzzz6:/opt/jacoco/report/profit-sharing/report.exec ~/Downloads/report.exec

curl --location 'https://api.dify.ai/v1/chat-messages' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer fake-token' \
--data '{
"query": "拉卡拉校园食堂开通分账业务审批填写规范",
"response_mode": "streaming",
"user": "yuhao",
"inputs": {}
}'


curl -w "总时间: %{time_total}s\n名称解析时间: %{time_namelookup}s\n连接时间: %{time_connect}s\nTLS握手时间: %{time_appconnect}s\n等待时间: %{time_starttransfer}s\n数据传输时间: %{time_total}s\nHTTP状态码: %{http_code}\n" -o /dev/null --location 'https://api.dify.ai/v1/chat-messages' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer fake-token' \
--data '{
"query": "拉卡拉校园食堂开通分账业务审批填写规范",
"response_mode": "streaming",
"user": "yuhao",
"inputs": {}
}'



curl -o /dev/null -s -w time_namelookup:"\t"%{time_namelookup}"\n"time_connect:"\t\t"%{time_connect}"\n"time_appconnect:"\t"%{time_appconnect}"\n"time_pretransfer:"\t"%{time_pretransfer}"\n"time_starttransfer:"\t"%{time_starttransfer}"\n"time_total:"\t\t"%{time_total}"\n"time_redirect:"\t\t"%{time_redirect}"\n" --location 'http://sg-dify.wosai-inc.com/v1/chat-messages' \
--header 'Content-Type: application/json' \
--header 'Authorization: Bearer fake-token' \
--data '{
"query": "拉卡拉校园食堂开通分账业务审批填写规范",
"response_mode": "streaming",
"user": "yuhao",
"inputs": {}
}'



curl -o /dev/null -s -w time_namelookup:"\t"%{time_namelookup}"\n"time_connect:"\t\t"%{time_connect}"\n"time_appconnect:"\t"%{time_appconnect}"\n"time_pretransfer:"\t"%{time_pretransfer}"\n"time_starttransfer:"\t"%{time_starttransfer}"\n"time_total:"\t\t"%{time_total}"\n"time_redirect:"\t\t"%{time_redirect}"\n" --location --request GET 'https://api.dify.ai/v1/datasets?page=1&limit=20' \
--header 'Authorization: Bearer fake-token'

