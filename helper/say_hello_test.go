package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// basic testing
// Fail -> lanjut sampai akhir
// FailNow -> langsung stop
// Error -> setara Fail
// Fatal -> setara FailNow

func TestSayHelloLutfi(t *testing.T) {
	actual := SayHello("Lutfi")
	expected := "Hello Lutfi"
	if actual != expected {
		t.Fail()
		// t.FailNow()
		// t.Errorf("Error")
		// t.Fatalf("Error")
	}
	// fmt.Println("apakah kode ini tereksekusi???")
}

func TestSayHelloAndriyanto(t *testing.T) {
	actual := SayHello("Andriyanto")
	expected := "Hello Andriyanto"
	if actual != expected {
		t.Fail()
	}
}

// using assert, require -- testify
// assert -> setara Fail
// require -> setara FailNow

func TestSayHelloAssertion(t *testing.T) {
	actual := SayHello("Lutfi")
	expected := "Hello Lutfi"
	assert.Equal(t, expected, actual)
	fmt.Println("ini akan dieksekusi")
}

func TestSayHelloRequire(t *testing.T) {
	actual := SayHello("Lutfi")
	expected := "Hello Lutfi"
	require.Equal(t, expected, actual)
	fmt.Println("ini tidak dieksekusi")
}

// skip -> biasanya untuk os, dll yang memerlukan kompatibilitas

func TestSkip(t *testing.T) {
	if runtime.GOOS == "linux" {
		t.Skip("can't run on linux")
	}
}

// main test function -> mengatur before dan after test

func TestMain(m *testing.M) {
	// before
	fmt.Println("===== sebelum semua test =====")

	// run the tests
	m.Run()

	// after
	fmt.Println("===== sesudah semua test =====")
}

// sub test -> satu fungsi test memiliki beberapa case testing

func TestSubTest(t *testing.T) {
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

// table test -> subtest dengan iterasi

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

// benchmark

func BenchmarkSayHello(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SayHello("Lutfi")
	}
}

func BenchmarkSub(b *testing.B) {
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

func BenchmarkTable(b *testing.B) {
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
