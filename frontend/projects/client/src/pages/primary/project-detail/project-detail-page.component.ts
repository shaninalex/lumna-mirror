import {Component, inject} from '@angular/core';
import {AppState} from '@client/shared/store';
import {Store} from '@ngrx/store';
import {selectProject} from '@client/entities/project/model/project.selectors';
import {ActivatedRoute} from '@angular/router';
import {AsyncPipe} from '@angular/common';
import {filter, map, switchMap, take, tap} from 'rxjs';
import {BoardViewComponent} from '@client/features/project';
import {CdkMenuModule} from '@angular/cdk/menu';
import {GetTasksActions} from '@client/entities/task';
import {GetStatusListActions} from '@client/entities/status';
import {OverlayModule} from '@angular/cdk/overlay';

@Component({
    selector: "fr-project-detail-page",
    imports: [
        AsyncPipe,
        BoardViewComponent,
        CdkMenuModule,
        OverlayModule,
    ],
    template: `
        @if (project$ | async; as project) {
            <div class="flex items-center gap-2 mb-4">
                <img src="/img/project.svg" class="w-6 rounded"/>
                <h3 class="font-bold text-xl">{{ project.title }}</h3>

                <button class="cursor-pointer"
                        cdkOverlayOrigin
                        #trigger="cdkOverlayOrigin"
                        (click)="isOpen = !isOpen">
                    <i class="i-dots-menu"></i>
                </button>

                <ng-template
                    cdkConnectedOverlay
                    cdkConnectedOverlayBackdropClass="bg-slate-50/25"
                    [cdkConnectedOverlayOrigin]="trigger"
                    [cdkConnectedOverlayOpen]="isOpen"
                    [cdkConnectedOverlayHasBackdrop]="true"
                    (backdropClick)="isOpen = false"
                >
                    <ul class="dropdown">
                        <li>Change view</li>
                        <li>Settings</li>
                    </ul>
                </ng-template>
            </div>

            <fr-board-view-feature [project]="project"/>
        }
    `
})
export class ProjectDetailPageComponent {
    private store = inject(Store<AppState>);
    private route = inject(ActivatedRoute);
    isOpen = false
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
