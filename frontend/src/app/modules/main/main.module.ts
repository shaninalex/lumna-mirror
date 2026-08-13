import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { routes } from './main.routes';
import { provideState } from '@ngrx/store';
import { provideEffects } from '@ngrx/effects';
import { mainEffects } from './store';
import { WorkspaceApi } from '@entities/workspace';
import { ProjectApi, projectFeature } from '@entities/project';
import { UserApi } from '@entities/user';
import { workspaceFeature } from '@entities/workspace';
import { userFeature } from '@entities/user';
import { LocalStorageService } from '@shared/services';


@NgModule({
    declarations: [],
    imports: [
        CommonModule,
        RouterModule.forChild(routes),
    ],
    providers: [
        WorkspaceApi,
        ProjectApi,
        UserApi,
        LocalStorageService,

        provideEffects(mainEffects),
        provideState(workspaceFeature),
        provideState(projectFeature),
        provideState(userFeature),
    ]
})
export class MainModule {}
