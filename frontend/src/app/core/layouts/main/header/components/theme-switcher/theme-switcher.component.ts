import { Component, inject } from '@angular/core';
import { CdkMenu, CdkMenuItem, CdkMenuTrigger } from '@angular/cdk/menu';
import { Theme } from '@shared/ui/theme.manager';
import { UiService } from '@shared/ui';
import { Observable } from 'rxjs';
import { AsyncPipe } from '@angular/common';

@Component({
    selector: 'app-theme-switcher',
    imports: [CdkMenuTrigger, CdkMenu, CdkMenuItem, AsyncPipe],
    template: `
        <button class="btn btn-sm" [cdkMenuTriggerFor]="theme">
            @if (themeMode | async; as themeMode) {
                @switch (themeMode) {
                    @case ('dark') {
                        <i class="fa-solid fa-moon"></i>
                    }
                    @case ('light') {
                        <i class="fa-solid fa-sun"></i>
                    }
                    @case ('auto') {
                        <i class="fa-solid fa-circle-half-stroke"></i>
                    }
                }
            }
        </button>
        <ng-template #theme>
            <div class="dropdown-menu d-block" cdkMenu>
                <button (click)="changeTheme('light')" class="dropdown-item" cdkMenuItem>
                    Light
                </button>
                <button (click)="changeTheme('dark')" class="dropdown-item" cdkMenuItem>
                    Dark
                </button>
                <button (click)="changeTheme('auto')" class="dropdown-item" cdkMenuItem>
                    Auto
                </button>
            </div>
        </ng-template>
    `,
})
export class ThemeSwitcher {
    private ui = inject(UiService);
    themeMode: Observable<Theme> = this.ui.getTheme();

    changeTheme(v: Theme) {
        this.ui.setTheme(v);
    }
}
