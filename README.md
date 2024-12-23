## walle

<br/>

## 🏷️ Description

This project is dedicated to implementation human-like generator traffic service. It's task for my institue practice. We can see task-requirements is contains in TASK.md.
Probably, I will add other useful and more complex features to this implementation in the future ⌛.

<p align="right"><a href="#walle">Back to top ⬆️</a></p>

## 🎯 Solutions and Techniques

- Layered architecture 
- Supports simple adding new traffic type module 
- Auto application configuration using config parser
- Multi-level logging using zap logger
- Flexibility deploy with docker
- Reproducibility thanks to config.yaml + configParser
- Obfuscation traffic with xray-core 
- Generation human-like traffic using selenium webdriver

<p align="right"><a href="#walle">Back to top ⬆️</a></p>


## 🗂️ Table of Contents
- [Description](#️-description)
- [Solutions and Techniques](#-solutions-and-techniques)
- [Table of Contents](#️-table-of-contents)
- [Working Tree](#-working-tree)
- [Сonfiguration file](#-сonfiguration-file)
- [Getting Started](#️--getting-started)
- [To do](#-to-do)
- [Contact](#-contact)

## 🌿 Working Tree
```
├── cmd
│   └── main.go
│      
├── config     
│   └── config.go
│   └── config.json
│
├── docker-compose.yaml
├── go.mod
├── go.sum
├── internal
│   ├── app
│   │   ├── app.go
│   │   └── interface.go
├── pkg
│   ├── habr
│   │   └── habr.go
│   ├── httpclient
│   │   └── httpclient.go
│   ├── jscript
│   │   └── jscript.go
│   ├── null
│   │   └── null.go
│   └── youtube
│       └── youtube.go
├── .gitignore
├── README.md
└── TASK.md
```
<p align="right"><a href="#walle">Back to top ⬆️</a></p>

## ⚙️ Сonfiguration file

```
app
{
  "UserAgent": "your_user_agent",
  "Proxy": {
    "Host": "xray_client_ip",
    "Port": "10800"
  },
  "Timeout": "1s",
  "Delay": "10s",
  "Deps": {
    "Habr": {
      "URLs": harb_urls
    },
    "YouTube": {
      "URLs": youtube_urls
    }
  }
}
```
```
xray
{
  "log": {
    "loglevel": "info",
    "access": "/var/log/xray/access.log",
    "error": "/var/log/xray/error.log"
  },
  "dns": {
    "servers": [
      {
        "address": "1.1.1.1",
        "domains": [
          "geosite:geolocation-!cn"
        ]
      },
      "localhost"
    ]
  },
  "routing": {
    "domainStrategy": "AsIs",
    "rules": [
      {
        "type": "field",
        "domain": ["geosite:category-ads-all"],
        "outboundTag": "block"
      },
      {
        "type": "field",
        "domain": ["geosite:category-ru"],
        "outboundTag": "direct"
      }
    ]
  },
  "inbounds": [
    {
      "tag": "socks-in",
      "protocol": "socks",
      "listen": "xray_client_ip",
      "port": 10800,
      "settings": {
        "udp": true
      }
    },
    {
      "tag": "http-in",
      "protocol": "http",
      "listen": "xray_client_ip",
      "port": 10801
    }
  ],
  "outbounds": [
    {
      "tag": "proxy",
      "protocol": "vless",
      "settings": {
        "vnext": [
          {
            "address": "xray_server_ip",
            "port": 443,
            "users": [
              {
                "id": "uuid",
                "flow": "xtls-rprx-vision",
                "encryption": "none",
                "level": 0
              }
            ]
          }
        ]
      },
      "streamSettings": {
        "network": "tcp",
        "security": "tls",
        "tlsSettings": {
          "serverName": "your_domain",
          "allowInsecure": false,
          "fingerprint": "chrome"
        }
      }
    },
    {
      "tag": "direct",
      "protocol": "freedom"
    },
    {
      "tag": "block",
      "protocol": "blackhole"
    }
  ]
}
```

<p align="right"><a href="#walle">Back to top ⬆️</a></p>

## 🛠️  Getting Started

0. Install all required dependencies(Go, Docker and etc)

1. Clone the repository

```bash
git clone https://github.com/vlpakhomov/walle
```   

2. You can build and run containers with default settings using the following commands:
```
# create and compose up with default settings
docker compose up -d
```

<p align="right"><a href="#walle">Back to top ⬆️</a></p>

## 📌 To do

- Implement mechanism for simple adding new traffic type module  ✅
- Generation human-like traffic by different human-patterns ✅
- Configure integration with xray-core ✅
- Multi-level logging ✅
- Containerization with docker ✅
- .....

<p align="right"><a href="#walle">Back to top ⬆️</a></p>



## 📫 Contact

Vladislav Pakhomov - [@gregmackool](https://t.me/gregmackool) - [vladislavpakhomov03@gmail.com](mailto:vladislavpakhomov03@gmail.com)
<br/>
Maxim Prokhorov - [@nottoxxic](https://t.me/nottoxxic) - [maksim78998@gmail.com](mailto:maksim78998@gmail.com)

Project Link: clone https://github.com/vlpakhomov/walle

<p align="right"><a href="#walle">Back to top ⬆️</a></p>
