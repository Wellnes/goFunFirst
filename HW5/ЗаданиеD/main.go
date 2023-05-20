package main

import (
	"fmt"
	"strings"
)

func main() {

	var исходноеСлово, новоеСлово string
	fmt.Scan(&исходноеСлово, &новоеСлово)

	for УсловиеИгрыВыполнено(исходноеСлово, новоеСлово) {
		исходноеСлово = новоеСлово
		fmt.Scan(&новоеСлово)
	}

	fmt.Println(новоеСлово)

}

func УсловиеИгрыВыполнено(исходноеСлово, новоеСлово string) bool {

	перваяБукваНовогоСлова := ПерваяБуква(новоеСлово)
	последняяБукваИсходногоСлова := ПоследняяБукваБезУчетаМягЗнака(исходноеСлово)

	return strings.ToUpper(перваяБукваНовогоСлова) == strings.ToUpper(последняяБукваИсходногоСлова)

}

func СловоКакСрез(вхСлово string) []rune {

	return []rune(вхСлово)

}

func ПерваяБуква(вхСлово string) string {
	return string(СловоКакСрез(вхСлово)[0])
}

func ПоследняяБукваБезУчетаМягЗнака(вхСлово string) string {

	var rez string

	СловоКакСрез := СловоКакСрез(вхСлово)

	for i, found := (len(СловоКакСрез) - 1), false; !found && i >= 0; i-- {
		rez = string(СловоКакСрез[i])
		found = strings.ToLower(rez) != "ь"
	}

	return rez
}
