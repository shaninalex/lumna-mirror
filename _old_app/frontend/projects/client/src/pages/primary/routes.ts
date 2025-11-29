import { Routes } from "@angular/router"
import { PrimaryRoot } from "./primary-root"
import { OverviewPage } from "./overview"
import { overviewResolver } from "./overview/overview.resolver"
import { ProjectsListPageComponent } from "@client/pages/primary/projects-list-page/projects-list-page.component"
import { projectListResolver } from "@client/pages/primary/projects-list-page/projects-list.resolver"
import { projectDetailRoutes } from "@client/pages/primary/project-detail-page"
import { ProjectsRootPageComponent } from "@client/pages/primary/projects-root-page"
import { SettingsPageComponent } from "@client/pages/primary/settings-page"
import { authGuard } from "./auth.guard"

export const mainRoutes: Routes = [
	{
		path: "",
		component: PrimaryRoot,
		canMatch: [authGuard],
		children: [
			{
				path: "",
				component: OverviewPage,
				resolve: { overview: overviewResolver },
				data: { breadcrumb: "Home" },
			},
			{
				path: "projects",
				component: ProjectsRootPageComponent,
				resolve: { projects: projectListResolver },
				data: { breadcrumb: "Projects" },
				children: [
					{
						path: "",
						component: ProjectsListPageComponent,
					},
					projectDetailRoutes,
				],
			},
			{
				path: "settings",
				component: SettingsPageComponent,
				data: { breadcrumb: "Settings" },
			},
		],
	},
]
