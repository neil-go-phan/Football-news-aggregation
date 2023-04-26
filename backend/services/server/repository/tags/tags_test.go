package tagsrepo

import (
	"fmt"
	"reflect"
	"server/entities"
	"server/repository"

	"testing"

	"github.com/stretchr/testify/assert"
)

var PATH = "./testJson/"
var PATH_FAIL = "./testJson/fail/"
var PATH_WRITE = "./testJson/writedTest"

type addTagTestCase struct {
	testName string
	input    string
	output   error
}

func assertAddTag(t *testing.T, addTagTestCase addTagTestCase, tagsRepo repository.TagRepository) {
	want := addTagTestCase.output
	got := tagsRepo.AddTag(addTagTestCase.input)
	if got != want {
		assert.EqualErrorf(t, got, want.Error(), "%s with input = '%s' is supose to %#v, but got %#v", addTagTestCase.testName, addTagTestCase.input, want, got)
	}

}

func TestAddTag(t *testing.T) {
	contructorTags := entities.Tags{
		Tags: []string{"tin tuc bong da", "premier league", "v-league"},
	}

	var addTagTestCases = []addTagTestCase{
		{testName: "tag input normal case", input: "bán kết", output: nil},
		{testName: "tag input capital", input: "Bán Kết", output: nil},
		{testName: "tag input all caps", input: "BÁN KẾT", output: nil},
		{testName: "tag input already format", input: "ban ket", output: nil},
		{testName: "tag input is hello ửold", input: "hello ửold", output: nil},
		{testName: "Tag already exist", input: "tin tuc bong da", output: fmt.Errorf("tag tin tuc bong da already exist")},
	}

	for _, c := range addTagTestCases {
		t.Run(c.testName, func(t *testing.T) {
			tagsRepo := NewTagsRepo(contructorTags, PATH)
			assertAddTag(t, c, tagsRepo)
		})
	}
}

type removeInput struct {
	slice []string
	index int
}

type removeTestCase struct {
	testName string
	input    removeInput
	output   []string
}

func assertRemove(t *testing.T, removeTestCases removeTestCase) {
	want := removeTestCases.output
	got := removeTag(removeTestCases.input.slice, removeTestCases.input.index)
	if !reflect.DeepEqual(want, got) {
		t.Errorf("%s with input = '%#v' is supose to %#v, but got %#v", removeTestCases.testName, removeTestCases.input, want, got)
	}

}

func TestRemove(t *testing.T) {
	var removeTestCases = []removeTestCase{
		{testName: "remove normal", input: removeInput{
			slice: []string{"tin tuc bong da", "premier league", "v-league"},
			index: 0,
		}, output: []string{"v-league", "premier league"}},
		{testName: "remove unexist index", input: removeInput{
			slice: []string{"tin tuc bong da", "premier league", "v-league"},
			index: -1,
		}, output: []string{"tin tuc bong da", "premier league", "v-league"}},
	}

	for _, c := range removeTestCases {
		t.Run(c.testName, func(t *testing.T) {
			assertRemove(t, c)
		})
	}
}

func TestAddTagWrongPath(t *testing.T) {
	assert := assert.New(t)
	contructorTags := entities.Tags{
		Tags: []string{"tin tuc bong da", "premier league", "v-league"},
	}

	newTag := "wrong path"

	tagsRepo := NewTagsRepo(contructorTags, PATH_FAIL)

	want := "file json not found"

	got := tagsRepo.AddTag(newTag)

	assert.Errorf(got, want, fmt.Sprintf("Method AddTag is supose to %#v, but got %#v", want, got))
}

func TestListTag(t *testing.T) {
	assert := assert.New(t)
	contructorTags := entities.Tags{
		Tags: []string{"tin tuc bong da", "premier league", "v-league"},
	}

	tagsRepo := NewTagsRepo(contructorTags, PATH)

	want := entities.Tags{
		Tags: []string{"tin tuc bong da", "premier league", "v-league"},
	}

	got := tagsRepo.ListTags()

	assert.Equal(got, want, fmt.Sprintf("Method ListTag is supose to %#v, but got %#v", want, got))
}

func TestReadTagsJSONSuccess(t *testing.T) {
	assert := assert.New(t)
	contructorTags := entities.Tags{
		Tags: []string{"tin tuc bong da", "premier league", "v-league"},
	}

	tagsRepo := NewTagsRepo(contructorTags, PATH)

	want := entities.Tags{
		Tags: []string{"tin tuc bong da", "premier league", "v-league", "hello uold"},
	}
	got, _ := tagsRepo.ReadTagsJSON()

	assert.Equal(want, got, fmt.Sprintf("Method ReadTagsJSON is supose to %#v, but got %#v", want, got))
}

func TestReadTagsJSONCantOpenFile(t *testing.T) {
	assert := assert.New(t)
	contructorTags := entities.Tags{
		Tags: []string{"tin tuc bong da", "premier league", "v-league"},
	}

	want := "file json not found"

	tagsRepo := NewTagsRepo(contructorTags, PATH_FAIL)

	_, got := tagsRepo.ReadTagsJSON()

	assert.Errorf(got, want, fmt.Sprintf("Method ReadTagsJSON is supose to %#v, but got %#v", want, got))
}

func TestDeleteTagSuccess(t *testing.T) {
	assert := assert.New(t)
	contructorTags := entities.Tags{
		Tags: []string{"tin tuc bong da", "premier league", "v-league"},
	}

	tagsRepo := NewTagsRepo(contructorTags, PATH_WRITE)

	got := tagsRepo.DeleteTag("tin tuc bong da")

	assert.Nilf(got, fmt.Sprintf("Method ReadTagsJSON is supose to nil, but got %#v", got))
}
