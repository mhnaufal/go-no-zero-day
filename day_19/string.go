package day_19

import "fmt"
import "strings"

func Strings() {
	var isContained = strings.Contains("Johny Depp", "Depp")
	fmt.Println(isContained)

  var prefix = strings.HasPrefix("Lufthy", "thy")
  fmt.Println(prefix)

  var howMany = strings.Count("lorem upsim dolor sit amet", "o")
  fmt.Println(howMany)

  var splitMe = strings.Split("The Batman", "")
  fmt.Println(splitMe)
}
