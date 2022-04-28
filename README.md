# SimpleOSCTransmitter
　SimpleOSCTransmitterは必要最低限の機能のみを持つOpenSoundControlクライアントです。軽量さと美麗なUIを特徴としています。
## 使い方
　"Address"に信号の送信先アドレスを、"Value"に送信したい値を入力した後、値のデータ型を選択して"Transmission"を押すことで値を送信できます。  
　VRChatアバターの操作への利用を想定しているため現状では値の送信先はlocalhost:9000のみになっています。  
　受け付ける値は数字か"true"または"false"の文字列だけです。それ以外の値を入力しても送信に失敗することはありませんが、正しいデータが送られません。
## 開発上の注意
　ビルドにはwailsの実行環境が必要です。
## ロードマップ
- localhost:9000以外のデバイスやポートへの値の送信のサポート
- ダークテーマの実装
## 使用したOSSのライセンス情報
- [Go](https://go.dev)
- [GoOSC](https://github.com/hypebeast/go-osc)
- [wails](https://github.com/wailsapp/wails)
- [Koruri](https://github.com/Koruri/Koruri)