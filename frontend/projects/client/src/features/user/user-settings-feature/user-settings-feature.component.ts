import {Component, inject, OnInit} from '@angular/core';
import { FormBuilder, FormControl, FormGroup, ReactiveFormsModule } from '@angular/forms';
import {Language, Settings, Theme, WeekStartDay} from '@client/entities/user';

@Component({
    selector: 'fr-settings-form-feature',
    standalone: true,
    imports: [ReactiveFormsModule],
    template: `
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
                    @for (d of weekDays; track d) {
                        <option [value]="d">{{ d }}</option>
                    }
                </select>
            </div>

            <button class="btn" type="submit">Save</button>
        </form>
    `
})
export class SettingsFormComponent implements OnInit {
    private fb = inject(FormBuilder);

    weekDays = ['Monday','Tuesday','Wednesday','Thursday','Friday','Saturday','Sunday'];
    form: FormGroup<{
        theme: FormControl<Theme>;
        language: FormControl<Language>;
        timezone: FormControl<string>;
        date_format: FormControl<string>;
        time_format: FormControl<string>;
        week_start_day: FormControl<WeekStartDay>;
    }>;

    ngOnInit() {
        this.form = this.fb.group({
            theme: this.fb.control<Theme>(Theme.Light, { nonNullable: true }),
            language: this.fb.control<Language>(Language.EN, { nonNullable: true }),
            timezone: this.fb.control<string>('', { nonNullable: true }),
            date_format: this.fb.control<string>('YYYY-MM-DD', { nonNullable: true }),
            time_format: this.fb.control<string>('HH:mm', { nonNullable: true }),
            week_start_day: this.fb.control<WeekStartDay>(WeekStartDay.Monday, { nonNullable: true }),
        });
    }

    save() {
        const settings: Settings = this.form.value as Settings;
        console.log(settings);
    }
}
