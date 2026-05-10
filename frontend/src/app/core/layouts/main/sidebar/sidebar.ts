import { Component } from "@angular/core";
import { RouterLink } from "@angular/router";
import packageJson from "@root/package.json";
import { WorkspaceDropdown } from "@entities/workspaces";

@Component({
    selector: "app-sidebar",
    imports: [RouterLink, WorkspaceDropdown],
    templateUrl: "./sidebar.html",
    styleUrl: "./sidebar.css"
})
export class Sidebar {
    varsion: string = packageJson.version;
}
