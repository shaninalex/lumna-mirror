import { BehaviorSubject, Observable } from 'rxjs'

export class SidebarClass {
    private _extend$: BehaviorSubject<boolean> = new BehaviorSubject(false)

    public state$(): Observable<boolean> {
        return this._extend$.asObservable()
    }

    public setState(v: boolean): void {
        this._extend$.next(v)
    }
}
