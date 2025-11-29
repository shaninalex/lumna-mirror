import { Actions, createEffect, ofType } from "@ngrx/effects"
import { inject } from "@angular/core"
import { Router } from "@angular/router"
import {
	ProjectCreateAction,
	ProjectListGetAction,
	ProjectListSetAction,
	ProjectPatchAction,
	ProjectSetAction,
	ProjectUpdateAction,
} from "@client/entities/project/model/project.actions"
import { exhaustMap, of, switchMap } from "rxjs"
import { ProjectService } from "@client/entities/project/api/project.service"

export const ProjectsGetEffect = createEffect(
	(actions$ = inject(Actions), service = inject(ProjectService), router = inject(Router)) => {
		return actions$.pipe(
			ofType(ProjectListGetAction.type),
			exhaustMap(() => service.List().pipe(switchMap(data => of(ProjectListSetAction({ payload: data })))))
		)
	},
	{ functional: true, dispatch: true }
)

export const ProjectCreateEffect = createEffect(
	(actions$ = inject(Actions), service = inject(ProjectService)) => {
		return actions$.pipe(
			ofType(ProjectCreateAction),
			exhaustMap(action => service.Create(action.payload).pipe(switchMap(result => of(ProjectSetAction({ payload: result })))))
		)
	},
	{ functional: true, dispatch: true }
)

export const ProjectPatchEffect = createEffect(
	(actions$ = inject(Actions), service = inject(ProjectService)) => {
		return actions$.pipe(
			ofType(ProjectPatchAction),
			exhaustMap(action =>
				service.Patch(action.projectId, action.payload).pipe(switchMap(data => of(ProjectUpdateAction({ payload: data }))))
			)
		)
	},
	{ functional: true, dispatch: true }
)
