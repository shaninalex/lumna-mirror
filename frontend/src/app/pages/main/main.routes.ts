import { Routes } from '@angular/router';
import { InboxPage } from './inbox';
import { MyWorkPage } from './my-work';


export const routes: Routes = [
    {
        path: 'inbox',
        component: InboxPage,
    },
    {
        path: 'my-work',
        component: MyWorkPage,
    },
    {
        path: '',
        pathMatch: 'full',
        redirectTo: '/inbox'
    }
];
