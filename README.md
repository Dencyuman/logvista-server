# logvista-server
Logvista server for On-Premise

## フロントエンドのビルド
1. `yarn build`する際は、`client`ディレクトリ内の`.env`ファイルの`VITE_API_URL`の値を以下のように設定すること
    - `{{.VITE_API_URL}}`は、サーバー起動時に、サーバー側の環境変数`VITE_API_URL`の値に置換される
    ```
    VITE_API_URL={{.VITE_API_URL}}
    ```
2. `client`ディレクトリ内で`yarn build`を実行した結果のdistディレクトリ内のファイルを、`static`ディレクトリにコピーする必要がある
    ```
    client/
    |- dist/
    |  |- assets/
    |  |  |- index-*.js
    |  |  |- index-*.css
    |  |  |- ...
    |  |- index.html
    |  |- ...
    |- ...
    ```
    - 静的リソースは`static`ディレクトリ内に配置することで、バックエンドAPIから提供される
    - 基本的には、`dist`ディレクトリ内のファイルをそのまま`static`ディレクトリ内にコピーする
    - `dist/assets/index-*.js`のようなjsファイルは、**`static`ディレクトリ直下に移動またはコピーする**
        - build時に環境変数でプレースホルダー`{{.VITE_API_URL}}`を埋め込んでいるため、`static`ディレクトリ内に配置したjsファイルはテンプレートとなりサーバー起動時に、`static/assets/`内にサーバー側の環境変数`VITE_API_URL`の値が埋め込まれたjsファイルが生成される
    ```
    static/
    |- assets/
    |  |- index-*.js  <-テンプレートエンジンによって環境変数が置換されたjsファイル
    |  |- index-*.css
    |  |- ...
    |- index.html
    |- index-*.js  <-dist/assets/index-*.jsを移動またはコピーしたテンプレートjsファイル
    |- ...
    ```


## バックエンドのビルド
1. `wire`というDIツールを使用しているため、`go build`する際は以下のコマンドで`wire_gen.go`をbuild対象に含める必要がある
    - icon.sysoが用意されているため、`-o`オプションを使用して生成されるexeファイルのアイコンとして指定すること
    ```
    go build -o main.go wire_gen.go
    ```