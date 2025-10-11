import {Injectable} from '@angular/core';
import {BehaviorSubject, Observable} from 'rxjs';
import {MessagesClass} from '@client/shared/ui/messages.class';
import {ThemeClass} from '@client/shared/ui/theme.class';

@Injectable({ providedIn: 'root' })
export class UiService extends MessagesClass {
    private extendSidebar$: BehaviorSubject<boolean> = new BehaviorSubject(false);

    public theme: ThemeClass

    constructor() {
        super();

        this.theme = new ThemeClass()
    }

    public extendSidebar(): Observable<boolean> {
        return this.extendSidebar$.asObservable();
    }

    public setExtendSidebar(v: boolean): void {
        this.extendSidebar$.next(v);
    }
}
