import type { ColumnModel } from '@entities/column';
import type { TaskModel } from '@entities/task';

export type KanbanColumn = ColumnModel & { tasks: KanbanCard[] };

export type KanbanCard = TaskModel & { column: number; position: number };
