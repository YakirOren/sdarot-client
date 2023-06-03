package main

import (
	"fmt"

	"github.com/YakirOren/sdarot-client/helpers"

	"github.com/AlecAivazis/survey/v2"
	"github.com/YakirOren/sdarot"
	"github.com/mgutz/ansi"
	log "github.com/sirupsen/logrus"
)

const (
	specificEpisodesMode = "Specific Episodes"
	specificSeasonsMode  = "Specific Seasons"
	everythingMode       = "Everything"
)

func main() {
	conf, err := helpers.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	seriesID := 0
	prompt := &survey.Input{
		Message: `Enter series ID`,
		Help:    `example: sdarot.tw/watch/` + ansi.Color(`id`, "red") + `:`,
	}

	err = survey.AskOne(prompt, &seriesID, survey.WithValidator(survey.ComposeValidators(survey.Required, helpers.IsInt)))
	if err != nil {
		return
	}

	client, err := sdarot.New(*conf)
	if err != nil {
		log.Fatal(err)
	}

	series, err := client.GetSeriesByID(seriesID)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("you choose ", series.EnglishName)

	Menu(series, client)
}

func Menu(series *sdarot.Series, client *sdarot.Client) {
	chosenMode, done := AskForMode()
	if done {
		return
	}

	var episodesToDownload []sdarot.VideoRequest

	switch chosenMode {
	case specificEpisodesMode:
		seasonNumber, done2 := AskForSeasonNumber(series)
		if done2 {
			return
		}

		eps := series.GetEpisodes(seasonNumber)

		var chosenEpisodes []int
		prompt2 := &survey.MultiSelect{
			Message: "Select Episodes",
			Options: helpers.IntToArr(len(eps)),
		}
		err := survey.AskOne(prompt2, &chosenEpisodes, survey.WithValidator(survey.Required))
		if err != nil {
			return
		}

		for _, an := range chosenEpisodes {
			episodesToDownload = append(episodesToDownload, eps[an])
		}

	case specificSeasonsMode:
		chosenSeasons := []int{0}

		if series.GetSeasons() != 1 {
			prompt2 := &survey.MultiSelect{
				Message: "Select Seasons",
				Options: helpers.IntToArr(series.GetSeasons()),
			}
			err := survey.AskOne(prompt2, &chosenSeasons, survey.WithValidator(survey.Required))
			if err != nil {
				return
			}
		} else {
			log.Println(series.EnglishName + " has only 1 season")
		}

		for _, seasonNumber := range chosenSeasons {
			episodesToDownload = append(episodesToDownload, series.Seasons[seasonNumber]...)
		}

	case everythingMode:
		for _, season := range series.Seasons {
			episodesToDownload = append(episodesToDownload, season...)
		}
	}

	log.Printf("Downloading %d episodes\n", len(episodesToDownload))

	d := Downloader{
		client: client,
	}

	for _, request := range episodesToDownload {
		d.Download(request, series.EnglishName)
	}
}

func AskForSeasonNumber(series *sdarot.Series) (int, bool) {
	seasonNum := 1

	if series.GetSeasons() != 1 {
		prompt3 := &survey.Input{
			Message: fmt.Sprintf("Enter Season Number [1-%d]", series.GetSeasons()),
		}
		validators := survey.ComposeValidators(survey.Required, helpers.InRange(int64(series.GetSeasons())))
		err := survey.AskOne(prompt3, &seasonNum, survey.WithValidator(validators))
		if err != nil {
			return 0, true
		}
	}
	return seasonNum, false
}

func AskForMode() (string, bool) {
	chosenMode := ""
	prompt2 := &survey.Select{
		Message: "Choose Mode:",
		Options: []string{
			specificEpisodesMode,
			specificSeasonsMode,
			everythingMode,
		},
	}
	err := survey.AskOne(prompt2, &chosenMode, survey.WithValidator(survey.Required))
	if err != nil {
		return "", true
	}
	return chosenMode, false
}
