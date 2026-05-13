import { createFeatureSelector, createSelector } from "@ngrx/store";
import { projectAdapter, ProjectState } from "./project.store";
import { ProjectModel } from "./project.model";

const feature = createFeatureSelector<ProjectState>("project");
const selectors = projectAdapter.getSelectors();
export const selectProjectList = createSelector(feature, (state) =>
    selectors.selectAll(state)
);

export const selectProject = (id: number) =>
    createSelector(selectProjectList, (list) =>
        list.filter((a: ProjectModel) => a.id === id)
    );
