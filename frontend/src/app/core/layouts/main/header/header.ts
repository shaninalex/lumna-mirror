import { Component } from '@angular/core';
import {ThemeSwitcher, UserMenuComponent, BreadcrumbsWidget} from './components';

@Component({
    selector: 'app-header',
    imports: [ThemeSwitcher, BreadcrumbsWidget, UserMenuComponent],
    templateUrl: './header.html',
    styleUrl: './header.css',
})
export class Header {
}
