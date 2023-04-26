package serverhelper

import (
	"reflect"
	serverhelper "server/helper"
	"testing"
	"time"
)

type formatVietnameseTestcase struct {
	testName string
	input    string
	output   string
}

func assertFormatVietnamese(t *testing.T, testName string, input string, output string) {
	want := output
	got := serverhelper.FormatVietnamese(input)
	if got != want {
		t.Errorf("%s with input = '%s' is supose to %v, but got %s", testName, input, want, got)
	}
}

func TestFormatVietnamese(t *testing.T) {
	var formatVietnameseTestcases = []formatVietnameseTestcase{
		{testName: "normal case", input: "bán kết", output: "ban ket"},
		{testName: "capital", input: "Bán Kết", output: "ban ket"},
		{testName: "All caps", input: "BÁN KẾT", output: "ban ket"},
		{testName: "Already format", input: "ban ket", output: "ban ket"},
		{testName: "Hello ửold", input: "hello ửold", output: "hello uold"},
	}
	for _, c := range formatVietnameseTestcases {
		t.Run(c.testName, func(t *testing.T) {
			assertFormatVietnamese(t, c.testName, c.input, c.output)
		})
	}
}

type formatElasticSearchIndexNameTestcase struct {
	testName string
	input    string
	output   string
}

func assertFormatElasticSearchIndexName(t *testing.T, testName string, input string, output string) {
	want := output
	got := serverhelper.FormatElasticSearchIndexName(input)
	if got != want {
		t.Errorf("%s with input = '%s' is supose to %v, but got %s", testName, input, want, got)
	}
}

func TestFormatElasticSearchIndexName(t *testing.T) {
	var formatElasticSearchIndexNameTestcases = []formatElasticSearchIndexNameTestcase{
		{testName: "normal case", input: "việt nam", output: "vietnam"},
		{testName: "capital", input: "Việt Nam", output: "vietnam"},
		{testName: "All caps", input: "Việt Nam", output: "vietnam"},
		{testName: "Already format", input: "vietnam", output: "vietnam"},
		{testName: "Hello ửold", input: "hello ửold", output: "hellouold"},
		{testName: "with ()", input: "viet nam (vietnam)", output: "vietnam"},
		{testName: "Lots of space", input: "   V i  ệ T   N  A   M", output: "vietnam"},
	}
	for _, c := range formatElasticSearchIndexNameTestcases {
		t.Run(c.testName, func(t *testing.T) {
			assertFormatElasticSearchIndexName(t, c.testName, c.input, c.output)
		})
	}
}

type fortmatTagsFromRequestTestcase struct {
	testName string
	input    string
	output   []string
}

func assertFormatTagsFromRequest(t *testing.T, testName string, input string, output []string) {
	want := output
	got := serverhelper.FortmatTagsFromRequest(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s with input = '%s' is supose to %v, but got %s", testName, input, want, got)
	}
}

func TestFortmatTagsFromRequest(t *testing.T) {
	var fortmatTagsFromRequestTestcases = []fortmatTagsFromRequestTestcase{
		{testName: "normal case", input: "vietnam,banket,goldenowl", output: []string{"vietnam", "banket", "goldenowl"}},
		{testName: "empty", input: ",", output: []string{}},
		{testName: "last coma", input: "vietnam,banket,goldenowl,", output: []string{"vietnam", "banket", "goldenowl"}},
		{testName: "first coma", input: ",vietnam,banket,goldenowl", output: []string{"vietnam", "banket", "goldenowl"}},
		{testName: "lots of coma", input: ",  , , ,,vietnam,banket,goldenowl, , , , ,", output: []string{"vietnam", "banket", "goldenowl"}},
	}
	for _, c := range fortmatTagsFromRequestTestcases {
		t.Run(c.testName, func(t *testing.T) {
			assertFormatTagsFromRequest(t, c.testName, c.input, c.output)
		})
	}
}

type fortmatDateVietnameseToDateStringTestcase struct {
	testName string
	input    time.Time
	output   string
}

func assertFormatDateToVietnamesDateSting(t *testing.T, testName string, input time.Time, output string) {
	want := output
	got := serverhelper.FormatDateToVietnamesDateSting(input)
	if !reflect.DeepEqual(got, want) {
		t.Errorf("%s with input = '%s' is supose to %v, but got %s", testName, input, want, got)
	}
}

func TestFormatDateToVietnamesDateSting(t *testing.T) {
	dateString := "2021-11-22"
	date, _ := time.Parse("2006-01-02", dateString)
	var fortmatDateVietnameseToDateStringTestcases = []fortmatDateVietnameseToDateStringTestcase{
		{testName: "normal case", input: date, output: "22/11/2021"},
	}
	for _, c := range fortmatDateVietnameseToDateStringTestcases {
		t.Run(c.testName, func(t *testing.T) {
			assertFormatDateToVietnamesDateSting(t, c.testName, c.input, c.output)
		})
	}
}

type formatCacheKeyTestcase struct {
	testName string
	input    string
	output   string
}

func assertFormatCacheKey(t *testing.T, testName string, input string, output string) {
	want := output
	got := serverhelper.FormatCacheKey(input)
	if got != want {
		t.Errorf("%s with input = '%s' is supose to %v, but got %s", testName, input, want, got)
	}
}

func TestFormatCacheKey(t *testing.T) {
	var formatCacheKeyTestcases = []formatCacheKeyTestcase{
		{testName: "normal case", input: "việt nam", output: "viet_nam"},
		{testName: "capital", input: "Việt Nam", output: "viet_nam"},
		{testName: "All caps", input: "Việt Nam", output: "viet_nam"},
		{testName: "Already format", input: "vietnam", output: "vietnam"},
		{testName: "Hello ửold", input: "hello ửold", output: "hello_uold"},
		{testName: "with +", input: "viet+nam ", output: "viet_nam"},
		{testName: "with -", input: "viet-nam ", output: "viet_nam"},
		{testName: "Lots of space", input: "   V i  ệ T   N  A   M", output: "v_i__e_t___n__a___m"},
	}
	for _, c := range formatCacheKeyTestcases {
		t.Run(c.testName, func(t *testing.T) {
			assertFormatCacheKey(t, c.testName, c.input, c.output)
		})
	}
}
