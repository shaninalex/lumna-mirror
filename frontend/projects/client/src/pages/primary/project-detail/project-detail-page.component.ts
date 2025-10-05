import {Component, inject, OnInit} from '@angular/core';
import {
    ActivatedRoute,
    NavigationEnd,
    NavigationStart,
    Router,
    RouterLink,
    RouterLinkActive,
    RouterOutlet
} from '@angular/router';
import {filter} from 'rxjs';
import {CdkMenuModule} from '@angular/cdk/menu';
import {OverlayModule} from '@angular/cdk/overlay';
import {Project} from '@client/entities/project';

@Component({
    selector: "fr-project-detail-page",
    imports: [
        CdkMenuModule,
        OverlayModule,
        RouterLink,
        RouterOutlet,
        RouterLinkActive,
    ],
    template: `
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
                    <a class="hover:underline"
                       [routerLink]="['/', 'projects', project.code, 'settings']"
                       [routerLinkActive]="'active-link'"
                    >
                        <i class="i-settings"></i> Project Settings
                    </a>
                </li>
            </ul>
        </ng-template>

        <router-outlet/>
    `
})
export class ProjectDetailPageComponent implements OnInit {
    private route = inject(ActivatedRoute);
    private router = inject(Router);
    project: Project
    isOpen = false;

    ngOnInit() {
        this.route.data.subscribe(data => {
            this.project = data['project']
        })

        this.router.events
            .pipe(filter(e => e instanceof NavigationStart))
            .subscribe(() => this.isOpen = false);
    }
}
