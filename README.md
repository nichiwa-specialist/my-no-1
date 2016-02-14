# 俺の世界一

利用者が、「俺にとって、この分野の世界一はここだ！」と思える店を紹介していくグルメSNSです。  
食べログなどのこれまでのグルメサイトとちがい、「誰が」オススメしているのかに焦点を当てるため、サクラや好みの合わないユーザーの評価で信憑性が下がる、ということがありません。

## 開発環境

開発環境として、[Google App Engine SDK for Go](https://cloud.google.com/appengine/downloads#Google_App_Engine_SDK_for_Go) が必要です。  
開発を行うマシンのOSに応じたSDKをダウンロードして展開し、パスを通してください。

SDKにパスを通せば、このプロジェクトをチェックアウトしたフォルダで

    goapp serve .

とgoapp serveコマンドを実行すれば、開発サーバが立ち上がり http://localhost:8080/top.html でトップページにアクセスできます。  
GAE/goの開発環境は変更を検出して自動的に更新してくれるので、開発サーバを立ち上げたままでコードを編集すれば、変更が即座に開発サーバに反映されます。

## デプロイ方法

クローンした開発ディレクトリ上で

    goapp deploy

と goapp deployコマンドを実行すれば、ブラウザが起動してgoogleの認証画面が開き、そこでデプロイ可能なユーザでログインすればGAEのmy-no-1プロジェクトにデプロイされ、 http://my-no-1.appspot.com/top.html でアクセスできます。  
デプロイを行った結果、ホームディレクトリに .appcfg_oauth2_tokens が作成されます。このファイルの内容はのちのci設定で必要になります。

## CI設定

このプロジェクトには Circle CI の設定ファイルが含まれており、masterへのコミットがそのままGAEにデプロイされるようになっています。
Circle CI上でこのプロジェクトをデプロイ対象とするには、Project settings の Enviroment Variables のページを開き、
REFRESH_TOKENに上記 .appcfg_oauth2_tokens ファイルに記載された refresh_token の値を設定します。


