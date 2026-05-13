export type AppStatus =
    | "bootstrapping"
    | "auth_checking"
    | "auth_required"
    | "authenticated"
    | "hydrating"
    | "ready";
