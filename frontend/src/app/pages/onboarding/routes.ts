import {Routes} from '@angular/router';
import { OnboardingContainer } from '@pages/onboarding/container';
import { WorkspaceOnboardingPage } from './workspace';
import { TeamOnboardingPage } from './team';

export const routes: Routes = [
    {
        path: 'onboarding',
        component: OnboardingContainer,
        children: [
            {
                path: '',
                redirectTo: 'workspace',
                pathMatch: 'full',
            },
            {
                path: 'workspace',
                component: WorkspaceOnboardingPage,
            },
            {
                path: 'team',
                component: TeamOnboardingPage,
            }
        ],
    }
];
