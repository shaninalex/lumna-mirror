import {Routes} from '@angular/router';

import {BoardPage} from './board-page/board-page';
import {BoardEditPage} from './board-edit-page/board-edit-page';

import {projectResolver} from '../resolvers/project.resolver';
import {boardResolver} from '../resolvers/board.resolver';

export const boardRoutes : Routes = [
    {
        path: ":id/board/:boardId",
        component: BoardPage,
        resolve: {
            project: projectResolver,
            board: boardResolver,
        }
    },
    {
        path: ":id/board/:boardId/edit",
        component: BoardEditPage,
        resolve: {
            project: projectResolver,
            board: boardResolver,
        }
    }
]
