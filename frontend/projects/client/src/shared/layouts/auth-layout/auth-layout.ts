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
                <div class="min-h-64 border border-slate-400 rounded p-4 w-xs">
                    <div class="text-center text-lg font-bold mb-4">{{ title }}</div>
                    <ng-content></ng-content>
                </div>
                <div class="text-xs text-slate-400 text-end">v{{ version }}</div>
            </div>
        </div>
    `
})
export class AuthLayout {
    @Input() title: string = "";
    version: string = version;
}
