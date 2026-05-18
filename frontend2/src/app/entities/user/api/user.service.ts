import { inject, Injectable } from "@angular/core";
import { map, Observable } from "rxjs";
import { UserModel } from "../model/user.model";
import { APIResponse } from "@shared/models";
import { HttpClient, HttpParams } from "@angular/common/http";

@Injectable({
    providedIn: "root"
})
export class UserApi {
    http = inject(HttpClient);

    me(): Observable<UserModel> {
        return this.http
            .get<
                APIResponse<UserModel>
            >(`/api/v1/user/me`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    invitationInit(token: string): Observable<any> {
        console.log("validate invitation");
        const params = new HttpParams().append("token", token);
        return this.http
            .get<
                APIResponse<any>
            >(`/api/v1/onboarding/invitation`, { params, withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
