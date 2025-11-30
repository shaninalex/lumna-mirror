import { BehaviorSubject, Observable } from 'rxjs'

export type tTheme = 'light' | 'dark'

export class ThemeClass {
    private _theme$: BehaviorSubject<tTheme> = new BehaviorSubject<tTheme>('light')

    constructor() {
        const lValue = localStorage.getItem('theme')
        const theme: tTheme = lValue ? (lValue as tTheme) : 'light'
        this._theme$.next(theme)
        this.set(theme)
    }

    public get(): Observable<tTheme> {
        return this._theme$.asObservable()
    }

    public set(v: tTheme): void {
        const html = document.documentElement
        html.removeAttribute('data-theme')
        if (v === 'light') {
            html.setAttribute('data-theme', 'light')
            this._theme$.next(v)
        } else if (v === 'dark') {
            html.setAttribute('data-theme', 'dark')
            this._theme$.next(v)
        } else {
            const isDark = window.matchMedia('(prefers-color-scheme: light)').matches
            const _v = isDark ? 'dark' : 'light'
            html.setAttribute('data-theme', _v)
            this._theme$.next(_v)
        }
        localStorage.setItem('theme', v)
    }

    public appTheme(): string {
        return localStorage.getItem('theme') ?? 'light'
    }
}
