import { createFeature } from '@ngrx/store';
import { projectReducer } from './model';

export const projectFeature = createFeature({
    name: 'project',
    reducer: projectReducer,
});
