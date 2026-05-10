import { Component } from "@angular/core";
import { AppLayoutHeader } from "./app-layout-header";
import { AppLayoutSidebar } from "./app-layout-sidebar";

@Component({
    selector: "app-layout",
    imports: [AppLayoutSidebar, AppLayoutHeader],
    templateUrl: "./app.layout.html",
    styleUrl: "./app.layout.css"
})
export class AppLayout {}
