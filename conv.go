// Упражнение 2.2:
// Напишите программу общего назначения для преобразования едениц, аналогичную ```cf```,
// которая считывает числа из аргументов командной строки
// (или из стандартного ввода, если аргументы командной строки отсутствуют)
// и преобразует каждое число в другие еденицы, как температуру — в градусы Цельсия и Фаренгейта,
// длину — в футы и метры, вес — в фунты и килограммы и.т.д.

package main

import (
	"fmt"
	"github.com/unixlinuxgeek/lenconv"
	"github.com/unixlinuxgeek/tempconv"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			t, err := strconv.ParseFloat(arg, 64)
			if err != nil {
				fmt.Fprintf(os.Stdout, "%v\n", err)
				os.Exit(1)
			}
			f := tempconv.Fahrenheit(t)
			c := tempconv.Celsius(t)
			fmt.Printf("%s = %s, %s = %s\n", f, tempconv.FahToCel(f), c, tempconv.CelToFah(c))

			ft := lenconv.Ft(t)
			met := lenconv.Meter(t)
			fmt.Printf("%s = %s, %s = %s\n", ft, lenconv.FtToMet(ft), met, lenconv.MetToFt(met))
		}
	}
}
