package main

import "testing"

// BenchmarkSayStruct undocummented.
func BenchmarkSayStruct(b *testing.B) {
	l := Lion{}
	for i := 0; i < b.N; i++ {
		sayStruct(l, "hello")
	}
}

// BenchmarkSayInterface undocummented.
func BenchmarkSayInterface(b *testing.B) {
	l := Lion{}
	for i := 0; i < b.N; i++ {
		sayInterface(l, "hello")
	}
}
