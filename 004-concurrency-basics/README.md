# 004-concurrency-basics

goroutineとchannelの基礎を体系的に理解するステップです。

## ゴール
- goroutineとOSスレッドの違いを説明できる
- channelの基本送受信を安全に使える
- デッドロックの典型パターンを回避できる

## 学習トピック
- goroutineの起動とライフサイクル
- unbuffered/buffered channel
- `select` の基礎
- channel closeのルール
- `sync.WaitGroup`

## 実装課題（自分で実装）
1. 複数workerでジョブ処理するプログラム
2. timeout付き処理（`select` + timer）
3. わざとデッドロックを作り、原因を言語化する

## 自分への確認質問
- closeするのは送信側か受信側か？
- buffered channelは何を解決し、何を隠す？
- goroutineリークはなぜ起こる？

## 完了条件
- worker処理が停止処理込みで安定して終了する
