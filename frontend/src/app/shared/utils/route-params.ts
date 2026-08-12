/**
 * Parses a numeric entity id out of a route param.
 *
 * parseInt() is not usable here: it stops at the first non-digit, so garbage
 * like "1werwerwer" would resolve to 1 and silently open an entity the URL
 * never referred to. Only a plain, unsigned integer is accepted, anything
 * else is reported as null so the caller can treat it as an invalid route.
 */
export function parseRouteId(value: string | null | undefined): number | null {
    if (!value || !/^\d+$/.test(value)) {
        return null;
    }

    const id = Number(value);
    return Number.isSafeInteger(id) ? id : null;
}
