https://qiita.com/api/v2/docs#投稿

```go
type QiitaItem struct {
	Title string
}

func GetQiitaItems(token string) []QiitaItem {
	req, err := http.NewRequest("GET", "http://qiita.com/api/v2/authenticated_user/items", nil)
	if err != nil {
		log.Print(err)
	}
	req.Header.Set("Authorization", "Bearer " + token)

	cli := &http.Client{}
	res, err := cli.Do(req)
	if err != nil {
		log.Print(err)
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	var result []QiitaItem
	err = json.Unmarshal(body, &result)

	return result
}

func main() {
	token := os.Getenv("QIITA_ACCESS_TOKEN")
	items := GetQiitaItems(token)

	for _, item := range items {
		println(item.Title)
	}
}
```

```
QIITA_ACCESS_TOKEN="..." go run main.go

Swiftでいい感じのViewModelを作るためのメモ
2017年におけるObjective-Cコミュニティの動向
git-grepで特定のディレクトリを除外する
オッ
Apple Pay アプリ開発、その前に
Hack the 転職ドラフト
test
PAY.JP iOS SDKでApple Pay アプリ内支払いをする
特定のSlackチャンネルに参加しているユーザーのメールアドレス一覧を取得する
Swift Package Manager で外部ライブラリを使いコマンドラインツールを開発する
Git: 特定の変更中のファイルの一覧を取得して何かする
mitmproxy 0.16 にリクエスト内容をcurlコマンドやPythonコードに変換できる便利なやつがついていた
コードレビューの担当者の配分を可視化するスクリプト
Node v6.0.0-preでReactをビルドする
GitHubの草のSVGをコマンドラインで取得する
手動でAirbrakeへエラー通知を送る
CircleCI上でFaux Pasを使って Objective-C コードの品質チェックをする
Beta by Crashlytics (fabric.io) のテスターのUDID一覧を取得する
jQuery#text() の childNodes 互換性問題
デバイスの空き容量をゴミファイルで埋める
```
