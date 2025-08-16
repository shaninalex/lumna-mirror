

import {Injectable} from "@angular/core";
import {BehaviorSubject} from 'rxjs';
import {IToast, ToastType} from '@client/shared/models';

@Injectable({
    providedIn: 'root'
})
export class ToastService {
    public notifications: BehaviorSubject<IToast[]> = new BehaviorSubject<IToast[]>([]);

    public add(message: string, type: ToastType, header?: string): void {
        const toast: IToast = {
            message,
            type,
        }
        if (header) {
            toast.header = header
        }
        const notifications = this.notifications.getValue()
        notifications.push(toast)
        this.notifications.next(notifications);
    }

    public remove(idx: number): void {
        let notifications = this.notifications.getValue()
        notifications = notifications.splice(idx, 1);
        this.notifications.next(notifications);
    }
}
