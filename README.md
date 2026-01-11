<!-- toc --><!-- markdown-toc start - Don't edit this section. Run M-x markdown-toc-refresh-toc -->
**Table of Contents**

- [1.はじめに](#1)
- [2.概要](#2)
  - [2.1.ver1](#21ver1)
- [3.画面遷移](#3)
  - [3.1.TOPページ (/)](#31top-)
  - [3.2.登録ページ (/register)](#32-register)
- [4.機能一覧](#4)
  - [4.1.TODOアイテム登録機能](#41todo)
  - [4.2.一覧表示機能](#42)

<!-- markdown-toc end -->

# 1.はじめに
Go言語とGORMを使用してTODOリストのWebアプリを作成する.  
練習用プロジェクト

# 2.概要
## 2.1.ver1
個人利用を想定. 登録したものを一覧表示するだけ. 


# 3.画面遷移
## 3.1.TOPページ (/)
- 登録したTODOリストを表示する

## 3.2.登録ページ (/register)
- TODOアイテムを登録する


# 4.機能一覧
## 4.1.TODOアイテム登録機能
以下の項目を記載してDBに登録する. (*: 必須)  
- item: item名* (自由記述式)
- priority: 優先度 (選択式: -, high, middle, low)
- status: ステータス (選択式: -, InProgress, Done)
- deadline: 締切日時 (カレンダーから選択)
- remarks: 備考 (自由記述式)


## 4.2.一覧表示機能
登録されたTODOアイテムを表示する.  
statusによる絞り込みができる.  
