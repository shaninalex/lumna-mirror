import {Component, inject} from '@angular/core';
import {BoardViewComponent} from '@client/features/project';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {ActivatedRoute} from '@angular/router';
import {filter, map, switchMap, take, tap} from 'rxjs';
import {selectProject} from '@client/entities/project/model/project.selectors';
import {GetStatusListActions} from '@client/entities/status';
import {GetTasksActions} from '@client/entities/task';
import {AsyncPipe} from '@angular/common';

@Component({
    selector: 'fr-board-view-page',
    imports: [
        BoardViewComponent,
        AsyncPipe,
    ],
    template: `
        @if (project$ | async; as project) {
            <fr-board-view-feature [project]="project"/>
        }
    `
})
export class BoardViewPageComponent {
    private store = inject(Store<AppState>);
    private route = inject(ActivatedRoute);

    project$ = this.route.params.pipe(
        filter(params => "projectKey" in params),
        map(params => params["projectKey"]),
        switchMap(code => this.store.select(selectProject(code)).pipe(
                take(1),
                filter(project => !!project),
                tap(project => {
                    this.store.dispatch(GetStatusListActions({projectId: project.id}))
                    this.store.dispatch(GetTasksActions({projectId: project.id}))
                })
            )
        )
    );
}
