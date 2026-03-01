import { Injectable } from '@angular/core';

@Injectable({
    providedIn: 'root',
})
export class CoreService {
    constructor() {
        console.log('CoreService initialized');
    }
}
