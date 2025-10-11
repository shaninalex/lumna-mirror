import {Component, OnInit} from '@angular/core';
import {OverlayModule} from '@angular/cdk/overlay';
import {NgClass} from '@angular/common';

@Component({
    selector: 'lu-theme-switcher',
    imports: [
        OverlayModule,
        NgClass,
    ],
    template: `
        <button tabindex="0" role="button"
                (click)="isOpen = !isOpen"
                cdkOverlayOrigin #trigger="cdkOverlayOrigin">
            <i class="i-sun" [ngClass]="{
                'i-sun': isDark,
                'i-moon': !isDark,
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
    `
})
export class ThemeSwitcherComponent implements OnInit {
    theme: string = 'device';
    isDark: boolean = false;
    isOpen = false;

    ngOnInit() {
        const saved = localStorage.getItem('theme');
        this.theme = saved ? saved : 'device';
        this.applyTheme(this.theme);
    }

    setTheme(theme: string) {
        this.theme = theme;
        localStorage.setItem('theme', this.theme);
        this.applyTheme(this.theme);
    }

    private applyTheme(theme: string) {
        const html = document.documentElement;
        html.removeAttribute('data-theme');
        if (theme === 'light') {
            html.setAttribute('data-theme', 'light');
            this.isDark = false;
        } else if (theme === 'dark') {
            html.setAttribute('data-theme', 'dark');
            this.isDark = true;
        } else {
            const isDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
            html.setAttribute('data-theme', isDark ? 'dark' : 'light');
            this.isDark = isDark;
        }
    }
}
