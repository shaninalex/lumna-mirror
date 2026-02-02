import { Component } from '@angular/core';
import { RouterOutlet } from '@angular/router';
import { MainLayout } from '@core';
@Component({
    selector: 'app-container',
    imports: [RouterOutlet, MainLayout],
    template: `
        <app-main-layout>
            <router-outlet />
        </app-main-layout>
    `,
})
export class DashboardContainer {}
