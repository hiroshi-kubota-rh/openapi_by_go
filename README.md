# 概要
This is a repository to study OpenAPI.
OpenAPIの勉強のため作成したレポジトリです．

M1マック環境の人はそのまま動くはずです．  
M1mac以外の人は，docker-compose.ymlのmysqlのplatform部分の記載(5行目)を削除して実行してください．

## 実行方法
1. レポジトリのrootでdocker compose up --build する
2. intern_pj_app_1, intern_pj_mysql_1, intern_pj_swagger-ui_1 の三つのcontainerがちゃんと立っているのを確認する．
3. `docker-compose exec mysql /bin/bash` でmysqlのたっているdbコンテナの中に入る．
4. mysqlのcontainer内で`mysql -u newview -pnewview` をして，mysqlの中にはいる．
5. `use newview;` でnew viewデータベースを使用
6. `show tables;` でaccountsのtalbeができていることを確認（今回はmysqlのimageの機能を使っているので，migrationの必要なし）
7. ここまでできいることが確認できたら,[http://localhost:8081](http://localhost:8081)にアクセスしてください．
swagger UIのページから各種APIを試すことが可能です．


## swagger UIについて
現状，実装しているAPI機能は以下の二つです．
1. GET /test
2. POST /impression_tags
3. GET /impression_tags/{id}
4. DELETE /impression_tags/{id}

### 実行方法
1. GET /testやPOST /accountsのタブを開く
2. 開いたタブの右上にある`Try it out`をクリック
3. GET /testに関してはそのまま`Execute`をクリック
4. POST /impression_tagsに関しては，Resquest Bodyのnameを好きに設定して，`Execute`を実行(jsonのkey部分はそのままにして，valueだけ変更してください)
5. きちんと実行できれば，どちらもResponseのServer responseの部分に200のレスポンスコードが表示されるはずです．impression_tagsのnameはuniqueである必要があります．
6. 本来であれば，adminのデータを使ったauthなどをしてみたかったのですが，その部分に関しては未実装です．
   
## dbについて
swaggerからacccountの登録を行なった場合，登録した情報がmysqlのtalbeに保存されます．
mysqlコンテナに入り，mysqlにログインした状態で
`select * from imression_tags;`
とかすると確認可能かと思います．
