package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Name = "translate app"
	app.Usage = "This app echo input arguments"
	app.Version = "0.0.1"

	app.Action = func(context *cli.Context) error {
		values := url.Values{}
		values.Add("q", context.Args().Get(0))
		values.Add("source", "en")
		values.Add("target", "ja")
		values.Add("format", "text")
		key := os.Getenv("GCP_API_KEY")
		values.Add("key", key)
		resp, err := http.Get("https://translation.googleapis.com/language/translate/v2" + "?" + values.Encode())

		if err != nil {
			fmt.Println(err)
			return nil
		}

		defer resp.Body.Close()

		var result struct {
			Data struct {
				Translations []struct {
					TranslatedText string `json:"translatedText"`
				} `json:"translations"`
			} `json:"data"`
		}

		if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
			panic(err)
		}

		fmt.Println(result.Data.Translations[0].TranslatedText)

		return nil
	}

	app.Run(os.Args)
}
