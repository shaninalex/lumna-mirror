import { type } from '@ngrx/signals';
import { ProjectModel, ProjectPayload } from './project.model';
import { eventGroup } from '@ngrx/signals/events';
import { ProjectEditModel } from '@features/project-edit';

export const projectEvents = eventGroup({
    source: 'Project',
    events: {
        getProjects: type<void>(),
        setProjects: type<ProjectModel[]>(),

        createProject: type<ProjectPayload>(),
        setProject: type<ProjectModel>(),

        deleteProject: type<number>(),
        _deleteProjectSuccess: type<number>(),

        patch: type<{id: number, data: ProjectEditModel}>(),

        updateProject: type<ProjectModel>(),

        failed: type<any>(), // TODO: proper error typing
    },
});
