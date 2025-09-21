import {Injectable} from '@angular/core';


@Injectable({ providedIn: 'root'})
export class TokenService {
    getAuthToken(): string {
        const token = localStorage.getItem("access_token")
        if (token) return token
        return ""
    }

    saveAuthToken(token: string): void {
        localStorage.setItem("access_token", token)
    }

    removeAuthToken(): void {
        localStorage.removeItem("access_token")
    }
}
