# 007-http-networking-internals

REST APIの理解を一段深めるために、HTTPとネットワーク基礎をGo視点で押さえるステップです。

## ゴール
- `net/http` のRequest/Response処理を説明できる
- connection reuse、timeout設定の要点を説明できる
- クライアント/サーバの両面からHTTPを理解できる

## 学習トピック
- `http.Server` と `http.Client`
- Transport設定（Keep-Alive、IdleConn）
- server timeout（read/write/idle）
- JSON encode/decodeの注意点
- middlewareの前段理解

## 実装課題（自分で実装）
1. カスタム`http.Client`で外部APIを呼ぶ
2. timeoutを細かく設定したHTTPサーバを作る
3. リトライ戦略（単純指数バックオフ）を試作する

## 自分への確認質問
- `http.Client` を毎回生成しない方がよい理由は？
- server timeoutが未設定だと何が危険？
- Keep-Aliveが効く条件は？

## 完了条件
- REST実装時にtimeout/connection設定を自分で決められる
