import { Component } from '@angular/core';
import {ThemeSwitcher, UserMenuComponent, BreadCrumbs} from './components';

@Component({
    selector: 'app-header',
    imports: [ThemeSwitcher, BreadCrumbs, UserMenuComponent],
    templateUrl: './header.html',
    styleUrl: './header.css',
})
export class Header {
}
