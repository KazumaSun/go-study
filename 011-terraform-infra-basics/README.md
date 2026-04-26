# 011-terraform-infra-basics

Goアプリを支えるインフラをコードで管理する入門ステップです。

## ゴール
- Terraformの基本概念（state、plan、apply）を説明できる
- 最小構成のクラウドリソースをIaCで作成できる
- 環境分離（dev/stg/prod）の考え方を説明できる

## 学習トピック
- provider / resource / variable / output
- state管理（ローカルとリモート）
- moduleの基礎
- 命名規約、タグ戦略
- セキュリティ基礎（secret管理の原則）

## 実装課題（自分で実装）
1. 最小VPCまたは相当ネットワーク構成を作る
2. Go APIを配置する想定のcompute基盤を定義
3. dev環境をmodule化して再利用可能にする

## 自分への確認質問
- stateが壊れると何が問題？
- 手動変更（ドリフト）をどう検出する？
- module化しすぎるデメリットは？

## 完了条件
- IaCの変更差分を安全にレビューできる
