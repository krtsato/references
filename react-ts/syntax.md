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
4. ロジックを抽出して `useHoge()`  という Custom Hook を定義する
5. できあがった Container Component の機能だけをテストする

詳細は repository : react-redux-saga-ts-prac

<br>

## Hooks

### State Hook

Class Component の Local State に相当するものを  
Functional Component でも使えるようにする．

```js
// 初期化
const [count, setCount] = useState(0)

// 状態更新
setCount(100)
setCount(prev => prev + 1) // 引数に関数も指定できる
```

`useState()` を使って複数の Local State を設定するとき  
関数定義は `useState()` をプレーンに連ねて書く．

```js
const [foo, setFoo] = useState(100)
const [bar, setBar] = useState('Inital Bar')
```

`useState()` はグローバルな配列に state 値を追加しており  
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
  doSomething()
  return clearSomething()
}, [watchVar])
```

- 挙動
  - 初回レンダリング時
    - `doSomething()` が実行される
  - 再レンダリング時
    - `watchVar` が変更 : `doSomething()` を実行
    - `watchVar` が不変 : `doSomething()` は未実行
  - アンマウント直前
    - `clearSomething()` を実行
- 第２引数
  - 省略した場合
    - レンダリングの度に `doSomething()` を実行
  - 空配列の場合
    - 初回レンダリング時に `doSomething()` を実行
  - 配列要素に変数を指定した場合
    - 初回レンダリング時に `doSomething()` を実行
    - 前回のレンダリング時と差分があれば `doSomething()` を実行

### useMemo

- 任意の処理結果を再計算せず保持しておける
- 再レンダリングによる不要な処理を避けることが目的
- 副作用処理には `useEffect()` を使う

```js
const memoVal = useMemo(() => {
  calculateSomething()
}, [watchVar])
```

- 挙動
  - 第１引数
    - 高コストな計算
    - コンポーネント
  - 第２引数
    - 省略した場合
      - レンダリングの度に `calculateSomething()`  の結果を返却
      - 無意味なので基本的に省略しない
    - 空配列の場合
      - 初回レンダリング時に `calculateSomething()`  の結果を返却
    - 配列要素に変数を指定した場合
      - 初回レンダリング時に `calculateSomething()`  の結果を返却
      - 前回のレンダリング時と差分があれば `calculateSomething()`  の結果を返却

<br>

Class Component における `shouldComponentUpdate()` の代替機能を実装可能．  
特定の Props が変更されたときだけ任意の子コンポーネントを再レンダリングする．

```js
// props.a/b が変更 -> childA/B を再レンダリング
const Parent = ({a, b}) => {
  const childA = useMemo(() => <ChildA a={a} />, [a])
  const childB = useMemo(() => <ChildB b={b} />, [b])

  return (
    <>
      {childA}
      {childB}
    </>
  )
}
```

### useCallback

- 任意の関数を再定義せず保持しておける
- 再レンダリングによる関数の再定義を避けることが目的
- 副作用処理には `useEffect()` を使う
- 基本的な挙動は `useMemo()` と同じ
  - 戻り値が関数である点が異なる

```js
const memoFunc = useCallback(() => {
  doSomething()
}, [watchVar])
```

例えば子コンポーネントに対しては  
`useCallback()` を通した関数の参照を渡すと良い．

```js
// 通常 props が変更される度に Parent は再レンダリングされる
const Parent = (props) => {  

  // その度に funcA は再定義される
  const funcA = () => {}

  // funcB は初回レンダリング時のみ定義される
  const funcB = useCallback(() => {
    doSomething()
  }, [])

  // funcA が更新 -> ChildA は不要に再レンダリングされる
  // funcB は存続 -> ChildB は再レンダリングされない
  return (
    <>
      <ChildA onClick={funcA} />
      <ChildB onClick={funcB} />
    </>
  )
}
```

- さらにコールバックによって state を管理する場合  
  - 第２引数は状況に応じて指定するか否かよく考える
  - 現状のベストプラクティスは
    - `setState()`  の引数に関数を指定する
    - 第２引数は空配列を指定する

```js
const [count, setCount] = useState(0)

// 直前の state を引数にとるという関数を
// 初回レンダリング時のみ定義する
const handleClick = useCallback(() => {
  setCount(prev => prev + 1)
}, [])

return (
  <>
   <p>count: {count}</p>
   <button onClick={handleClick}>+1</button>
  </>
)
```

以下はアンチパターン

```js
// メモ化されているが count の更新と共に再定義される
const handleClick = useCallback(() => {
  setCount(count + 1)
}, [count])

// ２回目以降のレンダリング時には定義されないが
// 初期状態を参照するため count = 1 のまま
const handleClick = useCallback(() => {
  setCount(count + 1)
}, [])
```

<br>

### useRef

- DOM を参照するための ref オブジェクトを取得する
- current プロパティで変更可能な値を保持する
  - current プロパティを変更しても再レンダリングされない

```js
// １つ前の state を表示する
const Counter = () => {
  const [count, setCount] = useState(0)
  const prevCountRef = useRef(0)
  const prevCount = prevCountRef.current
  
  useEffect(() => {
    prevCountRef.current = count
  })

  return <div>Now: {count}, before: {prevCount}</div>
}

// focus を割り当てる
const InputWithFocusBtn() {
  const inputRef = useRef(null)
  const onButtonClick = () => {
    inputRef.current.focus()
  }

  return (
    <>
      <input ref={inputRef} type="text" />
      <button onClick={onButtonClick}>Focus</button>
    </>
  )
}
```

- DOM ノードと ref の接続（解除）時に任意の処理を実行する場合
  - 前述の `useCallback()` と組み合わせて使用する
  - ref が異なるノードに割り当てられる度にコールバックによって `measuredRef` を再定義する
  - 他の影響でコンポーネントが再レンダリングされても `measuredRed` を不必要に定義せず済む

```js
// ref が DOM ノードに接続されたとき，その高さを表示する
const MeasuredTag = () => {
  const [height, setHeight] = useState(0)

  const measuredRef = useCallback(node => {
    if (node !== null) {
      setHeight(node.getBoundingClientRect().height);
    }
  }, [])

  return (
    <>
      <h1 ref={measuredRef}>H1 Tag</h1>
      <h2>H1 Height : {Math.round(height)}px</h2>
    </>
  )
}
```

<br>

### Other Hooks API

- useContext
- useReducer
- useImperativeHandle
- useLayoutEffect
- useDebugValue

<br>

## 参考文献

[Hooks API Reference](https://reactjs.org/docs/hooks-reference.html)  
[りあクト！ TypeScriptで始めるつらくないReact開発 第2版](https://github.com/oukayuka/ReactBeginnersBook-2.0)  
[雰囲気で使わない React hooks の useCallback/useMemo](https://qiita.com/seya/items/8291f53576097fc1c52a)  
[React Hooks、useStateの更新関数引数には関数を](https://qiita.com/Takepepe/items/7e62cc7d7d8b81ca50db)  
[useCallback in mrsekut-p's Scrapbox](https://scrapbox.io/mrsekut-p/useCallback)
