import { type } from '@ngrx/signals';
import {ProjectModel, ProjectPayload} from './project.model';
import { eventGroup } from '@ngrx/signals/events';

export const projectEvents = eventGroup({
    source: 'Project',
    events: {
        getProjects: type<void>(),
        setProjects: type<ProjectModel[]>(),

        createProject: type<ProjectPayload>(),
        setProject: type<ProjectModel>(),
        createProjectFailed: type<any>(), // TODO: proper error typing
    },
});
