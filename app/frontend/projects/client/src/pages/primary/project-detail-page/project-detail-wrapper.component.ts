import { Component, inject, OnInit } from '@angular/core'
import { ActivatedRoute, NavigationStart, Router, RouterLink, RouterLinkActive, RouterOutlet } from '@angular/router'
import { filter } from 'rxjs'
import { CdkMenuModule } from '@angular/cdk/menu'
import { OverlayModule } from '@angular/cdk/overlay'
import { Project } from '@client/entities/project'

@Component({
    selector: 'lu-project-detail-page',
    imports: [CdkMenuModule, OverlayModule, RouterLink, RouterOutlet, RouterLinkActive],
    template: `
        <div class="mb-4 flex items-center gap-4">
            <a class="flex items-center gap-2 hover:underline" [routerLink]="['/', 'projects', project.code]">
                <img src="/img/project.svg" class="w-6 rounded" />
                <h3 class="text-xl font-bold">{{ project.name }}</h3>
            </a>

            <button class="cursor-pointer" cdkOverlayOrigin #trigger="cdkOverlayOrigin" (click)="isOpen = !isOpen">
                <i class="i-dots-menu"></i>
            </button>
        </div>
        <ng-template
            cdkConnectedOverlay
            [cdkConnectedOverlayOrigin]="trigger"
            [cdkConnectedOverlayOpen]="isOpen"
            [cdkConnectedOverlayHasBackdrop]="true"
            (backdropClick)="isOpen = false"
        >
            <ul class="dropdown">
                <li>
                    <a
                        class="hover:underline"
                        [routerLink]="['/', 'projects', project.code, 'settings']"
                        [routerLinkActive]="'active-link'"
                    >
                        <i class="i-settings"></i> Project Settings
                    </a>
                </li>
            </ul>
        </ng-template>

        <router-outlet />
    `,
})
export class ProjectDetailWrapperComponent implements OnInit {
    private route = inject(ActivatedRoute)
    private router = inject(Router)
    project: Project
    isOpen = false

    ngOnInit() {
        this.route.data.subscribe(data => {
            this.project = data['project']
        })

        this.router.events.pipe(filter(e => e instanceof NavigationStart)).subscribe(() => (this.isOpen = false))
    }
}
