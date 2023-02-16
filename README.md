# CleanArchitectureHexagonal

## 預期

使用六角架構搭配DDD的貧血模型

## Tag

capter3_FolderStruct: 第三章基本的框，無程式碼  
v1: 連結main.go與company_gin_adapter.go  
    `http://localhost:8080/get_siteinfo`  
v2: 連結company_gin_adapter.go與get_site_info_service.go (透過port.in get_site_info_usecase.go  
    `http://localhost:8080/get_siteinfo?site_id=1`  
    `http://localhost:8080/get_siteinfo?site_id=2`  
v3: get_site_info_service.go使用domain的site.go的field  
v4: 連結get_site_info_service.go與company_leveldb_adapter.go (透過port.out load_company_port.go  
v5: 第四章，將原本由adapter.in傳入到application.port.in 的參數改成物件(command)傳遞，確保驗証
