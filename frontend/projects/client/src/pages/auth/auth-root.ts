import {Component} from '@angular/core';
import {RouterOutlet} from '@angular/router';

@Component({
    selector: 'fr-auth-root',
    imports: [RouterOutlet],
    template: `
        <router-outlet/>`
})
export class AuthRoot {
}
