import { ProjectModel } from "./project.model";
import {patchState, signalStore} from "@ngrx/signals";
import { withState } from "@ngrx/signals";
import {Events, withEventHandlers} from "@ngrx/signals/events";
import {inject} from '@angular/core';
import {switchMap, tap} from 'rxjs';
import {mapResponse} from '@ngrx/operators';
import {ProjectService} from '@entities/project/api/project.service';
import {projectEvents} from '@entities/project';
import {appEvents} from '@core/store/app.events';

type ProjectState = {
    projects: ProjectModel[];
};

const initialState: ProjectState = {
    projects: [],
};

export const ProjectStore = signalStore(
    { providedIn: 'root' },
    withState(initialState),
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
                                next: (projects) => {
                                    return projectEvents.setProjects(projects)
                                },
                                error: error => console.log(error)
                            })
                        )
                    )
                ),
            setProjects$: events
                .on(projectEvents.setProjects)
                .pipe(
                    tap(eventData => patchState(store, {
                        projects: eventData.payload,
                    }))
                )

        })
    )
)

