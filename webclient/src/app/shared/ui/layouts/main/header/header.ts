import {Component, inject, Input} from '@angular/core';
import {UserStore} from '@entities/user';

@Component({
    selector: 'app-header',
    imports: [],
    templateUrl: './header.html',
    styleUrl: './header.css',
})
export class Header {
    @Input() title: string
    readonly userStore = inject(UserStore);
    readonly user = this.userStore.user;
    readonly isAuthenticated = this.userStore.isAuthenticated;
}
