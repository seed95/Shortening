package i18n

import (
	"github.com/BurntSushi/toml"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/seed95/shortening/pkg/translate"
	"golang.org/x/text/language"
	"path/filepath"
)

type (
	translatorBundle struct {
		bundle    *i18n.Bundle
		localized map[string]*i18n.Localizer
	}
)

func New(path string) (translate.Translator, error) {
	translator := &translatorBundle{
		bundle:    i18n.NewBundle(language.English),
		localized: map[string]*i18n.Localizer{},
	}

	if err := translator.loadBundle(path); err != nil {
		return nil, err
	}

	return translator, nil
}

func (t *translatorBundle) loadBundle(path string) error {

	t.bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	messageFiles, err := filepath.Glob(filepath.Join(path, "*.toml"))
	if err != nil {
		return err
	}

	for _, messageFile := range messageFiles {
		_, err := t.bundle.LoadMessageFile(messageFile)
		if err != nil {
			return err
		}
	}

	return nil
}

func (t *translatorBundle) getLocalized(lang string) *i18n.Localizer {
	if _, ok := t.localized[lang]; !ok {
		t.localized[lang] = i18n.NewLocalizer(t.bundle, lang)
	}
	return t.localized[lang]
}

func (t *translatorBundle) Translate(key string, languages ...translate.Language) string {
	lang := translate.EN

	for _, l := range languages {
		switch l {
		case translate.EN:
			lang = translate.EN
			break
		case translate.FA:
			lang = translate.FA
			break
		}
	}
	message, err := t.getLocalized(string(lang)).Localize(&i18n.LocalizeConfig{MessageID: key})
	if err != nil {
		return key
	}

	return message
}
