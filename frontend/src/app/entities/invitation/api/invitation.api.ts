import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';
import { InvitationModel } from '../model/invitation.model';

@Injectable({
    providedIn: 'root',
})
export class InvitationApi {
    http = inject(HttpClient);

    List(): Observable<InvitationModel[]> {
        return this.http
            .get<APIResponse<InvitationModel[]>>(`/api/v1/invitation`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Create(payload: InvitationModel): Observable<InvitationModel> {
        return this.http
            .post<
                APIResponse<InvitationModel>
            >(`/api/v1/invitation`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Delete(invitationId: number): Observable<void> {
        return this.http
            .delete<
                APIResponse<void>
            >(`/api/v1/invitation/${invitationId}`, { withCredentials: true })
            .pipe(map((response) => response.data));
    }

    Patch(invitationId: number, payload: InvitationModel): Observable<InvitationModel> {
        return this.http
            .patch<
                APIResponse<InvitationModel>
            >(`/api/v1/invitation/${invitationId}`, payload, { withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
