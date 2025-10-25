import { Routes } from "@angular/router"
import { SessionExpiredPageComponent } from "@client/pages/system/session-expired-page"
import { Page403 } from "@client/pages/system/page-403-page"
import { Page404 } from "@client/pages/system/page-404-page"

export const systemRoutes: Routes = [
	{
		path: "session-expired",
		component: SessionExpiredPageComponent,
	},
	{
		path: "403",
		component: Page403,
	},
	{
		path: "404",
		component: Page404,
	},
	// {
	//     path: "**",
	//     redirectTo: "/404",
	// },
]
