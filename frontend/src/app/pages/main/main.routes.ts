import { Routes } from '@angular/router';
import { MainRoot } from './main.root';
import { SelectWorkspacePage } from './select-workspace';

export const routes: Routes = [
    {
        path: '',
        component: MainRoot,
        children: [
            {
                path: 'select-workspace',
                component: SelectWorkspacePage,
            }
        ]
    }
];
