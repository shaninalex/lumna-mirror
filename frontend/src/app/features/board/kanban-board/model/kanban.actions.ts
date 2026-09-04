import type { CdkDragDrop } from '@angular/cdk/drag-drop';
import type { ColumnModel } from '@entities/column';
import { createActionGroup, props } from '@ngrx/store';
import type { KanbanCard } from './kanban.models';

export const actionKanban = createActionGroup({
    source: 'Kanban',
    events: {
        dropColumn: props<{ event: unknown }>(),
        dropTask: props<{ event: CdkDragDrop<KanbanCard[]>; column: ColumnModel }>(),
    },
});
