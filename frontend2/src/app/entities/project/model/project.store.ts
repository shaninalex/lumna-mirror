import { createEntityAdapter, EntityState } from "@ngrx/entity";
import { createReducer, on } from "@ngrx/store";
import {
    actionProjectCreateSuccessful,
    actionProjectSetCurrentProjectId,
    actionProjectSetList
} from "./project.actions";
import { ProjectModel } from "./project.model";

export interface ProjectState extends EntityState<ProjectModel> {
    currentProjectId: number | undefined;
}
export const projectAdapter = createEntityAdapter<ProjectModel>();
const initialState = projectAdapter.getInitialState();

export const projectReducer = createReducer(
    initialState,
    on(actionProjectSetList, (state, { list }) =>
        projectAdapter.addMany(list, state)
    ),
    on(actionProjectCreateSuccessful, (state, { project }) =>
        projectAdapter.addOne(project, state)
    ),
    on(actionProjectSetCurrentProjectId, (state, { project_id }) => ({
        ...state,
        currentProjectId: project_id
    }))
);
