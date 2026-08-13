import { Injectable } from "@angular/core";

@Injectable()
export class LocalStorageService {
    get(key: string): string | null {
        const s = localStorage.getItem(key);
        return s === '' ? s: null;
    }

    set(key: string, s: string): void {
        localStorage.setItem(key, s);
    }

    // TODO:
    // setWithTimeout
}