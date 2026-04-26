# 006-context-cancellation-timeout

API/バッチで必須の`context`運用を習得するステップです。

## ゴール
- `context.Background` / `WithCancel` / `WithTimeout` を使い分ける
- キャンセル伝播を設計に組み込める
- context valueの使いどころを説明できる

## 学習トピック
- contextの木構造
- タイムアウト設計
- 早期終了とクリーンアップ
- `errgroup` の導入（必要なら）
- context valueのアンチパターン

## 実装課題（自分で実装）
1. 複数外部呼び出しをtimeout付きで並列化
2. 親処理中断時に子goroutineを全停止
3. request-idをcontext経由でログに埋める

## 自分への確認質問
- `context` を構造体に保持すると何が問題？
- timeoutはどの層で決めるべき？
- context valueは何でも入れてよい？

## 完了条件
- キャンセル漏れのない並行I/O処理を実装できる
