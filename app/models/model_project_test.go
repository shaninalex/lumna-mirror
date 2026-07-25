package models_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/utils"
)

func Test_ProjectMeta_GetLastEntityNumber(t *testing.T) {
	// nil receiver must not panic and defaults to 0
	var nilMeta *models.ProjectMeta
	assert.Equal(t, uint(0), nilMeta.GetLastEntityNumber("task"))

	// nil map defaults to 0
	m := &models.ProjectMeta{}
	assert.Equal(t, uint(0), m.GetLastEntityNumber("task"))

	// absent entity type defaults to 0
	m = models.NewProjectMeta()
	assert.Equal(t, uint(0), m.GetLastEntityNumber("task"))

	// present entity type returns stored value
	m.LastEntityNumber["task"] = 7
	assert.Equal(t, uint(7), m.GetLastEntityNumber("task"))
	// unrelated entity still defaults to 0
	assert.Equal(t, uint(0), m.GetLastEntityNumber("epic"))
}

func Test_ProjectMeta_SetLastEntityNumber(t *testing.T) {
	// nil receiver must not panic
	var nilMeta *models.ProjectMeta
	assert.NotPanics(t, func() { nilMeta.SetLastEntityNumber("task") })

	// nil map is lazily initialized, first set yields 1
	m := &models.ProjectMeta{}
	assert.NotPanics(t, func() { m.SetLastEntityNumber("task") })
	assert.Equal(t, uint(1), m.GetLastEntityNumber("task"))

	// increments are sequential
	m.SetLastEntityNumber("task")
	assert.Equal(t, uint(2), m.GetLastEntityNumber("task"))

	// entity types are tracked independently
	m.SetLastEntityNumber("epic")
	assert.Equal(t, uint(1), m.GetLastEntityNumber("epic"))
	assert.Equal(t, uint(2), m.GetLastEntityNumber("task"))
}

func Test_Project_GetMeta(t *testing.T) {
	// nil Meta returns a usable, empty meta (not nil)
	p := &models.Project{}
	m := p.GetMeta()
	assert.NotNil(t, m)
	assert.NotNil(t, m.LastEntityNumber)
	assert.Equal(t, uint(0), m.GetLastEntityNumber("task"))

	// invalid JSON falls back to a usable, empty meta
	p.Meta = utils.Pointer("not json")
	m = p.GetMeta()
	assert.NotNil(t, m)
	assert.NotNil(t, m.LastEntityNumber)

	// valid JSON without the map key still yields a usable map
	p.Meta = utils.Pointer(`{}`)
	m = p.GetMeta()
	assert.NotNil(t, m.LastEntityNumber)
	assert.Equal(t, uint(0), m.GetLastEntityNumber("task"))

	// valid JSON is parsed
	p.Meta = utils.Pointer(`{"last_entity_number":{"task":5}}`)
	m = p.GetMeta()
	assert.Equal(t, uint(5), m.GetLastEntityNumber("task"))
}

func Test_Project_SetMeta_RoundTrip(t *testing.T) {
	p := &models.Project{}
	m := models.NewProjectMeta()
	m.SetLastEntityNumber("task")

	assert.NoError(t, p.SetMeta(m))
	assert.NotNil(t, p.Meta)

	got := p.GetMeta()
	assert.Equal(t, uint(1), got.GetLastEntityNumber("task"))
}

// Test_Project_Meta_Persists mirrors the handleEventTaskCreated flow:
// read meta -> increment -> write back -> the increment survives a reload.
func Test_Project_Meta_Persists(t *testing.T) {
	p := &models.Project{}

	m := p.GetMeta()
	m.SetLastEntityNumber("task")
	assert.NoError(t, p.SetMeta(m))

	// simulating a fresh read of the persisted project
	reloaded := &models.Project{Meta: p.Meta}
	assert.Equal(t, uint(1), reloaded.GetMeta().GetLastEntityNumber("task"))
}
