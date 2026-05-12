export interface TaskModel {
    id: number;
    title: string;
    list_id: number;
    project_id: number;
    workspace_id: number;
    status: string; // 'column' in board view
}
