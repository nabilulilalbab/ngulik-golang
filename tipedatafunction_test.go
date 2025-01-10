package ngulik

import (
	"fmt"
	"testing"
)

type Greeter func(name string)

// cara 1
func (g Greeter) sayHello(name string) {
	g(name)
}

// cara 2
func (g Greeter) sayHello2(name string) {
	fmt.Println("hello, ", name)
}

func TestFunctionTipeTrick(t *testing.T) {
	greeter1 := Greeter(func(name string) {
		fmt.Println("Halo", name, "apa kabar?")
	})
	greeter2 := Greeter(func(name string) {
		fmt.Println("selamat Pagi, ", name, "!")
	})
	greeter1.sayHello("nabil")
	greeter2.sayHello("nabiel")
	greeter1.sayHello2("nabiel")
	greeter2.sayHello2("nabiel")
}

// penjelasan g name : g(name) mewakili anonymouse function jadi agar ketika mengirim function otomatis terpanggil contoh :

func TestFunctionTipeTrick2(t *testing.T) {
	// Ketika g(name) dipanggil, sebenarnya yang dijalankan adalah:
	func(name string) {
		fmt.Println("Halo,", name)
	}("Budi")
}
