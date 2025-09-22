import {Component, Input} from '@angular/core';
import { Logo } from './logo'
import {version} from '@root/package.json';

@Component({
    selector: 'fr-auth-layout',
    imports: [
        Logo
    ],
    template: `
        <div class="flex items-center justify-center h-screen">
            <div>
                <fr-logo />
                <div class="card border border-base-300 mb-2 bg-base-100">
                    <div class="card-body">
                        <div class="text-center text-lg font-bold mb-4">{{ title }}</div>
                        <ng-content></ng-content>
                    </div>
                </div>
                <div class="text-xs text-base-300 text-end">v{{ version }}</div>
            </div>
        </div>
    `
})
export class AuthLayout {
    @Input() title: string = "";
    version: string = version;
}
