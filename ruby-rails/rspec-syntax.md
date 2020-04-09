# RSpec

- [RSpec の思想](#rspec-の思想)
- [初期設定](#初期設定)
- [実行](#実行)
- [RSpec のメソッド](#rspec-のメソッド)
  - [example](#example)
  - [describe](#describe)
  - [expect](#expect)
  - [context](#context)
  - [before](#before)
  - [let](#let)
  - [let!](#let-1)
  - [subject](#subject)
  - [shared_examples](#shared_examples)
  - [shared_context](#shared_context)
  - [pending](#pending)
  - [xexample](#xexample)
  - [skip](#skip)
  - [get・post・patch・delete](#getpostpatchdelete)
- [計画的テストコード](#計画的テストコード)
  - [テストケースだけを先に書いておく](#テストケースだけを先に書いておく)
  - [shared_example・shared_context の分離](#shared_exampleshared_context-の分離)
- [Factory Bot](#factory-bot)
  - [導入](#導入)
  - [Factory Bot のメソッド](#factory-bot-のメソッド)
    - [attributes_for](#attributes_for)
    - [create](#create)
- [参考文献](#参考文献)

<br>

## RSpec の思想

- ビヘイビア駆動開発を Ruby で実践するために作られた
  - Behavior Driven Development : BDD
- BDD ではテストコードによってソフトウェアの仕様が定義される
  - テストコードが仕様書代わりになる
  - 仕様書を作り込むより実例を列挙する

<br>

## 初期設定

```zsh
% bundle exec rails g rspec:install
```

- ファイル名は `_spec.rb`
- spec ファイルは spec/ 配下に分類して配置する
  - 大まかな慣習がある
    - controllers : spec/controllers
    - models : spec/models
    - API 通信 : spec/requests

<br>

## 実行

```zsh
# ファイル指定
% bundle exec rspec spec/path/to/hoge_spec.rb

# 行番号指定
% bundle exec rspec spec/path/to/hoge_spec.rb:行番号

# spec/ 配下の spec ファイルをすべて実行
% bundle exec rspec

# タグ指定
% bundle exec rspec --tag=TAG_NAME
```

タグ指定の場合 spec ファイルの example にシンボルを付与する

```ruby
# tag_name がタグとして機能する
example "テストケース名", :tag_name do
  # ...
end
```

<br>

## RSpec のメソッド

### example

- いわゆるテストケースを作るメソッド
  - ある機能の前提条件と結果をコードで表現したもの
- 別名 : it, specify

```ruby
example "appends a character" do
  s = "ABC"
  s << "D"
  expect(s.size).to eq(4)
end
```

<br>

### describe

- example をまとめたもの
- describe メソッドの引数
  - クラス
  - 文字列
    - `#` はインスタンスメソッドであることを示す慣習
- 入れ子にすることができる

```ruby
# String クラスに関する仕様
RSpec.describe String do
  # その中の << メソッド
  describe "#<<" do
    # << メソッドによるテストケース
    example "文字の追加" do
      s = "ABC"
      s << "D"
      expect(s.size).to eq(4)
    end

    example ...
    example ...
  end
end
```

<br>

### expect

- オブジェクトを対象に expect する場合
  - `expect(T).to M`
    - T : ターゲットオブジェクト
      - response
      - など
    - M : マッチャーオブジェクト
      - eq
      - be
      - be_xxx
      - be_truthy / be_falsey
      - change + from / to / by
      - 配列 + include
      - raise_error
      - redirect_to
      - route_to
      - be_routable
      - be_within + of
      - など
    - T と M を比較するメソッド
      - to
      - not_to (= to_not)
- ブロックを対象に expect する場合
  - `expect {}.to M`
    - ブロック `{}` 内では式評価などを行う

<br>

### context

- describe と機能的には同じ
- describe より粒度が細かい分類に用いる場合が多い
  - 条件や状況に応じて結果が変化する場合など

```ruby
RSpec.describe User do
  describe '#greet' do
    context '12歳以下の場合' do
      example 'ひらがなで答える' do
        user = User.new(name: 'たろう', age: 12)
        expect(user.greet).to eq 'はじめまして'
      end
    end

    context '13歳以上の場合' do
      example '漢字で答える' do
        user = User.new(name: 'たろう', age: 13)
        expect(user.greet).to eq '初めまして'
      end
    end
  end
end
```

<br>

### before

- テストを実行前の共通処理やデータのセットアップ等を行う
- example の実行前に毎回呼ばれる
- ネストした describe や context ごとに用意できる
  - 親子関係に応じて before が順番に呼ばれる

```ruby
RSpec.describe User do
  describe '#greet' do
    before do
      @params = {name: 'たろう'}
    end

    context '12歳以下の場合' do
      example 'ひらがなで答える' do
        user = User.new(@params.merge(age: 12))
        expect(user.greet).to eq 'はじめまして'
      end
    end

    context '13歳以上の場合' do
      example '漢字で答える' do
        user = User.new(@params.merge(age: 13))
        expect(user.greet).to eq '初めまして'
      end
    end
  end
end
```

<br>

#### before の処理を適用しない

- example を引数としてメタデータを持っていたら適用しない
  - `example.metadata[:skip_before]`
  - before の中で `return if example.metadata[:skip_before]` はできない

```ruby
RSpec.describe 'User アカウント', type: :request do
  before do |example|
    unless example.metadata[:skip_before]
      # ...
    end
  end

  describe 'hoge' do
    example 'fuga', :skip_before do
      # ...
    end
  end
end
```

### let

- インスタンス変数やローカル変数を置き換える
- `let(:hoge) {fuga}`
  - fuga を hoge として参照できる
  - fuga がハッシュリテラルの場合, `{{param: val}}` のように二重カーリになる
- テストコードをトップダウンで構造化できる点が好ましい
  - 遅延評価されるため，必要になる瞬間まで呼ばれない
    - `expect(user.greet)` -> user とは
    - `let(:user) {User.new(name: 'たろう', age: age)}` -> age とは
    - `let(:age) {12}` が呼ばれる
    - `expect(User.new(name: 'たろう', age: 12).greet).to` が呼ばれる

```ruby
RSpec.describe User do
  describe '#greet' do
    let(:user) {User.new(name: 'たろう', age: age)}

    context '12歳以下の場合' do
      let(:age) {12}
      example 'ひらがなで答える' do
        expect(user.greet).to eq 'はじめまして'
      end
    end

    context '13歳以上の場合' do
      let(:age) {13}
      example '漢字で答える' do
        expect(user.greet).to eq '初めまして'
      end
    end
  end
end
```

<!-- markdownlint-disable no-trailing-punctuation -->

### let!

<!-- markdownlint-enable no-trailing-punctuation -->

- 事前実行される
  - let を before 内で定義するショートハンド
- let の遅延評価によるテスト失敗を回避するため

まず let を使いテストが失敗する例

```ruby
RSpec.describe Blog do
  let(:blog) {Blog.create(title: 'RSpec', content: 'やっていき')}

  example 'ブログの取得ができる' do
    # Blog.first が呼ばれた時点では let(:blog) が未実行
    # レコードが DB に保存されていない
    expect(Blog.first).to eq blog
  end
end
```

次に let! を使いテストが成功する例

```ruby
RSpec.describe Blog do
  let!(:blog) {Blog.create(title: 'RSpec', content: 'やっていき')}

  example 'ブログの取得ができる' do
    # let! によって事前実行されたレコードが比較される
    expect(Blog.first).to eq blog
  end
end
```

<br>

### subject

- テスト対象のオブジェクト / メソッドの実行結果が１つに定まる場合，テストコードを DRY にできる
- expect の形が `is_expected.to` に変化する
- example ~ do のネストを省略すると可読性が高まる
  - 小中規模のテストコードには有効
  - テストコードの肥大化が不可避の場合はつらそう
    - 記述省略分をネストをかけずに補える
    - [shared_examples](#shared_examples) を使えば良さげ

```ruby
RSpec.describe User do
  describe '#greet' do
    let(:user) {User.new(name: 'たろう', age: age)}
    subject {user.greet}

    context '12歳以下の場合' do
      let(:age) {12}
      example {is_expected.to eq 'はじめまして'}
    end

    context '13歳以上の場合' do
      let(:age) {13}
      example {is_expected.to eq '初めまして'}
    end
  end
end
```

<br>

### shared_examples

- example を再利用するメソッド
- 有効にはたらく場合
  - 同じ結果になる複数のテストケースがある
  - それらをより抽象的な区分で定義する
  - subject での記述省略分をネストをかけず補う
- `it_behaves_like 'SHARED_EXAMPLE_NAME'`
  - it_behaves_like が SHARED_EXAMPLE_NAME を呼び出す
  - SHARED_EXAMPLE_NAME に該当する shared_example を参照する
  - ここで it {is_expected.to ...} を実行する

まず shared_example を使用しない例

```ruby
RSpec.describe User do
  describe '#greet' do
    let(:user) {User.new(name: 'たろう', age: age)}
    subject {user.greet}

    context '2歳の場合' do
      let(:age) {2}
      it {is_expected.to eq 'はじめまして'}
    end
    # ... more context ...
    context '12歳の場合' do
      let(:age) {12}
      it {is_expected.to eq 'はじめまして'}
    end

    context '13歳の場合' do
      let(:age) {13}
      it {is_expected.to eq '初めまして'}
    end
    # ... more context ...
    context '113歳の場合' do
      let(:age) {113}
      it {is_expected.to eq '初めまして'}
    end
  end
end
```

次に shared_example を使用する例

```ruby
RSpec.describe User do
  describe '#greet' do
    let(:user) {User.new(name: 'たろう', age: age)}
    subject {user.greet}

    shared_examples 'ひらがなのあいさつ' do
      it {is_expected.to eq 'はじめまして'}
    end
    context '2歳の場合' do
      let(:age) {2}
      it_behaves_like 'ひらがなのあいさつ'
    end
    # ... more context ...
    context '12歳の場合' do
      let(:age) {12}
      it_behaves_like 'ひらがなのあいさつ'
    end

    shared_examples '漢字の挨拶' do
      it {is_expected.to eq '初めまして'}
    end
    context '13歳の場合' do
      let(:age) {13}
      it_behaves_like '漢字の挨拶'
    end
    # ... more context ...
    context '113歳の場合' do
      let(:age) {100}
      it_behaves_like '漢字の挨拶'
    end
  end
end
```

<br>

#### shared_examples のファイル分割

- spec/support 配下に共有エグザンプルを配置する
- rails_helper.rb で Rails が読み込むように設定する
- spec ファイルから `include_examples 'shared_examples_name', 'controller/name'` する
  - spec ファイルから `controller/name` を渡すことでエグザンプルを共有する
- 後述の shared_context も同様に `include_context 'shared_examples'` する

```ruby
Dir[Rails.root.join('spec/support/**/*.rb')].sort.each {|f| require f}
```

```ruby
RSpec.describe 'User アカウント', type: :request do
  context 'when ログイン前' do
    include_examples 'a blocked user controller', 'user/accounts'
  end
  # ...
end
```

```ruby
shared_examples 'a blocked user controller' do |controller|
  let(:args) {host: hoge, controller: controller}

  describe '#show' do
    example 'ログインフォームにリダイレクト' do
      get url_for(args.merge(action: :show))
      expect(response).to redirect_to(user_login_url)
    end
  end
end
```

### shared_context

- context を再利用するメソッド
- 有効にはたらく場合
  - 同じ条件による複数の describe (= example group) がある
  - それらをより広い区分で定義する
- `include_context 'SHARED_CONTEXT_NAME'`
  - include_context が SHARED_CONTEXT_NAME を呼び出す
  - SHARED_CONTEXT_NAME に該当する shared_context を参照する
  - ここに context 固有の処理を書いておく

まず shared_context を使用しない例

```ruby
RSpec.describe User do
  describe '#greet' do
    let(:user) {User.new(name: 'たろう', age: age)}
    subject {user.greet}

    context '12歳以下の場合' do
      let(:age) {12}
      it {is_expected.to eq 'はじめまして'}
    end
    context '13歳以上の場合' do
      let(:age) {13}
      it {is_expected.to eq '初めまして'}
    end
  end

  # 12 歳以下は true とするメソッドを
  # User クラスに作成したとする
  describe '#child?' do
    let(:user) {User.new(name: 'たろう', age: age)}
    subject {user.child?}

    context '12歳以下の場合' do
      let(:age) {12}
      it {is_expected.to eq true}
    end

    context '13歳以上の場合' do
      let(:age) {13}
      it {is_expected.to eq false}
    end
  end
end
```

次に shared_context を使用する例

```ruby
RSpec.describe User do
  let(:user) {User.new(name: 'たろう', age: age)}
  shared_context '12歳の場合' do
    let(:age) {12}
  end
  shared_context '13歳の場合' do
    let(:age) {13}
  end
  # ... more context ...

  describe '#greet' do
    subject {user.greet}
    context '12歳以下の場合' do
      include_context '12歳の場合'
      # ... more context ...
      it {is_expected.to eq 'はじめまして'}
    end
    context '13歳以上の場合' do
      include_context '13歳の場合'
      # ... more context ...
      it {is_expected.to eq '初めまして'}
    end
  end

  describe '#child?' do
    subject {user.child?}
    context '12歳以下の場合' do
      # ... more context ...
      include_context '12歳の場合'
      it {is_expected.to eq true}
    end
    context '13歳以上の場合' do
      # ... more context ...
      include_context '13歳の場合'
      it {is_expected.to eq false}
    end
  end
end
```

<br>

### pending

- 解消できないエラーを保留するために使う
- 保留とした理由などを書く

```ruby
example "appends nil" do
  pending("nil の仕様を調査中")
  s = "ABC"
  s << nil
  expect(s.size).to eq(4)
end
```

<br>

### xexample

- pending を書くのが面倒な場合 xexample メソッドを定義する
- テストケースを一時的に無効にできる

```ruby
xexample "appends nil" do
  s = "ABC"
  s << nil
  expect(s.size).to eq(4)
end
```

<br>

### skip

- 任意の箇所でテストケースを終了させる
- skip 以降は実行せず pending としてマークする

```ruby
describe '実行したくないクラス' do
  example '実行したくないテスト' do
    expect(1 + 2).to eq 3

    skip 'とりあえずここで実行を保留'

    # ここから先は実行されない
    expect(hoge).to eq fuga
  end
end
```

<br>

### get・post・patch・delete

- 第１引数  の URL のフルパスに対して，第２引数のパラメータを送信する
- `response`
  - `ActionController::TestResponse` オブジェクトを返すメソッド
  - アクションの実行結果に対する情報を保持している

```ruby
example 'fuga へリダイレクトする' do
  post hoge_full_url, params: {request_hash: {k: v, ...}}
  expect(response).to redirect_to(fuga_full_url)
end
```

- 認可のテストを書く場合 `before do ... end` と組み合わせるのが定石
- example 実行直前に email, password を送信して認可されたユーザをを作る

```ruby
before do
  post user_session_url, params: {
    user_login_form: {email: user.email, password: 'password'}
  }
end
```

<br>

## 計画的テストコード

### テストケースだけを先に書いておく

- TDD はしないがテストケースを緩く書いておく
  - テストケースの中で expect しない
- RSpec 上で仕様を設計することで，実装前の Todo リストになる

```ruby
RSpec.describe User do
  describe '#greet' do
    context '12歳以下の場合' do
      it 'ひらがなであいさつする'
    end

    context '13歳以上の場合' do
      it '漢字で挨拶する'
    end
  end
end
```

<br>

### shared_example・shared_context の分離

- example・context をファイルを跨いで共有する
- spec/rails_helper.rb に追記する
  - `Dir[Rails.root.join("spec", "support", "**", "*.rb")].each {|f| require f}`
  - spec/support/ 配下のファイルが自動的に読み込まれるようになる

```ruby
require 'rspec/rails'
+ Dir[Rails.root.join("spec", "support", "**", "*.rb")].each {|f| require f}
```

## Factory Bot

### 導入

- DB にテストデータを投入するための gem
- spec/rails_helper.rb に `config.include FactoryBot::Syntax::Methods` を追記する
- spec/factories/ 配下で factory ファイルを作成する
  - `FactoryBot.define do ... end` の中で factory を定義する
  - spec ファイルから `build` で呼び出せるようにシンボルを設定する

```ruby
FactoryBot.define do
  factory :user do # シンボルを設定する
    sequence(:email) {|n| "member#{n}@example.com"} # n 連番で作成される
    password {'password'}
    # ...
  end
end
```

```ruby
RSpec.describe User::Authenticator do
  describe '#authenticate' do
    example '正しいパスワードならば true を返す' do
      u = build(:user) # シンボルで呼び出す
      expect(described_class.new(u).authenticate('password')).to be_truthy
    end
  end
end
```

<br>

### Factory Bot のメソッド

#### attributes_for

- 定義した factory が持つハッシュオブジェクトを返却する
- Rails の controller で `assign_attributes` メソッドの引数としてそのまま使用できる
- update アクションで属性を一括設定する場合のテストで便利

```ruby
# params_hash は {email: 'member1@example.com', password: 'password', ...} となる
let(:params_hash) {attributes_for(:user)}
```

<br>

#### create

- 定義した factory を用いてモデルオブジェクトを作成する
- DB への保存に成功したそのオブジェクトを返す

```ruby
# user は定義済みの factory
let(:user) {create(:user)}
```

<br>

## 参考文献

[Relish Publisher RSpec](https://relishapp.com/rspec/)  
[使える RSpec 入門・その 1「RSpec の基本的な構文や便利な機能を理解する」](https://qiita.com/jnchito/items/42193d066bd61c740612)  
[使える RSpec 入門・その 2「使用頻度の高いマッチャを使いこなす」](https://qiita.com/jnchito/items/2e79a1abe7cd8214caa5)  
[Ruby on Rails 6 実践ガイド](https://www.oiax.jp/jissen_rails6)
