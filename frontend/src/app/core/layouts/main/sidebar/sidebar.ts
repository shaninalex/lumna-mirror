import {Component} from '@angular/core';
import {RouterLink} from '@angular/router';
import packageJson from '@root/package.json';


@Component({
    selector: 'app-sidebar',
    imports: [
        RouterLink
    ],
    templateUrl: './sidebar.html',
    styleUrl: './sidebar.css',
})
export class Sidebar {
    varsion: string = packageJson.version;
}
