package forismatic_api

const (
	LangRu languageOption = "ru"
	LangEn languageOption = "en"
)

type languageOption string

var currentLanguageOption = LangEn

func Lang() languageOption {
	return currentLanguageOption
}

func SetLang(l languageOption) {
	currentLanguageOption = l
}
