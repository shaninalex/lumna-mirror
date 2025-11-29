import { Route } from '@angular/router'
import { ProjectSettingsPageComponent, TaskDetailPageComponent } from '@client/pages/primary/project-detail-page/pages'
import { ProjectDetailWrapperComponent } from '@client/pages/primary/project-detail-page/project-detail-wrapper.component'
import { projectResolver } from './project.resolver'
import { taskResolver } from '@client/pages/primary/project-detail-page/task.resolver'
import { ProjectOverviewComponent } from './project-overview.component'
import { ViewModeWrapperComponent } from '@client/pages/primary/project-detail-page/view-mode-wrapper.component'

export const projectDetailRoutes: Route = {
    path: ':projectKey',
    component: ProjectDetailWrapperComponent,
    resolve: { project: projectResolver },
    data: {
        breadcrumb: (data: any, params: any) => {
            return data.project?.title ?? params['projectKey'] ?? 'Project'
        },
    },
    children: [
        {
            path: '',
            component: ProjectOverviewComponent,
            data: { breadcrumb: 'Overview' },
        },
        {
            path: 'settings',
            component: ProjectSettingsPageComponent,
            data: { breadcrumb: 'Settings' },
        },
        {
            path: ':viewMode',
            component: ViewModeWrapperComponent,
            data: { breadcrumb: (data: any, params: any) => params['viewMode'] },
            children: [
                {
                    path: ':taskCode',
                    component: TaskDetailPageComponent,
                    resolve: { task: taskResolver },
                    data: {
                        breadcrumb: (data: any, params: any) => {
                            return data.task?.title ?? params['taskCode'] ?? 'Task'
                        },
                    },
                },
            ],
        },
    ],
}
