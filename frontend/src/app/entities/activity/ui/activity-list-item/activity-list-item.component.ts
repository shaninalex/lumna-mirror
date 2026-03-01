import {Component, inject, Input, OnInit} from '@angular/core';
import {ActivityModel} from '@entities/activity';
import {Store} from '@ngrx/store';
import {AsyncPipe, DatePipe} from '@angular/common';
import {selectUser} from '@entities/user';

@Component({
    selector: 'app-activity-list-item',
    imports: [
        DatePipe,
        AsyncPipe
    ],
    template: `
        @if (user$ | async; as user) {
            <div class="d-flex align-items-start gap-1">
                <div>
                    <img class="rounded-circle" src="img/7.png" style="width: 32px; height: 32px;"/>
                </div>
                <div>
                    <div><b>{{ user.full_name }}</b> {{ activity.summary }}</div>
                    <div>{{ activity.created_at | date: 'd MMM, HH:mm' }}</div>
                </div>
            </div>
        }
    `
})
export class ActivityListItemComponent {
    @Input() activity: ActivityModel
    private store = inject(Store)
    user$ = this.store.select(selectUser)
}
