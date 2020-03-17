# Ruby on Rails

やりたいことを列挙しただけなので  
Rails のコードを書きながらこちらも編集していきます．

企業が運営する顧客管理サービスを開発する．  
ユーザは Admin / Staff / Customer を想定する．

- [環境構築](#環境構築)
- [トップページの作成](#トップページの作成)
  - [routes to root の定義](#routes-to-root-の定義)
  - [名前空間の controllers の作成](#名前空間の-controllers-の作成)
  - [view ファイルの分離](#view-ファイルの分離)
  - [部分テンプレートの表示](#部分テンプレートの表示)
  - [ヘルパーメソッドの定義](#ヘルパーメソッドの定義)
  - [アセットパイプライン](#アセットパイプライン)
  - [スタイルシートの分離](#スタイルシートの分離)
  - [アセットのプリコンパイル](#アセットのプリコンパイル)
  - [controllers のレイアウト選択](#controllers-のレイアウト選択)
  - [production モードでの起動](#production-モードでの起動)
- [エラーページの作成](#エラーページの作成)
- [サーバサイドにおけるユーザ認証の実装](#サーバサイドにおけるユーザ認証の実装)
- [フロントエンドにおけるユーザ認証の実装](#フロントエンドにおけるユーザ認証の実装)
- [ルーティングの設定](#ルーティングの設定)
- [Admin による Staff アカウント CRUD の実装](#admin-による-staff-アカウント-crud-の実装)
- [マスアサインメント脆弱性に対するセキュリティ強化](#マスアサインメント脆弱性に対するセキュリティ強化)
- [Staff アカウントによる自身の CRUD 実装](#staff-アカウントによる自身の-crud-実装)
- [Admin および Staff アカウントにおけるアクセス制御の実装](#admin-および-staff-アカウントにおけるアクセス制御の実装)
- [Admin による Staff アカウントの ログイン / ログアウト記録閲覧の実装](#admin-による-staff-アカウントの-ログイン--ログアウト記録閲覧の実装)
- [DB 格納前の正規化とバリデーションの実装](#db-格納前の正規化とバリデーションの実装)
- [プレゼンタによるフロントエンドのリファクタ](#プレゼンタによるフロントエンドのリファクタ)
- [Customer アカウントの CRUD 実装](#customer-アカウントの-crud-実装)
- [Capybara およびバリデーションによる Customer アカウントの CRUD リファクタ](#capybara-およびバリデーションによる-customer-アカウントの-crud-リファクタ)
- [ActiveSupport::Concern による機能共通化を目的としたリファクタ](#activesupportconcern-による機能共通化を目的としたリファクタ)
- [Customer アカウントにおける自宅住所と勤務先の任意入力の実装](#customer-アカウントにおける自宅住所と勤務先の任意入力の実装)
- [Customer アカウントにおける電話番号の CRUD 実装](#customer-アカウントにおける電話番号の-crud-実装)
- [参考文献](#参考文献)

ソースコード : [ruby-rails-prac](https://github.com/krtsato/ruby-rails-prac)

<br>

## 環境構築

- [ruby-rails-prac](https://github.com/krtsato/ruby-rails-prac) を要件に応じてアレンジする
- Rails API の構築情報は今後追加

<br>

## トップページの作成

### routes to root の定義

- routes.rb で名前空間 `namespace` を定義する
  - admin
  - staff
  - customer
- この定義に基づいて controllers や views のドメイン分割をしていく

<br>

### 名前空間の controllers の作成

- サービスがスケールした場合にクラス名やメソッド名の重複を防止する
- `bundle exec rails g controller admin/top`
- `class Admin::TopController < ApplicationController`
  - Admin モジュールにおける
  - Top についての controller クラスは
  - ApplicationController を継承する
- Rubocop に注意されるので `module ... end` を明記する
- staff / customer も同様

<br>

### view ファイルの分離

- views も DRY にする
- views/layouts/applications.rb を削除
- views/layouts/ 配下でドメインを分割
  - admin.html.erb (.slim)
  - staff.html.erb (.slim)
  - customer.html.erb (.slim)

<br>

### 部分テンプレートの表示

- 各ドメインの ERB からヘッダーとフッターを呼ぶ
- 部分テンプレートのディレクトリは views/shared/\_hoge.html.erb (.slim)
  - 接尾辞 `_` を付ける慣習がある

<br>

### ヘルパーメソッドの定義

- ERB などのテンプレート内で使用できるメソッド
  - views/helpers/ 配下に定義する
- head タグ内の title にアプリ名を表示する
- ブラウザのタブが表示中のページタイトルを反映する

<br>

### アセットパイプライン

- Rails アプリが JS, CSS, 画像ファイルを管理する
- ディレクティブ記法
  - rails new のあと app/assets/stylesheets/application.css に見られる
  - `*= require_tree .` : アセットとして見なす範囲を設定している
  - `*= require_self` : この宣言が書かれたファイルをアセットに含む

<br>

### スタイルシートの分離

- スタイリングも DRY にする
- application.css を削除して app/assets/stylesheets 配下でドメインを分割
  - admin.css が `*= require_tree ./admin` する
    - admin/hoge.scss でスタイリング
    - staff / customer も同様
- SCSS での変数定義を別ファイルで行う
  - e.g. 色を変数で表す
    - app/assets/stylesheets/admin/\_colors.scss
    - 接尾辞 `_` を付ける慣習がある
    - admin/hoge.scss で `@import "colors";` する

<br>

### アセットのプリコンパイル

- 分離したアセットのエントリポイントを含める
- config/initializers/assets.rb に追記

```ruby
Rails.application.config.assets.precompile += %w(staff.css admin.css customer.css)
```

<br>

### controllers のレイアウト選択

- ドメインごとのレイアウト表示を controllers に反映させる
- 通常は controller と同名のレイアウトが選択される
- それが無い場合 application という名前のレイアウトが選択される
  - e.g. 現状 admin/top という controller 名から app/views/layouts/admin/top.html.erb が選択される
  - 今後 admin/hoge という controller がアクションを実行するとき
  - app/views/layouts/admin.html.erb をエントリポイントにしたい
- app/controllers/application_controller.rb に追記
  - 正規表現を使ってレイアウト選択を一般化する
    - \A : 文字列の先頭
    - admin/ または sraff/ または customer/
    - Regexp.last_match : マッチした文字列情報を持つ MatchData オブジェクトを返す
      - 正規表現の１番目の括弧でマッチした文字列を指定
  - 結果的に admin または staff または customer を返す

```ruby
class ApplicationController < ActionController::Base
  layout :set_layout

  private

  def set_layout
    if params[:controller] =~ %r{\A(admin|staff|customer)/}
      Regexp.last_match[1]
    else
      'customer'
    end
  end
end
```

<br>

### production モードでの起動

- 早めに確認しておく
- DB の production 設定はデプロイ先の環境に応じて今後変更する

```bash
$ bundle exec rails db:create RAILS_ENV=production

% bundle exec rails assets:precompile

% bundle exec rails s -e production -b 0.0.0.0
```

<br>

## エラーページの作成

<br>

## サーバサイドにおけるユーザ認証の実装

<br>

## フロントエンドにおけるユーザ認証の実装

<br>

## ルーティングの設定

<br>

## Admin による Staff アカウント CRUD の実装

<br>

## マスアサインメント脆弱性に対するセキュリティ強化

<br>

## Staff アカウントによる自身の CRUD 実装

<br>

## Admin および Staff アカウントにおけるアクセス制御の実装

<br>

## Admin による Staff アカウントの ログイン / ログアウト記録閲覧の実装

<br>

## DB 格納前の正規化とバリデーションの実装

<br>

## プレゼンタによるフロントエンドのリファクタ

<br>

## Customer アカウントの CRUD 実装

<br>

## Capybara およびバリデーションによる Customer アカウントの CRUD リファクタ

<br>

## ActiveSupport::Concern による機能共通化を目的としたリファクタ

<br>

## Customer アカウントにおける自宅住所と勤務先の任意入力の実装

<br>

## Customer アカウントにおける電話番号の CRUD 実装

<br>

## 参考文献

[Ruby の Module の使い方とはいったい](https://qiita.com/shiopon01/items/fd6803f792398c5219cd)  
[Ruby on Rails 6 実践ガイド](https://www.oiax.jp/jissen_rails6)
