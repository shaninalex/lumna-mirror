import { inject, Injectable } from "@angular/core"
import { environment as env } from "@client/environments/environment.development"
import { catchError, map, Observable, throwError } from "rxjs"
import { HttpClient } from "@angular/common/http"
import { APIResponse } from "@client/shared/models"

@Injectable({ providedIn: "root" })
export class AuthService {
	http = inject(HttpClient)

	login(data: any): Observable<APIResponse<any>> {
		return this.http.post<APIResponse<any>>(`${env.API_ROOT}/api/v1/auth/login`, data, { withCredentials: true })
	}

	register(data: any): Observable<APIResponse<any>> {
		return this.http.post<APIResponse<any>>(`${env.API_ROOT}/api/v1/auth/register`, data, { withCredentials: true })
	}

	refresh(): Observable<void> {
		return this.http.get(`${env.API_ROOT}/api/v1/auth/refresh`, { withCredentials: true }).pipe(
			map(() => void 0), // emit something
			catchError(err => {
				return throwError(() => err) // propagate error
			})
		)
	}
}
