import {Component, DestroyRef, inject} from '@angular/core';
import {Store} from '@ngrx/store';
import {actionToggleSidebar} from '@core';
import { Actions, ofType } from '@ngrx/effects';
import { takeUntilDestroyed } from '@angular/core/rxjs-interop';
import { NgClass } from '@angular/common';

@Component({
    selector: 'lu-header',
    imports: [NgClass],
    styleUrl: './header.component.css',
    template: `
        <nav class="navbar navbar-expand-lg bg-body-tertiary">
            <div class="container-fluid">
                <button class="btn btn-sm btn-outline-secondary me-2" (click)="toggleSidebar()">
                    @if (sidebarHidden) {
                        <i class="fa-solid fa-chevron-right"></i>
                    } @else {
                        <i class="fa-solid fa-bars"></i>
                    }
                </button>
                <a href="#" class="me-4">
                    <img src="/images/logo-h.svg" alt="" style="width: 140px">
                </a>
                <button class="navbar-toggler" type="button" (click)="toggleMenu()">
                    <span class="navbar-toggler-icon"></span>
                </button>
                <div class="collapse navbar-collapse" [ngClass]="{'show': showMenu}">
                    <ul class="navbar-nav me-auto mb-2 mb-lg-0">
                        <li class="nav-item">
                            <a class="nav-link active" aria-current="page" href="#">Home</a>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link" href="#">Link</a>
                        </li>
                        <li class="nav-item dropdown">
                        <a class="nav-link dropdown-toggle" href="#" role="button" data-bs-toggle="dropdown" aria-expanded="false">
                            Dropdown
                        </a>
                        <ul class="dropdown-menu">
                            <li><a class="dropdown-item" href="#">Action</a></li>
                            <li><a class="dropdown-item" href="#">Another action</a></li>
                            <li><hr class="dropdown-divider"></li>
                            <li><a class="dropdown-item" href="#">Something else here</a></li>
                        </ul>
                        </li>
                        <li class="nav-item">
                            <a class="nav-link disabled" aria-disabled="true">Disabled</a>
                        </li>
                    </ul>
                    <form class="d-flex" role="search">
                        <input class="form-control me-2" type="search" placeholder="Search" aria-label="Search"/>
                        <button class="btn btn-outline-success" type="submit">Search</button>
                    </form>
                </div>
            </div>
        </nav>
    `,
})
export class HeaderComponent {
    private store = inject(Store);
    sidebarHidden = false;
    showMenu = false;

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

    toggleMenu(): void {
        this.showMenu = !this.showMenu;
    }
}
