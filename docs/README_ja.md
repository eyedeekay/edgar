エドガー
================================================================================================================================================================================================================================================================

読み出しで何でもホームページを生成します。 私の無限の makefile ナンセンスの交換。.
このツールは、README.mdファイルに基づいているプロジェクト用のページを作成するためのものです。
特に github のページには便利です。.

基本的には、単一のマークダウンファイルを取り、放射する本当に単純な静的サイトジェネレータ
hTMLページを適度に見直す.

STATUS: このプロジェクトは維持されます。 数日以内に問題やプルリクエスト、機能リクエストにお応えします。 それは
何をすべきか。.

使用方法
お問い合わせ

```md
Edgarの使用法:
-author 文字列
HTMLファイルの作成者 (デフォルトは eyedeekay)
-css 文字列
使用する CSS ファイルは、デフォルトは 1 が存在しない場合に生成されます (デフォルト style.css)
-donate 文字列
寄付セクションを暗号通貨ウォレットに追加します。 アドレス URL スキームを使用して、コンマ(スペースなし)で区切られます。 あなたがお金が私に行きたい場合は、実行前にそれらを変更. (既定のmonero:4A2BwLabGUiU65C5JRfwXqFTwWPYNSmuZRjbTDjsu9wT6wV6kMFyXn83ydnVjVcR7BCsWh8B5b4Z9b6cmqjfZiFd9sBUpWTD1DmyZAs5q2Lb8TBKJKJVJVe8B5B5B5b4Z9b6Z9b6cmqjfZiFd9sBUpWTD9sBUpWTD1DmyZASDmyZAS5qBJKJKJKJKJKJKJKJKFQJKFQJKFQJQJKFQJQJKFQJFQJVJFQJQJKFQJQJFQJFQJFQJFQJFQJFQJFQJFQJFQJFQJFQJFQJFQJFD
-filename 文字列
マークダウンファイルでHTMLに変換したり、コンマで区切られたファイルのリスト(デフォルト README.md,USAGE.md,index.html,docs/README.md)
-i2plink
ページのフッターに i2p リンクを追加します。 @Shoalsteedと@mark22k(デフォルトtrue)のロゴ提供
-寄付
donate セクションを無効にします(-donate ウォレットアドレスを true に設定する前に変更します) (デフォルトは true)
-out 入力ファイル。 ツイート
出力ファイルの名前(最初のファイルでのみ使用される、他の人はinputfile.html)(デフォルトindex.html)
-script 文字列
使用するスクリプトファイル。.
-スノーフレーク
ページフッターにスノーフレークを追加(デフォルトtrue)
-サポート文字列
寄付セクションのメッセージ/CTAを変更します。 (デフォルトでは「EDgarの独立した開発をサポート」)
-title 文字列
マークダウンファイルの最初のh1から生成される空白の場合、HTMLファイルのタイトル。.
```
