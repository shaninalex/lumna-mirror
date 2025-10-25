import { createAction, props } from "@ngrx/store"
import { Project, ProjectPatch } from "@client/entities/project/model/project.model"

export const ProjectListSetAction = createAction("[project] set list", props<{ payload: Project[] }>())

export const ProjectListGetAction = createAction("[project] get list")

export const ProjectSetAction = createAction("[project] set", props<{ payload: Project }>())

export const ProjectCreateAction = createAction("[project] create", props<{ payload: Record<string, string> }>())

export const ProjectPatchAction = createAction("[project] patch", props<{ projectId: number; payload: ProjectPatch }>())

export const ProjectUpdateAction = createAction("[project] update", props<{ payload: Project }>())
