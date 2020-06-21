# Golang Error

## Go > 1.13 のエラーハンドリング

- `fmt.Errorf("%w", err)` でエラー内容をラップして返却する
  - `fmt.Errorf("%s", err)` はエラー内容を単純な文字列でしか扱えない
  - ラップしたエラーを返却すれば，型やフィールド値を保持できる
- `errors.UnWrap(err)`
  - エラーの元情報を取得するためにラップを剥がす
  - 最後にラップされたエラーが返却される
- `errors.Is(errA, errB)`
  - errA が errB を同一インスタンスとしてラップしているか判定する
- `errors.As(errA, &errB)`
  - errA が errB の型をラップしているか判定する
  - いちばん使用機会が多そう
- 前提
  - Go の error 型を満たすために準備が必要
    - エラーの型 (構造体) を定義する
    - Error() メソッドをレシーバごとに定義する

```go
package main

import (
  "errors"
  "fmt"
)

// 独自のエラー構造型を準備する
type fooErrType struct {
  statusCode  int
  msg  string
}

// 独自のエラーメソッドを準備する
func (foo *fooErrType) Error() string {
  return fmt.Sprintf("statusCode: %d, msg: %s\n", foo.statusCode, foo.msg)
}

func main() {
  // 下位関数から err が返却される
  err := lowFunc()
  var fooErr *fooErrType

  // Unwrap を使って 1 次情報を取得する
  // Wrap した回数分 Unwrap する必要がある
  fmt.Printf("TopLevel: %v\n", err)
  fmt.Printf("Unwrap: %v\n", errors.Unwrap(err))

  // err が同一インスタンスをラップしているか判定する
  // この場合 false になる
  // fmt.Errorf("%w", fooErr) していないから
  if errors.Is(err, fooErr) {}

  // err が fooErrType 型をラップしているか判定する
  // true の場合 fooErr がその値を参照する
  if errors.As(err, &fooErr) {
    // Error() が文字列を返却する
    fmt.Printf("fooErr's Error(): %v\n", fooErr)
    // エラー構造体のフィールド値を取得する
    fmt.Printf("fooErr's field a: %v\n", fooErr.msg)
  }
}

// エラーの 2 次情報をラップして返却する
func lowFunc() error {
  err := lowerFunc()
  if err != nil {
    return fmt.Errorf("lowerFunc has a error -> %w", err)
  }

  return nil
}

// エラーの 1 次情報をラップして返却する
func lowerFunc() error {
  err := lowestFunc()
  if err != nil {
    return fmt.Errorf("lowestFunc has a error -> %w", err)
  }

  return nil
}

// 独自のエラー構造体にエラー情報を格納する
func lowestFunc() error {
  err := &fooErrType{statusCode: 404, msg: "foo not found"}
  return err
}
```

出力

```txt
TopLevel: lowerFunc has a error -> lowestFunc has a error -> statusCode: 404, msg: foo not found

Unwrap: lowestFunc has a error -> statusCode: 404, msg: foo not found

fooErr's Error(): statusCode: 404, msg: foo not found

fooErr's field a: foo not found
```

<br>

## 参考文献

[golang error handling (Go1.13)](https://qiita.com/yuukiyuuki327/items/4d31d62f6b70476082f3)  
[[go] エラーのラッピング](https://qiita.com/egawata/items/fcf3f5918f9a5284dc2d)
