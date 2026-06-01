import { createFeatureSelector, createSelector } from "@ngrx/store";
import { projectAdapter, ProjectState } from "./project.store";
import { ProjectModel } from "./project.model";
import { selectWorkspaceCurrentWorkspaceId, selectWorkspaceList } from "@entities/workspace";

const feature = createFeatureSelector<ProjectState>("project");
const selectors = projectAdapter.getSelectors();
export const selectProjectList = createSelector(feature, (state) =>
    selectors.selectAll(state)
);

export const selectProject = (id: number) =>
    createSelector(selectProjectList, (list) =>
        list.find((a: ProjectModel) => a.id === id)
    );

export const selectCurrentProjectList = createSelector(
    selectWorkspaceCurrentWorkspaceId,
    selectWorkspaceList,
    selectProjectList,
    (id, workspaces, projects) => {
        const workspace = workspaces.find((workspace) => workspace.id === id);
        if (!workspace) {
            return null;
        }
        return projects
            .filter((project) => project.workspace_id === workspace.id)
            .map((project) => ({
                ...project,
                appLink: `/app/${workspace.id}/project/${project.id}`
            }));
    }
);

export const selectProjectLink = (id: number) =>
    createSelector(
        selectProject(id),
        selectWorkspaceList,
        (project, workspaces) => {
            if (!workspaces) {
                return undefined;
            }
            if (!project) {
                return undefined;
            }
            const workspace = workspaces.find(
                (workspace) => workspace.id === project.workspace_id
            );
            if (!workspace) {
                return undefined;
            }
            return `/app/${workspace.id}/project/${project.id}`;
        }
    );

export const selectProjectCurrentId = createSelector(feature, (s) => {
    return s.currentProjectId;
});
