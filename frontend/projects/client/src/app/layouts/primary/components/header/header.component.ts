import {Component, inject} from '@angular/core';
import {MatIconModule} from '@angular/material/icon';
import {MatToolbarModule} from '@angular/material/toolbar';
import {MatFormFieldModule} from '@angular/material/form-field';
import {MatIconButton} from '@angular/material/button';
import {UiService} from '@client/shared/ui';
import {RouterLink} from '@angular/router';
import {CdkMenuModule} from '@angular/cdk/menu';

@Component({
    selector: "ts-header",
    template: `
        <mat-toolbar class="items-center border-b border-slate-200 relative z-100">
            <div class="flex gap-4 items-center">
                <a [routerLink]="['/']" class="flex gap-2 items-center">
                    <img src="/assets/img/logo-icon.svg" alt="" class="h-8">
                    <h3><span class="font-bold">Flowreon</span></h3>
                </a>
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

            <button matIconButton [cdkMenuTriggerFor]="header_menu" class="example-standalone-trigger">
                <img src="assets/img/1.png" class="rounded-full" alt="">
            </button>

            <ng-template #header_menu>
                <div class="bg-white flex flex-col gap-2 border p-2 rounded" cdkMenu>
                    <button cdkMenuItem>Refresh</button>
                    <button cdkMenuItem>Settings</button>
                    <button cdkMenuItem>Help</button>
                    <button cdkMenuItem>Sign out</button>
                </div>
            </ng-template>

        </mat-toolbar>
    `,
    imports: [
        MatIconModule,
        MatToolbarModule,
        MatFormFieldModule,
        MatIconButton,
        RouterLink,
        CdkMenuModule,
    ],
})
export class HeaderComponent {
    uiService = inject(UiService);
    isOpen = false;
    sidebarToggle(): void {
        this.uiService.toggleSidebar();
    }
}
