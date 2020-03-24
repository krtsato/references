# Golang

- [環境構築](#環境構築)
  - [Homebrew によるインストール](#homebrew-によるインストール)
  - [VSCode の追加設定](#vscode-の追加設定)
- [プロジェクト管理](#プロジェクト管理)
  - [モジュール](#モジュール)
  - [Makefile](#makefile)

## 環境構築

### Homebrew によるインストール

goenv は導入せず, グローバルで小規模に最新の Golang を動かす．  
「みんなのGo言語 改訂2版」を参照．  

開発時には Docker で環境構築をする．  
今後, VSCode Remote Container を使用する予定．  
[VSCodeとDockerでMacにGolangの開発環境を作成する](https://dev.classmethod.jp/devenv/vscode-remote-containers-golang/)

```zsh
% brew install go
% which go # /usr/local/bin/go
% go version # go version go1.13.5 darwin/amd64
% go env GOROOT # /usr/local/Cellar/go/1.13.5/libexec
% go env GOPATH # /Users/Name/go
```

<br>
.zshrc に PATH を通す

```zsh
# for modules
export GO111MODULE=on

# for executable binary, such as go, godoc and gofmt
GOROOT=$(go env GOROOT)
export PATH=$GOROOT/bin:$PATH

# for external go packages
GOPATH=$(go env GOPATH)
export PATH=$GOPATH/bin:$PATH
```

<br>

### VSCode の追加設定

IDE から拡張機能 [vscode-go](https://github.com/microsoft/vscode-go) をインストールする．

![vcode-go](/images/golang/vscode-go.png)

<br>

IDE から [golangci-lint](https://github.com/golangci/golangci) をインストールする．  
デフォルトの Lint ツールを厳格化するために，golangci-lint を採用した．

![golangci-lint](/images/golang/golangci-lint.png)

<br>

ターミナルから [gopls](https://github.com/golang/tools/blob/master/gopls/doc/user.md) をインストールする．  
Go module の導入によって gopls が gocode, godef を代替する動きがあるため．  
[gocode やめます(そして Language Server へ)](https://mattn.kaoriya.net/software/lang/go/20181217000056.htm)

```zsh
# install gopls
% go get golang.org/x/tools/gopls@latest
```

<br>
VSCode の settings.json は以下のように記述した．

```json
{
  "//": "Linter 厳格化のため golangci-lint を採用",
  "go.lintTool": "golangci-lint",

  "//": "ファイル保存時に Lint を実行",
  "go.lintOnSave": "file",

  "//": "golangci-lint の中でも高速な Linter を実行",
  "go.lintFlags": ["--fast"],

  "//": [
    "言語サーバに gopls を採用",
    "関数の引数や構造体のフィールドも補完対象",
    "パッケージ不使用時に import 文を削除"
  ],
  "go.useLanguageServer": true,
  "gopls": {
    "usePlaceholders": true,
    "completeUnimported": true
  }
}
```

<br>

## プロジェクト管理

### モジュール

Modules を使うことで以下が容易になる

- パッケージ管理
- 依存関係の定義 (go.mod)
- バージョンロック (go.sum)

go get / go build / go test などを実行すると  
依存モジュールが抽出され自動追加されていく．

```zsh
# GOPATH の外でプロジェクトを作成
% mkdir リポジトリ名 && cd リポジトリ名
% git init
% git add .
% git commit -m "initial commit"
% git remote add origin https://github.com/ユーザ名/リポジトリ名.git

# モジュール管理を初期化
% go mod init github.com/ユーザ名/リポジトリ名

# モジュールの追加
## Go のソースコード内に import 文を書く
% go get

## バージョン指定
% go get github.com/path/to/pkg@vX.X.X

# モジュールの更新
## プロジェクト全体のモジュールを更新
## -u : インターネットを利用して更新
% go get -u

## パッケージごとに更新
% go get -u pkg

## パッチレベルで更新
% go get -u=patch

# モジュールの削除
## 依存に不要な/必要なモジュールの削除/追加
## v : 削除したモジュール情報を表示
% go mod tidy -v
```

<br>

### Makefile

環境構築, テスト, デプロイなどにおいて定型的なタスクを記述する．  
`% make setup` で記述したコマンドが実行される．
詳細は追記予定．
