import {Component, inject, Input} from '@angular/core';
import {Project} from '@client/entities/project';
import {StatusColumn} from '@client/features/project/board-view-feature/board.model';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {EditStatusFormComponent} from '@client/entities/status';
import {Dialog} from '@angular/cdk/dialog';

@Component({
    selector: 'lu-column-header',
    imports: [
        ReactiveFormsModule,
        FormsModule,
    ],
    template: `
        <button (click)="openDialog()" class="cursor-pointer">
            <i class="i-dots-menu text-lg"></i>
        </button>
    `,
})
export class ColumnHeaderComponent {
    @Input() project: Project;
    @Input() column: StatusColumn;
    readonly dialog = inject(Dialog);

    openDialog(): void {
        this.dialog.open(EditStatusFormComponent, {
            data: {status: this.column.status},
        });
    }
}
