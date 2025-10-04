import {Component, Input} from '@angular/core';
import {version} from '@root/package.json';
import {MatCard, MatCardContent, MatCardHeader, MatCardTitle} from '@angular/material/card';

@Component({
    selector: 'fr-auth-layout',
    imports: [
        MatCard,
        MatCardHeader,
        MatCardTitle,
        MatCardContent
    ],
    template: `
        <div class="flex items-center justify-center h-screen">
            <div>
                <img src="img/logo.svg" alt="Lumna" class="mb-8 w-48 mx-auto"/>
                <mat-card appearance="outlined">
                    <mat-card-header>
                        <mat-card-title>{{ title }}</mat-card-title>
                    </mat-card-header>
                    <mat-card-content>
                        <ng-content></ng-content>
                    </mat-card-content>
                </mat-card>
                <div class="text-xs text-base-300 text-end">v{{ version }}</div>
            </div>
        </div>
    `
})
export class AuthLayout {
    @Input() title: string = "";
    version: string = version;
}
