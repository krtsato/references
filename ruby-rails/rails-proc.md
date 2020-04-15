# Ruby on Rails

Rails アプリケーションを構築する記録．

Rails のコードを書きながらこちらも編集していきます．  
長文のためインデックスからジャンプすることを勧めます．

企業が運営する顧客管理サービスを開発する．  
ユーザは Admin・Staff・Customer を想定する．

- [環境構築](#環境構築)
- [トップページの作成](#トップページの作成)
  - [routes to root の定義](#routes-to-root-の定義)
  - [名前空間の controllers の作成](#名前空間の-controllers-の作成)
  - [views ファイルの分離](#views-ファイルの分離)
  - [部分テンプレートの表示](#部分テンプレートの表示)
  - [ヘルパーメソッドの定義](#ヘルパーメソッドの定義)
  - [アセットパイプライン](#アセットパイプライン)
  - [スタイルシートの分離](#スタイルシートの分離)
  - [アセットのプリコンパイル](#アセットのプリコンパイル)
  - [controllers のレイアウト選択](#controllers-のレイアウト選択)
  - [production モードでの起動](#production-モードでの起動)
- [エラーページの作成](#エラーページの作成)
  - [raise メソッド](#raise-メソッド)
  - [例外処理の書き方](#例外処理の書き方)
  - [クラスメソッド rescue_from](#クラスメソッド-rescue_from)
  - [500 Internal Server Error](#500-internal-server-error)
  - [403 Forbidden](#403-forbidden)
  - [404 Not Found](#404-not-found)
    - [ActionController::RoutingError の処理](#actioncontrollerroutingerror-の処理)
    - [ActiveRecord::RecordNotFound の処理](#activerecordrecordnotfound-の処理)
  - [エラー処理の切り分け](#エラー処理の切り分け)
- [サーバサイドにおけるユーザ認証の前準備](#サーバサイドにおけるユーザ認証の前準備)
  - [初回マイグレーション](#初回マイグレーション)
  - [パスワードのハッシュ化](#パスワードのハッシュ化)
  - [seed データの投入](#seed-データの投入)
  - [認証後の session によるユーザ管理](#認証後の-session-によるユーザ管理)
  - [ログイン用のルーティング](#ログイン用のルーティング)
- [フロントエンドから流れに乗るユーザ認証の本実装](#フロントエンドから流れに乗るユーザ認証の本実装)
  - [ログイン・ログアウトのリンク](#ログインログアウトのリンク)
  - [form_with メソッド](#form_with-メソッド)
  - [ログインフォームの作成](#ログインフォームの作成)
  - [ログイン時の session の追加](#ログイン時の-session-の追加)
  - [ログアウト時の session 削除](#ログアウト時の-session-削除)
- [ルーティングのカスタマイズ](#ルーティングのカスタマイズ)
  - [アクション単位のルーティング](#アクション単位のルーティング)
  - [resources によるルーティング](#resources-によるルーティング)
  - [resource によるルーティング](#resource-によるルーティング)
  - [ルーティングにおける制約](#ルーティングにおける制約)
- [Admin による Staff アカウント CRUD の実装](#admin-による-staff-アカウント-crud-の実装)
  - [staff_members index アクション](#staff_members-index-アクション)
  - [staff_members show アクション](#staff_members-show-アクション)
  - [staff_members new アクション](#staff_members-new-アクション)
  - [staff_members edit アクション](#staff_members-edit-アクション)
  - [staff_members create アクション](#staff_members-create-アクション)
  - [staff_members update アクション](#staff_members-update-アクション)
  - [staff_members destroy アクション](#staff_members-destroy-アクション)
- [マスアサインメント脆弱性に対するセキュリティ強化](#マスアサインメント脆弱性に対するセキュリティ強化)
  - [Strong Parameters による防御](#strong-parameters-による防御)
- [Staff アカウントによる自身の閲覧・編集機能の実装](#staff-アカウントによる自身の閲覧編集機能の実装)
  - [staff_accounts show アクション](#staff_accounts-show-アクション)
  - [staff_accounts edit アクション](#staff_accounts-edit-アクション)
  - [staff_accounts update アクション](#staff_accounts-update-アクション)
- [Admin および Staff アカウントにおけるアクセス制御の実装](#admin-および-staff-アカウントにおけるアクセス制御の実装)
  - [ページアクセスにおける認証](#ページアクセスにおける認証)
  - [Admin による Staff の強制ログアウト](#admin-による-staff-の強制ログアウト)
  - [セッションタイムアウト](#セッションタイムアウト)
- [Admin による Staff アカウントのログイン・ログアウト記録閲覧の実装](#admin-による-staff-アカウントのログインログアウト記録閲覧の実装)
  - [StaffEvent モデルの作成](#staffevent-モデルの作成)
  - [StaffEvent・StaffMember の関連付け](#staffeventstaffmember-の関連付け)
  - [ログイン履歴の記録](#ログイン履歴の記録)
  - [ログイン履歴の表示](#ログイン履歴の表示)
  - [ページネーション](#ページネーション)
    - [kaminari のカスタマイズ](#kaminari-のカスタマイズ)
  - [StaffEvent による StaffMember 取得時の N + 1 問題](#staffevent-による-staffmember-取得時の-n--1-問題)
- [DB 格納前の正規化とバリデーションの実装](#db-格納前の正規化とバリデーションの実装)
  - [氏名・フリガナの正規化](#氏名フリガナの正規化)
  - [氏名・フリガナのバリデーション](#氏名フリガナのバリデーション)
  - [入社日・退職日のバリデーション](#入社日退職日のバリデーション)
  - [メールアドレスの正規化](#メールアドレスの正規化)
  - [メールアドレスのバリデーション](#メールアドレスのバリデーション)
- [Staff によるパスワードの変更](#staff-によるパスワードの変更)
  - [password を操作対象としたルーティングの設定](#password-を操作対象としたルーティングの設定)
  - [passwords show アクション](#passwords-show-アクション)
  - [passwords edit アクション](#passwords-edit-アクション)
  - [passwords update アクション](#passwords-update-アクション)
- [プレゼンタによるフロントエンドのリファクタ](#プレゼンタによるフロントエンドのリファクタ)
  - [プレゼンタ利用の準備](#プレゼンタ利用の準備)
  - [HtmlBuilder の使い方](#htmlbuilder-の使い方)
  - [StaffMember のモデルプレゼンタ](#staffmember-のモデルプレゼンタ)
  - [StaffEvent のモデルプレゼンタ](#staffevent-のモデルプレゼンタ)
  - [ggg](#ggg)
- [Customer アカウントの CRUD 実装](#customer-アカウントの-crud-実装)
- [Capybara およびバリデーションによる Customer アカウントの CRUD リファクタ](#capybara-およびバリデーションによる-customer-アカウントの-crud-リファクタ)
- [ActiveSupport::Concern による機能共通化を目的としたリファクタ](#activesupportconcern-による機能共通化を目的としたリファクタ)
- [Customer アカウントにおける自宅住所と勤務先の任意入力の実装](#customer-アカウントにおける自宅住所と勤務先の任意入力の実装)
- [Customer アカウントにおける電話番号の CRUD 実装](#customer-アカウントにおける電話番号の-crud-実装)
- [参考文献](#参考文献)

ソースコード : [ruby-rails-rspec-prac](https://github.com/krtsato/ruby-rails-rspec-prac)

<br>

## 環境構築

- [ruby-rails-rspec-prac](https://github.com/krtsato/ruby-rails-rspec-prac) の Shell を適宜アレンジする
- Rails API の構築情報は今後追加

<br>

## トップページの作成

### routes to root の定義

- routes.rb で名前空間 `namespace` を定義する
  - admin
  - staff
  - customer
- ルーティングにおける名前空間の影響
  - URL パスの先頭に `/admin` を付加する
  - controller 名の先頭に `/admin` を付加する
  - ルーティング名の先頭に `admin_` を付加する
- ルーティングの詳細は[後述](#ルーティングの設定)
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
- Staff・Customer も同様

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
  - 接頭辞辞 `_` を付ける慣習がある

<br>

### ヘルパーメソッドの定義

- ERB などのテンプレート内で使用できるメソッド
  - app/helpers/ 配下に定義する
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
    - Staff・Customer も同様
- SCSS での変数定義を別ファイルで行う
  - e.g. 色を変数で表す
    - app/assets/stylesheets/admin/\_colors.scss
    - 接頭辞 `_` を付ける慣習がある
    - admin/hoge.scss で `@import "colors";` する

<br>

### アセットのプリコンパイル

- 分離したアセットのエントリポイントを含める
- config/initializers/assets.rb に追記する

```ruby
+ Rails.application.config.assets.precompile += %w(admin.css staff.css customer.css)
```

<br>

### controllers のレイアウト選択

- ドメインごとのレイアウト表示を controllers に反映させる
- 通常は controller と同名のレイアウトが選択される
- それが無い場合 application という名前のレイアウトが選択される
  - e.g. 現状 admin/top という controller 名から app/views/layouts/admin/top.html.erb が選択される
  - 今後 admin/hoge という controller がアクションを実行するとき
  - app/views/layouts/admin.html.erb をエントリポイントにしたい
- app/controllers/application_controller.rb に追記する
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
- [ruby-rails-rspec-prac](https://github.com/krtsato/ruby-rails-rspec-prac) では Dockerfile の CMD でサーバを起動する
  - CMD で指定した start-rails-server.sh に `-e production` を追記する
  - コンテナを restart する．確認が終わったら追記箇所を元に戻す

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

- A で例外が発生した時点で A の処理は中断される
- そこで発生した例外オブジェクトが E1 のインスタンスならば e1 に格納
- 続けて B が実行される
- 同様に発生した例外オブジェクトが E2 のインスタンスならば e2 に格納
- 続けて C が実行される
- 例外オブジェクトが E1・E2 のインスタンスでない場合は SystemError
- A での例外発生に関わらず最後に D が実行される

<br>

### クラスメソッド rescue_from

- アクション内で発生した例外の処理方法を指定する
- `rescue_from Forbidden, with: :rescue403`
  - Forbidden・Forbidden の子孫が例外を例外を発生させたとき
    - アクションを中止する
    - rescue403 メソッドを実行する

```ruby
private

def rescue403(exception)
  @exception = exception # 渡された例外オブジェクト
  render template: "errors/forbidden", status: 403
end
```

<br>

### 500 Internal Server Error

- スタイリングの管理
  - app/assets/stylesheets/shared/errors.scss を作成する
  - app/assets/stylesheets/\*.css でアセットパイプラインの対象ディレクトリに設定する
    - `*=require_tree ./shared` する
- controller への追記
  - application_controller.rb に種々の例外処理メソッドを定義していく
  - `rescue500` は views/errors/internal_server_error.rb を表示する
- view ファイルを作成する
- 意図的に例外を発生させる
  - controller の index アクションにおいて `raise` する

```ruby
class ApplicationController < ActionController::Base
+ rescue_from StandardError, with: :rescue500

  private

+ def rescue500
+   render 'errors/internal_server_error', status: 500
+ end
end
```

<br>

### 403 Forbidden

- application_controller.rb に２種類の例外クラスを定義する
  - Forbidden : 権限不足によりリクエスト拒否
  - IpAddressRejected : IP アドレス制限によりリクエスト拒否
- ActionController::ActionControllerError
  - StandardError を継承する例外クラス
  - controller で発生する様々な例外の親・祖先クラス
- class 内部の class
  - ApplicationController が名前空間を提供する役割を持つ
  - 原則的には Application::Forbidden のように呼び出す
- `rescue_from 例外クラス` の注意事項
  - 親子関係にある例外を指定する場合，親・祖先の例外を先に指定する
- view ファイルを作成する
  - インスタンス変数 `@exception` を view 側で使う
- 意図的に例外を発生させる
  - ApplicationController を継承した controller では名前空間を省略して呼び出せる
  - e.g. `raise Forbidden` は `ApplicationController::` を省略している

```ruby
class ApplicationController < ActionController::Base

+ class Forbidden < ActionController::ActionControllerError; end
+ class IpAddressRejected < ActionController::ActionControllerError; end

  rescue_from StandardError, with: :rescue500 # 先に StandardError を指定
+ rescue_from Forbidden, with: :rescue403
+ rescue_from IpAddressRejected, with: :rescue403

  private

+ def rescue403(exception)
+   @exception = exception
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
  - `ErrorsController.action(action).call(env)` が例外用アクションを呼び出す

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
- errors_controller.rb で例外処理に応じた view ファイルを指定する
- view ファイルを作成する
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
- application_controller.rb に追記する

```ruby
class ApplicationController < ActionController::Base

  rescue_from StandardError, with: :rescue500 # 先に StandardError を指定
+ rescue_from ActiveRecord::RecordNotFound, with: :rescue404

  private

+ def rescue404
+   render 'errors/not_found', status: 404
+ end
end
```

<br>

### エラー処理の切り分け

- controllers/concerns/error_handlers.rb にモジュールとして切り分ける
- ActiveSupport::Concern によるモジュール化
  - `included do ... end` 内のメソッド
    - モジュールを読み込んだクラスのクラスメソッドになる
    - モジュールが読み込まれた直後に定義される
      - e.g. `scope ...`
  - `class_methods ... end` 内のメソッドは，そのモジュールを読み込んだクラスのクラスメソッドになる
  - ブロックで囲まず定義したメソッドは，そのモジュールを読み込んだクラスのインスタンスメソッドになる
  - application_controller.rb で定義した例外処理用のクラスは，名前空間付きで呼ぶ
    - e.g. `rescue_from ApplicationController::Forbidden`
- dev 環境ではデバッグ目的でエラー表示を加工しない
  - `if Rails.env.production?`

```ruby
class ApplicationController < ActionController::Base
  # 例外処理用のクラス
  class Forbidden < ActionController::ActionControllerError; end
  class IpAddressRejected < ActionController::ActionControllerError; end

  # モジュールの読み込み
+ include ErrorHandlers if Rails.env.production?
- # エラー関係のメソッドをすべて削除
end
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
  def rescue404
    render 'errors/not_found', status: 404
  end

  # rescue400, rescue403, rescue500 も同様
end
```

<br>

## サーバサイドにおけるユーザ認証の前準備

### 初回マイグレーション

- staff の会員情報を管理する DB テーブル staff_members を作成する
  - admin は同様の手順・異なる DB スキーマで実装する
- `bundle exec rails g model StaffMember` 単数形に注意
- マイグレーションスクリプトに追記
  - ブロック変数 `t` には TableDefinition オブジェクトがセットされる
  - このオブジェクトの各種メソッドがテーブルの定義を行う
  - index を設定することで検索・ソートを高速化する
    - メールアドレス
    - 苗字・名前
- `bundle exec rails db:migrate` する
- `bundle exec rails r "StaffMember.columns.each {|c| p [c.name, c.type]}"` でカラム構成を確認
  - 主キーはデフォルト設定される `["id", :integer]`

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

    add_index :staff_members, :email, unique: true
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
    - `password` を要素代入関数として定義する
    - 代入演算子 `=` を用いて引数を渡せる関数
    - `hoge = Hoge.new; hoge.password = 'fuga'`
  - `BCrypt::Password.create(raw_passward)`
    - gem bcrypt を使ってハッシュ値を生成する

<br>

### seed データの投入

- seed データも DRY にする
- db/seeds.rb で path を振り分ける
  - `%w()` : 配列の要素をスペース区切りで指定
  - `require` : 標準ライブラリ・外部ファイル・自作ファイルを読み込む関数
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
  email: "hoge@example.com",
  # ...
)
```

<br>

### 認証後の session によるユーザ管理

- 名前空間 Staff を DRY にするため Staff::Base クラスを作る
  - controllers/staff/top_controller.rb に継承させる
- controllers/staff/base.rb に共通処理を書く
  - 遅延初期化
    - StaffMember.find_by メソッドが多くても１回しか呼ばれない
  - session オブジェクトはクッキーの中に保持されている
  - `helper_method` でヘルパーメソッドとして登録する

```ruby
module Staff
- class TopController < ApplicationController
+ class TopController < Base # 同じ名前空間内のため Staff:: を省略
    # ...
  end
end
```

```ruby
module Staff
  class Base < ApplicationController
    private

    def current_staff_member
      return if session[:staff_member_id].blank?
      # 遅延初期化
      @current_staff_member ||= StaffMember.find_by(id: session[:staff_member_id])
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
- ルーティングの詳細は[後述](#ルーティングの設定)

| Task                           | HTTP method | URL path       | Controller     | Action  |
| ------------------------------ | ----------- | -------------- | -------------- | ------- |
| ログインフォーム<br>を表示する | GET         | /staff/login   | staff/sessions | new     |
| ログインする                   | POST        | /staff/session | staff/sessions | create  |
| ログアウトする                 | DELETE      | /staff/session | staff/sessions | destroy |

```ruby
namespace :staff do
  root 'top#index'
+ get 'login' => 'sessions#new', as: :login
+ resource :session, only: [:create, :destroy]
end
```

<br>

## フロントエンドから流れに乗るユーザ認証の本実装

### ログイン・ログアウトのリンク

- views/shared/\_header.html.erb を DRY にする
  - ユーザ認証はユーザの種類によって処理が異なる
  - \_footer.html.erb は共通のまま
  - 各ドメインに shared/ を作成する
    - e.g. staff/shared/\_header.html.erb
- views/layouts/staff.html.erb を編集する
  - `current_staff_member` は登録済みの helper_method であり `@current_staff_member` を返す

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
      link_to 'ログイン', :staff_login, method: :get
    end
  %>
</header>
```

<br>

### form_with メソッド

- デフォルトでリモートフォーム化される
  - 入力内容を Ajax で送信するため, レスポンスに応じてブラウザ状態の更新が必要
  - 要件に応じてオフにする
    - config/initializers/action_view.rb に追記する
    - `config.action_view.form_with_generates_remote_forms = false`
- オプション
  - `model:` モデルオブジェクト・フォームオブジェクトを指定する
    - モデルオブジェクト : ActiveRecord::Base を継承したクラスのインスタンス
    - フォームオブジェクト : フォームで指定される非 ActiveRecord モデル
  - `url:` フォームの入力データを送信する URL・そのシンボルを指定する

<br>

### ログインフォームの作成

- app/forms/staff/login_form.rb でフォームオブジェクトを作成する
  - モデルオブジェクトではないので ActiveRecord::Base を継承しない
  - `include ActiveModel::Model` : `form_with` の `model:` として指定できる
  - `attr_accessor` : 指定した属性はフォームのフィールド名になる
- セッション周りのコントローラを作成 `bundle exec rails g controller staff/sessions`
  - フォームオブジェクトを作成しインスタンス変数に格納して view へ渡す
- view ファイルを作成する

```ruby
module Staff
  class LoginForm
    include ActiveModel::ActiveModel
    attr_accessor :email, :password
  end
end
```

```ruby
module Staff
  class SessionsController < Base
    def new
      if current_staff_member
        redirect_to :staff_root
      else
        @form = LoginForm.new
        render action: 'new'
      end
    end
  end
end
```

<br>

### ログイン時の session の追加

- サービスオブジェクトとして app/services/staff/authenticator.rb を作成する
  - controller のインスタンスメソッドではなく，独立したクラスとして実装される
  - 7 Patterns to Refactor Fat ActiveRecord Models の思想に基づく
  - `BCrypt::Password.new(@staff_member.hashed_password) == raw_password`
    - ハッシュパスワードのインスタンスを作成する
    - BCrypt のインスタンスメソッド `==` で平文パスワードをハッシュ化する
      - BCrypt では比較演算子がオーバーライドされている
    - インスタンスが保持しているハッシュ値と同じならば true を返す
- sessions_controller.rb に session 追加の機能を書く
  - 本来はフォームから送信された params オブジェクトを直に取り回すべきでない
  - 今後 [Strong parameters](#マスアサインメント脆弱性に対するセキュリティ強化) で置換する
- 認証の手順
  - email から staff_member を取得する
  - `suspended` に関係なくパスワードのハッシュ値比較などを行う
  - `suspended` かどうか確認する
- flash・session オブジェクトに値を設定する
  - `flash.now.alert` : alert 属性にセットされた値がアクション終了時に削除される
    - 通常は次のアクセス時まで flash を保持している
    - flash にセットしたメッセージを当該アクションでのみ使用する場合に有効
- view ファイルで flash を受け取る
  - `now` : はフロント側では指定しなくて良い

```ruby
module Staff
  class Authenticator
    def initialize(staff_member)
      @staff_member = staff_member
    end

    def authenticate(raw_password)
      @staff_member&.hashed_password &&
        @staff_member.start_date <= Date.today &&
        (@staff_member.end_date.nil? || @staff_member.end_date > Date.today) &&
        BCrypt::Password.new(@staff_member.hashed_password) == raw_password
    end
  end
end
```

```ruby
module Staff
  class SessionsController < Base
+   def create
+     @form = LoginForm.new(params[:staff_login_form])
+     if @form.email.present?
+       staff_member = StaffMember.find_by(email: @form.email.downcase)
+     end
+
+     if Authenticator.new(staff_member).authenticate(@form.password)
+       if staff_member.suspended?
+         flash.now.alert = 'アカウントが停止されています'
+         render action: 'new'
+       else
+         session[:staff_member_id] = staff_member.id
+         flash.notice = 'ログインしました'
+         redirect_to :staff_root
+       end
+     else
+       flash.now.alert = 'メールアドレスまたはパスワードが正しくありません'
+       render action: 'new'
+     end
+   end
+ end
end
```

```erb
<header>
  <%# ... %>
+ `<%= content_tag(:span, flash.notice, class: 'notice') if flash.notice %>`
+ `<%= content_tag(:span, flash.alert, class: 'alert') if flash.alert %>`
</header>
```

<br>

### ログアウト時の session 削除

```ruby
module Staff
  class SessionsController < Base
    #...
+   def destroy
+     session.delete(:staff_member_id)
+     redirect_to :staff_root
+   end
  end
end
```

<br>

## ルーティングのカスタマイズ

### アクション単位のルーティング

- ルーティングに名前を与える
  - `as: hoge`
  - URL パスを誤指定すると 404 表示前にエラーが発生する
  - URL パスの変更に強くなる
- ヘルパーメソッド
  - `hoge_path` : URL のパス部分を返す
    - クエリパラメタを付加できる `hoge_path(k1: v1, k2: v2, ...)`
  - `hoge_url` : URL 全体を返す
- パラメタに制約を設ける
  - `get "hoge/:year" => "hoge#show", constraints: {year: /20\d\d}`
- `namespace` のオプション
  - path : URL パスの先頭文字列を変更する
    - e.g. `namespace :fuga, path: 'piyo' do ... end`
  - module : controller 名の先頭文字列を変更する
    - e.g. `namespace :fuga, module: 'piyo' do ... end`
  - as : ルーティング名の先頭文字列を変更する
    - e.g. `namespace :fuga, as: 'piyo' do ... end`

<br>

### resources によるルーティング

- リソースとはアクションによる操作の対象物を意味する
- 複数リソースを CRUD する場合に用いる
  - e.g. 管理者 Admin が 職員 Staff を一覧表示がする
- `resources :controller_names`
  - 複数形で指名する 
  - 複数形の controller 名に繋げる
- 7 つの基本アクションに対するルーティングを一括指定する
  - index, show, new, edit, create, update, destroy
- オプション
  - only : 基本アクションの一部にルーティングを設定
    - e.g. `resources :staff_members, only: [:index, :new, :create]`
  - except : 基本アクションの一部をルーティングから除外
    - e.g. `resources :staff_members, except: [:show, :destroy]`
  - controller : controller を変更する
    - e.g. `resources :staff_members, controller: 'employees'`
  - path : URL のパスを変更する
    - e.g. `resources :staff_members, path: 'staff'`

```ruby
namespace :admin do
+ resources :staff_members
end
```

| アクション内容         | HTTP メソッド | アクション名 | URL パス                      | ルーティング名           |
| ---------------------- | ------------- | ------------ | ----------------------------- | ------------------------ |
| 職員のリスト表示       | GET           | index        | /admin/staff_members          | :admin_staff_members     |
| 職員の詳細表示         | GET           | show         | /admin/staff_members/:id      | :admin_staff_member      |
| 職員の登録フォーム表示 | GET           | new          | /admin/staff_members/:id/new  | :new_admin_staff_member  |
| 職員の編集フォーム表示 | GET           | edit         | /admin/staff_members/:id/edit | :edit_admin_staff_member |
| 職員の追加             | POST          | create       | /admin/staff_members          | :admin_staff_members     |
| 職員の更新             | PATCH         | update       | /admin/staff_members/:id      | :admin_staff_member      |
| 職員の削除             | DELETE        | destroy      | /admin/staff_members/:id      | :admin_staff_member      |

<br>

### resource によるルーティング

- 単数リソースを CRUD する場合に用いる
  - e.g. 職員 Staff がアカウントページを確認する
- `resource :controller_name`
  - 単数形で指名する
  - 複数形の controller 名に繋げる
- URL パスに id パラメタを埋め込む必要はない
  - 職員が自身のアカウントを管理できる = ログインしている
  - id は session オブジェクトから取得できる
  - 管理者が複数の職員の中から１人を指名する = `/admin/staff_members/:id`

```ruby
namespace :staff do
  # controllers/staff/accounts_controller.rb における
  # Staff::AccountsController に繋げる
+ resource :account, except: [:new, :create, :destroy]
end
```

| アクション内容               | HTTP メソッド | アクション名 | URL パス            | ルーティング名      |
| ---------------------------- | ------------- | ------------ | ------------------- | ------------------- |
| アカウントの詳細表示         | GET           | show         | /staff/account      | :staff_account      |
| アカウントの編集フォーム表示 | GET           | edit         | /staff/account/edit | :edit_staff_account |
| アカウントの更新             | PATCH         | update       | /staff/account      | :staff_account      |

<br>

### ルーティングにおける制約

- トップページのホスト名と URL パスを変更する
- config/initializers/rrrp.rb に設定を書く
  - `config` は `Rails::Application::Configuration` のインスタンスを返すメソッド
  - Rails 本体または Gem パッケージの各種設定を編集・追加できる
- `Rails.application.config.hoge` で設定した `config` の中身にアクセスする

```ruby
Rails.application.configure do
  config.rrrp = {
    admin: {host: ENV['ADMIN_STAFF_HOST_NAME'], path: 'admin'},
    staff: {host: ENV['ADMIN_STAFF_HOST_NAME'], path: ''},
    customer: {host: ENV['CUSTOMER_HOST_NAME'], path: 'mypage'}
  }
end
```

```ruby
Rails.application.routes.draw do
+ config = Rails.application.config.rrrp

+ constraints host: config[:admin][:host] do
-   namespace :admin
+   namespace :admin, path: config[:admin][:path] do
      # ...
    end
  end

+ constraints host: config[:staff][:host] do
-   namespace :staff
+   namespace :staff, path: config[:staff][:path] do
      # ...
    end
  end
  # ...
end
```

<br>

## Admin による Staff アカウント CRUD の実装

### staff_members index アクション

- 一覧表示する seed データを db/seeds/development/staff_members.rb に用意する
- `bundle exec rails db:migrate:reset` の後に `db:seed` することで seed の重複エラーを回避

```ruby
+ family_names = %w(佐藤:サトウ:sato ...)
+ given_names = %w(二郎:ジロウ:jiro ...)

+ fn_size = family_names.size
+ gn_size = given_names.size
+ all_combinations = fn_size * gn_size

+ all_combinations.times do |n|
+   fn = family_names[n % fn_size].split(":")
+   gn = given_names[n % gn_size].split(":")

+   StaffMember.create!(
+     email: "#{fn[2]}.#{gn[2]}@example.com",
+     family_name: fn[0],
+     given_name: gn[0],
+     family_name_kana: fn[1],
+     given_name_kana: gn[1],
+     password: "password",
+     start_date: (100 - n).days.ago.to_date,
+     end_date: n == 0 ? Time.zone.today : nil,
+     suspended: n == 1
+   )
+ end
```

- Admin が StaffMember を扱う admin/staff_members_controller.rb を作成する
  - `bundle exec rails g controller admin/staff_members`
  - フリガナを姓・名の順にソートしつつ，staff_members テーブルの全レコードを取得

```ruby
module Admin
  class StaffMembersController < Base
    def index
      @staff_members = StaffMember.order(:family_name_kana, :given_name_kana)
    end
  end
end
```

- view ファイル admin/staff_members/index.html.erb を作成する
  - `end_date.try(:strftime, %Y/%m/%d)`
    - Date クラスのインスタンスメソッド `strftime` で日付をフォーマットする
    - `end_date` が nil の場合 `try` メソッドが nil を返す
      - 第１引数 : レシーバが nil でないとき実行するメソッド
      - 第２引数 : メソッドに渡す引数
  - エスケープ処理を抑制する場合は `raw` メソッドを使う

<br>

### staff_members show アクション

- データを閲覧する場合に限らず，edit アクションにリダイレクトするだけの場合もある
  - e.g. update に失敗して編集フォームが再表示される場合
  - `http://rrrp.example.com/admin/staff_members/123` のような URL が提供される
  - このページはお気に入り登録・リンクのコピペによって再表示され得る
  - show アクションにアクセスさせて，即座に edit アクションへリダイレクトする
- `redirect_to [:edit, :admin, staff_member]`
  - 引数が配列の場合 redirect_to は配列要素からルーティング名を推定する
  - ルーティング名 : edit_admin_staff_memnber
  - URL パス : `/admin/staff_members/123/edit`

```ruby
module Admin
  class StaffMembersController < Base
+   def show
+     staff_member = StaffMember.find(params[:id])
+     redirect_to [:edit, :admin, staff_member]
+   end
  end
end
```

<br>

### staff_members new アクション

- インスタンスを生成して admin/staff_members/new.html.erb を表示する
- `<%= form_with ... do |f| %>` : ブロック変数 `f` にフォームビルダーがセットされる
- `<%= render 'form', f: f %>` : 部分テンプレート \_form.html.erb 内で `f` を参照する

```ruby
module Admin
  class StaffMembersController < Base
+   def new
+     @staff_member = StaffMember.new
+   end
  end
end
```

```erb
<%= form_with model: @staff_member, url: [:admin, @staff_member] do |f| %>
  <%= render 'form', f: f %>
  <%# ... %>
<% end %>
```

```erb
<div>
  <%= f.label :password, 'パスワード', class: 'required' %>
  <%= f.password_field :password, size: 32, required: true %>
</div>
```

<br>

### staff_members edit アクション

- レコードを取得して admin/staff_members/edit.html.erb を表示する
- 編集フォームは新規作成フォームと共通で利用する
- パスワードの変更は分離して[後述](#staff-によるパスワードの変更)
  - `f.object.new_record?` : DB に未保存ならばフォームを表示
  - 表示のためにハッシュをデコードする必要がある
  - 職員アカウントを更新する度にパスワードをデコード・ハッシュ化するのは実用的でない
    - 漏洩・盗聴リスク
    - 計算コスト

```ruby
module Admin
  class StaffMembersController < Base
+   def edit
+     @staff_member = StaffMember.find(params[:id])
+   end
  end
end
```

```erb
+ <% if f.object.new_record? %>
    <div>
      <%= f.label :password, 'パスワード', class: 'required' %>
      <%= f.password_field :password, size: 32, required: true %>
    </div>
+ <% end %
```

<br>

### staff_members create アクション

- バリデーションの実装は[後述](#db-格納前の正規化とバリデーションの実装)
- 本来はフォームから送信された params オブジェクトを直に取り回すべきでない
  - 今後 [Strong parameters](#マスアサインメント脆弱性に対するセキュリティ強化) で置換する

```ruby
module Admin
  class StaffMembersController < Base
+   def create
+     @staff_member = StaffMember.new(params[:staff_member])
+     if @staff_member.save
+       flash.notice = '職員アカウントを新規登録しました'
+       redirect_to :admin_staff_members
+     else
+       render action: 'new'
+     end
+   end
  end
end
```

<br>

### staff_members update アクション

- `assign_attributes`
  - モデルオブジェクトの属性を一括設定する
  - オブジェクトの変更をするだけで DB には保存しない
- 本来はフォームから送信された params オブジェクトを直に取り回すべきでない
  - 今後 [Strong parameters](#マスアサインメント脆弱性に対するセキュリティ強化) で置換する

```ruby
module Admin
  class StaffMembersController < Base
+   def update
+     @staff_member = StaffMember.find(params[:id])
+     @staff_member.assign_attributes(params[:staff_member])
+     if @staff_member.save
+       flash.notice = '職員アカウントを更新しました'
+       redirect_to :admin_staff_members
+     else
+       render action: 'edit'
+     end
+   end
  end
end
```

<br>

### staff_members destroy アクション

```ruby
module Admin
  class StaffMembersController < Base
+   def destroy
+     staff_member = StaffMember.find(params[:id])
+     staff_member.destroy!
+     flash.notice = '職員アカウントを削除しました'
+     redirect_to :admin_staff_members
+   end
  end
end
```

<br>

## マスアサインメント脆弱性に対するセキュリティ強化

### Strong Parameters による防御

- admin/sessions_controller.rb を編集
- `params.require(:admin_login_form).permit(:email, :password)`
  - params オブジェクトが :admin_login_form キーを持つか確認する
    - 持たない場合は例外 `ActionController::ParameterMissing` が発生する
  - `その中で permit されていないパラメータを除去する
    - 悪意のあるパラメータを含んだ，フォームからのリクエストを受け付けないため
- staff/sessions_controller.rb・admin/staff_members_controller.rb も同様

```ruby
module Admin
  class SessionsController < Base
    def create
-     @form = LoginForm.new(params[:admin_login_form])
+     @form = LoginForm.new(login_form_params)
      # ...
    end

    private

+   def login_form_params
+     params.require(:admin_login_form).permit(:email, :password)
+   end
  end
end
```

```ruby
module Admin
  class StaffMembersController < Base
    def create
-     @staff_member = StaffMember.new(params[:staff_member])
+     @staff_member = StaffMember.new(staff_member_params)
      # ...
    end

    def update
-     @staff_member.assign_attributes(params[:staff_member])
+     @staff_member.assign_attributes(staff_member_params)
      # ...
    end

    private

    def staff_member_params
      params.require(:staff_member).permit(:email, :password, ...)
    end
  end
end
```

<br>

## Staff アカウントによる自身の閲覧・編集機能の実装

### staff_accounts show アクション

- `bundle exec rails g controller staff/accounts` する
- controllers/staff/base.rb を継承させて，ログイン中の職員データを返す
- view ファイルをファイルを作成する
  - show ページ
  - show ページへ誘導する header リンク

```ruby
module Staff
  class AccountsController < Base
    def show
      @staff_member = current_staff_member
    end
  end
end
```

<br>

### staff_accounts edit アクション

```ruby
module Staff
  class AccountsController < Base
+   def edit
+     @staff_member = current_staff_member
+   end
  end
end
```

<br>

### staff_accounts update アクション

- [staff_members update アクション](#staff_members-update-アクション)との違い
  - `@staff_member = current_staff_member`
    - 自分自身のアカウント情報を編集するため session から id を取得する
    - 理由は [resource によるルーティング](#resource-によるルーティング)を参照
  - `redirect_to :staff_account`
    - シンボルが表す URL は [ルーティング表](#resource-によるルーティング)で確認
    - この場合は show ページに戻る

```ruby
module Staff
  class AccountsController < Base
+   def update
+     @staff_member = current_staff_member
+     @staff_member.assign_attributes(staff_member_params)
+     if @staff_member.save
+       flash.notice = 'アカウント情報を更新しました'
+       redirect_to :staff_account
+     else
+       render action: 'edit'
+     end
+   end

+   private

+   def staff_member_params
+     params.require(:staff_member).permit(
+       :email, :family_name, :given_name,
+       :family_name_kana, :given_name_kana
+     )
+   end
  end
end
```

<br>

## Admin および Staff アカウントにおけるアクセス制御の実装

### ページアクセスにおける認証

- Admin の認証状態に応じてアクセスページを制限する
- `before_action :func` : controller に書かれた各アクションの実行直前にメソッドを呼ぶ
- `skip_before_action :func` : `before_action` に指定したメソッドを controller 内で実行しない
  - SessionsController : そもそも認証を行う必要があるため
  - TopController : 未ログインでアクセスするトップページのため
- Staff も同様

```ruby
module Admin
  class Base < ApplicationController
+   before_action :authorize

    private

+   def authorize
+     return if current_administrator.present?
+     flash.alert = '管理者としてログインして下さい'
+     redirect_to :admin_login
+   end
  end
end
```

```ruby
module Admin
  class TopController < Base
    skip_before_action :authorize
    # ...
  end
end
```

```ruby
module Admin
  class SessionsController < Base
    skip_before_action :authorize
    # ...
  end
end
```

<br>

### Admin による Staff の強制ログアウト

- Admin が Staff に対して `suspended = true` する場合
  - Staff が自主的にログアウトするまでアカウントを停止できない
- Staff が退職する場合
  - アカウント終了日を迎えても利用を継続できてしまう
- models/staff_member.rb にアカウントが active か確認するメソッドを追記する
- controllers/staff/base.rb の `before_action` でアカウント状態を確認する
- Admin も同等
  - ただし `active?` メソッドは使えないので `!current_administrator.suspended?` とする

```ruby
class StaffMember < ApplicationRecord
+ def active?
+   !suspended? && start_date <= Time.zone.today && (end_date.nil? || end_date > Time.zone.today)
+ end
end
```

```ruby
module Staff
  class Base < ApplicationController
+   before_action :check_account

    private

+   def check_account
+     return if current_staff_member.blank? || current_staff_member.active?
+
+     session.delete(:staff_member_id)
+     flash.alert = 'アカウントが無効になりました'
+     redirect_to :staff_root
+   end
  end
end
```

<br>

### セッションタイムアウト

- ログイン時刻を session オブジェクトに格納する
- controllers/staff/base.rb の `before_action` で最終アクセス時間を確認する
  - 変数 `TIMEOUT` はテストコードでも使用するため private にしない
  - タイムアウト前にアクセスした場合，最終アクセス時間を現在時刻に更新する

```ruby
module Staff
  class SessionsController < Base
    def create
      if Authenticator.new(staff_member).authenticate(@form.password)
        if staff_member.suspended?
          # ...
        else
          session[:staff_member_id] = staff_member.id
+         session[:staff_last_access_time] = Time.current
          go_to_staff_root('ログインしました')
        end
        # ...
      end
    end
  end
end
```

```ruby
module Staff
  class Base < ApplicationController
+   before_action :check_timeout

    private

+   TIMEOUT = 60.minutes
+   def check_timeout
+     return if current_staff_member.blank?
+
+     if session[:staff_last_access_time] >= TIMEOUT.ago
+       session[:staff_last_access_time] = Time.current
+     else
+       session.delete(:staff_member_id)
+       flash.alert = 'セッションがタイムアウトしました'
+       redirect_to :staff_login
+     end
+   end
  end
end
```

<br>

## Admin による Staff アカウントのログイン・ログアウト記録閲覧の実装

### StaffEvent モデルの作成

- StaffEvent : Staff のログイン・ログアウトを記録するモデル
  - id : 整数型．主キー
  - staff_member_id : 整数型．外部キー
  - type : 文字列型．logged_in・logged_out・reject のいすれか
  - created_at : 日時型．モデルの性質上 updated_at は起こり得ない
- StaffMember : StaffEvent = 1 : 多
- `bundle exec rails g model staff_event` する
- マイグレーションスクリプトを編集
  - `t.references :staff_member` : シンボル名末尾に `_id` を付与した整数カラムを作成する．
    - `null: false` : null を許容しない
    - `index: false` : references はデフォルトで index を設定するためオフにする．
    - `foreign_key: true` : staff_members・staff_events のテーブル間に外部キー制約を設ける
  - `add_index :staff_events, [:staff_member_id, :created_at]`
    - staff_member_id・created_at の組み合わせに対して，複合インデックスを設定する
    - イベントを職員別に・発生時間順に並べて取得する場合に便利
- web コンテナで確認 `psql -U rrrp_user -h db rrrp_dev_db`

```ruby
class CreateStaffEvents < ActiveRecord::Migration[6.0]
  def change
    create_table :staff_events do |t|
      t.references :staff_member, null: false, index: false, foreign_key: true # 職員レコードへの外部キー
      t.string :type, null: false # イベントタイプ
      t.datetime :created_at, null: false # 発生時刻
    end

    add_index :staff_events, :created_at
    add_index :staff_events, [:staff_member_id, :created_at]
  end
end
```

### StaffEvent・StaffMember の関連付け

- models/staff_member.rb に追記
  - `has_many :events` : 一対多となる StaffMember のインスタンスメソッド `events` を定義する
  - `class_name: 'StaffEvent'` : `:event` だけでは不明瞭なのでクラスを指定する
    - 関連付けからクラス名を推定できるとき省略可能 `has_many :staff_events`
    - `:events` とした場合 `@staff_member.events` のように呼び出せる
  - `dependent: :destroy`: StaffMember より前に StaffEvent を削除する
- models/staff_event.rb に追記
  - `self.inheritance_column = nil` : `type` というカラムが STI (Single Table Inheritance) と見なされるのを無効化
  - `belongs_to :member` : StaffEvent モデルが `member` というインスタンスメソッドで StaffMember を参照する
    - `class_name: 'StaffMember'` : `:member` だけでは不明瞭なのでクラスを指定する
    - `:member` とした場合 `@event.member` のように呼び出せる
    - `foreign_key: staff_member_id` : このままでは外部キーが `:member_id` になってしまうので指定する
    - `inverse_of: :events` : 同値データの参照時に，経由するインスタンスが異なると `==` で false になる問題を回避

```ruby
class StaffMember < ApplicationRecord
+ has_many :events, class_name: 'StaffEvent', dependent: :destroy
end
```

```ruby
class StaffEvent < ApplicationRecord
  self.inheritance_column = nil

  belongs_to :member, class_name: 'StaffMember', foreign_key: 'staff_member_id',  inverse_of: :events
  alias_attribute :occurred_at, :created_at
end
```

<br>

### ログイン履歴の記録

- staff/sessions_controller.rb に追記
  - ログイン・ログアウト時に履歴イベントを作成する
  - このとき `type` を文字列型で指定する

```ruby
module Staff
  class SessionsController < Base
    def create
      # ...
      if Staff::Authenticator.new(staff_member).authenticate(@form.password)
        if staff_member.suspended?
+         staff_member.events.create!(type: 'rejected')
          # ...
        else
+         staff_member.events.create!(type: 'logged_in')
          # ...
        end
      # ...
      end
    end

    def destroy
+     current_staff_member&.events&.create!(type: 'logged_out')
      # ...
    end
  end
end
```

<br>

### ログイン履歴の表示

- ルーティングを追加する
  - 特定の職員についてログイン履歴を表示する
  - すべての職員のログイン履歴を表示する
- view ファイルに `:admin_staff_events` へのリンクを追加する
- view ファイルを作成する
  - `:admin_staff_events` から `:admin_staff_member_staff_events` に遷移させる

| アクション内容                 | HTTP メソッド | アクション名 | URL パス                                           | ルーティング名                   |
| ------------------------------ | ------------- | ------------ | -------------------------------------------------- | -------------------------------- |
| ある職員のログイン履歴一覧表示 | GET           | index        | /admin/staff_members/:staff_member_id/staff_events | :admin_staff_member_staff_events |
| 全職員のログイン履歴一覧表示   | GET           | index        | /admin/staff_events                                | :admin_staff_events              |

```ruby
Rails.application.routes.draw do
  config = Rails.application.config.rrrp

  constraints host: config[:admin][:host] do
    namespace :admin, path: config[:admin][:path] do
      # ...
+     resources :staff_members do
+       resources :staff_events, only: [:index]
+     end
+    resources :staff_events, only: [:index]
    end
  end
end
```

- `bundle exec rails g controller admin/staff_events` して index アクションを提供する
- `params[:staff_member_id]` の有無によって２種類の index 表示を区別する

```ruby
module Admin
  class StaffEventsController < Base
    def index
      if params[:staff_member_id]
        @staff_member = StaffMember.find(params[:staff_member_id])
        @events = @staff_member
      else
        @events = StaffEvent
      end
      @events = @events.order(occurred_at: :desc)
    end
  end
end
```

### ページネーション

- seed を投入する
- ページネーションに使う gem kaminari をセットアップする
  - `bundle exec rails g kaminari:config`
  - `bundle exec rails g kaminari:views default`
  - config/locales/views/paginate.ja.yml を作成する
  - コンテナを再起動して反映させる
    - config/initializers/・config/locales/ 配下の編集後は忘れず実行

```yml
ja:
  views:
    pagination:
      first: "先頭"
      last: "末尾"
      previous: "前"
      next: "次"
      truncate: "..."
```

- admin/staff_events_controller.rb に追記する
  - `page` は kaminari が提供するメソッド
  - 引数に指定された整数をページ番号と見なす
  - 引数が nil の場合は `page(1)` となる
- view ファイルに paginate の処理を追記する
- N + 1 問題への対処は[後述](#staffevent-による-staffmember-取得時の-n--1-問題)

```ruby
module Admin
  class StaffEventsController < Base
    def index
      if params[:staff_member_id]
        # ...
      end
+     @events = @events.order(occurred_at: :desc).page(params[:page])
    end
  end
end
```

<br>

#### kaminari のカスタマイズ

- １ページ目を表示するとき「先頭」「前」のクリック不可リンクを表示させる
- 最終ページを表示するとき「次」「末尾」のクリック不可リンクを表示させる
- views/kaminari/\_paginator.html.erb を編集
  - ページネーションリンクの列配置を決定するファイル
  - `unless` の条件を取り除く

```ruby
# ...
- <%= first_page_tag unless current_page.first? %>
+ <%= first_page_tag %>
- <%= prev_page_tag unless current_page.first? %>
+ <%= prev_page_tag %>
# ...
- <%= next_page_tag unless current_page.last? %>
+ <%= next_page_tag %>
- <%= last_page_tag unless current_page.last? %>
+ <%= last_page_tag %>
# ...
```

- リンクを１つずつカスタマイズする
  - views/kaminari/\_first_page.html.erb
  - `link_to_unless` : hoge
    - 第１引数 : リンク表示の条件式
    - 第２引数
      - リンクの文字列
      - 国際化に関するヘルパーメソッド `translate` のエイリアス `t` を使う
      - [前項](#ページネーション)の YAML ファイルを参照する
    - 第３引数 : リンク先の URL．先頭ページの URL を返すヘルパーメソッド `url` を使う
    - オプション `remote: remote` : Ajax でリクエストするかどうか
  - ブロック表記 `link_to_unless(...) do |name| ... end`
    - 条件式が偽のとき，ブロック戻り値をリンク文字列とする
    - ブロック変数 `name` には第２引数が渡される
  - \_last_page.html.erb・\_prev.html.erb・\_next_page.html.erb も同様

```ruby
<span class="first">
- <%= link_to_unless current_page.first?, t('views.pagination.first').html_safe, url, remote: remote %>
+ <%=
+   link_to_unless(current_page.first?, t('views.pagination.first').html_safe, url) do |name|
+     content_tag(:span, name, class: 'disabled')
+   end
+ %>
</span>
```

<br>

### StaffEvent による StaffMember 取得時の N + 1 問題

- N + 1 問題によって不要なクエリが発行される
  - 表示する StaffEvent レコードを一括取得する
  - StaffEvent に紐付いた StaffMember レコードを一つずつ取得する
- staff_member_id をまとめて１つのクエリを発行する
  - `includes` メソッドの引数に[関連付けの名前](#staffeventstaffmember-の関連付け)を与える
  - クエリの回数は２回に減る

```ruby
module Admin
  class StaffEventsController < Base
    def index
      # ...
-     @events = @events.order(occurred_at: :desc).page(params[:page])
+     @events = @events.order(occurred_at: :desc).includes(:member).page(params[:page])
    end
  end
end
```

<br>

## DB 格納前の正規化とバリデーションの実装

- 正規化 : 規則に従うように情報を変換する
  - models/concerns/ 配下に切り出す
  - Ruby 標準ライブラリ nkf を使う
  - models/ 配下のファイル  で読み込む
  - `before_validation do ... end` で正規化する
- バリデーション : 規則に従っているか検証する

<br>

### 氏名・フリガナの正規化

- models/concerns/string_normalizer.rb に正規化関数を作成
- `nkf` メソッドのオプション
  - `-W` : UTF-8 で入力を受け付ける
  - `-w` : UTF-8 で出力する
  - `-Z1` : 全角の英数字・スペース・記号を半角にする
  - `--katakana` : ひらがなをカタカナにする
- `strip` : 文字列の先頭・末尾の空文字を削除するメソッド

```ruby
require 'nkf'

module StringNormalizer
  extend ActiveSupport::Concern

  def normalize_as_name(text)
    NKF.nkf('-WwZ1', text).strip if text
  end

  def normalize_as_furigana(text)
    NKF.nkf('-WwZ1 --katakana', text).strip if text
  end
end
```

```ruby
class StaffMember < ApplicationRecord
+ include StringNormalizer
  # ...
+ before_validation do
+   self.family_name = normalize_as_name(family_name)
+   self.given_name = normalize_as_name(given_name)
+   self.family_name_kana = normalize_as_furigana(family_name_kana)
+   self.given_name_kana = normalize_as_furigana(given_name_kana)
+ end
end
```

<br>

### 氏名・フリガナのバリデーション

- `presence: true` : 値の入力を必須とする
- 正規表現
  - `\p{han}` : 任意の漢字一文字にマッチする
  - `\p{hiragana\}` : 任意のひらがな１文字にマッチする
  - `\u{30fc\}` : 長音符１文字にマッチする
- 詳細なフォーマットを定める `format` バリデーション
  - `with: 正規表現` : ここでカタカナを指定する
  - `allow_blank: true` : 値が空のときバリデーションしない
    - エラーメッセージの重複を回避するため
    - false のとき「email が入力されていない」「email が不正な値です」のようになる

```ruby
class StaffMember < ApplicationRecord
  # ..,
  HUMAN_NAME_REGEXP = /\A[\p{han}\p{hiragana}\p{katakana}\u{30fc}A-Za-z]+\z/.freeze
+ KATAKANA_REGEXP = /\A[\p{katakana}\u{30fc}]+\z/.freeze
+ validates :family_name, :given_name, presence: true, format: {with: HUMAN_NAME_REGEXP, allow_blank: true}
+ validates :family_name_kana, :given_name_kana, presence: true, format: {with: KATAKANA_REGEXP, allow_blank: true}
end
```

<br>

### 入社日・退職日のバリデーション

- Date 型のバリデーションを提供する gem date_validator を使用する
- バリデーション内容
  - 入社日は 2020/01/01 以降かつ本日から１年以内．空値を禁止する
  - 退職日は入社日以降かつ本日から１年以内．空値を許容する
- `date` オプションで以下のキーを指定する
  - `after` : 指定された日付より後．その日付は含まない
  - `after_or_equal_to` : 指定された日付より後．その日付を含む
  - `before` : 指定された日付より前．その日付は含まない
  - `before_or_equal_to` : 指定された日付より前．その日付を含む
  - `allow_blank` : 空値を許可するかどうか
- `before: -> (_obj) {1.year.from_now.to_date}` で動的に日付を指定する
  - `-> (_obj)`
    - 名無し関数 Proc オブジェクトを作成する
    - 接頭辞 `_` が付いた引数は関数内で使用されない
  - `before: 1.year.from_now.to_date` と書いた場合
    - production モードでの起動時に１回だけクラスが読み込まれる
    - その起動時を基準とした１年後の日付に固定されてしまう

```ruby
class StaffMember < ApplicationRecord
  # ...
+ validates :start_date, presence: true, date: {
+   after_or_equal_to: Time.zone.local(2020, 1, 1),
+   before: -> (_obj) {1.year.from_now.to_date},
+   allow_blank: true
+ }
+ validates :end_date, date: {
+  after: :start_date,
+   before: -> (_obj) {1.year.from_now.to_date},
+   allow_blank: true
+ }
end
```

<br>

### メールアドレスの正規化

- `nkf` メソッドのオプションは[氏名・フリガナの正規化](#氏名フリガナの正規化)と同様
- `downcase` することで email アドレスを小文字化している
  - 通常のメールアドレスは大文字・小文字を区別しない
  - PostgreSQL は区別するため大文字入力を正規化する

```ruby
module StringNormalizer
  # ...
+ def normalize_as_email(text)
+   NKF.nkf('-WwZ1', text).strip.downcase if text
+ end
end
```

```ruby
class StaffMember < ApplicationRecord
  # ...
  before_validation do
+   self.email = normalize_as_email(email)
  end
end
```

<br>

### メールアドレスのバリデーション

- gem valid_email2 を使ってバリデーションする
- `uniqueness: {case_sensitive: false}`
  - デフォルトでは大文字・小文字を区別して unique とする
  - 小文字に統一して unique と見なしたいので大文字を区別しない

```ruby
class StaffMember < ApplicationRecord
  # ...
+ validates :email, presence: true, 'valid_email_2/email': true, uniqueness: {case_sensitive: false}
end
```

<br>

## Staff によるパスワードの変更

### password を操作対象としたルーティングの設定

- 単数リソースによるルーティング
- アカウント閲覧ページにパスワード変更ページへのリンク `:edit_staff_password` を設置する
- `bundle exec rails g controller staff/passwords` して controller を作成する

```ruby
constraints host: config[:staff][:host] do
  namespace :staff, path: config[:staff][:path] do
    # ...
+   resource :password, only: [:show, :edit, :update]
  end
end
```

<br>

### passwords show アクション

- 即座に edit アクションへリダイレクトさせる
- 理由は [staff_mambers show アクション](#staff_members-show-アクション)と同様

```ruby
module Staff
  class PasswordsController < Base
    def show
      redirect_to :edit_staff_password
    end
  end
end
```

<br>

### passwords edit アクション

- フォームフォブジェクト app/forms/staff/change_password_form.rb を作成する
- `include ActiveModel::Model` と `attr_accessor` の効能は[ログインフォームの作成](#ログインフォームの作成)を参照
- バリデーション
  - `validates :hoge, ... confirmation: true` : `:hoge` と `:hoge_confirmation` が一致するとき検証成功
  - `validate do ... end` : 組み込みの `presence` や `form` 以外で検証する場合に使用する
  - `errors` : Error オブジェクトを返すメソッド．詳しくは後述
- `save` メソッドの定義
  - `include ActiveModel::Model` することで非 ActiveRecord モデルとなっている
  - インスタンスがデフォルトで save メソッドを持たないため定義する

```ruby
module Staff
  class ChangePasswordForm
    include ActiveModel::Model
    attr_accessor :object, :current_password, :new_password, :new_password_confirmation

    validates :new_password, presence: true, confirmation: true

    validate do
      unless Staff::Authenticator.new(object).authenticate(current_password)
        errors.add(:current_password, :wrong)
      end
    end

    def save
      return unless valid?

      object.password = new_password
      object.save!
    end
  end
end
```

- staff/passwords_controller.rb に追記する
  - `object` 属性に `current_staff_member` を渡すことで，その後のバリデーション・保存ができる
  - フォームオブジェクトをインスタンス変数に格納して view へ渡す
- view ファイルを作成する

```ruby
module Staff
  class PasswordsController < Base
    # ...
+   def edit
+     @change_password_form = ChangePasswordForm.new(object: current_staff_member)
+   end
  end
end
```

<br>

### passwords update アクション

- 自作した `save` メソッドのため，属性値 `object` にパスワードを変更する Staff を格納する
- passwords_controller.rb という切り口から StaffMember モデルの属性値を更新する

```ruby
module Staff
  class PasswordsController < Base
+   def update
+     @change_password_form = ChangePasswordForm.new(staff_member_params)
+     @change_password_form.object = current_staff_member
+     if @change_password_form.save
+       flash.now.notice = 'パスワードを変更しました'
+       redirect_to :staff_account
+     else
+       flash.now.alert = '入力に誤りがあります'
+       render action: 'edit'
+     end
+   end
  end
end
```

<br>

## プレゼンタによるフロントエンドのリファクタ

### プレゼンタ利用の準備

- view ファイルの可読性を高めるためコードを切り出す
- view で使うメソッドをモジュール内で定義する
- 状況に応じて特別な手続きが必要になる
  - view のヘルパーメソッドを使えるようにする
  - HTML ビルダを用意する : lib/html_builder.rb
    - HTML 解析・生成ができる gem nokogiri を使う
  - モデルプレゼンタを用意する : app/presenters/model_presenter.rb
    - `attr_reader` : 読み出し専用にする
    - `delegate`
      - 第１引数 : 委譲元のメソッド
      - オプション `to` : 指定されたメソッドが返す，委譲先のオブジェクト
    - `view_context` : すべてのヘルパーメソッドを持っているオブジェクト
  - フォームプレゼンタを用意する : app/preseters/form_presenter.rb
- ApplicationHelper モジュールに定義しない
  - グローバルに定義されて名前衝突が起きるから
- モデルクラスのインスタンスメソッドに定義しない
  - モデルが肥大化するから
  - 本来の責務 : 正規化・バリデーション・属性値の確認

```ruby
module HtmlBuilder
  def markup(tag_name = nil, options = {})
    root = Nokogiri::HTML::DocumentFragment.parse('')

    Nokogiri::HTML::Builder.with(root) do |doc|
      if tag_name
        doc.method_missing(tag_name, options) do
          yield(doc)
        end
      else
        yield(doc)
      end
    end

    # 許可するタグ・属性を適宜追加する
    sanitize(root.to_html, tags: %w[a table th tr td])
  end
end
```

```ruby
require 'html_builder'

class ModelPresenter
  include HtmlBuilder
  attr_reader :object, :view_context

  delegate :sanitize, :link_to, to: :view_context

  def initialize(object, view_context)
    @object = object
    @view_context = view_context
  end
end
```

```ruby
require 'html_builder'

class FormPresenter
  include HtmlBuilder
  attr_reader :form_builder, :view_context

  delegate :label, :text_field, :date_field, :password_field, :check_box, :radio_button, :text_area, :object, to: :form_builder

  def initialize(form_builder, view_context)
    @form_builder = form_builder
    @view_context = view_context
  end
end
```

<br>

### HtmlBuilder の使い方

- `markup` メソッドは引数なし，または引数にタグ名を指定して呼び出す
- いずれの場合でもブロックを伴う

```ruby
# 引数なし・ネストなし
# <span class='mark'>*</span>印の付いた項目は入力必須です
markup do |m|
  m.span '*', class: 'mark'
  m.text '印の付いた項目は入力必須です'
end
```

```ruby
# 引数なし・ネストあり
# <div class="notes"><span> ... </span></div>
markup do |m|
  m.div(class: 'notes') do
    m.span '*', class: 'mark'
    m.text '印の付いた項目は入力必須です'
  end
end
```

```ruby
# 引数あり・ネストあり
# <div class="notes"><span> ... </span></div>
markup(:div, class: 'notes') do |m|
  m.span '*', class: 'mark'
  m.text '印の付いた項目は入力必須です'
end
```

```ruby
# 生成済みの HTML コードを加える
markup(:div, class: 'notes') do |m|
  m << "<span class='mark'></span>"
  m.text '印の付いた項目は入力必須です'
end
```

<br>

### StaffMember のモデルプレゼンタ

- `object` への委譲
  - 親クラスの `initialize` で StaffMember オブジェクトが渡される
  - `object.family_name` が `family_name` に短縮できる
- `sanitize`
  - 第１引数 : HTML ドキュメント
  - `tags` オプション : 許可する HTML タグ名
  - `attributes` オプション : 許可する HTML 属性名
  - `scrubber` オプション : より自在で強力なタグ・属性操作を行う
  - `raw` および `html_safe` で HTML を表示すると rubocop に注意される
- views/admin/staff_members/index.html.erb からメソッドを呼ぶ
  - 疑似変数 `self` : view ファイルで呼ぶと view_context を参照する

```ruby
class StaffMemberPresenter < ModelPresenter
  delegate :suspended?, :family_name, :given_name, :family_name_kana, :given_name_kana, to: :object

  def full_name
    family_name + ' ' + given_name
  end

  def full_name_kana
    family_name_kana + ' ' + given_name_kana
  end

  # 職員の停止フラグの On/Off を表現する記号を返す
  def suspended_mark
    suspended? ? sanitize('&#x2611;') : sanitize('&#x2610;')
  end
end
```

```erb
<table>
  <% @staff_members.each do |m| %>
    <% p = StaffMemberPresenter.new(m, self) %>
    <tr>
      <td><%= p.full_name %></td>
      <%# ... %>
    </tr>
  <% end %>
</table>
```

<br>

### StaffEvent のモデルプレゼンタ

- 継承した `markup` メソッドで HTML ドキュメントをビルドする
  - `delegate :member, :description, :occurred_at, to: :object`
    - object には event が入る
    - `event.member.family_name` を `member.family_name` と書ける
  - `instance_variable_get` : レシーバが持っているインスタンス変数の値を取得するメソッド
- views/admin/staff_evens/index.html.erb から `table_row` メソッドを呼ぶ
- プレゼンテータによって代替された，不要な部分テンプレートは削除する

```ruby
class StaffEventPresenter < ModelPresenter
  delegate :member, :description, :occurred_at, to: :object

  def table_row
    markup(:tr) do |m|
      unless view_context.instance_variable_get(:@staff_member)
        m.td do
          m << link_to(member.family_name + member.given_name, [:admin, member, :staff_events])
        end
      end
      m.td description
      m.td(class: 'date') do
        m.text occurred_at.strftime('%Y/%m/%d %H:%M:%S')
      end
    end
  end
end
```

```erb
<table>
  <%# ... %>
  <% @events.each do |event| %>
    <%= StaffEventPresenter.new(event, self).table_row %>
  <% end %>
  <%# ... %>
</table>
```

<br>

### StaffMember のフォームプレゼンタ

```ruby

```

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
[Rails アプリの例外ハンドリングとエラーページの表示についてまとめてみた](https://qiita.com/upinetree/items/273ae574f1c021d24c37)  
[Rails の rescue_from で拾えない例外を exceptions_app で処理する](https://qiita.com/ma2ge/items/938d9f8f4839eb336318)  
[ActionDispatch ってなんだろう？](https://blog.eiel.info/blog/2014/03/30/action-dispatch/)  
[Rails のリクエストのライフサイクルと Rack を理解する（翻訳）](https://techracho.bpsinc.jp/hachi8833/2019_10_03/77493)  
[ActiveSupport::Concern でハッピーなモジュールライフを送る](https://www.techscore.com/blog/2013/03/22/activesupportconcern-%E3%81%A7%E3%83%8F%E3%83%83%E3%83%94%E3%83%BC%E3%81%AA%E3%83%A2%E3%82%B8%E3%83%A5%E3%83%BC%E3%83%AB%E3%83%A9%E3%82%A4%E3%83%95%E3%82%92%E9%80%81%E3%82%8B/)  
[Rails 4.2 からは module ClassMethods ではなく Concern#class_methods を使おう](https://blog.yujigraffiti.com/2015/01/rails-42module-classmethodsconcernclass.html)  
[Rails 5.1〜6: ‘form_with’ API ドキュメント完全翻訳](https://techracho.bpsinc.jp/hachi8833/2017_05_01/39502)
[Method: BCrypt::Password#==](https://www.rubydoc.info/github/codahale/bcrypt-ruby/BCrypt%2FPassword:==)  
[7 Patterns to Refactor Fat ActiveRecord Models](https://codeclimate.com/blog/7-ways-to-decompose-fat-activerecord-models/)  
[Ruby と Rails における Time, Date, DateTime, TimeWithZone の違い](https://qiita.com/jnchito/items/cae89ee43c30f5d6fa2c#activesupporttimewithzone%E3%82%AF%E3%83%A9%E3%82%B9)  
[Active Record の関連付け](https://railsguides.jp/association_basics.html)  
[【初心者】Rails の validates の presence でエラーメッセージが重複するのを防ぐ方法](https://qiita.com/lasershow/items/0229855720aaf2be5fc8)  
[Rails ビューの HTML エスケープは#link_to などのヘルパーメソッドで解除されることがある](https://techracho.bpsinc.jp/hachi8833/2016_08_31/25326)  
[Rails で raw HTML を sanitize する](https://fiveteesixone.lackland.io/2015/01/25/sanitize-raw-html-in-rails/)  
[Ruby on Rails 6 実践ガイド](https://www.oiax.jp/jissen_rails6)
