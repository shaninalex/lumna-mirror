import { inject, Injectable } from '@angular/core';
import { BehaviorSubject, Observable, tap } from 'rxjs';
import { Title } from '@angular/platform-browser';
import { Theme, ThemeManager } from './theme.manager';

@Injectable({
    providedIn: 'root',
})
export class UiService {
    pageTitle: BehaviorSubject<string> = new BehaviorSubject('');
    theme: BehaviorSubject<Theme> = new BehaviorSubject<Theme>('auto');
    private themeManager = new ThemeManager(this.theme);

    titleService = inject(Title);

    constructor() {
        this.pageTitle.pipe(tap((value) => this.titleService.setTitle(value))).subscribe();
        this.themeManager.init();
    }

    /** Page title API */
    public setPageTitle(t: string) {
        this.pageTitle.next(t);
    }

    public getPageTitle(): string {
        return this.pageTitle.value;
    }

    /** Theme API */
    public toggleTheme() {
        this.themeManager.toggleTheme();
    }

    public setTheme(theme: Theme) {
        this.themeManager.applyTheme(theme);
    }

    public getTheme(): Observable<Theme> {
        return this.theme;
    }
}
