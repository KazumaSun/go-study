# 002-struct-interface-design

Goらしい設計の中心である`struct`と`interface`を使いこなすステップです。

## ゴール
- メソッドレシーバ（値/ポインタ）を適切に選べる
- `interface`の設計を「最小」にできる
- 埋め込み（embedding）とコンポジションを使い分けられる

## 学習トピック
- `struct`、コンストラクタ風関数、ゼロ値設計
- メソッドセット（value receiver / pointer receiver）
- `interface`の暗黙実装
- `any`、型アサーション、型スイッチ
- 埋め込みと責務分離

## 実装課題（自分で実装）
1. 支払いドメイン（`PaymentProcessor` interface）を作る
2. `Repository` interface を使ったメモリ実装を作る
3. ログ出力を差し替え可能な設計にする

## 自分への確認質問
- なぜ「受け側がinterfaceを定義する」が推奨される？
- pointer receiverが必要なケースは？
- interfaceを大きくしすぎると何が起きる？

## 完了条件
- 3課題ともにテストしやすい依存分離ができている
