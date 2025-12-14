package repositories_test

// func Test_RepositoryBoardListCreate(t *testing.T) {
// 	ctx := tests.TestContext()
// 	tests.ResetDatabase()

// 	repo := repositories.NewProjectRespository()
// 	listRepo := repositories.NewBoardListRepository()

// 	project := models.Project{Name: "test"}
// 	_ = repo.Create(ctx, &project)

// 	entry := models.BoardList{
// 		Name: "todo",
// 		BoardId: pro,
// 	}
// 	listRepo.Create(ctx)

// 	count, err := repo.Count(ctx)
// 	assert.NoError(t, err, "Should count project without errors")
// 	assert.Equal(t, 1, count)
// }

// func Test_RepositoryBoardListCount(t *testing.T) {
// 	ctx := tests.TestContext()
// 	tests.ResetDatabase()

// 	repo := repositories.NewProjectRespository()

// 	_ = repo.Create(ctx, &models.Project{Name: "test"})
// 	_ = repo.Create(ctx, &models.Project{Name: "test 2"})

// 	count, err := repo.Count(ctx, db.Option{Key: "name", Value: "test 2"})
// 	assert.NoError(t, err, "Should count project without errors")
// 	assert.Equal(t, 1, count)

// 	count, err = repo.Count(ctx, db.Option{Key: "name", Value: "none existed project"})
// 	assert.NoError(t, err, "Should count project without errors")
// 	assert.Equal(t, 0, count)

// 	// We do not implement "OR" condition.
// 	// count, err = repo.Count(ctx, db.Option{Key: "name", Value: "test"}, db.Option{Key: "name", Value: "test 2"})
// 	// assert.NoError(t, err, "Should count project without errors")
// 	// assert.Equal(t, 2, count)
// }

// func Test_RepositoryBoardListDelete(t *testing.T) {
// 	ctx := tests.TestContext()
// 	tests.ResetDatabase()

// 	repo := repositories.NewProjectRespository()
// 	projectA := models.Project{Name: "A"}
// 	_ = repo.Create(ctx, &projectA)
// 	projectB := models.Project{Name: "B"}
// 	_ = repo.Create(ctx, &projectB)

// 	count, _ := repo.Count(ctx)
// 	assert.Equal(t, 2, count)

// 	err := repo.Delete(ctx, projectA.GetId())
// 	assert.NoError(t, err, "Should delete project without errors")
// 	count, _ = repo.Count(ctx)
// 	assert.Equal(t, 1, count)

// 	count, _ = repo.Count(ctx, db.Option{Key: "name", Value: "A"})
// 	assert.Equal(t, 0, count)

// 	count, _ = repo.Count(ctx, db.Option{Key: "name", Value: "B"})
// 	assert.Equal(t, 1, count)
// }

// func Test_RepositoryBoardListGet(t *testing.T) {
// 	ctx := tests.TestContext()
// 	tests.ResetDatabase()

// 	repo := repositories.NewProjectRespository()
// 	projectA := models.Project{Name: "A"}
// 	_ = repo.Create(ctx, &projectA)
// 	projectB := models.Project{Name: "B"}
// 	_ = repo.Create(ctx, &projectB)

// 	dbProjectA, err := repo.Get(ctx, projectA.GetId())
// 	assert.NoError(t, err, "Should get project without errors")
// 	assert.Equal(t, dbProjectA.Name, projectA.Name)

// 	dbProjectB, err := repo.Get(ctx, projectB.GetId())
// 	assert.NoError(t, err, "Should get project without errors")
// 	assert.Equal(t, dbProjectB.Name, projectB.Name)

// 	dbProjectC, err := repo.Get(ctx, 123)
// 	assert.Error(t, err, "Should NOT get non existed project")
// 	assert.Nil(t, dbProjectC)
// }

// func Test_RepositoryBoardListList(t *testing.T) {
// 	ctx := tests.TestContext()
// 	tests.ResetDatabase()

// 	repo := repositories.NewProjectRespository()
// 	projectA := models.Project{Name: "A"}
// 	_ = repo.Create(ctx, &projectA)
// 	projectB := models.Project{Name: "B"}
// 	_ = repo.Create(ctx, &projectB)

// 	projects, err := repo.List(ctx)
// 	assert.NoError(t, err, "Should list projects without errors")
// 	assert.Equal(t, 2, len(projects))

// 	testProjects := []string{projectA.Name, projectB.Name}
// 	queriedProjects := make([]string, len(projects))
// 	for i, u := range projects {
// 		queriedProjects[i] = u.Name
// 	}

// 	for _, e := range testProjects {
// 		assert.Contains(t, queriedProjects, e, fmt.Sprintf("Project %s does not exists in queriedProjects %v", e, queriedProjects))
// 	}
// }

// func Test_RepositoryBoardListListWithOption(t *testing.T) {
// 	ctx := tests.TestContext()
// 	tests.ResetDatabase()

// 	repo := repositories.NewProjectRespository()
// 	projectA := models.Project{Name: "A"}
// 	_ = repo.Create(ctx, &projectA)
// 	projectB := models.Project{Name: "B"}
// 	_ = repo.Create(ctx, &projectB)

// 	projects, err := repo.List(ctx, db.Option{Key: "name", Value: projectA.Name})
// 	assert.NoError(t, err, "Should list projects without errors")
// 	assert.Equal(t, 1, len(projects))
// 	assert.Equal(t, projectA.Name, projects[0].Name)
// }

// func Test_RepositoryBoardListUpdate(t *testing.T) {
// 	ctx := tests.TestContext()
// 	tests.ResetDatabase()

// 	repo := repositories.NewProjectRespository()
// 	projectA := models.Project{Name: "A"}
// 	_ = repo.Create(ctx, &projectA)

// 	newName := "New Name"

// 	err := repo.Update(ctx, projectA.GetId(), db.Option{Key: "name", Value: newName})
// 	assert.NoError(t, err, "Should update project without errors")

// 	updatedProject, _ := repo.Get(ctx, projectA.GetId())
// 	assert.Equal(t, newName, updatedProject.Name)
// }
