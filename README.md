# baseHttp
base on: https://github.com/bxcodec/go-clean-arch#readme  
http api server example  

using cobra to implement micro service  
using gin to implement http service  

```
├── README.md
├── app
│   ├── api                      // api main
│   └── main.go                  // app main
├── config                       // config 模組資料夾
│   ├── db_configuration.go
│   └── env_configuration.go
├── config.yaml                  // 環境設定檔
├── domain                       // 領域資料夾
│   └── user
├── entities                     // core
│   ├── delivery                 // response
│   └── user.go
├── go.mod
├── go.sum
├── middleware                   
│   └── default.go
├── utils                        // 共用工具
    ├── common.go
    ├── cronjob_handle.go
    ├── curl_request.go
    └── logging.go
├── docker-compose.yaml                  
└── mysql.sql                    // for init db
```
