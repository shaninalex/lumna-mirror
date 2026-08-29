import { Injectable } from "@angular/core";

@Injectable({
    providedIn: 'root'
})
export class LocalStorageService {
    get(key: string): string | null {
        const s = localStorage.getItem(key);
        if (s === "") {
            return null
        }
        return s
    }

    set(key: string, s: string): void {
        localStorage.setItem(key, s);
    }

    // TODO:
    // setWithTimeout
}