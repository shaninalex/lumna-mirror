import { BehaviorSubject } from 'rxjs';

export type Theme = 'light' | 'dark' | 'auto';

export class ThemeManager {
    private readonly storageKey = 'theme';

    constructor(private theme$: BehaviorSubject<Theme>) {}

    public init() {
        const saved = this.getStoredTheme();
        const theme: Theme = saved ?? 'auto';
        this.applyTheme(theme);
    }

    private getStoredTheme(): Theme | null {
        const stored = localStorage.getItem(this.storageKey);
        if (stored === 'light' || stored === 'dark' || stored === 'auto') {
            return stored;
        }
        return null;
    }

    private getSystemTheme(): 'light' | 'dark' {
        return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
    }

    public applyTheme(theme: Theme) {
        this.theme$.next(theme);
        const resolvedTheme = theme === 'auto' ? this.getSystemTheme() : theme;
        document.documentElement.setAttribute('data-bs-theme', resolvedTheme);
        localStorage.setItem(this.storageKey, theme);
    }

    public toggleTheme() {
        const next: Theme =
            this.theme$.value === 'light'
                ? 'dark'
                : this.theme$.value === 'dark'
                  ? 'auto'
                  : 'light';
        this.applyTheme(next);
    }
}
