import { Component } from '@angular/core';
import { MainLayout } from '@core/layouts/main/mainLayout';
import { RouterOutlet } from '@angular/router';

@Component({
    selector: 'app-projects',
    imports: [MainLayout, RouterOutlet],
    template: `
        <app-main-layout pageTitle="Projects">
            <router-outlet />
        </app-main-layout>
    `,
    styleUrl: './projects.css',
})
export class Projects {}
