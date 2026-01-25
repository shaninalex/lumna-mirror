import {Component, Input} from '@angular/core';
import {Header} from '@core/layouts/main/header/header';
import {Sidebar} from '@core/layouts/main/sidebar/sidebar';

@Component({
    selector: 'app-main-layout',
    imports: [
        Header,
        Sidebar
    ],
    templateUrl: './main.html',
    styleUrl: './main.css',
})
export class MainLayout {
    @Input() title: string
}
