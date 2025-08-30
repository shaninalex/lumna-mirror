import { Component, Input } from '@angular/core';
import { MatCard } from '@angular/material/card';

@Component({
    selector: 'auth-layout',
    imports: [MatCard],
    template: `
        <div class="h-screen flex items-center justify-center bg-base-100">
            <div class="w-92">
                <img src="/assets/img/logo.svg" class="w-2xs mx-auto mb-8" />
                <mat-card appearance="outlined" class="p-4">
                    <h1 class="font-bold text-2xl mb-4 text-center">{{ title }}</h1>
                    <ng-content></ng-content>
                </mat-card>
            </div>
        </div>
    `,
})
export class AuthLayoutComponent {
    @Input() title: string;
}
