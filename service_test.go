package main

import (
	"testing"

	"github.com/magiconair/properties/assert"
)

func TestAdd(t *testing.T) {
	type args struct {
		model *model
	}

	testCases := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "test add",
			args: args{
				model: &model{
					Name: "test",
				},
			},
			err: nil,
		},
	}

	s := NewService()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := s.add(tc.args.model.Name)

			assert.Equal(t, err, tc.err)
		})
	}
}

func TestRemove(t *testing.T) {
	type args struct {
		id int
	}

	testCases := []struct {
		name string
		args args
		err  error
	}{
		{
			name: "test remove",
			args: args{
				id: 1,
			},
			err: nil,
		},
	}

	s := NewService()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			err := s.remove(tc.args.id)

			assert.Equal(t, err, tc.err)
		})
	}
}

func TestGetAll(t *testing.T) {
	testCases := []struct {
		name     string
		expedted []model
		err      error
	}{
		{
			name:     "test get all",
			expedted: []model{},
			err:      nil,
		},
	}

	s := NewService()

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual, err := s.getAll()

			assert.Equal(t, tc.expedted, actual)
			assert.Equal(t, tc.err, err)
		})
	}
}
