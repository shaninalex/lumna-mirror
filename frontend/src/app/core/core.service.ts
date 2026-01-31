import { inject, Injectable } from '@angular/core';
import { Events, Dispatcher } from '@ngrx/signals/events';

@Injectable({
    providedIn: 'root',
})
export class CoreService {
    readonly events = inject(Events);
    readonly dispatcher = inject(Dispatcher);

    constructor() {
        console.log('CoreService initialized');
    }
}
