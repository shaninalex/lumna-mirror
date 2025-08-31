import {Component, inject} from '@angular/core';
import {Store} from '@ngrx/store';
import {AppState} from '@client/shared/store';
import {map, Observable} from 'rxjs';
import {selectSession} from '@client/entities/session';
import {AsyncPipe} from '@angular/common';
import {CdkMenu, CdkMenuItem, CdkMenuModule} from '@angular/cdk/menu';

@Component({
    selector: 'fr-header',
    imports: [
        AsyncPipe,
        CdkMenu,
        CdkMenuItem,
        CdkMenuModule,
    ],
    template: `
        <div class="bg-white py-2 px-4 flex items-center justify-between">
            <div class="flex items-center gap-2 ms-auto">
                @if (email$ | async; as email) {
                    <div>{{ email }}</div>
                    <button [cdkMenuTriggerFor]="header_menu" class="flex items-center gap-2 cursor-pointer">
                        <img src="img/1.png" class="rounded-full w-8" alt="">
                    </button>
                    <ng-template #header_menu>
                        <div class="bg-white flex flex-col gap-2 border p-2 rounded mt-2" cdkMenu>
                            <button cdkMenuItem>Refresh</button>
                            <button cdkMenuItem>Settings</button>
                            <button cdkMenuItem>Help</button>
                            <button cdkMenuItem>Sign out</button>
                        </div>
                    </ng-template>
                }
            </div>
        </div>
    `
})
export class Header {
    private store: Store<AppState> = inject(Store<AppState>);
    email$: Observable<string> = this.store.select(selectSession).pipe(
        map(session => session?.identity?.traits?.email)
    )
}
