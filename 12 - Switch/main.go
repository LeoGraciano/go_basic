package main

import "fmt"

func dayOfWeek(day int) string {
	switch day {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Dia inválido"
	}
}

func switchDayOfWeek(day int) string {
	switch day {
	case 1, 2, 3, 4, 5, 6:
		return fmt.Sprintf("%d - %s", day, "Dia da semana")
	case 7:
		return fmt.Sprintf("%d - %s", day, "Domingo")
	default:
		return fmt.Sprintf("%d - %s", day, "Dia inválido")
	}
}

func switchDayOfWeekWithFallthrough(day int) string {
	switch day {
	case 1, 2, 3, 4, 5:
		fallthrough
	case 6:
		return fmt.Sprintf("%d - %s", day, "Dia da semana")
	case 7:
		return fmt.Sprintf("%d - %s", day, "Domingo")
	default:
		return fmt.Sprintf("%d - %s", day, "Dia inválido")
	}
}
func main() {
	fmt.Println("Switch")
	println("---------------------")
	fmt.Println(dayOfWeek(0))
	println("---------------------")
	fmt.Println(dayOfWeek(5))
	println("---------------------")
	fmt.Println(switchDayOfWeek(5))
	println("---------------------")
	fmt.Println(switchDayOfWeekWithFallthrough(5))
}
