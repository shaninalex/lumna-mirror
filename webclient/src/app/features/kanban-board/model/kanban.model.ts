import { ListModel } from '@entities/list';
import {TaskModel} from '@entities/task';
 
export interface KanbanModel {
    list: ListModel
    tasks: TaskModel[]
}