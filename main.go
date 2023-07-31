package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func main() {
	// .envファイルから環境変数を読み込む
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err)
	}

	accessToken := os.Getenv("LINE_ACCESS_TOKEN")

	// LINE Notify APIのエンドポイントURL
	apiURL := "https://notify-api.line.me/api/notify"

	// POSTリクエストのフォームデータを定義
	formData := url.Values{}
	formData.Set("message", "テスト")
	formData.Set("stickerPackageId", "8515")
	formData.Set("stickerId", "16581242")
	// formData.Set("notificationDisabled", "true") // 通知を送らない場合true

	// POSTリクエストを作成
	req, err := http.NewRequest("POST", apiURL, strings.NewReader(formData.Encode()))
	if err != nil {
		log.Fatalf("Error creating request: %s", err)
	}

	// ヘッダーを設定
	req.Header.Set("Authorization", "Bearer "+accessToken)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	// HTTPクライアントを作成
	client := &http.Client{}

	// リクエストを送信してレスポンスを取得
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Error sending request: %s", err)
	}
	defer resp.Body.Close()

	// レスポンスを表示
	fmt.Println("Status Code:", resp.StatusCode)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %s", err)
	}
	fmt.Println("Response Body:", string(body))
}
