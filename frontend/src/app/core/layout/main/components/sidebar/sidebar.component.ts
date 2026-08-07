import {Component, DestroyRef, inject} from '@angular/core';
import {Actions, ofType} from '@ngrx/effects';
import {actionToggleSidebar} from '@core';
import {takeUntilDestroyed} from '@angular/core/rxjs-interop';
import {NgClass} from '@angular/common';
import { RouterLink } from "@angular/router";

@Component({
    selector: 'lu-sidebar',
    imports: [NgClass, RouterLink],
    styleUrl: './sidebar.component.css',
    template: `
        <nav class="sidebar h-100 bg-body-tertiary" [ngClass]="{ 'sidebar-closed': hideSidebar }">
            <ul class="nav flex-column">
                <li class="nav-item">
                    <a class="nav-link text-body" routerLink="/inbox">
                        <i class="fa-solid fa-chart-column"></i>
                        <span class="sidebar-nav-link">Inbox</span>
                    </a>
                </li>
                <li class="nav-item">
                    <a class="nav-link text-body" routerLink="/boards">
                        <i class="fa-solid fa-chart-simple"></i>
                        <span class="sidebar-nav-link">Boards</span>
                    </a>
                </li>
                <li class="nav-item">
                    <a class="nav-link text-body" routerLink="/backlog">
                        <i class="fa-regular fa-file"></i>
                        <span class="sidebar-nav-link">Tasks</span>
                    </a>
                </li>
                <li class="nav-item">
                    <a class="nav-link text-body" href="#">
                        <i class="fa-regular fa-file"></i>
                        <span class="sidebar-nav-link">Files</span>
                    </a>
                </li>
                <li class="nav-item">
                    <a class="nav-link text-body" href="#">
                        <i class="fa-regular fa-envelope"></i>
                        <span class="sidebar-nav-link">Messages</span>
                    </a>
                </li>
                <li class="nav-item">
                    <a class="nav-link text-body" href="#">
                        <i class="fa-solid fa-bell"></i>
                        <span class="sidebar-nav-link">Notifications</span>
                    </a>
                </li>
            </ul>

            <ul class="nav flex-column">
                <li class="nav-item">
                    <a class="nav-link text-body" href="#">
                        <i class="fa-regular fa-circle-question"></i>
                        <span class="sidebar-nav-link">Help</span>
                    </a>
                </li>
            </ul>
        </nav>
    `,
})
export class SidebarComponent {
    hideSidebar = false;

    private actions$ = inject(Actions);
    private ref = inject(DestroyRef);

    constructor() {
        this.actions$
            .pipe(
                ofType(actionToggleSidebar),
                takeUntilDestroyed(this.ref),
            ).subscribe(() => this.hideSidebar = !this.hideSidebar);
    }
}
