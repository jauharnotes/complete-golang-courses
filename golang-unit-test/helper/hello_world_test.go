package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

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