import {Component, Input, OnChanges, SimpleChanges} from '@angular/core';
import {Header} from '@shared/ui/layouts/main/header/header';
import {Sidebar} from '@shared/ui/layouts/main/sidebar/sidebar';

@Component({
    selector: 'app-main-layout',
    imports: [
        Header,
        Sidebar
    ],
    templateUrl: './main.html',
    styleUrl: './main.css',
})
export class MainLayout implements OnChanges {
    @Input() title: string

    ngOnChanges(changes: SimpleChanges) {
        console.log(changes['title'])
    }
}
