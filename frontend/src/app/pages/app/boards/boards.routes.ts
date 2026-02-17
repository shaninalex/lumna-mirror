import { Routes } from '@angular/router';

import { BoardPage } from './board-page/board-page';
import { BoardEditPage } from './board-edit-page/board-edit-page';

export const boardRoutes: Routes = [
    {
        path: 'board/:id',
        component: BoardPage,
    },
    {
        path: 'board/:id/edit',
        component: BoardEditPage,
    },
];
