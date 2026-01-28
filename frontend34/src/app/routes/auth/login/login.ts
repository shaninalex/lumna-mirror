import {Component} from '@angular/core';

@Component({
    selector: 'app-login',
    imports: [],
    template: `
        <p>
            login works!
        </p>
    `,
})
export class Login {
    // 1
    // after form submission we get login_token that give us ability to authorize in the system
    //
    // 2.01 - build PKCE codes
    // 2.02 - generate state code
    // 2
    // build /oauth/authorize url to get a code
    //
    // 3
    // make POST /oauth/token request to get access token (and refresh in cookies)
}
