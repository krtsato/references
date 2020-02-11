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

### re-ducks

ドメインごとに以下のファイルを持つ
actions と action が本来の責務に集中できる．

- index.ts
- types.ts
- actions.ts
- reducers.ts
- operations.ts
- selectors.ts

<br>

## 参考文献

[Scaling your Redux App with ducks](https://www.freecodecamp.org/news/scaling-your-redux-app-with-ducks-6115955638be/)  
[re-ducks-examples by jthegedus](https://github.com/jthegedus/re-ducks-examples)  
[Re-ducksパターン：React + Redux のディレクトリ構成ベストプラクティス](https://noah.plus/blog/021/)  
[React/Reduxで秩序あるコードを書く](https://speakerdeck.com/naoishii/reduxde-zhi-xu-arukodowoshu-ku)  
[React/Redux約三年間書き続けたので知見を共有します](https://tech.enigmo.co.jp/entry/2018/12/04/140027)  
[ReactをTypeScriptで書く4: Redux編](https://www.dkrk-blog.net/javascript/react_ts04)
