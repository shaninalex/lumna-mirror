import { inject, Injectable } from "@angular/core"
import { HttpErrorResponse, HttpEventType, HttpHandler, HttpRequest, HttpResponse } from "@angular/common/http"
import { BehaviorSubject, catchError, EMPTY, switchMap, take, tap, throwError } from "rxjs"
import { Router } from "@angular/router"
import { AuthService } from "@client/entities/auth"
import { UiService } from "@client/shared/ui/ui.service"

interface apiResponseMessage {
	messages: string[]
}

@Injectable({ providedIn: "root" })
export class GlobalInterceptor {
	private router = inject(Router)
	private authService = inject(AuthService)
	private isRefreshing = false
	private refreshSubject = new BehaviorSubject<boolean>(true)
	private ui = inject(UiService)

	intercept(req: HttpRequest<any>, next: HttpHandler) {
		return next.handle(req).pipe(
			tap(event => {
				if (event.type === HttpEventType.Response) {
					const resp: HttpResponse<apiResponseMessage> = event
					if (resp.body?.messages) {
						for (let i = 0; i < resp.body.messages.length; i++) {
							this.ui.messages.add(resp.body.messages[i])
						}
					}
				}
			}),
			catchError((err: HttpErrorResponse) => {
				if (err.status === 401) {
					return this.handle401(req, next)
				}
				return throwError(() => err)
			})
		)
	}

	private handle401(req: HttpRequest<any>, next: HttpHandler) {
		if (!this.isRefreshing) {
			this.isRefreshing = true
			this.refreshSubject.next(false)

			return this.authService.refresh().pipe(
				switchMap(() => {
					this.isRefreshing = false
					this.refreshSubject.next(true)
					return next.handle(req)
				}),
				catchError(err => {
					this.isRefreshing = false
					this.refreshSubject.next(false) // unblock waiting requests
					this.router.navigate(["/auth/login"])
					return EMPTY
				})
			)
		} else {
			return this.refreshSubject.pipe(
				take(1),
				switchMap(success => {
					if (success) {
						return next.handle(req)
					} else {
						this.router.navigate(["/auth/login"])
						return EMPTY
					}
				})
			)
		}
	}
}
