import {Component, Input} from '@angular/core';
import {version} from '@root/package.json';

@Component({
    selector: 'fr-auth-layout',
    template: `
        <div class="flex items-center justify-center h-screen">
            <div>
                <img src="img/logo.svg" alt="Lumna" class="mb-8 w-48 mx-auto" />
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
