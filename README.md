# Forismatic Api
##### Allows you to get quotes from forismatic.com in your app

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-81%25-brightgreen.svg?longCache=true&style=flat)</a>

#### Example
```golang
forismatic.SetLang(forismatic.LangRu)
q, err := forismatic.QuoteJson()
log.Print(q.Text)
```