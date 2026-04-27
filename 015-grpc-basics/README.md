# 015-grpc-basics

gRPCの基礎から実運用で必要な要素までを学ぶカリキュラムです。

## ゴール
- Protocol Buffersの設計とコード生成を説明できる
- unary / streaming RPCを実装できる
- deadline、retry、status codeの使い分けを説明できる

## 学習トピック
- `.proto` 設計
- サーバ/クライアント実装
- context deadline と cancellation
- interceptor（認証、ログ、メトリクス）
- gRPC Gatewayの概念

## 実装課題
1. UserServiceのunary RPC（GetUser）を実装する
2. server streaming RPC（ListUsers）を実装する
3. interceptorで認証チェックとログ出力を実装する

## 完了条件
- RESTとgRPCのトレードオフを非機能要件込みで説明できる
