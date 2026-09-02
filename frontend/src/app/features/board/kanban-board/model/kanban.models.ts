import type { ColumnModel } from '@entities/column';
import type { TaskModel } from '@entities/task';

export type KanbanColumn = ColumnModel & { tasks: KanbanCard[] };

export interface KanbanCard {
    id: number;
    column: number;
    position: number;

    task: TaskModel;
}
