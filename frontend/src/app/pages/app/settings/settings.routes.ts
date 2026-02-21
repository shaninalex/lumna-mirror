import { Routes } from '@angular/router';
import { SettingsContainer } from './settings-container';
import { SettingsPage } from './settings-page/settings-page';

export const settingsRoutes: Routes = [
    {
        path: 'settings',
        component: SettingsContainer,
        children: [
            {
                path: '',
                component: SettingsPage,
                data: { title: 'Settings page' }
            },
        ],
    },
];
