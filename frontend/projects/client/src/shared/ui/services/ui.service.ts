import {Injectable, OnDestroy} from "@angular/core";
import {BehaviorSubject, Subscription} from "rxjs";

@Injectable({
    providedIn: 'root'
})
export class UiService implements OnDestroy {
    public title: BehaviorSubject<string> = new BehaviorSubject<string>("Login");
    public headerTitle: BehaviorSubject<string> = new BehaviorSubject<string>("");
    public darkTheme: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(false);
    public loading: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(true);
    public appLoading: BehaviorSubject<boolean> = new BehaviorSubject<boolean>(true);
    private subscriptions: Subscription = new Subscription();

    constructor() {
        const theme: string | null = localStorage.getItem("theme");
        this.darkTheme.next(theme === "dark")

        this.subscriptions.add(
            this.title.subscribe({
                next: (title: string) => {
                    document.title = `JaJira - ${title}`;
                }
            })
        );
        this.subscriptions.add(
            this.darkTheme.subscribe(isDarkTheme => {
                if (isDarkTheme) {
                    // @ts-ignore
                    document.querySelector("html").dataset.bsTheme = "dark"
                    localStorage.setItem("theme", "dark");
                } else {
                    // @ts-ignore
                    document.querySelector("html").dataset.bsTheme = "light"
                    localStorage.setItem("theme", "light");
                }
            })
        );
    }

    ngOnDestroy() {
        this.subscriptions.unsubscribe();
    }

    public changeTheme(): void { this.darkTheme.next(!this.darkTheme.value) }

    public setTitle(label: string): void {
        this.title.next(label);
        this.headerTitle.next(label);
    }
}

export enum NotificationType {
    INFO = 'info',
    WARNING = 'warning',
    DANGER = 'danger',
}

export interface Notification {
    id: string
    title: string
    description?: string
    level: NotificationType
}
