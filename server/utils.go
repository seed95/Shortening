package server

import (
	"espad_task/pkg/translate"
	"sort"
	"strconv"
	"strings"
)

type langQ struct {
	lang string
	q    float64
}

func GetLanguage(acceptLanguages string) []translate.Language {

	var lqs []langQ
	languages := strings.Split(acceptLanguages, ",")

	for _, language := range languages {
		language = strings.Trim(language, " ")
		langWithQ := strings.Split(language, ";")

		if len(langWithQ) == 1 {
			lq := langQ{langWithQ[0], 1}
			lqs = append(lqs, lq)
		} else {
			valueQ := strings.Split(langWithQ[1], "=")
			q, err := strconv.ParseFloat(valueQ[1], 64)
			if err != nil {
				continue
			}
			lq := langQ{langWithQ[0], q}
			lqs = append(lqs, lq)
		}
	}

	sort.SliceStable(lqs, func(i, j int) bool {
		return lqs[i].q > lqs[j].q
	})

	var result []translate.Language
	for _, lq := range lqs {
		result = append(result, translate.GetLanguage(lq.lang))
	}

	return result
}
