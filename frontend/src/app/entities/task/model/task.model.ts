export interface TaskModel {
    id: number;
    title: string;
    order: number;
    done: boolean;
    body: string;
    code: string;

    status_id: number;
    project_id: number;

    created_at: Date;
    updated_at: Date;
}

export function makeListLabel(task: TaskModel): string {
    return `${task.code}  ${task.title}`;
}

export interface TaskCreateModel {
    title: string;
    body: string;
    order: number;
    column_id: number;
    board_id: number;
}

export interface TaskListQueryModel {
    board_id: number
}
