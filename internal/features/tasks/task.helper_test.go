package tasks

import (
	"testing"
	entities "todo-go-fiber/internal/db/entities"
)

func TestValidateId(t *testing.T) {
	tests := []struct {
		idStr    string
		expected int64
		exectErr bool
	}{
		{"0", 0, false},
		{"1", 1, false},
		{"-1", -1, false},
		{"123", 123, false},
		{"-123", -123, false},
		{"1231231231231231322", 1231231231231231322, false},
		{"qwerty", 0, true},
		{"q", 0, true},
		{"12.3", 0, true},
	}

	for _, test := range tests {
		result, err := validateId(test.idStr)

		if (err != nil) != test.exectErr {
			t.Errorf("validateId(%v): expected error: '%v' GOT '%v'", test.idStr, test.exectErr, err != nil)
		}

		if result != test.expected {
			t.Errorf("validateId(%v): expected value: '%v' GOT '%v'", test.idStr, test.expected, result)
		}
	}
}

func TestGetCreateTaskQuery(t *testing.T) {
	desc := "description test"
	tests := []struct {
		task      *entities.Task
		expected  string
		expectErr bool
	}{
		{
			task:      &entities.Task{Title: "title test"},
			expected:  "INSERT INTO tasks (\n title \n)\nVALUES ( 'title test' )\nRETURNING id;",
			expectErr: false,
		},
		{
			task:      &entities.Task{Title: "title test", Description: nil},
			expected:  "INSERT INTO tasks (\n title \n)\nVALUES ( 'title test' )\nRETURNING id;",
			expectErr: false,
		},
		{
			task:      &entities.Task{Title: "title test", Description: &desc},
			expected:  "INSERT INTO tasks (\n title, description \n)\nVALUES ( 'title test', 'description test' )\nRETURNING id;",
			expectErr: false,
		},
		{
			task:      &entities.Task{Title: "", Description: nil},
			expected:  "",
			expectErr: true,
		},
	}

	for _, test := range tests {
		result, err := getCreateTaskQuery(test.task)

		if (err != nil) != test.expectErr {
			t.Errorf("getCreateTaskQuery(%v) expected error: %v, got: %v", test.task, test.expectErr, err)
		}
		if result != test.expected {
			t.Errorf("getCreateTaskQuery(%v) expected result: %v, got: %v", test.task, test.expected, result)
		}
	}
}
func TestGetUpdateTaskQuery(t *testing.T) {
	desc := "description test"
	tests := []struct {
		task      *entities.Task
		id        int64
		expected  string
		expectErr bool
	}{
		{
			task:      &entities.Task{Title: "title test"},
			id:        1,
			expected:  "UPDATE tasks set title = 'title test',updated_at = now()\nwhere id = '1'\nRETURNING id;",
			expectErr: false,
		},
		{
			task:      &entities.Task{Title: "title test", Description: nil},
			id:        2,
			expected:  "UPDATE tasks set title = 'title test',updated_at = now()\nwhere id = '2'\nRETURNING id;",
			expectErr: false,
		},
		{
			task:      &entities.Task{Title: "title test", Description: &desc, Status: "in_progress"},
			id:        3,
			expected:  "UPDATE tasks set title = 'title test', description = 'description test', status = 'in_progress',updated_at = now()\nwhere id = '3'\nRETURNING id;",
			expectErr: false,
		},
		{
			task:      &entities.Task{Title: "", Description: nil, Status: ""},
			id:        4,
			expected:  "",
			expectErr: true,
		},
		{
			task:      &entities.Task{Title: "", Description: &desc, Status: ""},
			id:        5,
			expected:  "UPDATE tasks set description = 'description test',updated_at = now()\nwhere id = '5'\nRETURNING id;",
			expectErr: false,
		},
		{
			task:      &entities.Task{Title: "", Description: nil, Status: "completed"},
			id:        6,
			expected:  "UPDATE tasks set status = 'completed',updated_at = now()\nwhere id = '6'\nRETURNING id;",
			expectErr: false,
		},
	}

	for _, test := range tests {
		result, err := getUpdateTaskQuery(test.task, test.id)

		if (err != nil) != test.expectErr {
			t.Errorf("getUpdateTaskQuery(%v, %d) expected error: %v, got: %v", test.task, test.id, test.expectErr, err)
		}
		if result != test.expected {
			t.Errorf("getUpdateTaskQuery(%v, %d) expected result: %v, got: %v", test.task, test.id, test.expected, result)
		}
	}
}
