import {Component, inject, Input} from '@angular/core';
import {Project} from '@client/entities/project';
import {StatusColumn} from '@client/features/project/board-view-feature/board.model';
import {FormsModule, ReactiveFormsModule} from '@angular/forms';
import {EditStatusFormComponent, Status} from '@client/entities/status';
import {MatButtonModule} from '@angular/material/button';
import {MatIconModule} from '@angular/material/icon';
import {MatDialog,} from '@angular/material/dialog';

@Component({
    selector: 'fr-column-header',
    imports: [
        ReactiveFormsModule,
        FormsModule,
        MatButtonModule,
        MatIconModule,
    ],
    template: `
        <div class="mb-2 flex justify-between">
            <div class="flex justify-between gap-2">
                <button matIconButton (click)="openDialog()">
                    <mat-icon>more_vert</mat-icon>
                </button>
            </div>
        </div>
    `,
})
export class ColumnHeaderComponent {
    @Input() project: Project;
    @Input() column: StatusColumn;
    readonly dialog = inject(MatDialog);

    openDialog(): void {
        const dialogRef = this.dialog.open(EditStatusFormComponent, {
            data: {status: this.column.status},
        });

        dialogRef.afterClosed().subscribe(result => {
            if (result !== undefined) {
                console.log(result)
            }
        });
    }
}
