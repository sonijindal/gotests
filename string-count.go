package main

import (
  "fmt"
  "strings"
  "unicode"
  "unicode/utf8"
)

func count(content string) {
  f := func(c rune) bool {
        return !unicode.IsLetter(c) && !unicode.IsNumber(c)
          }
  words := strings.FieldsFunc(content, f)
    //kv := make([]mapreduce.KeyValue, len(words))
      for i, word := range words {
            fmt.Println(i, "=>", word)
      }
}
func count_rune() {
  const nihongo = "日本語"
      for i, w := 0, 0; i < len(nihongo); i += w {
                runeValue, width := utf8.DecodeRuneInString(nihongo[i:])
                        fmt.Printf("%#U starts at byte position %d\n", runeValue, i)
                                w = width
                                fmt.Println("isletter:", unicode.IsLetter(runeValue))
                                fmt.Println("isDigit:", unicode.IsDigit(runeValue))
                                    }
}
func main() {
  str := "This is my house.   I stay here;"
  count(str)
  count_rune()
}
