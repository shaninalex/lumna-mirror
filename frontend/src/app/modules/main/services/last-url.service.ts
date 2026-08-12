import { Injectable } from "@angular/core";

const LastUrlServiceKEY = "last_url"

@Injectable({
    providedIn: 'root'
})
export class LastUrlService {
    get(): string | null {
        const u = localStorage.getItem(LastUrlServiceKEY);
        return u === '' ? u: null;
    }

    set(u: string): void {
        localStorage.setItem(LastUrlServiceKEY, u);
    }
}