import {Component} from '@angular/core';
import {version} from '@root/package.json';
import {RouterLink, RouterLinkActive} from '@angular/router';

@Component({
    selector: 'fr-sidebar',
    imports: [
        RouterLink,
        RouterLinkActive
    ],
    template: `
        <div class="h-full flex flex-col space-between w-48 bg-base-100 border-e border-base-300">
            <a [routerLink]="['/']" class="flex items-center gap-2 p-4">
                <img src="img/logo-icon.svg" class="w-10" alt="Flowreon"/>
                <span class="font-bold text-lg">Flowreon</span>
            </a>
            <div class="p-4 flex flex-col flex-grow">
                <div class="flex flex-col gap-2">
                    @for (nav of menu; track $index) {
                        <a class="hover:underline"
                           [routerLinkActive]="'text-sky-500'"
                           [routerLink]="[nav.url]"
                           [routerLinkActiveOptions]="{exact: nav.exact}">
                            {{ nav.label }}
                        </a>
                    }
                </div>
                <div class="flex-grow"></div>
                <div class="text-xs p-1 text-gray-600">v{{ version }}</div>
            </div>
        </div>
    `
})
export class Sidebar {
    version: string = version;

    // Note: permission server should give proper amount of available links for a user.
    // if user does not allowed to see some url - it should not be in the menu.
    // even if user manually accessed it - permission server should return 403
    menu = [
        {
            label: "Home",
            url: "/",
            exact: true,
        },
        {
            label: "Projects",
            url: "/projects",
            exact: false,
        }
    ]
}
