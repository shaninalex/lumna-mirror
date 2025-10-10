import {Route} from '@angular/router';
import {
    BoardViewPageComponent,
    ProjectSettingsPageComponent, TaskDetailPageComponent
} from '@client/pages/primary/project-detail-page/pages';
import {ProjectDetailPageComponent} from '@client/pages/primary/project-detail-page/project-detail-page.component';
import {projectResolver} from './project.resolver';
import {taskResolver} from '@client/pages/primary/project-detail-page/task.resolver';

export const projectDetailRoutes: Route = {
    path: ":projectKey",
    component: ProjectDetailPageComponent,
    resolve: {project: projectResolver},
    data: {
        breadcrumb: (data: any, params: any) => {
            return data.project?.title ?? params['projectKey'] ?? 'Project'
        }
    },
    children: [
        {
            path: "",
            component: BoardViewPageComponent,
            data: { breadcrumb: 'Board' },
        },
        {
            path: "settings",
            component: ProjectSettingsPageComponent,
            data: { breadcrumb: "Settings"},
        },
        {
            path: ":taskCode",
            component: TaskDetailPageComponent,
            resolve: {task: taskResolver},
            data: {
                breadcrumb: (data: any, params: any) => {
                    return data.task?.title ?? params['taskCode'] ?? 'Task'
                }
            },
        },
    ]
}
