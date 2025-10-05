import {Injectable} from '@angular/core';
import {BehaviorSubject, Observable} from 'rxjs';

@Injectable({ providedIn: 'root' })
export class UiService {
    private extendSidebar$: BehaviorSubject<boolean> = new BehaviorSubject(false);

    public extendSidebar(): Observable<boolean> {
        return this.extendSidebar$.asObservable();
    }

    public setExtendSidebar(v: boolean): void {
        this.extendSidebar$.next(v);
    }
}
