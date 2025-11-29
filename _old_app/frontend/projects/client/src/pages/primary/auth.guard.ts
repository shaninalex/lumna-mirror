import { inject } from "@angular/core"
import { CanMatchFn, Router } from "@angular/router"
import { Store } from "@ngrx/store"
import { AppState } from "@client/shared/store"
import { selectUserState, UserGetAction } from "@client/entities/user"
import { filter, map, take, tap } from "rxjs"

export const authGuard: CanMatchFn = () => {
	const store = inject(Store<AppState>)
	const router = inject(Router)

	return store.select(selectUserState).pipe(
		tap(state => {
			if (!state.loaded) {
				store.dispatch(UserGetAction())
			}
		}),
		filter(state => state.loaded),
		take(1),
		map(state => {
			return state.user ? true : router.createUrlTree(["/auth/login"])
		})
	)
}
