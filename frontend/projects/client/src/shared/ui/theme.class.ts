

export class ThemeClass {
    public appTheme(): string {
        return localStorage.getItem('theme') ?? 'light';
    }
}
