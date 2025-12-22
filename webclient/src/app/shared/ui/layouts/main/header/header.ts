import {Component, inject, Input} from '@angular/core';
import {UserStore} from '@entities/user';

@Component({
    selector: 'app-header',
    imports: [],
    template: `
        <header class="header">
            <div class="font-medium">{{ title }}</div>
            @if(isAuthenticated()) {
                <div class="flex items-center gap-3">
                    <span class="text-sm">@BeeBeeKing17</span>
                    <img src="img/7.png" alt="" class="rounded-full size-8">
                </div>
            }
        </header>
    `,
    styleUrl: './header.css',
})
export class Header {
    @Input() title: string
    readonly userStore = inject(UserStore);
    readonly user = this.userStore.user;
    readonly isAuthenticated = this.userStore.isAuthenticated;
}
