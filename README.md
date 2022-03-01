-
## Server結構：
```
├── README.md
│
│
├── config                        # 相關設定
│        ├── app                       # 各環境相關連線設定
│        └── notify                    # telegrem通知設定
│
├── delivery                      # 接入
│
├── models                        # 資料儲存體
│
├── process                       # 共通邏輯
│        
│
├── service                       # Server提供的不同服務
│        └── server                    # 伺服器
│
├── utils
│        └── utils.go                  # 共用工具
│
├── .env                          # 環境變數定義
├── go.mod                        # 套件模組管理的檔案 (自動生成的)
├── go.sum                        # 套件模組管理的檔案 (自動生成的)
├── main.go                       # 程式進入點
 