export function makeLink(link: string): string {
    const url = new URL(link, window.location.origin);

    if (url.origin === window.location.origin) {
        return url.pathname + url.search;
    }

    return link;
}
