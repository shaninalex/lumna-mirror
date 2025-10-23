import {Component, Input, OnInit} from '@angular/core';
import {version} from '@root/package.json';

@Component({
    selector: 'lu-auth-layout',
    imports: [],
    template: `
        <div class="flex items-center justify-center h-screen bg-slate-50 dark:bg-gray-700">
            <div>
                <img src="img/logo.svg" class="mb-8 w-48 mx-auto"/>
                <div class="card">
                    <div class="card-title">{{ title }}</div>
                    <ng-content></ng-content>
                </div>
                <div class="text-xs text-base-300 text-end">v{{ version }}</div>
            </div>
        </div>
    `
})
export class AuthLayout implements OnInit {
    @Input() title: string = "";
    version: string = version;

    ngOnInit() {
        const html = document.documentElement;
        const isDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
        const _v = isDark ? 'dark' : 'light'
        html.setAttribute('data-theme', _v);
    }
}
