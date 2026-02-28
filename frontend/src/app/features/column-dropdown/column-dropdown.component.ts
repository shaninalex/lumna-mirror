import { Component, Input } from '@angular/core';
import { ColumnModel } from '@entities/column';
import { ColumnDeleteFeature } from '@features/column-delete/column-delete';
import { CdkMenu, CdkMenuTrigger } from '@angular/cdk/menu';

@Component({
    selector: 'app-column-dropdown-feature',
    template: `
        <button [cdkMenuTriggerFor]="menu" class="btn btn-sm">
            <i class="fa-solid fa-ellipsis"></i>
        </button>

        <ng-template #menu>
            <div class="dropdown-menu d-block" cdkMenu>
                <app-column-delete-feature [columnId]="column.id" [columnName]="column.title" />
            </div>
        </ng-template>
    `,
    imports: [ColumnDeleteFeature, CdkMenu, CdkMenuTrigger],
})
export class ColumnDropdownFeature {
    @Input() column: ColumnModel;
}
