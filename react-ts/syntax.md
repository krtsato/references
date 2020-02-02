# React

## 基礎知識

### Local State

- コンポーネント自身が内部に持つ状態
- どのコンポーネントからもアクセスできる State と区別するための呼称

<br>

### ライフサイクル

- Mounting
  - コンポーネントが生成され DOM ノードに挿入される段階
- Updating
  - 変更を検知してコンポーネントが再レンダリングされる段階
  - Props の変更
  - Local State の変更
- Unmounting
  - コンポーネントが DOM ノードから削除される段階
- Error Handling
  - そのコンポーネント自身および子孫コンポーネントのエラーを補足する

![react-process](/images/react-process.png)

詳細追加予定

<br>

### Functional Component

- Stateless Functional Component として台頭
- Hooks を用ると state を持てるため FC と呼ばれる

cf. Class Component

- クラス内の this の挙動が難解
- 記述が冗長
- ライフメソッドの時系列が複雑化しがち
- 導入予定の各種最適化が難しい

<br>

### コンポーネントの作り方

Class Component / Function Component とは別の切り口で  
コンポーネントを分類する  

- Presentational Component : 見た目を担うコンポーネント
- Container Component : 処理を担うコンポーネント

![two-kinds-of-comp](/images/pc-comp.png)

闇雲に作らないための１つの指標

1. FC で見た目だけを整えた Presentational Component を作る
2. Presentational Component からスタイルガイドを作る
3. Presentational Component を import して Hooks や HOC で機能を追加する
4. できあがった Container Component の機能だけをテストする

<br>

## Hooks

### State Hook

Class Component の Local State に相当するものを  
Functional Component でも使えるようにする機能

```js
// 初期化
const [count, setCount] = useState(0);

// 状態更新
setCount(100);
setCount(prev => prev + 1);
```

useState() を使って複数の Local State を設定するとき  
関数定義は useState() をプレーンに連ねて書く．

```js
const [foo, setFoo] = useState(100);
const [bar, setBar] = useState('Inital Bar');
```

useState() はグローバルな配列に state 値を追加しており，
条件文などでくるんでしまうと配列の順番がおかしくなるため．

<br>

### Effect Hook

- Class Component のライフサイクルメソッドに相当するものを実現する
- 副作用を扱う
  - API データの取得
  - 手動での DOM の改変
  - ログの記録

```js
useEffect(() => {
  doSomething();
  return clearSomething();
}, [watchVar]);
```

- 挙動
  - 初回レンダリング直後
    - doSomething() が実行される
  - 再レンダリング時
    - watchVar が変更 : doSomething() を実行
    - watchVar が不変 : doSomething() は未実行
  - アンマウント直前
    - clearSomething() を実行
- 第２引数
  - 省略した場合
    - レンダリングの度に doSomething() を実行
  - 空配列の場合
    - 初回レンダリング直後でのみ doSomething() を実行
