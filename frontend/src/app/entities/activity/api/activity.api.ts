import { inject, Injectable } from '@angular/core';
import { map, Observable } from 'rxjs';
import { APIResponse } from '@shared/models';
import { HttpClient } from '@angular/common/http';
import {ActivityModel} from '@entities/activity/model';

@Injectable({
    providedIn: 'root',
})
export class ActivityApi {
    http = inject(HttpClient);

    List(activity_id: number, activity_type: string): Observable<ActivityModel[]> {
        return this.http
            .get<
                APIResponse<ActivityModel[]>
            >(`/api/v1/activity/logs`, { params: {activity_id, activity_type}, withCredentials: true })
            .pipe(map((response) => response.data));
    }
}
