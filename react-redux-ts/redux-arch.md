# Redux

記述途中です．  
repository: react-redux-saga-ts-prac で作業してます．

<br>

## 設計方式

- rails way
- ducks
- re-ducks

Redux は中規模以上のアプリを堅く作ることに長けている．
しかし，考えなしに rails way を採用するとディレクトリ管理が大変．

ducks で構成してもスケールしない．  
１ファイルの記述量が増えすぎる上，非同期処理を組み込みにくい．  

<br>

## re-ducks

ドメインごとに以下のファイルを持つ．
actions と action が本来の責務に集中できる．

- index.ts
- types.ts
- actions.ts
- reducers.ts
- operations.ts
- selectors.ts

<br>

### Index の責務

- いろいろまとめる
- Containers に提供する
<br>

### Operation の責務

- １つ以上の Action を組み合わせた関数
- 必要に応じて複雑な Operation を作る場合もある
  - e.g. redux-thunk/saga などの非同期通信
- 同じドメインの Action は参照して良い
- 別のドメインの Action は Operation を経由して取得する
- 疑問
  - 非同期処理が Operation にまとめられるならば，わざわざ redux-thunk ではなく redux-saga を用いるメリットはあるのか？

<br>

### Selector の責務

- state から必要な値を算出する関数
  - インターフェース `(state) => return value`
- state を扱う各所をシンプルに保つことが目的
- 既存の state から算出できる値は state に保存せず Selector から取得する
- 別のドメインの Selector から値を直接参照しても良い
- 疑問
  - Redux Hooks の `useSelector()` では？
  - `useSelector` を使った場合，Container と Selector の役割が重複する？

<br>

### Reducer のデザイン

モノリシックな reducers を作らず分割する．

- AppState
  - ドメインデータとは別の Reducer として用意する
  - アプリ全体の state を管理する
    - e.g. isLoading など
- DomainState
  - ドメイン特有の state
    - e.g. Task, User など
- UIState
  - UI特有のstate
  - e.g. Modal, DisplayToggler など
  - DomainState との境界が曖昧になる場合もあるので注意

<br>

### Action のデザイン

１つの Reducer に対して１つの Action 群を定義すると見通しが良い．
<br>

### Types のデザイン

- Store / Dispatch
  - Containers での度重なる型付けに備えて予め定義する
- ActionTypes
  - 各ドメインの全 Action が登録される
  - Reducer の switch 文で比較するため
  - 新しい Action を作成したら追加していく
  - フォーマットは `DOMAIN_NAME/ACTION_NAME`
- Action 型の集合体
  - 複数の Action の型が登録される
  - HogeActions : HogeReducer = 1 : 1
  - 様々な場所から参照される
    - e.g. Action Creator, Reducer, Containers 配下など

<br>

### 設計順序（自己流）

1. ドメイン設計
2. Types に落とし込む
3. Action Creator を書く
4. Reducer を書く
5. Operation を書く
<br>

## 参考文献

[Scaling your Redux App with ducks](https://www.freecodecamp.org/news/scaling-your-redux-app-with-ducks-6115955638be/)  
[re-ducks-examples by jthegedus](https://github.com/jthegedus/re-ducks-examples)  
[Re-ducksパターン：React + Redux のディレクトリ構成ベストプラクティス](https://noah.plus/blog/021/)  
[React/Reduxで秩序あるコードを書く](https://speakerdeck.com/naoishii/reduxde-zhi-xu-arukodowoshu-ku)  
[React/Redux約三年間書き続けたので知見を共有します](https://tech.enigmo.co.jp/entry/2018/12/04/140027)  
[ReactをTypeScriptで書く4: Redux編](https://www.dkrk-blog.net/javascript/react_ts04)
