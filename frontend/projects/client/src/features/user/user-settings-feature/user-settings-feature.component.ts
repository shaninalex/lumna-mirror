import {Component, inject, OnInit} from '@angular/core';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import {selectUser, Settings, UserSetAction, UserUpdateSettingsAction, UserModel} from '@client/entities/user';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {filter} from 'rxjs';
import {LoaderComponent} from '@client/shared/ui/loader';
import {Actions, ofType} from '@ngrx/effects';

@Component({
    selector: 'fr-settings-form-feature',
    standalone: true,
    imports: [ReactiveFormsModule, LoaderComponent],
    template: `
        @if (user) {
            <div class="flex items-center gap-2">
                <div>Username:</div>
                <div class="text-lg font-bold">{{ user.code }}</div>
            </div>
        }

        <hr class="my-4 border-gray-300">

        <form [formGroup]="form" (ngSubmit)="save()">
            <div class="mb-4">
                <label for="theme">Theme</label>
                <select id="theme" formControlName="theme" class="input">
                    <option [value]="0">Light</option>
                    <option [value]="1">Dark</option>
                    <option [value]="2">Device</option>
                </select>
            </div>

            <div class="mb-4">
                <label for="language">Language</label>
                <select id="language" formControlName="language" class="input">
                    <option [value]="0">English</option>
                    <option [value]="1">Ukrainian</option>
                    <option [value]="2">German</option>
                </select>
            </div>

            <div class="mb-4">
                <label for="timezone">Timezone</label>
                <input id="timezone" type="text" formControlName="timezone" class="input"/>
            </div>

            <div class="mb-4">
                <label for="date_format">Date format</label>
                <input id="date_format" type="text" formControlName="date_format" class="input"/>
            </div>

            <div class="mb-4">
                <label for="time_format">Time format</label>
                <input id="time_format" type="text" formControlName="time_format" class="input"/>
            </div>

            <div class="mb-4">
                <label for="week_start_day">Week Start Day</label>
                <select id="week_start_day" formControlName="week_start_day" class="input">
                    @for (d of weekDays; track $index) {
                        <option [value]="$index">{{ d }}</option>
                    }
                </select>
            </div>

            @if (loading) {
                <ui-loader />
            } @else {
                <button class="btn btn-primary" type="submit">Save</button>
            }
        </form>
    `
})
export class SettingsFormComponent implements OnInit {
    private fb = inject(FormBuilder);
    private actions$ = inject(Actions);
    private store = inject(Store<AppState>);
    loading: boolean = false;
    user: UserModel
    weekDays = ['Monday','Tuesday','Wednesday','Thursday','Friday','Saturday','Sunday'];
    form: FormGroup = new FormGroup({
        theme: new FormControl("light"),
        language: new FormControl("en"),
        timezone: new FormControl(''),
        date_format: new FormControl('YYYY-MM-DD'),
        time_format: new FormControl('HH:mm'),
        week_start_day: new FormControl("0"),
    });

    constructor() {
        this.actions$.pipe(ofType(UserSetAction)).subscribe(() => this.loading = false)
    }

    ngOnInit() {
        this.store.select(selectUser).pipe(filter(user => !!user)).subscribe(user => {
            if (!user) return
            this.user = user
            this.form.setValue({
                theme: user.settings.theme,
                language: user.settings.language,
                timezone: user.settings.timezone,
                date_format: user.settings.date_format,
                time_format: user.settings.time_format,
                week_start_day: user.settings.week_start_day,
            })
        })
    }

    save() {
        this.loading = true
        const settings: Settings = this.form.value as Settings;
        this.store.dispatch(UserUpdateSettingsAction({ payload: settings }))
    }
}
