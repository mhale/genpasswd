package main

import (
	"log"
	"os/exec"
	"testing"
)

// Testing table to match inputs with expected outputs.
var tests = []struct {
	in  string
	out int
}{
	{"", 12},
	{"10", 10},
	{"12", 12},
	{"0", 12},
	{"-10", 12},
	{"foo", 12},
}

// Build the executable so we can run it in TestLength().
func init() {
	if err := exec.Command("go", "build").Run(); err != nil {
		log.Fatal(err)
	}
}

// Remove the compiled executable after the tests have run.
func cleanup() {
	if err := exec.Command("rm", "./genpasswd").Run(); err != nil {
		log.Fatal(err)
	}
}

// Test for correct password length.
func TestLength(t *testing.T) {
	defer cleanup()
	for _, tt := range tests {
		out, err := exec.Command("./genpasswd", tt.in).Output()
		if err != nil {
			log.Fatal(err)
		}
		password := out[:len(out)-1]
		if x := len(password); x != tt.out {
			t.Errorf("genpasswd %s = %s: length is %d, want %d", tt.in, password, x, tt.out)
		}
	}
}

// Benchmark to find run time of the generate method.
// Run with: go test -bench=.
func BenchmarkGenerate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		generate(defaultLength)
	}
}
