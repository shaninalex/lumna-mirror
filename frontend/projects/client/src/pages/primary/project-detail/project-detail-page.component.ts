import {Component, inject} from '@angular/core';
import {AppState} from '@client/shared/store';
import {Store} from '@ngrx/store';
import {selectProject} from '@client/entities/project/model/project.selectors';
import {ActivatedRoute} from '@angular/router';
import {AsyncPipe} from '@angular/common';
import {filter, map, switchMap, tap} from 'rxjs';
import {BoardViewComponent} from '@client/features/project/board-view-feature';
import {MatButtonModule} from '@angular/material/button';
import {CdkMenuModule, CdkMenuTrigger} from '@angular/cdk/menu';
import {GetTasksActions} from '@client/entities/task';


@Component({
    selector: "fr-project-detail-page",
    imports: [
        AsyncPipe,
        BoardViewComponent,
        CdkMenuTrigger,
        CdkMenuModule,
    ],
    template: `
        @if (project$ | async; as project) {
            <div class="flex items-center gap-2 mb-4">
                <img src="/img/project.svg" class="w-6 rounded"/>
                <h3 class="font-bold text-xl">{{ project.title }}</h3>
                <button class="btn" [cdkMenuTriggerFor]="project_detail_menu">
                    menu
                </button>
                <ng-template #project_detail_menu>
                    <div class="bg-white flex flex-col gap-2 border p-2 rounded" cdkMenu>
                        <button cdkMenuItem>Refresh</button>
                        <button cdkMenuItem>Settings</button>
                        <button cdkMenuItem>Help</button>
                        <button cdkMenuItem>Sign out</button>
                    </div>
                </ng-template>
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
