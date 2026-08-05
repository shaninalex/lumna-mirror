import { Component } from '@angular/core';
import { SidebarComponent, HeaderComponent } from './components';

@Component({
    selector: 'lu-main-layout',
    imports: [SidebarComponent, HeaderComponent],
    styleUrl: './main.layout.css',
    template: `
        <div class="dashboard">
            <div class="dashboard-header">
                <lu-header />
            </div>
            <div class="dashboard-sidebar">
                <lu-sidebar />
            </div>
            <div class="dashboard-content">
                <ng-content />
            </div>
        </div>
    `,
})
export class MainLayout {}
