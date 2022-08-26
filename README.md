# 部活支援アプリ(web)

## 実行の流れ

```
docker-compose build
docker-compose up 
```

結構時間かかる

## 構築メモ
### go(gin)のdocker構築
yml fileのcommand: /bin/bashのコメントを外す。
Dockerfileの"added"以降をコメントアウト。
``` 
docker-compose build
docker-compose up go
# cmdで実行
docker exec -it go_container /bin/sh
# bash
go mod init test
go get -u github.com/gin-gonic/gin
```
yml fileのcommand:~をコメントアウト
Dockerfileの"added"以降のコメントを外す
```
docker-compose build
docker-compose up go
```


#### メモ
React側でGinのREST APIを叩く設定
1. 以下の拡張機能追加
https://chrome.google.com/webstore/detail/cors-unblock/lfhmikememgdcahcdlaciloancbhjino/related
2. 有効(cがオレンジの状態)にする

- 基本的には別ポート番号にchromeからアクセスすることはセキュリティ上できないらしい
- この拡張機能を使ってそれを可能にしてるっぽい
- ユーザー数は世界に10万人以上いるっぽいし、評価もそんなに低くないからおそらく大丈夫



