package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
)

func Test_ProjectMeta_GetNextEntityNumber(t *testing.T) {
	// nil receiver must not panic and defaults to 1
	var nilMeta *models.ProjectMeta
	assert.Equal(t, uint(1), nilMeta.GetNextEntityNumber("task"))

	// nil map defaults to 1
	m := &models.ProjectMeta{}
	assert.Equal(t, uint(1), m.GetNextEntityNumber("task"))

	// absent entity type defaults to 1
	m = models.NewProjectMeta()
	assert.Equal(t, uint(1), m.GetNextEntityNumber("task"))

	// present entity type returns stored value
	m.NextEntityNumber["task"] = 7
	assert.Equal(t, uint(7), m.GetNextEntityNumber("task"))
	// unrelated entity still defaults to 1
	assert.Equal(t, uint(1), m.GetNextEntityNumber("epic"))
}

func Test_ProjectMeta_SetNextEntityNumber(t *testing.T) {
	// nil receiver must not panic
	var nilMeta *models.ProjectMeta
	assert.NotPanics(t, func() { nilMeta.SetNextEntityNumber("task") })

	// nil map is lazily initialized, first set yields next number 2
	m := &models.ProjectMeta{}
	assert.NotPanics(t, func() { m.SetNextEntityNumber("task") })
	assert.Equal(t, uint(2), m.GetNextEntityNumber("task"))

	// increments are sequential
	m.SetNextEntityNumber("task")
	assert.Equal(t, uint(3), m.GetNextEntityNumber("task"))

	// entity types are tracked independently
	m.SetNextEntityNumber("epic")
	assert.Equal(t, uint(2), m.GetNextEntityNumber("epic"))
	assert.Equal(t, uint(3), m.GetNextEntityNumber("task"))
}

func Test_Project_GetMeta(t *testing.T) {
	// nil Meta returns a usable, empty meta (not nil)
	p := &models.Project{}
	m := p.GetMeta()
	assert.NotNil(t, m)
	assert.NotNil(t, m.NextEntityNumber)
	assert.Equal(t, uint(1), m.GetNextEntityNumber("task"))

	// invalid JSON falls back to a usable, empty meta
	p.Meta = utils.Pointer("not json")
	m = p.GetMeta()
	assert.NotNil(t, m)
	assert.NotNil(t, m.NextEntityNumber)

	// valid JSON without the map key still yields a usable map
	p.Meta = utils.Pointer(`{}`)
	m = p.GetMeta()
	assert.NotNil(t, m.NextEntityNumber)
	assert.Equal(t, uint(1), m.GetNextEntityNumber("task"))

	// valid JSON is parsed
	p.Meta = utils.Pointer(`{"next_entity_number":{"task":5}}`)
	m = p.GetMeta()
	assert.Equal(t, uint(5), m.GetNextEntityNumber("task"))
}

func Test_Project_SetMeta_RoundTrip(t *testing.T) {
	p := &models.Project{}
	m := models.NewProjectMeta()
	m.SetNextEntityNumber("task")

	assert.NoError(t, p.SetMeta(m))
	assert.NotNil(t, p.Meta)

	got := p.GetMeta()
	assert.Equal(t, uint(2), got.GetNextEntityNumber("task"))
}

// Test_Project_Meta_Persists mirrors the handleEventTaskCreated flow:
// read meta -> increment -> write back -> the increment survives a reload.
func Test_Project_Meta_Persists(t *testing.T) {
	p := &models.Project{}

	m := p.GetMeta()
	m.SetNextEntityNumber("task")
	assert.NoError(t, p.SetMeta(m))

	// simulating a fresh read of the persisted project
	reloaded := &models.Project{Meta: p.Meta}
	assert.Equal(t, uint(2), reloaded.GetMeta().GetNextEntityNumber("task"))
}
