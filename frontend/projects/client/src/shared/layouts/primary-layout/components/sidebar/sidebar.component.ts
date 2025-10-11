import {Component, inject, OnInit} from '@angular/core';
import {version} from '@root/package.json';
import {RouterLink, RouterLinkActive} from '@angular/router';
import {AsyncPipe, NgClass} from '@angular/common';
import {UiService} from '@client/shared/ui/ui.service';
import {Observable} from 'rxjs';

@Component({
    selector: 'lu-sidebar',
    imports: [
        RouterLink,
        RouterLinkActive,
        NgClass,
    ],
    template: `
        <div class="h-full flex flex-col space-between bg-base-100 border-e border-gray-300 dark:bg-gray-800 dark:border-gray-600">
            <a [routerLink]="['/']" class="flex items-center gap-2 p-4">
                @if (!extendSidebar) {
                    <img src="img/logo-h.svg" class="w-full" alt="Lumna"/>
                } @else {
                    <img src="img/logo-icon.svg" class="w-full" alt="Lumna"/>
                }
            </a>
            <div class="p-4 flex flex-col flex-grow"
                [ngClass]="{'text-center': extendSidebar}">
                <div class="flex flex-col gap-2">
                    @for (nav of menu; track $index) {
                        <a class="hover:underline"
                           [routerLinkActive]="'text-sky-500'"
                           [routerLink]="[nav.url]"
                           [routerLinkActiveOptions]="{exact: nav.exact}">
                            <i class="text-lg" [ngClass]="nav.icon"></i>
                            @if (!extendSidebar) {
                                <span class="ms-2">{{ nav.label }}</span>
                            }
                        </a>
                    }
                </div>
                <div class="flex-grow"></div>
                <div class="flex flex-col gap-2">
                    <a class="hover:underline"
                       [routerLinkActive]="'text-sky-500'"
                       [routerLink]="['settings']"
                       [routerLinkActiveOptions]="{exact: false}">
                        <i class="i-settings text-lg"></i>
                        @if (!extendSidebar) {
                            <span class="ms-2">Settings</span>
                        }
                    </a>
                </div>
                @if (!extendSidebar) {
                    <div class="text-xs p-1 text-gray-600">v{{ version }}</div>
                }
            </div>
        </div>
    `
})
export class SidebarComponent implements OnInit {
    private uiService = inject(UiService);
    extendSidebar: boolean;

    version: string = version;

    // Note: permission server should give proper amount of available links for a user.
    // if user does not allowed to see some url - it should not be in the menu.
    // even if user manually accessed it - permission server should return 403
    menu = [
        {
            label: "Home",
            url: "/",
            exact: true,
            icon: "i-home"
        },
        {
            label: "Projects",
            url: "/projects",
            exact: false,
            icon: "i-board"
        }
    ]

    ngOnInit() {
        this.uiService.extendSidebar().subscribe(data => this.extendSidebar = data)
    }
}
