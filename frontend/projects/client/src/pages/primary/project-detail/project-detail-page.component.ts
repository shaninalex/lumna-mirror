import {Component, inject} from '@angular/core';
import {AppState} from '@client/shared/store';
import {Store} from '@ngrx/store';
import {selectProject} from '@client/entities/project/model/project.selectors';
import {ActivatedRoute} from '@angular/router';
import {AsyncPipe} from '@angular/common';
import {filter, map, switchMap, tap} from 'rxjs';
import {BoardViewComponent} from '@client/features/project/board-view-feature';
import {CdkMenuModule} from '@angular/cdk/menu';
import {GetTasksActions} from '@client/entities/task';


@Component({
    selector: "fr-project-detail-page",
    imports: [
        AsyncPipe,
        BoardViewComponent,
        CdkMenuModule,
    ],
    template: `
        @if (project$ | async; as project) {
            <div class="flex items-center gap-2 mb-4">
                <img src="/img/project.svg" class="w-6 rounded"/>
                <h3 class="font-bold text-xl">{{ project.title }}</h3>
                <div class="dropdown">
                    <div tabindex="0" role="button" class="btn btn-primary btn-sm">
                        menu
                    </div>
                    <ul tabindex="0" class="dropdown-content left-0 menu bg-base-100 rounded-box z-1 w-36 p-2 shadow-sm">
                        <li>Change view</li>
                        <li>Settings</li>
                    </ul>
                </div>
            </div>

            <fr-board-view-feature [project]="project" />
        }
    `
})
export class ProjectDetailPageComponent {
    private store = inject(Store<AppState>);
    private route = inject(ActivatedRoute);

    project$ = this.route.params.pipe(
        filter(params => "projectKey" in params),
        map(params => params["projectKey"]),
        switchMap(code => {
            this.store.dispatch(GetTasksActions({ projectCode: code }))
            return this.store.select(selectProject(code))
        })
    );
}
