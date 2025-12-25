import { ProjectModel } from "./project.model";
import {patchState, signalStore, withHooks, withMethods} from "@ngrx/signals";
import { withState } from "@ngrx/signals";
import {Events, withEventHandlers} from "@ngrx/signals/events";
import {inject} from '@angular/core';
import {switchMap, tap} from 'rxjs';
import {mapResponse} from '@ngrx/operators';
import {ProjectService} from '@entities/project/api/project.service';
import {projectEvents} from '@entities/project';
import {addEntities, addEntity, withEntities} from '@ngrx/signals/entities';

export const ProjectStore = signalStore(
    { providedIn: 'root' },
    withEntities<ProjectModel>(),
    withHooks({
        onInit() {
            console.log('ProjectStore initialized');
        },
    }),
    withEventHandlers(
        (
            store,
            events = inject(Events),
            projectService = inject(ProjectService)
        ) => ({
            getProjects$: events
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

            createProject$: events
                .on(projectEvents.createProject)
                .pipe(
                    tap(() => console.log(projectEvents.createProject.type)),
                    switchMap(e =>
                        projectService.CreateProject(e.payload).pipe(
                            mapResponse({
                                next: project => projectEvents.setProject(project),
                                error: err => projectEvents.createProjectFailed(err)
                            })
                        )
                    )
                ),

            setProjects$: events
                .on(projectEvents.setProjects)
                .pipe(
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

