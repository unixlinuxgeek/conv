// Упражнение 2.2:
// Напишите программу общего назначения для преобразования едениц, аналогичную ```cf```,
// которая считывает числа из аргументов командной строки
// (или из стандартного ввода, если аргументы командной строки отсутствуют)
// и преобразует каждое число в другие еденицы, как температуру — в градусы Цельсия и Фаренгейта,
// длину — в футы и метры, вес — в фунты и килограммы и.т.д.

package main

import (
	"bufio"
	"fmt"
	"github.com/unixlinuxgeek/dimsconv"
	"github.com/unixlinuxgeek/freqconv"
	"github.com/unixlinuxgeek/lenconv"
	"github.com/unixlinuxgeek/tempconv"
	"github.com/unixlinuxgeek/wtconv"
	"os"
	"strconv"
	"strings"
)

func main() {
	if len(os.Args) > 1 {
		for _, arg := range os.Args[1:] {
			conv(arg)
		}
	} else {
		fmt.Print("Введите значения: ")
		s := bufio.NewReader(os.Stdin)
		defer os.Stdin.Close() // prevent memory leak

		a, _ := s.ReadString('\n')

		w := strings.Fields(a)
		for _, a := range w {
			conv(a)
		}
	}
}

//	Using go modules:
//
// "github.com/unixlinuxgeek/dimsconv"
// "github.com/unixlinuxgeek/lenconv"
// "github.com/unixlinuxgeek/tempconv"
// "github.com/unixlinuxgeek/wtconv"
func conv(arg string) {
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

	kg := wtconv.Kg(t)
	lb := wtconv.Lb(t)
	fmt.Printf("%s = %s, %s = %s\n", kg, wtconv.KgToLb(kg), lb, wtconv.LbToKg(lb))

	cm := dimsconv.Cm(t)
	i := dimsconv.Inch(t)
	fmt.Printf("%s = %s, %s = %s\n", cm, dimsconv.CmToInch(cm), i, dimsconv.InchToCm(i))

	kh := freqconv.Kilohertz(t)
	hz := freqconv.Hertz(t)
	fmt.Printf("%s = %s, %s = %s\n", kh, freqconv.KHzToHz(kh), hz, freqconv.HzToKHz(hz))
	fmt.Println()

}
