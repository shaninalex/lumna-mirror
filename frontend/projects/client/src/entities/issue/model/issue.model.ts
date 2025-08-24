export type IssueType = "Story" | "Task" | "Bug";

export interface Issue {
    id: number;
    epicId?: number;
    sprintId?: number;
    type: IssueType;
    title: string;
    description?: string;
    status: string;
    assignee?: string;
    createdAt: Date;
    updatedAt: Date;
    deletedAt?: Date;
}
