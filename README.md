# zipcode2prefs

![test](https://github.com/soh335/zipcode2prefs/workflows/test/badge.svg)

zipcode2prefs return Japanese Prefectures from zipcode.
zipcode2prefs.gen.go is generated from ken_all.zip

## HOW TO USE

```go
prefs := zipcode2prefs.Get("1500013") // []string{"東京都"}
```
