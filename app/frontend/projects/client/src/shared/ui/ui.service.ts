import { Injectable } from '@angular/core'
import { MessagesClass } from '@client/shared/ui/messages.class'
import { ThemeClass } from '@client/shared/ui/theme.class'
import { SidebarClass } from '@client/shared/ui/sidebar.class'

@Injectable({ providedIn: 'root' })
export class UiService {
    public theme: ThemeClass
    public messages: MessagesClass
    public sidebar: SidebarClass

    constructor() {
        this.theme = new ThemeClass()
        this.messages = new MessagesClass()
        this.sidebar = new SidebarClass()
    }
}
