import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { routes } from './main.routes';
import { provideState } from '@ngrx/store';
import { provideEffects } from '@ngrx/effects';
import { mainEffects } from './store';
import { WorkspaceApi, workspaceFeature  } from '@entities/workspace';
import { ProjectApi, projectFeature } from '@entities/project';
import { UserApi, userFeature  } from '@entities/user';
import { TaskApi, taskFeature } from '@entities/task';
import { AppRoutes } from '@core/routes';



@NgModule({
    declarations: [],
    imports: [
        CommonModule,
        RouterModule.forChild(routes),
    ],
    providers: [
        AppRoutes,
        WorkspaceApi,
        ProjectApi,
        UserApi,
        TaskApi,

        provideEffects(mainEffects),
        provideState(workspaceFeature),
        provideState(projectFeature),
        provideState(userFeature),
        provideState(taskFeature),
    ]
})
export class MainModule {}
