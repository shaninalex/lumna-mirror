import { Component, inject, Input, OnInit } from '@angular/core';
import { Header } from '@core/layouts/main/header/header';
import { Sidebar } from '@core/layouts/main/sidebar/sidebar';
import { UiService } from '@shared/ui';

@Component({
    selector: 'app-main-layout',
    imports: [Header, Sidebar],
    templateUrl: './main.html',
    styleUrl: './main.css',
})
export class MainLayout implements OnInit {
    @Input() pageTitle: string;

    private ui = inject(UiService);

    ngOnInit(): void {
        this.ui.setPageTitle(this.pageTitle);
    }
}
