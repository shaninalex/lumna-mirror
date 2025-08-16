import {Routes} from '@angular/router';
import {PageHomeComponent} from '@client/pages/main/home.component';
import {HomeWrapperComponent} from '@client/pages/main/home-wrapper.component';

export const mainRoutes: Routes = [
    {
        path: '',
        component: HomeWrapperComponent,
        children: [
            {
                path: '',
                component: PageHomeComponent,
            },
        ],
    },
];
