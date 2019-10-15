import cfscrape
import sys
import requests

# 此脚本需要安装nodejs
# apt-get install nodejs

url = sys.argv[1]
userAgent = sys.argv[2]

proxies = {}
if len(sys.argv) == 4:
  proxy = sys.argv[3]
  protocol = proxy.split(":")[0]
  proxies = {protocol:proxy}


headers = {
  "User-Agent":userAgent
}


tokens = cfscrape.get_tokens(url,proxies=proxies,user_agent=userAgent,timeout=5)
print(tokens[0])






