import { Routes } from '@angular/router';

import { BoardPage } from './board-page/board-page';
import { BoardEditPage } from './board-edit-page/board-edit-page';
import {boardResolver} from '@pages/app/boards/resolver';

export const boardRoutes: Routes = [
    {
        path: 'board/:id',
        component: BoardPage,
        resolve: {
            board: boardResolver
        }
    },
    {
        path: 'board/:id/edit',
        component: BoardEditPage,
    },
];
