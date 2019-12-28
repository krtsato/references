# Golang

## 環境構築

### Homebrew によるインストール

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

### VSCode の追加設定

IDE から拡張機能 [vscode-go](https://github.com/microsoft/vscode-go) をインストールする．

![vcode-go](/images/vscode-go.png)

IDE から [golangci-lint](https://github.com/golangci/golangci) をインストールする．  
デフォルトの Lint ツールを厳格化するために，golangci-lint を採用した．

![golangci-lint](/images/golangci-lint.png)

ターミナルから [gopls](https://github.com/golang/tools/blob/master/gopls/doc/user.md) をインストールする．  
Go module の導入によって gopls が gocode, godef を代替する動きがあるため．  
[gocode やめます(そして Language Server へ)](https://mattn.kaoriya.net/software/lang/go/20181217000056.htm)

```zsh
# install gopls
% GO111MODULE=on go get golang.org/x/tools/gopls@latest
```

VSCode の settings.json は以下のように記述した．

```json
{
  "go.lintTool": "golangci-lint", // Linter 厳格化のため golangci-lint を採用
  "go.lintOnSave": "file", // ファイル保存時に Lint を実行
  "go.lintFlags": [
    "--fast" // golangci-lint の中でも高速な Linter を VSCode 上で実行
  ],
  "go.useLanguageServer": true, // 言語サーバに gopls を採用
  "gopls": {
    "usePlaceholders": true, // 関数の引数や構造体のフィールドも補完対象
    "completeUnimported": true // パッケージ不使用時に import 文を削除
  }
}
```

次は Makefile を作る
