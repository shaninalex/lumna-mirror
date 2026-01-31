import { Routes } from '@angular/router';

import { BoardPage } from './board-page/board-page';
import { BoardEditPage } from './board-edit-page/board-edit-page';

export const boardRoutes: Routes = [
    {
        path: ':id/board/:boardId',
        component: BoardPage,
    },
    {
        path: ':id/board/:boardId/edit',
        component: BoardEditPage,
    },
];
