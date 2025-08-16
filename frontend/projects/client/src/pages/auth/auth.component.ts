import { Component } from '@angular/core';
import {RouterOutlet} from '@angular/router';

@Component({
    selector: 'jr-auth',
    imports: [
        RouterOutlet
    ],
    template: `
        <div class="d-flex vh-100 align-items-center justify-content-center">
            <router-outlet/>
        </div>
    `
})
export class AuthComponent {}
