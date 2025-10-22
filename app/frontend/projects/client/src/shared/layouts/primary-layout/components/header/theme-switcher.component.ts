import {Component, inject} from '@angular/core';
import {OverlayModule} from '@angular/cdk/overlay';
import {AsyncPipe, NgClass} from '@angular/common';
import {Observable} from 'rxjs';
import {tTheme} from '@client/shared/ui/theme.class';
import {UiService} from '@client/shared/ui/ui.service';

@Component({
    selector: 'lu-theme-switcher',
    imports: [
        OverlayModule,
        NgClass,
        AsyncPipe,
    ],
    template: `
        @if (theme$ | async; as theme) {
            <button tabindex="0" role="button" (click)="isOpen = !isOpen"
                    cdkOverlayOrigin #trigger="cdkOverlayOrigin">
                <i class="i-sun" [ngClass]="{
                    'i-sun': theme === 'dark',
                    'i-moon': theme === 'light',
                }"></i>
            </button>
            <ng-template
                cdkConnectedOverlay
                [cdkConnectedOverlayOrigin]="trigger"
                [cdkConnectedOverlayOpen]="isOpen"
                [cdkConnectedOverlayHasBackdrop]="true"
                (backdropClick)="isOpen = false"
            >
                <div class="dropdown">
                    <button type="button" class="block cursor-pointer" (click)="setTheme('light')">light</button>
                    <button type="button" class="block cursor-pointer" (click)="setTheme('dark')">dark</button>
                </div>
            </ng-template>
        }
    `
})
export class ThemeSwitcherComponent {
    private ui: UiService = inject(UiService);

    isOpen = false;
    theme$: Observable<tTheme> = this.ui.theme.get()

    setTheme(v: tTheme) {
        this.ui.theme.set(v)
    }
}
