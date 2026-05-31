import { createFeatureSelector, createSelector } from "@ngrx/store";
import { projectAdapter, ProjectState } from "./project.store";
import { ProjectModel } from "./project.model";
import { selectWorkspaceCurrentWorkspaceSlug, selectWorkspaceList } from "@entities/workspace";

const feature = createFeatureSelector<ProjectState>("project");
const selectors = projectAdapter.getSelectors();
export const selectProjectList = createSelector(feature, (state) =>
    selectors.selectAll(state)
);

export const selectProject = (id: number) =>
    createSelector(selectProjectList, (list) =>
        list.filter((a: ProjectModel) => a.id === id)
    );

export const selectCurrentProjectList = createSelector(
    selectWorkspaceCurrentWorkspaceSlug,
    selectWorkspaceList,
    selectProjectList,
    (slug, workspaces, projects) => {
        const workspace = workspaces.find(
            (workspace) => workspace.slug === slug
        );
        if (!workspace) {
            return null;
        }
        return projects.filter(
            (project) => project.workspace_id === workspace.id
        );
    }
);
