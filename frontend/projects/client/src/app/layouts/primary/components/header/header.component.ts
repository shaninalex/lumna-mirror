import {Component, inject} from '@angular/core';
import {MatIconModule} from '@angular/material/icon';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatIconButton} from '@angular/material/button';
import {UiService} from '@client/shared/ui';
import {MatMenu, MatMenuItem, MatMenuTrigger} from '@angular/material/menu';

@Component({
    selector: "ts-header",
    imports: [
        MatIconModule,
        MatToolbarModule,
        MatFormFieldModule,
        MatIconButton,
        MatMenuTrigger,
        MatMenu,
        MatMenuItem,
    ],
    template: `
        <mat-toolbar class="items-center shadow">
            <div class="flex gap-4 items-center">
                <div class="flex gap-2 items-center">
                    <img src="/assets/img/logo-simple.png" alt="" class="h-8">
                    <h3><span class="font-bold">Taskiro</span></h3>
                </div>
                <button matIconButton (click)="sidebarToggle()">
                    <mat-icon>menu</mat-icon>
                </button>
            </div>
            <span class="flex-1"></span>
            <div class="text-sm">
                center text
            </div>
            <span class="flex-1"></span>
            <button matIconButton>
                <mat-icon>notifications</mat-icon>
            </button>
            <button matIconButton [matMenuTriggerFor]="beforeMenu">
                <img src="assets/img/1.png" class="rounded-full" alt="">
            </button>
            <mat-menu #beforeMenu="matMenu" xPosition="before">
                <button mat-menu-item>Item 1</button>
                <button mat-menu-item>Item 2</button>
            </mat-menu>
        </mat-toolbar>
    `
})
export class HeaderComponent {
    uiService = inject(UiService);

    sidebarToggle(): void {
        this.uiService.toggleSidebar();
    }
}
