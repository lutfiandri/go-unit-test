# Rangkuman Golang Unit Test

Referensi belajar: [PZN](https://www.youtube.com/watch?v=t9QJPE5vwhs) \
Menggunakan library: [github.com/stretchr/testify](https://pkg.go.dev/github.com/stretchr/testify)

## Sedikit Tentang Testing

Testing berfungsi untuk mengecek apa yang kita coding itu benar atau tidak. Konsep cara mengeceknya sederhana, kita cukup membuat **actual contition** dan **expected condition**. Jika actual dan expected condition sama, maka test berhasil.

Testing itu ada 3:
Nama|Keterangan
-----|-----
End to End Testing|Pengetesan seluruh app, termasuk frontend (tidak ada mock)
Integration Testing|Pengetesan satu app (menggunakan mock server jika memakai service dari app lain)
Unit Testing|Pengetesan unit terkecil (function) dalam dalam app

## Unit Test

Sekali lagi, unit test dilakukan untuk mengetest setiap fungsi dalam app. Contoh hasil run unit test:
```
=== RUN   TestSayHello
--- PASS: TestSayHello (0.00s)
PASS
```

### Info

*|Contoh
---|---
Unit File|`helper/say_hello.go`
Nama File|`helper/say_hello_test.go`
Fungsi Test|`TestSayHello(t *testing.T)`
Run All|`go test -v`
Run Specific|`go test -v -run=TestSayHello`

*Keterangan: Sebelum `go test`, masuk ke directory `helper` dulu.*

### Contoh

**Basic Testing**
```go
func TestSayHello(t *testing.T) {
	actual := SayHello("Lutfi")
	expected := "Hello Lutfi"
	assert.Equal(t, expected, actual)
	fmt.Println("ini akan dieksekusi")
}
```

**Sub Test**
```go
func TestSayHello(t *testing.T) {
	t.Run("Lutfi", func(t *testing.T) {
		actual := SayHello("Lutfi")
		expected := "Hello Lutfi"
		assert.Equal(t, expected, actual)
	})
	t.Run("Andriyanto", func(t *testing.T) {
		actual := SayHello("Andriyanto")
		expected := "Hello Andriyanto"
		assert.Equal(t, expected, actual)
	})
}
```

**Table Test**
```go
func TestTableTest(t *testing.T) {
	tests := []struct {
		name     string
		request  string
		expected string
	}{
		{
			name:     "0",
			request:  "Lutfi",
			expected: "Hello Lutfi",
		},
		{
			name:     "1",
			request:  "Andriyanto",
			expected: "Hello Andriyanto",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := SayHello(test.request)
			assert.Equal(t, test.expected, actual)
		})
	}
}
```

### Skip Test

Terkadang, ada fungsi yang tidak kompatibel pada beberapa kondisi. Contohnya adalah kita memiliki suatu fungsi yang tidak bisa dipakai di linux, maka kita bisa skip test tersebut khusus untuk linux.

Penggunaan: `t.Skip("message")`

```go
func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("can't run on linux")
	}
}

```

## MainTest

Kita bisa mengatur sebelum dan sesudah melakukan testing.

### Info

*|Contoh
---|---
Unit File|`helper/say_hello.go`
Nama File|`helper/say_hello_test.go`
Fungsi Test|`TestMainSayHello(m *testing.M)`

### Contoh

```go
func TestMain(m *testing.M) {
	// before
	fmt.Println("===== sebelum semua test =====")

	// run the tests
	m.Run()

	// after
	fmt.Println("===== sesudah semua test =====")
}
```

## Benchmark

Jika testing berfungsi untuk mengetest kebenaran fungsi yang dibuat, benchmark hanya mengetest seberapa cepat fungsi tersebut dieksekusi. Contoh hasil run benchmark
```
BenchmarkSayHello-4    48538173        21.44 ns/op
```

### Info

*|Contoh
---|---
Unit File|`helper/say_hello.go`
Nama File|`helper/say_hello_test.go`
Fungsi Test|`BenchmarkSayHello(b *testing.B)`
Run All Benchmarks and Tests|`go test -v -bench=.`
Run All Benchmarks|`go test -v -run=TestNothing -bench=.`
Run Specific|`go test -v -run=TestNothing -bench=BenchmarkSayHello`

### Contoh

**Basic Benchmark**
```go
func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Lutfi")
	}
}
```

**Sub Benchmark**
```go
func BenchmarkSayHello(b *testing.B) {
	b.Run("Lutfi", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Lutfi")
		}
	})
	b.Run("Lutfi Andriyanto", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			SayHello("Lutfi Andriyanto")
		}
	})
}
```

**Table Benchmark**
```go
func BenchmarkSayHello(b *testing.B) {
	benchmarks := []struct {
		name    string
		request string
	}{
		{
			name:    "0",
			request: "Lutfi",
		},
		{
			name:    "1",
			request: "Lutfi Andriyanto",
		},
	}

	for _, benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				SayHello(benchmark.request)
			}
		})
	}
}
```
