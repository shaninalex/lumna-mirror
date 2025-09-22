import {Component, OnInit} from '@angular/core';

@Component({
    selector: 'fr-theme-switcher',
    imports: [],
    template: `
        <select name="theme" id="theme" [value]="theme" (change)="onThemeChange($event)">
            <option value="device">device</option>
            <option value="light">light</option>
            <option value="dark">dark</option>
        </select>
    `
})
export class ThemeSwitcherComponent implements OnInit {
    theme: string = 'device';

    ngOnInit() {
        const saved = localStorage.getItem('theme');
        this.theme = saved ? saved : 'device';
        this.applyTheme(this.theme);
    }

    onThemeChange(event: Event) {
        const select = event.target as HTMLSelectElement;
        this.theme = select.value;
        localStorage.setItem('theme', this.theme);
        this.applyTheme(this.theme);
    }

    private applyTheme(theme: string) {
        const html = document.documentElement;
        html.removeAttribute('data-theme');
        if (theme === 'light') {
            html.setAttribute('data-theme', 'retro');
        } else if (theme === 'dark') {
            html.setAttribute('data-theme', 'halloween');
        }
    }
}
