import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';

@Component({
    selector: 'auth-container',
    imports: [RouterOutlet],
    template: `<router-outlet />`,
})
export class AuthContainer {}
