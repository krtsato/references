# Web Terms

Web 技術を説明できるようになるための記述  

## HTTP (Hyper Text Transfer Protocol)

Web 上のデータをサーバとクライアントの間で通信するときのプロトコル。

## HTTPS (Hyper Text Transfer Protocol Secure)

HTTPS では SSL というプロトコルによって通信内容が暗号化されていた。現在はSSLではなくTLS (Transport Layer Security) というプロトコルが使われている。SSL という言葉が長い間使われてきたため，現在でも SSL や SSL/TLS と呼ばる場合がある。

### SSL (Secure Socket Layer)

公開鍵・秘密鍵を使って送受信者が同じ共通鍵を用意し，その共通鍵でデータの暗号・復号化を行う。

## ソケット通信

インターネットは TCP/IP という通信プロトコルを利用する。そのTCP/IPをプログラムから利用するには，プログラムの世界と TCP/IP の世界を結ぶ特別な出入り口 (ソケット) が必要となる。  
ソケットを使用することにより，アプリケーションはトランスポート層やネットワーク層と対話できる。
