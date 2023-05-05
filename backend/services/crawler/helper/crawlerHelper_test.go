package crawlerhelpers

import (
	"testing"
)

type formatToSearchTestcase struct {
	testName string
	input    string
	output   string
}

func assertFormatToSearch(t *testing.T, testName string, input string, output string) {
	want := output
	got := FormatToSearch(input)
	if got != want {
		t.Errorf("%s with input = '%s' is supose to %v, but got %s", testName, input, want, got)
	}
}

func TestFormatToSearch(t *testing.T) {
	var formatToSearchTestcases = []formatToSearchTestcase{
		{testName: "normal case", input: "bán kết", output: "ban+ket"},
		{testName: "capital", input: "Bán Kết", output: "ban+ket"},
		{testName: "All caps", input: "BÁN KẾT", output: "ban+ket"},
		{testName: "Already format", input: "ban ket", output: "ban+ket"},
		{testName: "Hello ửold", input: "hello ửold", output: "hello+uold"},
	}
	for _, c := range formatToSearchTestcases {
		t.Run(c.testName, func(t *testing.T) {
			assertFormatToSearch(t, c.testName, c.input, c.output)
		})
	}
}

type formatClassNameTestcase struct {
	testName string
	input    string
	output   string
}

func assertFormatClassName(t *testing.T, testName string, input string, output string) {
	want := output
	got := FormatClassName(input)
	if got != want {
		t.Errorf("%s with input = '%s' is supose to %v, but got %s", testName, input, want, got)
	}
}

func TestFormatClassName(t *testing.T) {
	var formatClassNameTestcases = []formatClassNameTestcase{
		{testName: "normal case", input: "testclass", output: ".testclass"},
		{testName: "two class", input: "test class", output: ".test.class"},
	}
	for _, c := range formatClassNameTestcases {
		t.Run(c.testName, func(t *testing.T) {
			assertFormatClassName(t, c.testName, c.input, c.output)
		})
	}
}

type formatDateTestcase struct {
	testName string
	input    string
	output   string
}

func assertFormatDate(t *testing.T, testName string, input string, output string) {
	want := output
	got := FormatDate(input)
	if got != want {
		t.Errorf("%s with input = '%s' is supose to %v, but got %s", testName, input, want, got)
	}
}

func TestFormatDate(t *testing.T) {
	var formatDateTestcases = []formatDateTestcase{
		{testName: "normal case", input: "monday, 20/3", output: "20/3"},
		{testName: "multi date", input: "monday, 20/3, 21/3", output: "20/3"},
		{testName: "only one date", input: "20/3", output: "20/3"},
	}
	for _, c := range formatDateTestcases {
		t.Run(c.testName, func(t *testing.T) {
			assertFormatDate(t, c.testName, c.input, c.output)
		})
	}
}
