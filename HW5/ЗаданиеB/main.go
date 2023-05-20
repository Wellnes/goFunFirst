package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {

	var вхСтрока string // = "дома"

	ввод := bufio.NewScanner(os.Stdin)

	if ввод.Scan() {
		вхСтрока = ввод.Text()
	} else {
		return
	}

	ВхСтрокаКакСрез := []rune(вхСтрока)

	первыйСимвол := string(ВхСтрокаКакСрез[0])
	последнийСимвол := string(ВхСтрокаКакСрез[len(ВхСтрокаКакСрез)-1])

	первыйСимволЭтоД := strings.Compare(strings.ToUpper(первыйСимвол), "Д") == 0
	последнийСимволЭтоА := strings.Compare(strings.ToUpper(последнийСимвол), "А") == 0

	if utf8.RuneCountInString(вхСтрока) > 1 && (первыйСимволЭтоД && последнийСимволЭтоА) {
		fmt.Println("СОГЛАСЕН")
	} else {
		fmt.Println("НЕ СОГЛАСЕН")
	}

}
