import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MainLayout } from '@core/layout';

@Component({
    selector: 'lu-main',
    imports: [MainLayout, RouterOutlet],
    template: `
        <lu-main-layout>
            <router-outlet />
        </lu-main-layout>
    `,
})
export class MainRoot {}
