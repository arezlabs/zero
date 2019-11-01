package ci_test

import (
	"bytes"
	"io/ioutil"
	"os"
	"sync"
	"testing"

	"github.com/commitdev/commit0/internal/config"
	"github.com/commitdev/commit0/internal/generate/ci"
	"github.com/commitdev/commit0/internal/templator"
	"github.com/gobuffalo/packr/v2"
)

var testData = "../../test_data/ci/"

// setupTeardown removes all the generated test files before and after
// the test runs to ensure clean data.
func setupTeardown(t *testing.T) func(t *testing.T) {
	os.RemoveAll("../../test_data/ci/actual")
	return func(t *testing.T) {
		os.RemoveAll("../../test_data/ci/actual")
	}
}

func TestGenerateJenkins(t *testing.T) {
	teardown := setupTeardown(t)
	defer teardown(t)

	templates := packr.New("templates", "../../../templates")
	testTemplator := templator.NewTemplator(templates)

	var waitgroup *sync.WaitGroup

	testConf := &config.Commit0Config{
		Language: "go",
		CI: config.CI{
			System: "jenkins",
		},
	}

	err := ci.Generate(testTemplator.CI, testConf, testData+"/actual", waitgroup)
	if err != nil {
		t.Errorf("Error when executing test. %s", err)
	}

	actual, err := ioutil.ReadFile(testData + "actual/Jenkinsfile")
	if err != nil {
		t.Errorf("Error reading created file: %s", err.Error())
	}
	expected, err := ioutil.ReadFile(testData + "/expected/Jenkinsfile")
	if err != nil {
		t.Errorf("Error reading created file: %s", err.Error())
	}

	if !bytes.Equal(expected, actual) {
		t.Errorf("want:\n%s\n\n, got:\n%s\n\n", string(expected), string(actual))
	}
}

func TestGenerateCircleCI(t *testing.T) {
	teardown := setupTeardown(t)
	defer teardown(t)

	templates := packr.New("templates", "../../../templates")
	testTemplator := templator.NewTemplator(templates)

	var waitgroup *sync.WaitGroup

	testConf := &config.Commit0Config{
		Language: "go",
		CI: config.CI{
			System: "circleci",
		},
	}

	err := ci.Generate(testTemplator.CI, testConf, testData+"/actual", waitgroup)
	if err != nil {
		t.Errorf("Error when executing test. %s", err)
	}

	actual, err := ioutil.ReadFile(testData + "actual/.circleci/config.yml")
	if err != nil {
		t.Errorf("Error reading created file: %s", err.Error())
	}
	expected, err := ioutil.ReadFile(testData + "/expected/.circleci/config.yml")
	if err != nil {
		t.Errorf("Error reading created file: %s", err.Error())
	}

	if !bytes.Equal(expected, actual) {
		t.Errorf("want:\n%s\n\ngot:\n%s\n\n", string(expected), string(actual))
	}
}

func TestGenerateTravisCI(t *testing.T) {
	teardown := setupTeardown(t)
	defer teardown(t)

	templates := packr.New("templates", "../../../templates")
	testTemplator := templator.NewTemplator(templates)

	var waitgroup *sync.WaitGroup

	testConf := &config.Commit0Config{
		Language: "go",
		CI: config.CI{
			System: "travisci",
		},
	}

	err := ci.Generate(testTemplator.CI, testConf, testData+"/actual", waitgroup)
	if err != nil {
		t.Errorf("Error when executing test. %s", err)
	}

	actual, err := ioutil.ReadFile(testData + "actual/.travis.yml")
	if err != nil {
		t.Errorf("Error reading created file: %s", err.Error())
	}
	expected, err := ioutil.ReadFile(testData + "/expected/.travis.yml")
	if err != nil {
		t.Errorf("Error reading created file: %s", err.Error())
	}

	if !bytes.Equal(expected, actual) {
		t.Errorf("want:\n%s\n\n, got:\n%s\n\n", string(expected), string(actual))
	}
}
