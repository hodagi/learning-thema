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

* 準備

```
# cueの初期化
MODPATH="github.com/example/thema_example"
cue mod init $MODPATH
# 念のため
go mod tidy
```

* 実行

```sh
go test
```

* リザルト

```sh
soharaki@DESKTOP:~/work/tutorial-thema$ go test
#######################################
version0.0:{Firstfield:foo}
version1.0:{Firstfield:foo Secondfield:-1}
Lacuna(欠落したフィールド情報)
{[] [{secondfield <nil>}] 0 -1 used as a placeholder value - replace with a real value before persisting!}
#######################################
PASS
ok      github.com/example/thema_example        0.191s
```