import {Component, inject} from '@angular/core';
import {AppState} from '@client/shared/store';
import {Store} from '@ngrx/store';
import {selectProject} from '@client/entities/project/model/project.selectors';
import {ActivatedRoute} from '@angular/router';
import {AsyncPipe} from '@angular/common';
import {filter, map, switchMap} from 'rxjs';

@Component({
    selector: "fr-project-detail-page",
    imports: [
        AsyncPipe,
    ],
    template: `
        @if (project$ | async; as project) {
            <div>{{ project.title }}</div>
        }
    `
})
export class ProjectDetailPageComponent {
    private store = inject(Store<AppState>);
    private route = inject(ActivatedRoute);

    project$ = this.route.params.pipe(
        filter(params => "projectKey" in params),
        map(params => params["projectKey"]),
        switchMap(code => this.store.select(selectProject(code)))
    );
}
