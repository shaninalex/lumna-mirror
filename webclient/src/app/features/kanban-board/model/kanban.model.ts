import {TaskModel} from '@entities/task';
 
export interface KanbanModel {
    id: number
    title: string
    order: number
    tasks: TaskModel[]
}