package tests

import (
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"testing"
)

type reqMap map[string]struct{}

func newReqMap(tasks ...string) reqMap {
	result := make(reqMap)
	for _, key := range tasks {
		result[key] = struct{}{}
	}
	return result
}

func (r reqMap) String() string {
	toSlice := make([]string, 0, len(r))
	for k := range r {
		toSlice = append(toSlice, k)
	}
	return strings.Join(toSlice, ",")
}

var requiredTasks = newReqMap("test", "cover", "build")

func validateTaskFile(taskMap map[string]interface{}) error {
	t, ok := taskMap["tasks"]
	if !ok {
		return errors.New("tasks field not found in Taskfile")
	}

	tasks := t.(map[string]interface{})

	for k := range tasks {
		_, has := requiredTasks[k]
		if has {
			delete(requiredTasks, k)
		}
	}

	if len(requiredTasks) > 0 {
		return errors.Errorf("not required tasks in Taskfile: %v", requiredTasks)
	}

	return nil
}

func TestTaskFile(t *testing.T) {
	taskFile, err := os.ReadFile("../Taskfile.yaml")
	assert.NoError(t, err)
	taskMap := make(map[string]interface{})
	err = yaml.Unmarshal(taskFile, &taskMap)
	assert.NoError(t, err)
	err = validateTaskFile(taskMap)
	assert.NoError(t, err)
}
