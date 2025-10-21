import {ToastMessage} from '@client/shared/models';
import { v4 as uuidV4} from 'uuid'
import {BehaviorSubject, Observable} from 'rxjs';


export class MessagesClass {
    private _messages$: BehaviorSubject<ToastMessage[]> = new BehaviorSubject<ToastMessage[]>([]);

    public list(): Observable<ToastMessage[]> {
        return this._messages$.asObservable()
    }

    public add(msg: string): void {
        const messages = [...this._messages$.value]
        messages.push({
            id: uuidV4(),
            message: msg,
        })
        this._messages$.next(messages)
    }

    public remove(id: string): void {
        let messages = [...this._messages$.value]
        this._messages$.next(messages.filter(m => m.id !== id))
    }
}
