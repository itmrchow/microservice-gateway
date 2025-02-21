# microservice-gateway

## Architecture
- delivery (對外層)
  - 負責處理 HTTP/gRPC 請求
  - 處理請求參數驗證
  - 回應格式轉換
  - 不包含業務邏輯
- usecase (業務邏輯)
  - 實現所有業務邏輯
  - 協調不同 repository 的操作
  - 處理事務管理
  - 不依賴具體實現細節
- entities (領域模型)
  - 定義核心業務實體
  - 包含實體的基本驗證規則
  - 不包含資料庫相關邏輯
  - 最純粹的業務規則
- repository（資料存取層）
  - 處理資料持久化
  - 封裝資料庫操作細節
  - 提供資料存取介面
  - 可以切換不同的資料來源
- infrastructure（基礎設施）
  - 提供通用功能
  - 不包含業務邏輯

### Gateway
- gateway
  - pkg
    - [x] mux
    - [ ] log
    - [x] config
    - [ ] swagger
    - [ ] jwt
    - [ ] grpc
    - [ ] wire
  - feature
    - [x] health check API
    - [ ] error handler
      - [ ] 400 api
      - [ ] 500 api
    - [ ] input output log
    - [ ] log init
    - [ ] login API
      - jwt
      - test
    - [ ] log output
    - [ ] rate limit
    - [ ] OAuth?

### Micro Services
- user
  - pkg
    - [ ] mysql
    - [ ] gorm
    - [ ] grpc
  - feature
    - [ ] login func
      - test

### Common pkgs
- common
- protobuf

