# Golang

## 環境構築

goenv は導入せず, グローバルで小規模に Golang を動かす．  
開発時には Docker で環境構築をする．  
「みんなのGo言語 改訂2版」を参照．

```zsh
% brew install go
% which go # /usr/local/bin/go
% go version # go version go1.13.5 darwin/amd64
% go env GOROOT # /usr/local/Cellar/go/1.13.5/libexec
% go env GOPATH # /Users/Name/go
```

.zshrc に PATH を通す

```zsh
# for executable binary, such as go, godoc and gofmt
GOROOT=$(go env GOROOT)
export PATH=$GOROOT/bin:$PATH

# for external go packages
GOPATH=$(go env GOPATH)
export PATH=$GOPATH/bin:$PATH
```

---
VSCode の設定をどうするか．  
以下の機能をほぼ満たせる環境にしたい．

- goimports
- go vet
- golint
- gorename
- guru
- gopls (gocode, godefの代替?)
