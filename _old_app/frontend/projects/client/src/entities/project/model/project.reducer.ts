import { createEntityAdapter, EntityAdapter, EntityState } from "@ngrx/entity"
import { Project } from "./project.model"
import { createReducer, on } from "@ngrx/store"
import { ProjectListSetAction, ProjectSetAction, ProjectUpdateAction } from "./project.actions"

export interface ProjectState extends EntityState<Project> {}
export const projectsAdapter: EntityAdapter<Project> = createEntityAdapter<Project>()
export const projectsReducer = createReducer(
	projectsAdapter.getInitialState(),
	on(ProjectListSetAction, (state, action) => projectsAdapter.addMany(action.payload, state)),
	on(ProjectSetAction, (state, action) => projectsAdapter.addOne(action.payload, state)),
	on(ProjectUpdateAction, (state, action) =>
		projectsAdapter.updateOne(
			{
				id: action.payload.id,
				changes: action.payload,
			},
			state
		)
	)
)
