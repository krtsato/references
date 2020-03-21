# Ruby on Rails

やりたいことを列挙しただけなので  
Rails のコードを書きながらこちらも編集していきます．  
長文のためインデックスからジャンプすることを勧めます．

企業が運営する顧客管理サービスを開発する．  
ユーザは Admin / Staff / Customer を想定する．

- [環境構築](#環境構築)
- [トップページの作成](#トップページの作成)
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

### views ファイルの分離

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
+ Rails.application.config.assets.precompile += %w(staff.css admin.css customer.css)
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
- database.yml の production 設定はデプロイ先の環境に応じて今後変更する
- [ruby-rails-prac](https://github.com/krtsato/ruby-rails-prac) では Dockerfile の CMD でサーバを起動する
  - CMD で指定した start-rails-server.sh に追記してコンテナを restart する
  - 確認が終わったら追記箇所を元に戻す

```bash
$ bundle exec rails db:create RAILS_ENV=production

$ bundle exec rails assets:precompile

$ vim init_proj/start-rails-server.sh
# - bundle exec rails s -b "0.0.0.0" -p 3000
# + bundle exec rails s -b "0.0.0.0" -p 3000 -e production

$ docker-compose restart web
```

<br>

## エラーページの作成

### raise メソッド

- 例外を発生させる
- `raise 例外クラス名, 説明文`
  - 引数なしの場合は StandardError を継承する RuntimeError が発生する

<br>

### 例外処理の書き方

A ~ D は任意の Ruby コード  
E1, E2 は Exception クラスの子孫  
e1, e2 は 変数

```ruby
begin
  # A
rescue E1 => e1
  # B
rescue E2 => e2
  # C
ensure
  # D
end
```

- A で例外は発生した時点で A の処理は中断される
- そこで発生した例外オブジェクトが E1 のインスタンスならば e1 に格納
- 続けて B が実行される
- 同様に発生した例外オブジェクトが E2 のインスタンスならば e2 に格納
- 続けて C が実行される
- 例外オブジェクトが E1 / E2 のインスタンスでない場合は SystemError
- A での例外発生に関わらず最後に D が実行される

<br>

### クラスメソッド rescue_from

- アクション内で発生した例外の処理方法を指定する
- `rescue_from Forbidden, with: :rescue403`
  - Forbidden / Forbidden 子孫の例外が発生したとき
    - アクションを中止する
    - rescue403 メソッドを実行する

```ruby
private

def rescue403(e)
  @exception = e # 渡された例外オブジェクト
  render template: "errors/forbidden", status: 403
end
```

<br>

### 500 Internal Server Error

- スタイリングの管理
  - app/assets/stylesheets/shared/errors.scss を作成する
  - app/assets/stylesheets/*.css でアセットパイプラインの対象ディレクトリに設定する
    - `*=require_tree ./shared` する
- controller への追記
  - application_controller.rb に種々の例外処理メソッドを定義していく
  - `rescue500(e)` は views/errors/internal_server_error.rb を表示する
- view ファイルを作成する
- 意図的に例外を発生させる
  - controller の index アクションにおいて `raise` する

```ruby
class ApplicationController < ActionController::Base
+ rescue_from StandardError, with: :rescue500

  private

+ def rescue500(e)
+   render 'errors/internal_server_error', status: 500
+ end
end
```

<br>

### 403 Forbidden

- application_controller.rb に２種類の例外クラスを定義する
  - Forbidden : 権限不足により拒否
  - IpAddressRejected : IP アドレス制限により拒否
- ActionController::ActionControllerError
  - StandardError を継承する例外クラス
  - controller で発生する様々な例外の親 / 祖先クラス
- class 内部の class
  - ApplicationController がモジュールとしての役割を持ち名前空間を提供する
  - 原則的には Application::Forbidden のように呼び出す
- `rescue_from 例外クラス` の注意事項
  - 親子関係にある例外を指定する場合，親 / 祖先の例外を先に指定する
- view ファイルを作成する
  - インスタンス変数 `@exception` を view 側で使う
- 意図的に例外を発生させる
  - ApplicationController を継承した controller では省略形で呼び出せる
  - e.g. `raise Forbidden`

```ruby
class ApplicationController < ActionController::Base

+ class Forbidden < ActionController::ActionControllerError; end
+ class IpAddressRejected < ActionController::ActionControllerError; end

  rescue_from StandardError, with: :rescue500 # 先に StandardError を指定
+ rescue_from Forbidden, with: :rescue403
+ rescue_from IpAddressRejected, with: :rescue403

  private

+ def rescue403(e)
+   @exception = e
+   render template: 'errors/forbidden', status: 403
+ end
end
```

<br>

### 404 Not Found

- リソースが見つからない
  - ルーティングが存在しない
    - ActionController::RoutingError
  - DB に指定された条件に合うレコードが存在しない
    - ActiveRecord::RecordNotFound
- [rescue_from](#クラスメソッド-rescue_from) はアクションにおける例外を捕捉するメソッド
  - ルーティングの段階で発生する例外は捕捉できない

<br>

#### ActionController::RoutingError の処理

- `config.exceptions_app = -> (env)`
  - ミドルウェア ActionDispatch::ShowExceptions が Rails 外で発生した例外を env に投げる
  - env ハッシュオブジェクト
    - HTTP リクエストの情報がすべて含まれている
      - e.g. path_info, request_method
- `ActionDispathch::Request.new(env)`
  - ミドルウェア ActionDispathch の Request クラスが env を元に新たなハッシュオブジェクトを作る
  - `ErrorsController.action(action).call(env)` が例外用アクションを呼ぶ

```ruby
Rails.application.configure do
  config.exceptions_app = -> (env) do
    request = ActionDispatch::Request.new(env)

    action =
      case request.path_info
      when '/404'; :not_found
      when '/422'; :unprocessable_entity
      else; :internal_server_error
      end

    ErrorsController.action(action).call(env)
  end
end
```

- `bundle exec rails g controller errors` する
- error_controller.rb で例外処理に応じた view ファイルを指定する
- view ファイルを作る
- 存在しない URL を入力して意図的な例外を発生させる

```ruby
class ErrorsController < ApplicationController
  layout "staff"

  def not_found
    render status: 404
  end

  def unprocessable_entity
    render status: 422
  end

  def internal_server_error
    render status: 500
  end
end
```

<br>

#### ActiveRecord::RecordNotFound の処理

- DB にリクエストしたレコードがなかった場合も 404 にしてみる
- application_controller.rb に追記

```ruby
class ApplicationController < ActionController::Base

  rescue_from StandardError, with: :rescue500 # 先に StandardError を指定
+ rescue_from ActiveRecord::RecordNotFound, with: :rescue404

  private

+ def rescue404(e)
+   @exception = e
+   render 'errors/not_found', status: 404
+ end
```

<br>

### エラー処理の切り分け

- controllers/concerns/error_handlers.rb にモジュールとして切り分ける
- ActiveSupport::Concern によるモジュール化
  - `include do ... end` 内のメソッド
    - モジュールが読み込まれた直後に定義される
      - e.g. `scope ...` の定義
    - モジュールを読み込んだクラスのクラスメソッドになる
  - `class_methods ... end` 内のメソッドは，そのモジュールを読み込んだクラスのクラスメソッドになる
  - ブロックで囲まず定義したメソッドは，そのモジュールを読み込んだクラスのインスタンスメソッドになる
  - application_controller.rb で定義した例外処理用のクラスは，名前空間付きで呼ぶ
    - e.g. `rescue_from ApplicationController::Forbidden`
- dev 環境ではデバッグ目的でエラー表示を加工しない
  - `if Rails.env.production?`

```ruby
class ApplicationController < ActionController::Base
  # ...

  # 例外処理用のクラス
  class Forbidden < ActionController::ActionControllerError; end
  class IpAddressRejected < ActionController::ActionControllerError; end

  # モジュールの読み込み
+ include ErrorHandlers if Rails.env.production?
```

```ruby
module ErrorHandlers
  extend ActiveSupport::Concern

  # rescue_from は ApplicationController クラスのクラスメソッド
  included do
    rescue_from StandardError, with: :rescue500

    # 名前空間をつける
    rescue_from ApplicationController::Forbidden, with: :rescue403
    rescue_from ApplicationController::IpAddressRejected, with: :rescue403
    rescue_from ActiveRecord::RecordNotFound, with: :rescue404
    rescue_from ActionController::ParameterMissing, with: :rescue400
  end

  private

  # rescue404 は ApplicationController クラスのインスタンスメソッド
  def rescue404(e)
    render 'errors/not_found', status: 404
  end

  # rescue400, rescue403, rescue500 も同様
end
```

<br>

## サーバサイドにおけるユーザ認証の実装

### 初回マイグレーション

- staff の会員情報を管理する DB テーブル staff_members を作成する
- `bundle exec rails g model StaffMember` 単数形に注意
- マイグレーションスクリプトに追記
  - ブロック変数 `t` には TableDefinition オブジェクトがセットされる
  - このオブジェクトの各種メソッドがテーブルの定義を行う
  - index の設定
    - 検索 / ソートの高速化
      - メールアドレス
        - PostgreSQL の仕様でインデックスは大文字 / 小文字の区別あり
        - SQLの関数 `LOWER(email)` で小文字にする
        - 通常は `:email` のように指定する
      - 苗字，名前
        - フリガナでソートして一覧表示するとき効果的
- `bundle exec rails db:migrate` する
- `bundle exec rails r "StaffMember.columns.each {|c| p [c.name, c.type]}"` でカラム構成を確認
  - 主キーはデフォルト設定される `["id", :integer`

```ruby
class CreateStaffMembers < ActiveRecord::Migration[6.0]
  def change
    create_table :staff_members do |t|
      t.string :email, null: false                      # メールアドレス
      t.string :family_name, null: false                # 姓
      t.string :given_name, null: false                 # 名
      t.string :family_name_kana, null: false           # 姓（カナ）
      t.string :given_name_kana, null: false            # 名（カナ）
      t.string :hashed_password                         # パスワード
      t.date :start_date, null: false                   # 開始日
      t.date :end_date                                  # 終了日
      t.boolean :suspended, null: false, default: false # 無効フラグ

      t.timestamps
    end

    add_index :staff_members, "LOWER(email)", unique: true
    add_index :staff_members, [:family_name_kana, :given_name_kana]
  end
end
```

### パスワードのハッシュ化

- application_record.rb
  - `self.abstruct_class = true` で自身を抽象クラスにする
    - インスタンス化されない
  - models/staff_member.rb
    - `def password=(raw_passeord) ... end`
      - 要素代入関数の定義の仕方
      - 代入演算子 `=` を用いて引数を渡せる関数
      - `hoge = Hoge.new; hoge.raw_password = 'fuga'`
    - `BCrypt::Password.create(raw_passward)`
      - gem bcrypt を使ってハッシュ値を生成する

<br>

### seed データの投入

- seed データも DRY にする
  - db/seeds.rb で path を振り分ける
    - `%w()`  : 配列の要素をスペース区切りで指定
    - `require`  : 標準ライブラリ / 外部ファイル / 自作ファイルを読み込む
  - db/seeds/development/staff_members.rb に seed を書く
  - `bin/rails r "puts StaffMember.count"` で seed 投入を確認

```ruby
table_names = %w(staff_members)

table_names.each do |table_name|
  path = Rails.root.join("db", "seeds", Rails.env, "#{table_name}.rb")
  if File.exist?(path)
    puts "Creating #{table_name}..."
    require(path)
  end
end
```

```ruby
StaffMember.create!(
  email: "taro@example.com",
  family_name: "山田",
  given_name: "太郎",
  family_name_kana: "ヤマダ",
  given_name_kana: "タロウ",
  password: "password",
  start_date: Date.today
)
```

<br>

### session によるユーザ管理

- 名前空間 Staff を DRY にするため Staff::Base クラスを作る
  - controllers/staff/base.rb に共通処理を書く
  - controllers/staff/top_controller.rb に継承させる
  - 遅延初期化
    - StaffMember.find_by メソッドが多くても１回しか呼ばれない
  - session オブジェクトはクッキーの中に保持されている
  - `helper_method` でヘルパーメソッドとして登録する

```ruby
module Staff
- class TopController < ApplicationController
+ class TopController < Staff::Base
    # ...
  end
end
```

```ruby
module Staff
  class Base < ApplicationController
    private

    def current_staff_member
      if session[:staff_member_id]
        # 遅延初期化
        @current_staff_member ||= StaffMember.find_by(id: session[:staff_member_id])
      end
    end

    helper_method :current_staff_member
  end
end
```

<br>

### ログイン用のルーティング

- ログインする = session を新たに開始する
  - session というリソースを追加する = POST
- ログアウトする = session を終了する
  - session というリソースを削除する = DELETE
- `as: :login`
  - ルーティングに名前を付ける
  - :staff_login というシンボルを用いて URL パスを参照できる
- `resource :session, only: [:create, :destroy]` は以下のショートハンド
  - `post 'session' => 'session#create', as: :session`
  - `delete 'session' => 'session#destroy'`

| Task | HTTP method | URL path | Controller | Action |
| --- | --- | --- | --- | --- |
| ログインフォーム<br>を表示する | GET | /staff/login | staff/sessions | new |
| ログインする | POST | /staff/session | staff/sessions | create |
| ログアウトする | DELETE | /staff/session | staff/sessions | destroy |

```ruby
namespace :staff do
  root 'top#index'
+ get 'login' => 'sessions#new', as: :login
+ resource :session, only: [:create, :destroy]
end
```

<br>

## フロントエンドにおけるユーザ認証の実装

### ログイン / ログアウトリンクの設置

- views/shared/_header.html.erb を DRY にする
  - ユーザ認証はユーザの種類によって処理が異なる
  - _footer.html.erb は共通のまま
  - 各ドメインに shared/ を作成する
    - e.g. staff/shared/_header.html.erb
- views/layouts/staff.html.erb を編集する
  - `current_staff_member` は helper_method なので `@current_staff_member` を返す

```erb
- <%= render 'shared/header' %>
+ <%= render 'staff/shared/header' %>
```

```erb
<header>
  <span class="logo-mark">Ruby-Rails-RSpec-Prac</span>
  <%=
    if current_staff_member
      link_to 'ログアウト', :staff_session, method: :delete
    else
      link_to 'ログイン', :staff_login # デフォルトで GET 通信
    end
  %>
</header>
```

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
[Railsアプリの例外ハンドリングとエラーページの表示についてまとめてみた](https://qiita.com/upinetree/items/273ae574f1c021d24c37)  
[Rails の rescue_from で拾えない例外を exceptions_app で処理する](https://qiita.com/ma2ge/items/938d9f8f4839eb336318)  
[ActionDispatch ってなんだろう？](https://blog.eiel.info/blog/2014/03/30/action-dispatch/)  
[RailsのリクエストのライフサイクルとRackを理解する（翻訳）](https://techracho.bpsinc.jp/hachi8833/2019_10_03/77493)  
[ActiveSupport::Concern でハッピーなモジュールライフを送る](https://www.techscore.com/blog/2013/03/22/activesupportconcern-%E3%81%A7%E3%83%8F%E3%83%83%E3%83%94%E3%83%BC%E3%81%AA%E3%83%A2%E3%82%B8%E3%83%A5%E3%83%BC%E3%83%AB%E3%83%A9%E3%82%A4%E3%83%95%E3%82%92%E9%80%81%E3%82%8B/)  
[Rails 4.2からはmodule ClassMethodsではなくConcern#class_methodsを使おう](https://blog.yujigraffiti.com/2015/01/rails-42module-classmethodsconcernclass.html)  
[Ruby on Rails 6 実践ガイド](https://www.oiax.jp/jissen_rails6)
