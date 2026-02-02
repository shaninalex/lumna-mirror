import { BehaviorSubject } from 'rxjs';

export type Theme = 'light' | 'dark' | 'auto';

export class ThemeManager {
    private readonly storageKey = 'theme';

    constructor(private theme$: BehaviorSubject<Theme>) {}

    /** Initialize theme on app start */
    public init() {
        const saved = this.getStoredTheme();
        const theme: Theme = saved ?? 'auto';
        this.applyTheme(theme);
    }

    /** Get theme from localStorage, validate value */
    private getStoredTheme(): Theme | null {
        const stored = localStorage.getItem(this.storageKey);
        if (stored === 'light' || stored === 'dark' || stored === 'auto') {
            return stored;
        }
        return null;
    }

    /** Detect system dark mode */
    private getSystemTheme(): 'light' | 'dark' {
        return window.matchMedia('(prefers-color-scheme: dark)').matches ? 'dark' : 'light';
    }

    /**
     * Apply theme:
     * - if 'auto', resolve to system theme for html[data-theme]
     * - always save the actual value in BehaviorSubject and localStorage
     */
    public applyTheme(theme: Theme) {
        this.theme$.next(theme);
        const resolvedTheme = theme === 'auto' ? this.getSystemTheme() : theme;
        document.documentElement.setAttribute('data-theme', resolvedTheme);
        localStorage.setItem(this.storageKey, theme);
    }

    /** Toggle between light/dark/auto */
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
