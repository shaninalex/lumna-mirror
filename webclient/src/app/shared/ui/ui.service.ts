import {inject, Injectable} from '@angular/core';
import {BehaviorSubject, tap} from 'rxjs';
import {Title} from '@angular/platform-browser';

@Injectable({
    providedIn: 'root'
})
export class UiService {
    pageTitle: BehaviorSubject<string> = new BehaviorSubject('')
    titleService = inject(Title);

    constructor() {
        this.pageTitle.pipe(
            tap(value => this.titleService.setTitle(value))
        ).subscribe()
    }

    public setPageTitle(t: string) {
        this.pageTitle.next(t)
    }

    public getPageTitle(): string {
        return this.pageTitle.value
    }
}
