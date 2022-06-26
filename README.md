# learning-thema

[Thema](https://github.com/grafana/thema) のチュートリアルに従って作ってみた。

# 前提
cue, golangがインストールされていること

# 環境

```sh
soharaki@DESKTOP:~/work/learning-thema$ cue version
cue version v0.4.3 linux/amd64
soharaki@DESKTOP:~/work/learning-thema$ go version
go version go1.18.3 linux/amd64
```

# 動かし方

```
# cueの初期化
MODPATH="github.com/example/thema_example"
cue mod init $MODPATH
# 念のため
go mod tidy
```