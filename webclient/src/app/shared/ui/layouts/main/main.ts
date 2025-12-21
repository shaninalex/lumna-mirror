import {Component, Input} from '@angular/core';
import {Header} from '@shared/ui/layouts/main/header/header';
import {Sidebar} from '@shared/ui/layouts/main/sidebar/sidebar';

@Component({
    selector: 'app-main',
    imports: [
        Header,
        Sidebar
    ],
    templateUrl: './main.html',
    styleUrl: './main.css',
})
export class Main {
    @Input() title: string
}
