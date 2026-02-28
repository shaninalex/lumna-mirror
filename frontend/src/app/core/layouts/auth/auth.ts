import {Component, Input} from '@angular/core';

@Component({
    selector: 'app-auth-layout',
    imports: [],
    template: `
        <section class="hero is-fullheight is-bold">
            <div class="hero-body">
                <div class="container">
                    <div class="columns is-vcentered">
                        <div class="column is-4 is-offset-4">
                            @if (hasLogo) {
                                <figure class="image is-64x64 mb-4">
                                    <img src="img/logo-icon.svg" />
                                </figure>
                            }
                            <ng-content/>
                        </div>
                    </div>
                </div>
            </div>
        </section>
    `,
})
export class AuthLayout {
    @Input() hasLogo: boolean;
}
