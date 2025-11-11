import { Component, Input } from "@angular/core"
import { DatePipe } from "@angular/common"
import { Comment } from "../../model"

@Component({
	selector: "lu-comment",
	imports: [DatePipe],
	template: `
		<div class="flex items-start gap-2">
			<div>
				<img src="img/1.png" alt="" class="w-8 rounded-full" />
			</div>

			<div>
				<div class="flex items-center gap-2">
					<div>John Doe</div>
					<div class="text-xs">{{ comment.created_at | date: "MMM d, HH:mm" }}</div>
				</div>
				<div>
					{{ comment.content }}
				</div>
			</div>
		</div>
	`,
})
export class CommentComponent {
	@Input() comment: Comment
}
