package main

import (
	"log"

	"net/http"

	"github.com/mileusna/useragent"
	"github.com/spf13/viper"
)

func handler(w http.ResponseWriter, r *http.Request) {
	redirect := func(url string) {
		http.Redirect(w, r, url, http.StatusSeeOther)
	}

	userAgent := useragent.Parse(r.UserAgent())

	switch {
	case userAgent.IsIOS() && viper.IsSet(k_ios_url):
		redirect(viper.GetString(k_ios_url))
	case userAgent.IsAndroid() && viper.IsSet(k_android_url):
		redirect(viper.GetString(k_android_url))
	default:
		redirect(viper.GetString(k_fallback_url))
	}
}

const k_ios_url string = "iosurl"
const k_android_url string = "androidurl"
const k_fallback_url string = "fallbackurl"

func main() {

	// set viper defaults
	viper.SetDefault(k_fallback_url, "https://github.com/bemayr/app-forward")

	// set viper config file handling
	viper.SetConfigFile(".env")
	viper.WatchConfig()

	// set viper environment variable handling
	viper.SetEnvPrefix("appforward")
	viper.BindEnv(k_ios_url)
	viper.BindEnv(k_android_url)
	viper.BindEnv(k_fallback_url)

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
