import { type } from '@ngrx/signals';
import { ProjectModel } from './project.model';
import { eventGroup } from '@ngrx/signals/events';

export const projectEvents = eventGroup({
    source: 'Project',
    events: {
        getProjects: type<void>(),
        setProjects: type<ProjectModel[]>(),
    },
});
