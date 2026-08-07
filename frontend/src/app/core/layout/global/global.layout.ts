import { Component } from '@angular/core';
import { ThemeSwitcherComponent } from '../main/components/theme-switcher';
import { UserMenuComponent } from '../main/components/user-menu';
import { RouterLink } from "@angular/router";

@Component({
    selector: 'lu-global-layout',
    imports: [ThemeSwitcherComponent, UserMenuComponent, RouterLink],
    template: `
        <nav class="navbar navbar-expand-lg border-bottom bg-body">
            <div class="container-fluid">
                <a routerLink="/">
                    <img src="images/logo-h.svg" alt="" style="width: 160px">
                </a>

                <div class="flex align-items-center">
                    <lu-theme-switcher />
                    <button class="btn btn-sm">
                        <i class="fa-solid fa-bell"></i>
                    </button>
                    <lu-user-menu />
                </div>
            </div>
        </nav>

        <ng-content />
    `,
})
export class GlobalLayout {
}
