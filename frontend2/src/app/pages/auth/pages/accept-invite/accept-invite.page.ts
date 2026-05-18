import { AsyncPipe } from "@angular/common";
import { Component, inject } from "@angular/core";
import { ActivatedRoute } from "@angular/router";
import { filter, map } from "rxjs";
import { AcceptInviteFeature } from "@features/onboarding";

@Component({
    selector: "app-auth-login-page",
    imports: [AsyncPipe, AcceptInviteFeature],
    template: `
        @if (token$ | async; as token) {
            <app-accept-invite-feature [token]="token" />
        }
    `
})
export class AcceptInvitePage {
    private route = inject(ActivatedRoute);
    token$ = this.route.params.pipe(
        filter((params) => !!params["token"]),
        map((params) => params["token"])
    );
}
