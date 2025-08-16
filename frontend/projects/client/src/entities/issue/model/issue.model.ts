export type IssueType = "Story" | "Task" | "Bug";
export type IssueStatus = "Backlog" | "In Progress" | "Done";

export interface Issue {
    id: number;
    epicId?: number;
    sprintId?: number;
    type: IssueType;
    title: string;
    description: string;
    status: IssueStatus;
    assignee?: string;
    createdAt: string;
    updatedAt: string;
}
