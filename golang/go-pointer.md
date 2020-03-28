# Golang Pointer

- [値の代入と参照の代入](#値の代入と参照の代入)
  - [Java の例](#java-の例)
  - [Go の例](#go-の例)
- [メモリとアドレス](#メモリとアドレス)
  - [メモリとは](#メモリとは)
  - [値とアドレスの違い](#値とアドレスの違い)
- [メモリ配置とデータへのアクセス](#メモリ配置とデータへのアクセス)
  - [型サイズ](#型サイズ)
  - [オフセット](#オフセット)
  - [コピー代入](#コピー代入)
  - [直接参照と間接参照](#直接参照と間接参照)
    - [ポインタ渡し](#ポインタ渡し)
    - [データへのアクセス](#データへのアクセス)
- [変数とアドレスの関係](#変数とアドレスの関係)
- [バグの回避](#バグの回避)
  - [グローバル変数の書き換え](#グローバル変数の書き換え)
  - [null の実現](#null-の実現)
  - [null 構造体を定義する](#null-構造体を定義する)
- [参考文献](#参考文献)

64 bit マシンを前提とする

<br>

## 値の代入と参照の代入

### Java の例

- プリミティブ型は代入元に影響を与えない
- 参照型は代入元に影響を与える

```java
public class Main {
  public static void main(String[] args) {
    int a = 1;
    int b = a;
    b = 100; // b は プリミティブ int 型
    System.out.printf("a: %d\n", a); // a: 1
    System.out.printf("b: %d\n", b); // b: 100
  }
}
```

```java
class A {
  public int i;
  public A(int i) {
    this.i = i;
  }
}

public class Main {
  public static void main(String[] args) {
    A a = new A(1);
    A b = a;
    b.i = 100; // b は 参照 int 型
    System.out.printf("a.i: %d\n", a.i); // a.i: 100
    System.out.printf("b.i: %d\n", b.i); // b.i: 100
  }
}
```

このように Java では

- 値の代入
  - ある値をコピーして別の値を確保し，それを代入する
- 参照の代入
  - インスタンスなどを通じて参照元を書き換える

によって代入の型を区別している

<br>

### Go の例

- slice / map / chan 以外の型は値の代入
- slice / map / chan 型は参照の代入
- ポインタ型を明示すると参照の代入
  - e.g. int 型のポインタ型は \*int 型と呼ぶ
  - e.g. int a の 参照値は &a (アンパサンド a) と呼ぶ

```go
type A struct {
  i int
}

func main() {
  a := A{i: 1}
}

b := a // 通常の A 型をコピー
b.i = 100 // Java と異なる値の代入
fmt.Printf("a.i: %d\n", a.i) // a.i: 1
fmt.Printf("b.i: %d\n", b.i) // b.i: 100
```

```go
type A struct {
  i int
}

func main() {
  a := A{i: 1}
  b := &a // *A 型である a の参照が &a という形で呼ばれて代入
  b.i = 100 // 参照元を書き換える
  fmt.Printf("a.i: %d\n", a.i) // a.i: 100
  fmt.Printf("b.i: %d\n", b.i) // b.i: 100
}
```

このように Go では

- 基本的には値の代入
- 参照の代入は明示する
- slice / map / chan 型は明示せずとも参照の代入

によって代入の型が区別される

<br>

## メモリとアドレス

### メモリとは

- プログラム実行時にデータが格納される場所
- ポインタはメモリの位置を特定する
  - point : 〜を指差す
- メモリの位置をアドレスという
  - アドレスは単なる正の整数

<br>

### 値とアドレスの違い

- 通常型のメモリ配置
  - `var a int` の場合 `a` のメモリサイズは 8 byte
  - `a` が占有する領域の先頭アドレスは 0xc00009e008
  - アドレス 0xc00009e008 から始まる 8 byte サイズの領域に 0 が格納される

```go
func main() {
  var a int // 通常の int 型
  fmt.Printf("value: %d\n", a) // value: 0
  fmt.Printf("address: %p\n", &a) // address: 0xc00009e008
}
```

- ポインタ型のメモリ配置
  - `var b *int` の場合，`b` のメモリサイズは 8 byte
  - `b` が占有する領域の先頭アドレスは 0xc00000e020
  - アドレス 0xc00000e020 から始まる 8 byte サイズの領域に 0x0 が格納される
    - 実はアドレスを値として格納している
    - `b` には何も代入されていないため，代わりに 0 が代入される
    - 8 byte = 64 bit つまり 0 が 64 個並ぶ
  - ポインタ型の変数の値が 0x0 のとき，nil であることと同義

```go
func main() {
  var b *int // ポインタ int 型
  fmt.Printf("value: %p\n", b) // value: 0x0
  fmt.Printf("address: %p\n", &b) // address: 0xc00000e020
}
```

<br>

## メモリ配置とデータへのアクセス

### 型サイズ

- メモリ上に確保されるサイズは型によって決まる
  - 値によって決まるわけではない
- コンポジット型の比較
  - 非参照型
    - 要素数に応じて型サイズが増減する
    - e.g. array
  - 参照型
    - 要素数に関係なく型サイズは不変
    - e.g. slice, map, chan
- ポインタの型サイズはすべて 8 byte

```go
type A struct {
  i int // 型サイズ 8 byte
  f float64 // 型サイズ 8 byte
  s string // 型サイズ 16 byte
}

type B struct {
  s string // 型サイズ 16 byte
  a A // 型サイズ 8 + 8 + 16 byte
}

func main() {
  a := A{}
  fmt.Printf("A: %d\n", unsafe.Sizeof(a)) // A: 32
  fmt.Printf("A pointer: %d\n", unsafe.Sizeof(&a)) // A pointer: 8
  b := B{}
  fmt.Printf("B: %d\n", unsafe.Sizeof(b)) // B: 48
  fmt.Printf("B pointer: %d\n", unsafe.Sizeof(&b)) // B pointer: 8
}
```

<br>

### オフセット

- 先頭アドレスからのサイズ距離
  - e.g. 構造体の先頭アドレスから N byte 目にポインタがアクセスしてデータを取得する

```go
type A struct {
  i int // 型サイズ 8 byte
  f float64 // 型サイズ 8 byte
  s string  // 型サイズ 16 byte
}

func main() {
  a := &A{}
  fmt.Printf("a.i starts: %d\n", unsafe.Offsetof(a.i)) // a.i starts: 0
  fmt.Printf("a.f starts: %d\n", unsafe.Offsetof(a.f)) // a.f starts: 8
  fmt.Printf("a.s starts: %d\n", unsafe.Offsetof(a.s)) // a.s starts: 16
}
```

<br>

### コピー代入

- 変数に値を代入するとき必ずコピーが作成される
  - まず型サイズの領域を確保する
  - その領域にコピーした値を代入する
  - 値は同じだがアドレスが異なる
  - ポインタ型の変数であってもコピー代入は発生する

![copy-subst-01](/images/golang/copy-subst-01.png)  
![copy-subst-02](/images/golang/copy-subst-02.png)

```go
type A struct {
  i int
}

func main() {
  var a1 A
  a2 := a1
  fmt.Printf("a1: %p\n", &a1) // a1: 0xc00001e030
  fmt.Printf("a2: %p\n", &a2) // a2: 0xc00001e038
}
```

<br>

### 直接参照と間接参照

メモリ上のデータにアクセスする方法のこと．  
計算機はこれらを組み合わせて処理を行なっている．

- 直接参照
  - 通常型を指定すると使われる
  - 直接データへアクセスして読み書きする
- 間接参照
  - ポインタ型を指定すると使われる
  - アドレス値を経由して間接的にデータへアクセスし，読み書きする

```go
type A struct  {
  i int
}

func main() {
  a1 := A{}
  a2 := new(A)
  var a3 A
  var a4 *A
  fmt.Printf("a1: %T\n", a1) // a1: main.A 通常型
  fmt.Printf("a2: %T\n", a2) // a2: *main.A ポインタ型
  fmt.Printf("a3: %T\n", a3) // a3: main.A 通常型
  fmt.Printf("a4: %T\n", a4) // a4: *main.A ポインタ型
}
```

<br>

#### ポインタ渡し

- ポインタ型の変数をある関数に渡すこと
- アドレスを値として格納する

![pointer-pass](/images/golang/pointer-pass.png)

```go
type A struct {
  i int
}

func test(a2 *A) {
  fmt.Printf("a2 value: %p\n", a2) // a2 value: 0xc00008e000
  fmt.Printf("a2 address: %p\n", &a2) // a2 address: 0xc000096000
}

func main() {
  a1 := new(A)
  fmt.Printf("a1 value: %p\n", a1) // a1 value: 0xc00008e000
  fmt.Printf("a1 address: %p\n", &a1) // a1 address: 0xc000088010
  test(a1)
}
```

<br>

#### データへのアクセス

- 直接参照
  - `a1` のメモリ領域 16 byte を確保する
  - `f` の 8 byte を書き換える

![direct-ref](/images/golang/direct-ref.png)

```go
type A struct {
  i int
  f float64
}

func main() {
  a1 := A{}
  a1.f = 2.4
  fmt.Printf("a1.f: %f\n", a1.f) // a1.f: 2.400000
  fmt.Printf("a1 address: %p\n", &a1) // a1 address: 0xc00001e030
  fmt.Printf("a1 size: %d\n", unsafe.Sizeof(a1)) // a1 size: 16
}
```

![indirect-ref](/images/golang/indirect-ref.png)

- 関節参照
  - `a2` のメモリ領域 8 byte を確保する
  - 構造体 `A` のメモリ領域 16 byte を確保する
  - `A` の先頭アドレスの値を `a2` の領域に格納する
  - `a2` の値から `A` の先頭アドレスを特定して `f` の 8 byte を書き換える

```go
type A struct {
  i int
  f float64
}

func main() {
  a2 := new(A)
  a2.f = 2.4
  fmt.Printf("a2.f: %f\n", a2.f) // a2.f: 2.400000
  fmt.Printf("a2 address: %p\n", &a2) // a2 address: 0xc000088018
  fmt.Printf("a2 value: %p\n", a2) // a2 value: 0xc000092020
  fmt.Printf("a2 size: %d\n", unsafe.Sizeof(*a2)) // a2 size: 16
}
```

<br>

## 変数とアドレスの関係

- ポインタ型の変数代入時に，その変数アドレスと代入元の先頭アドレスは一致しない
- どの領域の先頭アドレスを見ているか明確に把握することが大切

アンチパターン

![var-add-rel-01](/images/golang/var-add-rel-01.png)  

![var-add-rel-02](/images/golang/var-add-rel-02.png)  

![var-add-rel-03](/images/golang/var-add-rel-03.png)  

![var-add-rel-04](/images/golang/var-add-rel-04.png)

```go
type A struct {
  i int
}

func main() {
  list := []A{{i: 1}, {i: 2}, {i: 3}}
  pList := make([]*A, 0, len(list))

  for _, v := range list {
    // ポインタ型のアドレスを Slice に格納する
    pList = append(pList, &v)
  }

  // 1 2 3 という出力にならない
  // 格納値ではなく最後に経由した v の値を出力している
  for _, v := range pList {
    fmt.Println(v.i) // 3 3 3
  }
}
```

正しい実装

![var-add-rel-05](/images/golang/var-add-rel-05.png)

```go
type A struct {
  i int
}

func main() {
  list := []A{{i: 1}, {i: 2}, {i: 3}}
  pList := make([]*A, 0, len(list))

  for i := range list {
    // インデックスで要素を指定して Slice に格納する
    pList = append(pList, &list[i])
  }

  for _, v := range pList {
    fmt.Println(v.i) // 1 2 3
  }
}
```

<br>

## バグの回避

### グローバル変数の書き換え

- ある goroutine でデータを書き換えると，その後にアクセスする goroutine は書き換えられたデータを参照してしまう
- e.g. Go で net/http パッケージを利用する場合
  - リクエスト ごとに goroutine が作成される
    - goroutine はスレッド型の並行処理機構
    - メモリ空間が共有される
  - グローバル変数は複数の goroutine からアクセスされる
  - 意図しないデータの書き換えが発生し得る

```go
type Data struct {
  number int
}

var onMemoryData *Data

func initData() {
  onMemoryData = &Data {
    number: 100,
  }
}

func getData() Data {
  return *onMemoryData
}

func main() {
  initData()

  list := make([]Data, 5)
  for i := range list {
    data := getData()
    data.number += i // 間接参照でデータ領域を書き換えてしまう
    list[i] = *data
  }

  fmt.Printf("%v", list) // [{100} {101} {102} {103} {104}]

}
```

正しい実装

```go
type Data struct {
  number int
}

var onMemoryData *Data

func initData() {
  onMemoryData = &Data {
    number: 100,
  }
}

func getData() *Data {
  return onMemoryData
}

func main() {
  initData()

  list := make([]Data, 5)
  for i := range list {
    data := getData()
    data.number += i // 間接参照でデータ領域を書き換えてしまう
    list[i] = *data
  }

  // [{100} {101} {102} {103} {104}] という出力にならない
  fmt.Printf("%v", list) // [{100} {101} {103} {106} {110}]
}
```

アドレスを返却するとき，データが書き換えられる可能性があることを意識する

<br>

### null の実現

- 通常型では null を表現できない
  - 通常型の値はゼロ値がセットされる
  - ゼロ値であるか null であるか区別できない
    - e.g. ゼロ値は int 型で 0，string 型で空文字になる
- ポインタ型では表現できる
  - `<nil>` と表示される

下記では JSON データをパースするときに null を表現する

```go
import (
  "encoding/json"
  "fmt"
)

type Data struct {
  Number1 int  ‘json:"number1"‘ // 通常型
  Number2 *int ‘json:"number2"‘ // ポインタ型
}

func main() {
  b1 := []byte(‘{"number1":100, "number2":200}‘) // JSON データを仮作成する
  data1 := Data{}
  json.Unmarshal(b1, &data1) // JSON データを構造体にパースする
  fmt.Printf("data1.Number1: %d\n", data1.Number1) // data1.Number1: 100
  fmt.Printf("data1.Number2: %d\n", *data1.Number2) // data1.Number2: 200

  b2 := []byte("{}") // null の JSON をを仮作成する
  data2 := Data{}
  json.Unmarshal(b2, &data2) // JSON データを構造体にパースする
  fmt.Printf("data2.Number1: %d\n", data2.Number1) // data2.Number1: 0
  fmt.Printf("data2.Number2: %v\n", data2.Number2) // data2.Number2: <nil>
}
```

<br>

### null 構造体を定義する

- しかし本当は通常型で null を表現したい
  - ポインタ型は意図せずデータの書き換えが発生し得るから
- `database/sql` パッケージに `NullString` という構造体が定義されている
  - `bool` フィールドが false : null
  - `bool` フィールドが true : null でない
    - `string` フィールドを読み込んでデータを取得する

```go
package sql

type NullString struct {
  String string
  Valid  bool
}
```

- この構造体を参考にして int 通常型で null を表現する NullInt を独自定義できる
- `UnmarshalJSON` で独自定義の型に対応させる
  - json パッケージの `json.Unmarshaler` インターフェースを満たすようにする
  - 関数 `json.Unmarshal` 内部で `UnmarshalJSON` が呼ばれるようになる

```go
package json

type Unmarshaler interface {
  UnmarshalJSON([]byte) error // これを満たす独自定義の型を作成する
}
```

```go
import (
  "encoding/json"
  "fmt"
)

type NullInt struct {
  Int int
  Valid bool
}

// 独自定義の型に対応させる
func (n *NullInt) UnmarshalJSON(b []byte) error {
  i, err := strconv.ParseInt(string(b), 10, 64)
  if err != nil {
    return err
  }
  n.Int = int(i)
  n.Valid = true
  return nil
}

type Data struct {
  Number1 int ‘json:"number1"‘
  Number2 NullInt ‘json:"number2"‘
}

func main() {
  b1 := []byte(‘{"number1":100, "number2":200}‘)
  data1 := Data{}
  json.Unmarshal(b1, &data1)
  fmt.Printf("data1.Number1: %d\n", data1.Number1) // data1.Number1: 100
  fmt.Printf("data1.Number2: %v\n", data1.Number2) // data1.Number2: {200 true}

  b2 := []byte("{}")
  data2 := Data{}
  json.Unmarshal(b2, &data2)
  fmt.Printf("data2.Number1: %d\n", data2.Number1) // data2.Number1: 0
  fmt.Printf("data2.Number2: %v\n", data2.Number2) // data2.Number2: {0 false}

  if data1.Number2.Valid {
    fmt.Printf("data1.Number2 value: %d\n", data1.Number2.Int) // data1.Number2 value: 200
  } else {
    fmt.Println("data1.Number2 is null")
  }

  if data2.Number2.Valid {
    fmt.Printf("data2.Number2 value: %d\n", data2.Number2.Int)
  } else {
    fmt.Println("data2.Number2 is null") // data2.Number2 is null
  }
}
```

- この手法は手間がかかる
- ゼロ値を null と見なした方が安全な場合もある
  - e.g. int 型の値が 0 を取らないの場合 0 を null と見なせる

<br>

## 参考文献

[JSON and Go](https://blog.golang.org/json)  
[GO のポインタを完全に理解する本](https://techbookfest.org/product/5039923363053568)
