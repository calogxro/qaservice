package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestService(t *testing.T) {
	// Setup

	es := NewEventStoreStub()
	service := NewQAService(es)
	rr := NewReadRepositoryStub()
	projector := NewProjector(rr)
	projection := NewQAProjection(rr)

	es.Subscribe(func(event *Event) {
		projector.Project(event)
	})

	//Create

	answer := testAnswer
	service.CreateAnswer(answer)
	projAnswer, _ := projection.GetAnswer("name")

	assert.Equal(t, &answer, projAnswer)

	// Update

	answer = Answer{Key: answer.Key, Value: answer.Value + "_2"}
	service.UpdateAnswer(answer)
	projAnswer, _ = projection.GetAnswer("name")

	assert.Equal(t, &answer, projAnswer)

	// Delete

	service.DeleteAnswer(answer.Key)
	projAnswer, err := projection.GetAnswer("name")

	assert.Nil(t, projAnswer)
	assert.NotNil(t, err)
	assert.IsType(t, &KeyNotFound{}, err)

	// History

	events, _ := service.GetHistory(answer.Key)
	assert.Equal(t, 3, len(events))
}

// sequence allowed:
// create → delete → create → update
func TestCreateDeleteCreateUpdate(t *testing.T) {
	service := NewQAService(NewEventStoreStub())

	var testAnswer = Answer{Key: "name", Value: "John"}
	var err error

	// Create
	_, err = service.CreateAnswer(testAnswer)
	assert.Nil(t, err)

	// Delete
	_, err = service.DeleteAnswer(testAnswer.Key)
	assert.Nil(t, err)

	// Create
	_, err = service.CreateAnswer(testAnswer)
	assert.Nil(t, err)

	// Update
	_, err = service.UpdateAnswer(testAnswer)
	assert.Nil(t, err)
}

// sequence allowed:
// create → update → delete → create → update
func TestCreateUpdateDeleteCreateUpdate(t *testing.T) {
	service := NewQAService(NewEventStoreStub())

	var testAnswer = Answer{Key: "name", Value: "John"}
	var err error

	// Create
	_, err = service.CreateAnswer(testAnswer)
	assert.Nil(t, err)

	// Update
	_, err = service.UpdateAnswer(testAnswer)
	assert.Nil(t, err)

	// Delete
	_, err = service.DeleteAnswer(testAnswer.Key)
	assert.Nil(t, err)

	// Create
	_, err = service.CreateAnswer(testAnswer)
	assert.Nil(t, err)

	// Update
	_, err = service.UpdateAnswer(testAnswer)
	assert.Nil(t, err)
}

// sequence not allowed:
// create → delete → update
func TestCreateDeleteUpdate(t *testing.T) {
	service := NewQAService(NewEventStoreStub())

	var testAnswer = Answer{Key: "name", Value: "John"}
	var err error

	// Create
	_, err = service.CreateAnswer(testAnswer)
	assert.Nil(t, err)

	// Delete
	_, err = service.DeleteAnswer(testAnswer.Key)
	assert.Nil(t, err)

	// Update
	_, err = service.UpdateAnswer(testAnswer)
	assert.NotNil(t, err)
}

// sequence not allowed:
// create → create
func TestCreateCreate(t *testing.T) {
	service := NewQAService(NewEventStoreStub())

	var testAnswer = Answer{Key: "name", Value: "John"}
	var err error

	// Create
	_, err = service.CreateAnswer(testAnswer)
	assert.Nil(t, err)

	// Create
	_, err = service.CreateAnswer(testAnswer)
	assert.NotNil(t, err)
}

// sequence not allowed:
// create → delete → delete
func TestCreateDeleteDelete(t *testing.T) {
	service := NewQAService(NewEventStoreStub())

	var testAnswer = Answer{Key: "name", Value: "John"}
	var err error

	// Create
	_, err = service.CreateAnswer(testAnswer)
	assert.Nil(t, err)

	// Delete
	_, err = service.DeleteAnswer(testAnswer.Key)
	assert.Nil(t, err)

	// Delete
	_, err = service.DeleteAnswer(testAnswer.Key)
	assert.NotNil(t, err)
}