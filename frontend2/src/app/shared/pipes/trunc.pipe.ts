import { Pipe, PipeTransform } from "@angular/core";

@Pipe({
    name: "trunc"
})
export class TruncPipe implements PipeTransform {
    transform(value: string): string {
        if (!value) return "";

        const words = value.split(" ");
        if (words.length > 1) {
            return `${words[0][0]}${words[1][0]}`.toUpperCase();
        }
        if (value.length > 1) {
            return `${value[0]}${value[1]}`.toUpperCase();
        }
        return value.toUpperCase();
    }
}
