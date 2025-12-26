import { ProjectModel } from "./project.model";
import { patchState, signalStore, withHooks, withMethods } from "@ngrx/signals";
import { withState } from "@ngrx/signals";
import { Events, withEventHandlers } from "@ngrx/signals/events";
import { inject } from '@angular/core';
import { filter, switchMap, tap } from 'rxjs';
import { mapResponse } from '@ngrx/operators';
import { ProjectService } from '@entities/project/api/project.service';
import { projectEvents } from '@entities/project';
import { addEntities, addEntity, removeEntity, updateEntity, withEntities } from '@ngrx/signals/entities';

export const ProjectStore = signalStore(
    { providedIn: 'root' },
    withEntities<ProjectModel>(),
    withHooks({
        onInit() {
            console.log('ProjectStore initialized');
        },
    }),
    withMethods((store) => ({
        byId(id: number): ProjectModel | undefined {
            console.log(store.entities())
            return store.entities().find(p => p.id === id)
        }
    })),
    withEventHandlers(
        (
            store,
            events = inject(Events),
            projectService = inject(ProjectService)
        ) => ({
            actionGetProjects$: events
                .on(projectEvents.getProjects)
                .pipe(
                    tap(() => console.log(projectEvents.getProjects.type)),
                    switchMap(() =>
                        projectService.GetProjects().pipe(
                            mapResponse({
                                next: projects => projectEvents.setProjects(projects),
                                error: error => console.log(error)
                            })
                        )
                    )
                ),

            actionCreateProject$: events
                .on(projectEvents.createProject)
                .pipe(
                    tap(() => console.log(projectEvents.createProject.type)),
                    switchMap(e =>
                        projectService.CreateProject(e.payload).pipe(
                            mapResponse({
                                next: project => projectEvents.setProject(project),
                                error: err => projectEvents.failed(err)
                            })
                        )
                    )
                ),

            actionDeleteProject$: events
                .on(projectEvents.deleteProject)
                .pipe(
                    tap(() => console.log(projectEvents.deleteProject.type)),
                    switchMap(e =>
                        projectService.DeleteProject(e.payload).pipe(
                            mapResponse({
                                next: () => projectEvents._deleteProjectSuccess(e.payload),
                                error: err => projectEvents.failed(err)
                            })
                        )
                    )
                ),

            actionPatchProject$: events
                .on(projectEvents.patch)
                .pipe(
                    tap(() => console.log(projectEvents.deleteProject.type)),
                    switchMap(e =>
                        projectService.Patch(e.payload.id, e.payload.data).pipe(
                            mapResponse({
                                next: data => projectEvents.updateProject(data),
                                error: err => projectEvents.failed(err)
                            })
                        )
                    )
                ),

            updateProject$: events
                .on(projectEvents.updateProject)
                .pipe(
                    tap(e => patchState(store, updateEntity({
                        id: e.payload.id,
                        changes: e.payload,
                    })))
                ),


            deleteProject$: events
                .on(projectEvents._deleteProjectSuccess)
                .pipe(
                    tap(e => patchState(store, removeEntity(e.payload)))
                ),

            setProjectList$: events
                .on(projectEvents.setProjects)
                .pipe(
                    filter(data => !!data.payload),
                    tap(e => patchState(store, addEntities(e.payload)))
                ),

            setProject$: events
                .on(projectEvents.setProject)
                .pipe(
                    tap(e => patchState(store, addEntity(e.payload)))
                ),
        })
    )
)

