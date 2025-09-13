import { Routes } from '@angular/router';
import {SessionExpiredComponent} from './session-expired/session-expired.component';
import {Page403} from '@client/pages/system/page-403/page-403';
import {Page404} from '@client/pages/system/page-404/page-404';
import {NewOrgComponent} from '@client/pages/system/org';

export const systemRoutes: Routes = [
    {
        path: "session-expired",
        component: SessionExpiredComponent,
    },
    {
        path: "set-organization",
        component: NewOrgComponent,
    },
    {
        path: "403",
        component: Page403,
    },
    {
        path: "404",
        component: Page404,
    },
    {
        path: "**",
        redirectTo: "/404",
    },
];
