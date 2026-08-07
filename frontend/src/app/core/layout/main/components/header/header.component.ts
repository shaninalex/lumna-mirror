import { Component, DestroyRef, inject } from '@angular/core';
import { Store } from '@ngrx/store';
import { actionToggleSidebar } from '@core';
import { Actions, ofType } from '@ngrx/effects';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { UserMenuComponent } from "../user-menu";
import { ProjectsMenuComponent } from "../projects-menu";
import { ThemeSwitcherComponent } from "../theme-switcher";


@Component({
    selector: 'lu-header',
    imports: [UserMenuComponent, ProjectsMenuComponent, ThemeSwitcherComponent],
    styleUrl: './header.component.css',
    template: `
        <nav class="navbar navbar-expand-lg bg-body-tertiary">
            <div class="container-fluid">
                <div class="d-flex align-items-center gap-2">

                    <button class="btn btn-sm btn-outline-secondary" (click)="toggleSidebar()">
                        @if (sidebarHidden) {
                            <i class="fa-solid fa-chevron-right"></i>
                        } @else {
                            <i class="fa-solid fa-bars"></i>
                        }
                    </button>
                    <lu-projects-menu />
                </div>

                <div class="flex align-items-center">
                    <lu-theme-switcher />
                    <button class="btn btn-sm">
                        <i class="fa-solid fa-bell"></i>
                    </button>
                    <lu-user-menu />
                </div>
            </div>
        </nav>
    `,
})
export class HeaderComponent {
    private store = inject(Store);
    sidebarHidden = false;

    private actions$ = inject(Actions);
    private ref = inject(DestroyRef);

    constructor() {
        this.actions$
            .pipe(
                ofType(actionToggleSidebar),
                takeUntilDestroyed(this.ref),
            ).subscribe(() => this.sidebarHidden = !this.sidebarHidden);
    }

    toggleSidebar(): void {
        this.store.dispatch(actionToggleSidebar());
    }
}
