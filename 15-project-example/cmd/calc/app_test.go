//go:build !urfave

package main

import (
	"testing"

	"github.com/hankmor/calc/pkg/calculator"
)

func TestParseArguments(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		wantA   float64
		wantOp  string
		wantB   float64
		wantErr bool
	}{
		{name: "valid", args: []string{"10", "+", "20"}, wantA: 10, wantOp: "+", wantB: 20},
		{name: "invalid first number", args: []string{"bad", "+", "20"}, wantErr: true},
		{name: "invalid second number", args: []string{"10", "+", "bad"}, wantErr: true},
		{name: "missing argument", args: []string{"10", "+"}, wantErr: true},
		{name: "extra argument", args: []string{"10", "+", "20", "extra"}, wantErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			a, operator, b, err := parseArguments(test.args)
			if test.wantErr {
				if err == nil {
					t.Fatal("parseArguments() error = nil, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("parseArguments() error = %v", err)
			}
			if a != test.wantA || operator != test.wantOp || b != test.wantB {
				t.Fatalf("parseArguments() = %v, %q, %v; want %v, %q, %v", a, operator, b, test.wantA, test.wantOp, test.wantB)
			}
		})
	}
}

func TestCalculate(t *testing.T) {
	tests := []struct {
		name    string
		a       float64
		op      string
		b       float64
		want    float64
		wantErr bool
	}{
		{name: "add", a: 10, op: "+", b: 2, want: 12},
		{name: "subtract", a: 10, op: "-", b: 2, want: 8},
		{name: "multiply", a: 10, op: "*", b: 2, want: 20},
		{name: "divide", a: 10, op: "/", b: 2, want: 5},
		{name: "unsupported operator", a: 10, op: "%", b: 2, wantErr: true},
		{name: "division by zero", a: 10, op: "/", b: 0, wantErr: true},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got, err := calculate(calculator.New(), test.a, test.op, test.b)
			if test.wantErr {
				if err == nil {
					t.Fatal("calculate() error = nil, want error")
				}
				return
			}
			if err != nil {
				t.Fatalf("calculate() error = %v", err)
			}
			if got != test.want {
				t.Fatalf("calculate() = %v, want %v", got, test.want)
			}
		})
	}
}
