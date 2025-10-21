interface timedObject {
    created_at: Date | string
}

export function byMostRecent<T extends timedObject>(a: T, b: T) {
    return new Date(b.created_at).getTime() - new Date(a.created_at).getTime();
}
