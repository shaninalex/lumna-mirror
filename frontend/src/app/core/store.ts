import { projectsReducer } from '@entities/project';
import * as projectEffects from '@entities/project/model/project.effects';

export const effects = [projectEffects];

export const reducers = {
    project: projectsReducer,
};
