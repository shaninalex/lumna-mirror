import {Component, inject} from '@angular/core';
import {AppState} from '@client/shared/store';
import {Store} from '@ngrx/store';
import {selectProject} from '@client/entities/project/model/project.selectors';
import {ActivatedRoute, RouterLink, RouterOutlet} from '@angular/router';
import {AsyncPipe} from '@angular/common';
import {filter, map, switchMap, take, tap} from 'rxjs';
import {CdkMenuModule} from '@angular/cdk/menu';
import {OverlayModule} from '@angular/cdk/overlay';

@Component({
    selector: "fr-project-detail-page",
    imports: [
        AsyncPipe,
        CdkMenuModule,
        OverlayModule,
        RouterLink,
        RouterOutlet,
    ],
    template: `
        @if (project$ | async; as project) {
            <div class="flex items-center gap-4 mb-4">
                <a class="hover:underline gap-2 flex items-center" [routerLink]="['/', 'projects', project.code]">
                    <img src="/img/project.svg" class="w-6 rounded"/>
                    <h3 class="font-bold text-xl">{{ project.title }}</h3>
                </a>

                <button class="cursor-pointer"
                        cdkOverlayOrigin
                        #trigger="cdkOverlayOrigin"
                        (click)="isOpen = !isOpen">
                    <i class="i-dots-menu"></i>
                </button>
            </div>
            <ng-template
                cdkConnectedOverlay
                cdkConnectedOverlayBackdropClass="bg-slate-50/25"
                [cdkConnectedOverlayOrigin]="trigger"
                [cdkConnectedOverlayOpen]="isOpen"
                [cdkConnectedOverlayHasBackdrop]="true"
                (backdropClick)="isOpen = false"
            >
                <ul class="dropdown">
                    <li>
                        <a class="hover:underline" [routerLink]="['/', 'projects', project.code, 'settings']">
                            <i class="i-settings"></i> Project Settings
                        </a>
                    </li>
                </ul>
            </ng-template>
            <router-outlet/>
        }
    `
})
export class ProjectDetailPageComponent {
    private store = inject(Store<AppState>);
    private route = inject(ActivatedRoute);

    isOpen = false;
    project$ = this.route.params.pipe(
        filter(params => "projectKey" in params),
        map(params => params["projectKey"]),
        switchMap(code => this.store.select(selectProject(code))
            .pipe(
                take(1),
                filter(project => !!project)
            )
        )
    );
}
