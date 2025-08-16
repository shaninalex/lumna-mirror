import {Component, inject} from '@angular/core';
import {UiService} from '@client/shared/ui';
import {BehaviorSubject,} from 'rxjs';
import {NgClass} from '@angular/common';
import {MatListModule} from '@angular/material/list';
import {MatIconModule} from '@angular/material/icon';
import {RouterLink} from '@angular/router';
import { version } from '../../../../../../../../package.json';


@Component({
    selector: "ts-sidebar",
    imports: [
        NgClass,
        MatListModule,
        MatIconModule,
        RouterLink,
    ],
    template: `
        <div class="shadow h-full transition-[width] flex flex-col space-between"
             [ngClass]="{
                'w-3xs': !sidebarFolded.value,
                'w-10': sidebarFolded.value,
            }">
            <mat-list>
                @for (item of sideNav; track $index) {
                    <mat-list-item [routerLink]="[item.link]">
                        <mat-icon class="cursor-pointer" matListItemIcon>{{ item.icon }}</mat-icon>
                        <div class="cursor-pointer" matListItemTitle>{{ item.label }}</div>
                    </mat-list-item>
                }
            </mat-list>
            <div class="flex-grow"></div>
            <div class="text-xs p-1 text-slate-600">v{{Version}}</div>
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
            link: "",
            active: true
        },
        {
            label: "Recent",
            icon: "browse_gallery",
            link: "",
            active: true
        },
        {
            label: "Starred",
            icon: "star",
            link: "",
            active: true
        },
        {
            label: "Projects",
            icon: "workspaces",
            link: "",
            active: true
        },
        {
            label: "Teams",
            icon: "groups",
            link: "",
            active: false
        }
    ]
}
