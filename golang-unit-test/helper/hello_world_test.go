package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestTableHelloWorld(t *testing.T) {
	tests := []struct {
		name string
		request string
		expected string
	} {

		{
			name: "Jauhar",
			request: "Jauhar",
			expected: "Hallo Jauhar",
		},
		{
			name: "Juned",
			request: "Juned",
			expected: "Hallo Juned",
		},
		{
			name: "Agung",
			request: "Agung",
			expected: "Hallo Agung",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := HelloWorld(test.request)
			assert.Equal(t, result, test.expected)
		})
	}
}

func TestSubTest(t *testing.T) {
	t.Run("Johar", func(t *testing.T) {
		result := HelloWorld("Johar")
		require.Equal(t, "Hallo Johar", result)
	})
	t.Run("Uddin", func(t *testing.T) {
		result := HelloWorld("Uddin")
		require.Equal(t, "Hallo Uddin", result)
	})
}

func TestMain(m *testing.M) {
	//Before
	fmt.Println("BEFORE UNIT TEST")

	m.Run()

	//After
	fmt.Println("AFTER UNIT TEST")
}

func TestSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Can not run on Linux")
	}

	result := HelloWorld("Jojon")
	require.Equal(t, "Hallo Jojon", result, "Result must be 'Hallo Jojon'")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Juned")
	assert.Equal(t, "Hallo Juned", result, "Result must be 'Hallo Juned'")
	fmt.Println("HelloWorld assert done")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Imam")
	require.Equal(t, "Hallo Imam", result, "Result must be 'Hallo Imam'")
	fmt.Println("HelloWorld Require done")
}

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Johar")
	if result != "Hallo Johar"{
		// init test failed
		t.Error("Result must be 'Hallo Johar'")

		fmt.Println("TestHelloWorld done")
	}
}

func TestHelloWorld2(t *testing.T) {
	result := HelloWorld("Agung")
	if result != "Hallo Agung"{
		// init test failed
		t.Fatal("Result must be 'Hallo Agung'")

		fmt.Println("TestHelloWorld2 done")
	}
}

// benchmarks
func BenchmarkHelloWorld(b *testing.B) {
	for i := 0; i <b.N; i++ {
		HelloWorld("johar")
	}
}

func BenchmarkHelloWorldJohar(b *testing.B) {
	for i := 0; i <b.N; i++ {
		HelloWorld("joharuddin")
	}
}

// sub benchmarks
func BenchmarkSub(b *testing.B) {
	b.Run("johar", func(b *testing.B) {
		for i := 0; i <b.N; i++ {
			HelloWorld("johar")
		}
	})
	b.Run("joharuddin", func(b *testing.B) {
		for i := 0; i <b.N; i++ {
			HelloWorld("joharuddin")
		}
	})
}

// benchmarks table
func BenchmarkHelloWorldTable(b *testing.B) {
	benchmarks := []struct{
		name string
		request string
	}{
		{
			name: "HelloWorld(johar)",
			request: "johar",
		},
		{
			name: "HelloWorld(juned)",
			request: "juned",
		},
	}
	for _,benchmark := range benchmarks {
		b.Run(benchmark.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(benchmark.request)
			}
		})
	}
}