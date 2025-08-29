import {Component, inject} from '@angular/core';
import {UiService} from '@client/shared/ui';
import {BehaviorSubject,} from 'rxjs';
import {NgClass} from '@angular/common';
import {MatListModule} from '@angular/material/list';
import {MatIconModule} from '@angular/material/icon';
import {RouterLink, RouterLinkActive} from '@angular/router';
import {version} from '../../../../../../../../package.json';


@Component({
    selector: "ts-sidebar",
    imports: [
        NgClass,
        MatListModule,
        MatIconModule,
        RouterLink,
        RouterLinkActive,
    ],
    template: `
        <div class="border-r border-slate-200 h-full flex flex-col space-between"
             [ngClass]="{
                'w-3xs': !sidebarFolded.value,
                'w-10': sidebarFolded.value,
            }">
            <div class="flex flex-col">
                @for (item of sideNav; track $index) {
                    <a [routerLink]="[item.active ? item.link : '/404']"
                       [routerLinkActive]="'bg-slate-100 text-teal-500'"
                       [ngClass]="{
                            'text-slate-400 pointer-events-none': !item.active,
                       }"
                       class="flex gap-2 whitespace-nowrap p-2 hover:bg-slate-100">
                        <mat-icon matListItemIcon>{{ item.icon }}
                        </mat-icon>
                        @if (!sidebarFolded.value) {
                            <span>{{ item.label }}</span>
                        }
                    </a>
                }
            </div>
            <div class="flex-grow"></div>
            <div class="text-xs p-1 text-slate-600">v{{ Version }}</div>
        </div>`
})
export class SidebarComponent {
    uiService = inject(UiService);
    sidebarFolded: BehaviorSubject<boolean> = this.uiService.sidebar
    public Version: string = version
    sideNav = [
        {
            label: "For You",
            icon: "account_circle",
            link: "/for-you",
            active: true
        },
        {
            label: "Recent",
            icon: "browse_gallery",
            link: "/tasks/recent",
            active: true
        },
        {
            label: "Starred",
            icon: "star",
            link: "/starred",
            active: true
        },
        {
            label: "Projects",
            icon: "workspaces",
            link: "/projects",
            active: true
        },
        {
            label: "Teams",
            icon: "groups",
            link: "/teams",
            active: false
        }
    ];
}
