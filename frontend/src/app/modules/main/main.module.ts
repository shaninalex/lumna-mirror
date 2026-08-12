import { NgModule } from '@angular/core';
import { CommonModule } from '@angular/common';
import { RouterModule } from '@angular/router';
import { routes } from './main.routes';
import { StoreModule } from '@ngrx/store';
import { EffectsModule } from '@ngrx/effects';
import { mainEffects, mainReducers } from './store';
import { WorkspaceApi } from '@entities/workspace';



@NgModule({
    declarations: [],
    imports: [
        CommonModule,
        
        StoreModule.forFeature('main', mainReducers),
        EffectsModule.forFeature(mainEffects),

        RouterModule.forChild(routes),
    ],
    providers: [
        WorkspaceApi,
    ]
})
export class MainModule {}
