import {Component, Input} from '@angular/core';
import {version} from '@root/package.json';

@Component({
    selector: 'fr-auth-layout',
    imports: [],
    template: `
        <div class="flex items-center justify-center h-screen bg-slate-100">
            <div>
                <img src="img/logo.svg" alt="Lumna" class="mb-8 w-48 mx-auto"/>
                <div class="card">
                    <div class="card-header">{{ title }}</div>
                    <ng-content></ng-content>
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
