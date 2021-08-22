# Procon2021RegistUserActivity

## 基本仕様
---
### 概要
ユーザーの作業開始、休憩開始、休憩終了の動作を記録としてDBに格納します。

### 特徴
ステータスをDBに入れます。
ユーザーIDが正しいかどうかは確認していません。(すでにログイン済みのため)

## 機能仕様
- HTTP method:POST
- endpoint:/activity/regist

---
### 利用法

- endpoint:/activity/regist

####  request
```cassandraql
{
    "userID":"0001",
    "timestamp":1629625499,
    "status":3
}
```

####  resqonse
```cassandraql
{
  "Message": ""
}
or
{
  "Message": "request is wrong"
}
```
※サーバーエラーの時はステータスコードなどを返します。

---
### アーキテクチャ
MVCもどき

### 依存環境
- AWS (Lambda)
- AWS (API Gateway)

### 使用ライブラリ
主要なライブラリのみを示す
- aws-sdk()
- Gin(Webフレームワーク)

### デプロイ


### セットアップ
対象となるソースコードをzipに圧縮してアップロードする必要があります
2. GOOS=linux GOARCH=amd64 go build -o hello main.go
3. zip function.zip hello

### 注意点
動作確認はAWSのLambdaにアップロードする必要があります
zipはGitにあげないこと

### 作成者
Taketo Wakamatsu (若松丈人)
