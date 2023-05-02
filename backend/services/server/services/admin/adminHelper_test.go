package adminservices

import (
	"fmt"
	"server/entities"
	"testing"

	"github.com/stretchr/testify/assert"
)

type checkRegexpTestCase struct {
	testName string
	input    string
	output   bool
}

func assertCheckRegexp(t *testing.T, testName string, input string, output bool) {
	want := output
	got := checkRegexp(input)
	if got != want {
		t.Errorf("%s with input = '%s' is supose to %v, but got %v", testName, input, want, got)
	}
}

func TestCheckRegexp(t *testing.T) {
	var checkRegexpTestCases = []checkRegexpTestCase{
		{testName: "normal case", input: "test1", output: true},
		{testName: "vietnamese case", input: "tét1", output: false},
		{testName: "have space", input: "test 1", output: false},
		{testName: "capital letter", input: "Test1", output: true},
		{testName: "special letter", input: "Test1", output: true},
	}
	for _, c := range checkRegexpTestCases {
		t.Run(c.testName, func(t *testing.T) {
			assertCheckRegexp(t, c.testName, c.input, c.output)
		})
	}
}

func TestGenerateToken(t *testing.T) {
	username := "test"
	assert := assert.New(t)
	_, err := generateToken(username)
	assert.Nil(err)
}

type checkIsAdminCorrectTestCase struct {
	testName     string
	inputAdmin   Admin
	inputEnAdmin entities.Admin
	output       error
}

func assertCheckIsAdminCorrect(t *testing.T, testName string, inputAdmin Admin, inputEnAdmin entities.Admin, output error) {
	want := output
	got := checkIsAdminCorrect(&inputAdmin, inputEnAdmin)
	assert := assert.New(t)
	if want != nil {
		assert.EqualError(got, want.Error(), "%s with input = '%v' and '%v' is supose to %v, but got %v", testName, inputAdmin, inputEnAdmin, want, got)
	} else {
		assert.Nil(got, "%s with input = '%v' and '%v' is supose to %v, but got %v", testName, inputAdmin, inputEnAdmin, want, got)
	}

}

func TestCheckIsAdminCorrect(t *testing.T) {
	var checkIsAdminCorrectTestCases = []checkIsAdminCorrectTestCase{
		{testName: "normal case", inputAdmin: Admin{Username: "Test1", Password: "test1"}, inputEnAdmin: entities.Admin{Username: "Test1", Password: "test1"}, output: nil},
		{testName: "username incorrect case", inputAdmin: Admin{Username: "Test1", Password: "test1"}, inputEnAdmin: entities.Admin{Username: "Test2", Password: "test1"}, output: fmt.Errorf("username is incorrect")},
		{testName: "password incorrect case", inputAdmin: Admin{Username: "Test1", Password: "test1"}, inputEnAdmin: entities.Admin{Username: "Test1", Password: "test2"}, output: fmt.Errorf("password is incorrect")},
	}
	for _, c := range checkIsAdminCorrectTestCases {
		t.Run(c.testName, func(t *testing.T) {
			assertCheckIsAdminCorrect(t, c.testName, c.inputAdmin, c.inputEnAdmin, c.output)
		})
	}
}

type validateAdminTestCase struct {
	testName string
	input    Admin
	output   error
}

func assertValidateAdmin(t *testing.T, testName string, input Admin, output error) {
	want := output
	got := validateAdmin(&input)
	assert := assert.New(t)
	if want != nil {
		assert.EqualError(got, want.Error(), "%s with input = '%v' is supose to %v, but got %v", testName, input, want, got)
	} else {
		assert.Nil(got, "%s with input = '%v' is supose to %v, but got %v", testName, input, want, got)
	}
}

func TestValidateAdmin(t *testing.T) {
	var validateAdminTestCases = []validateAdminTestCase{
		{testName: "normal case", input: Admin{Username: "admintest", Password: "test1"}, output: nil},
		{testName: "username have space case", input: Admin{Username: "admin test", Password: "test1"}, output: fmt.Errorf("username must not contain special character")},
		{testName: "username contain spectial letter", input: Admin{Username: "admin*test", Password: "test1"}, output: fmt.Errorf("username must not contain special character")},
		{testName: "username vietnamese", input: Admin{Username: "admintét", Password: "test1"}, output: fmt.Errorf("username must not contain special character")},
		{testName: "password contain spectial letter", input: Admin{Username: "admintest", Password: "test@1"}, output: fmt.Errorf("password must not contain special character")},
		{testName: "password have space case", input: Admin{Username: "admintest", Password: "test 1"}, output: fmt.Errorf("password must not contain special character")},
		{testName: "password vietnamese", input: Admin{Username: "admintest", Password: "admintét"}, output: fmt.Errorf("password must not contain special character")},
		{testName: "username too short", input: Admin{Username: "admin", Password: "test1"}, output: fmt.Errorf("Key: 'Admin.Username' Error:Field validation for 'Username' failed on the 'min' tag")},
		{testName: "username too long", input: Admin{Username: "admintest123213123", Password: "test1"}, output: fmt.Errorf("Key: 'Admin.Username' Error:Field validation for 'Username' failed on the 'max' tag")},
	}
	for _, c := range validateAdminTestCases {
		t.Run(c.testName, func(t *testing.T) {
			assertValidateAdmin(t, c.testName, c.input, c.output)
		})
	}
}
