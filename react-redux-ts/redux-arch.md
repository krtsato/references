# React Redux

- [Redux の設計方式](#redux-の設計方式)
- [Re-ducks](#re-ducks)
  - [Index の責務](#index-の責務)
  - [Operation の責務](#operation-の責務)
  - [Selector の責務](#selector-の責務)
  - [Reducer のデザイン](#reducer-のデザイン)
  - [Action のデザイン](#action-のデザイン)
  - [Types のデザイン](#types-のデザイン)
  - [Store の責務](#store-の責務)
  - [設計順序（半自己流）](#設計順序半自己流)
    - [ドメイン設計での注意](#ドメイン設計での注意)
    - [Actions / Reducers での注意](#actions--reducers-での注意)
    - [Types での注意](#types-での注意)
    - [Component での注意](#component-での注意)
    - [Operations での注意](#operations-での注意)
    - [Selectors での注意](#selectors-での注意)
- [参考文献](#参考文献)

ソースコード : [react-redux-ts-prac](https://github.com/krtsato/react-redux-ts-prac)

<br>

## Redux の設計方式

- Rails Way
- Ducks
- Re-ducks

Redux は中規模以上のアプリを堅く作ることに長けている．  
しかし，rails way を採用するとディレクトリ管理が大変．
Ducks Pattern で構成してもスケールしない．  
１ファイルの記述量が増えすぎる上，非同期処理を組み込みにくい．

<br>

## Re-ducks

- Pros
  - 一言で表すと Reducer ドリブンという印象
  - 責務分散により Reducer と Action が本来の責務に集中できる
    - ドメインごとにディレクトリを掘る
    - ドメインごとに以下のファイル群を持つ
      - index.ts
      - types.ts
      - actions.ts
      - reducers.ts
      - operations.ts
      - selectors.ts
  - 非同期処理を行う Operations を Custom Hook として再利用できる
    - middleware 不要で依存や学習コストを低減できる

- Cons
  - 設計段階からドメイン定義をある程度固める必要がある
  - ディレクトリ管理はしやすいがファイル数が増える
    - Re-ducks というより Redux の記述量の問題でもある
  - Reducer を簡潔にしていくと各所の命名が難しくなる
    - 前提
      - 同ドメイン下でも [Reducer は分割されるべき](#reducer-のデザイン)
      - [Action 集合型 : Reducer = 1 : 1 の法則](#action-のデザイン)
    - 前提により，Types において Action 集合型の命名が難しくなる
    - 型や変数を適切に命名しないと [Selectors にしわ寄せが来る](#selectors-での注意)
    - [Custom Hook で非同期処理をする場合](#operations-での注意)，その中で dispatch する必要がある
      - Container Component に返されるのは dispatch でなく dispatch を実行する関数 `Promise<void>`
      - つまり dispatch の実行場所を完全には統一できない

- 改善案
  - まず接尾辞にルールを設けるところから始める
  - これだけでもかなり楽になる
    - Reducer : HogeRed
    - Action 集合型 : HogeActions
    - Action : FugaAct
    - Operations : FugaOpe
    - Selectors : FugaSel
    - Container で定義するイベントハンドラ : FugaHdl
  - Custom Hook の中で dispatch する関数を `const dipatchFugaRes` と名付ける
    - res : result あるいは response の意味
    - Container 側では `dispatchFugaRes()`  に加えて `useDispatch()` による `dispatch(fugaOpe)` が見られる
    - 直感的に dispatch されたと分かるため悪くない

Cons の数だけ見ると微妙だが，細かい話が中心の Cons に比べて Pros の効果が絶大なので，少なくとも個人開発では Re-ducks パターンを採用したい．

<br>

### Index の責務

- Re-ducks における各ドメインに配置される
- Containers に Operations や Selectors などを re-export する
- TypeScript 3.8 で `export * as Hoge from "path/to/hoge"` がサポートされた
  - @typescript-eslint/parser や prettier の対応待ち

<br>

### Operation の責務

- １つ以上の Action を組み合わせた関数
- 必要に応じて複雑な Operation を作る場合もある
  - e.g. redux-thunk/saga などの非同期通信
- 同じドメインの Action は参照して良い
- 別のドメインの Action は Operation を経由して取得する
- 疑問
  - 非同期処理が Operation にまとめられるならば，わざわざ redux-thunk ではなく redux-saga を用いるメリットはあるのか？
  - 個人的には そもそも middleware を使わない選択肢が丸い印象
    - 非同期処理を Custom Hook に閉じ込めて取り回せば良いのでは
  - [非同期処理の middleware を用いるメリットは一応ある](#operations-での注意)

<br>

### Selector の責務

- state から必要な値を算出する関数
  - インターフェース `(state) => return value`
  - state の型は Store State (RootState) である
  - [定義の仕方](#store-の責務)は少し特殊になる
- state を扱う各所をシンプルに保つことが目的
- 既存の state から算出できる値は state に保存せず Selector から取得する
- 別のドメインの Selector から値を直接参照しても良い
- 疑問
  - Redux Hooks の `useSelector()` では？
  - `useSelector` を使った場合，Container と Selector の役割が重複する？
    - 責務が異なるため重複しない
      - Containers では関数のメモ化や dispatch を記述する
      - import された Operations と Selectors が連携して記述される

<br>

### Reducer のデザイン

モノリシックな reducers を作らず分割する．

- Domain State
  - ドメイン特有の state
    - e.g. Task, User など
  - ドメイン共通の state
    - app ドメインに切り分けて管理する
      - e.g. errorMsg, SuccessInfo など
- UI State
  - UI 特有の state
  - e.g. Modal, DisplayToggler など
  - DomainState との境界が曖昧になる場合もあるので注意

<br>

### Action のデザイン

- １つの Reducer に対して１つの Action 集合型を定義すると見通しが良い
  - Action 集合型とは Types で定義される Action の Lookup Type
    - e.g. AddTodo, DelTodo 型があるとき `TodoManage["AddTodo"]` を集合型と呼ぶことにする
  - [前述](#reducer-のデザイン)したとおり，Reducer の分割に伴って命名に迷いがち
    - 単に todoReducer としても Todo の何なのか分からない
    - Reducer を簡潔に書くことが最優先なので多少の不便には目を瞑る
    - DRY の観点から言えば「怪我の功名」的な良さでもある

<br>

### Types のデザイン

- Store / Dispatch
  - 必要に応じて記述する
  - Redux Hook を利用する場合, これらの型を定義する必要はない
- ActionTypes
  - ドメインごとの全 Action が登録される
  - Reducer の switch 文で比較するため
  - 新しい Action を作成したら追加していく
  - フォーマットは `APP_NAME/DOMAIN_NAME/ACTION_NAME`
    - アプリの規模に合わせて考える
- HogeLiteral
  - ActionTypes とは異なり，action.type の比較ではなく State の取得に備えるため
  - 各所で参照が必要となった際に定義していけば良い
    - e.g. Selectors, Containers 配下など
- Action 型の集合型
  - 複数の Action の型が登録される
    - e.g. `HogeAction["Fuga"]` と参照できるよう, 型は Lookup Type にまとめる
  - (type HogeActions) : (const HogeReducer) = 1 : 1
  - 各所から参照される
    - e.g. Action Creator, Reducer, Containers 配下など

<br>

### Store の責務

ここまではドメイン内の設計について記述したが  
Store ではドメインを統合して RootState を定義するため  
ファイルは以下のように配置される．

```zsh
.
├── src
│   ├── components
│   │    └── ...
│   │
│   ├── containers
│   │    └── ...
│   │
│   └── reducks
│        ├── store.ts
│        ├── types.d.ts
│        ├── domainHoge/
│        ├── domainFuga/
│        └── ...
```

- store.ts
  - Redux Toolkit の組み込み
  - ミドルウェアの組み込み
  - `combineReducers()` による Reducers の統合
    - 引数となる Reducer オブジェクトの key にドメイン名を指定すると，Selectors で明快な呼び出しが可能になる
      - e.g. state.users.id
        - state : RootState
        - users : store.ts で定義する key
        - id : ドメイン内の Reducer で同様に定義する key

```js
export const rootReducer = combineReducers({
  domainHoge: domainHogeReducers,
  domainFuga: domainFugaReducers
  // ...
})
```

- types.d.ts
  - Selectors で多用される RootState 型を定義する
  - store.ts 内で定義すると循環インポートが発生するため
  - 型定義ファイルに名前空間を宣言することで解決

```js
// types.d.ts
// import the above rootReducer

declare namespace Root {
  export type State = ReturnType<typeof rootReducer>
}

// selectors.ts
// import the above Root namespace
const hogeSelector = (state: Root.State): someReturn => {}
```

<br>

### 設計順序（半自己流）

重要なのは 1 ~ 6 の作業．

1. ドメイン設計
2. Action Creator を書く
3. Types を書く
4. Components を書く
5. Operations を途中まで書く
6. Selectors を途中まで書く
7. Reducers を書く
8. Store を書く
9. Containers を書く
10. Operations と Selectors を仕上げる
11. 全体を仕上げる

<br>

#### ドメイン設計での注意

- app ドメインは共通の状態管理を行う
  - エラーや通知処理を担う
- 疑問 `isLoading` を管理させる必要があるか?
  - 今のところ必要性を感じない
  - `isLoading` を用いる場面は非同期処理中
  - 非同期処理を Custom Hooks に閉じ込めると, その中で `isLoading` も取り回せる
    - Container において `[isLoading, asynchResult] = useCustomHook` のように結果を取得する
  - あとは `isLoading` に応じて Component を振り分けるだけ

<br>

#### Actions / Reducers での注意

- types.js を見ながら簡潔に書くのみ
  - actions.ts : plain object を返すのみ
  - reducers.ts : 更新された state object を返すのみ
- これが Re-ducks の責務分散による偉功
  - 複雑性は Selectors と Operations が吸収する

<br>

#### Types での注意

- Action 型の命名は比較的難しいので, 何度も見直す
- d.ts ファイルにすると, Reducers で便利な [`const ActionTypes`](#types-のデザイン) による比較ができなくなる

<br>

#### Component での注意

- 変数や型が取得できない場合, 適当に定義して後で差し替える
- 時間をかけるのはロジック部分が完成してから

<br>

#### Operations での注意

- Operations 内では, 基本的に dispatch させず シンプルに Actions を返す
  - Operations はどうしても肥大化しがち
  - Re-ducks でもたらされた責務の分散をいかしたい
- ただし, 非同期処理を行う Operation は Custom Hook にする
  - このとき, その Operation の中で dispatch する
  - 理由: dispatch は Promise object でなく plain object を返す必要があるから
    - Container 側で `dispatch(asyncResult)` をするとエラー
    - 非同期処理の middleware は，dispatch を拡張するという点に置いて利用価値がある
    - しかし Custom Hook を作れば事足りる場合が多い
    - 依存が減るという意味でも Custom Hook を作れば当面は良さそう

<br>

#### Selectors での注意

- コードを書くのは簡単だが, 命名が下手だとここにしわ寄せが来る
- `(state: Root.State) => state.domainName.reducerName.dataName` を綺麗に書きたい
  - `state.todos.todos` とかになりがち
  - [Action 集合型 : Reducer = 1 : 1 の法則](#types-のデザイン)
    - domainName : store.ts で Reducers をまとめるときに指定
    - reducerName : reducers.ts で `combineReducers()` するときに指定
    - dataName : types.ts で payload オブジェクト内に指定

<br>

## 参考文献

[Scaling your Redux App with ducks](https://www.freecodecamp.org/news/scaling-your-redux-app-with-ducks-6115955638be/)  
[re-ducks-examples by jthegedus](https://github.com/jthegedus/re-ducks-examples)  
[Re-ducksパターン：React + Redux のディレクトリ構成ベストプラクティス](https://noah.plus/blog/021/)  
[React/Reduxで秩序あるコードを書く](https://speakerdeck.com/naoishii/reduxde-zhi-xu-arukodowoshu-ku)  
[React/Redux約三年間書き続けたので知見を共有します](https://tech.enigmo.co.jp/entry/2018/12/04/140027)  
[ReactをTypeScriptで書く4: Redux編](https://www.dkrk-blog.net/javascript/react_ts04)  
[Reduxの非同期処理にReact Hooksを使う](https://yo7.dev/articles/redux-async-hook)  
[非同期処理にredux-thunkやredux-sagaは必要無い](https://qiita.com/Naturalclar/items/6157d0b031bbb00b3c73)  
[おすすめ自作 React hooks集2 (useRouter)](https://qiita.com/pikohideaki/items/4238dd17818e58c33799)
