# 005-concurrency-patterns

実務で頻出する並行処理パターンを身につけるステップです。

## ゴール
- fan-out / fan-in を実装できる
- producer-consumerのバックプレッシャを設計できる
- `sync.Mutex` / `sync.RWMutex` / `sync.Map` の使い分けができる

## 学習トピック
- worker pool設計
- fan-out / fan-in
- rate limitの基本
- 共有メモリと排他
- race conditionの検出（`-race`）

## 実装課題（自分で実装）
1. 画像変換ジョブの疑似worker pool
2. 集約処理（複数ソースからfan-in）
3. `-race` で失敗するコードと修正版を作る

## 自分への確認質問
- channel中心設計とmutex中心設計の違いは？
- バックプレッシャを無視すると何が起きる？
- `RWMutex` は常に速いわけではないのはなぜ？

## 完了条件
- 競合がない並行処理実装を自力で設計できる
