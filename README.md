# 部活支援アプリ(web)

## 実行の流れ

```
docker-compose build
docker-compose run --rm react sh -c "create-react-app ./"
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