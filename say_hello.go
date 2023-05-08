package go_say_hello

func SayHello(name string) string {
	return "Hello " + name
}

// Release tag
// 1. git tag v1.0.0
// 2. git push origin v1.0.0

// Jika terjadi major upgrade
// 1. misal function ada tambahan porameter
// 2. Ubah nama module di akhir menjadi /v2