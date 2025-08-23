
export interface AppError {
    id: string
    key: string
    message: string
}

// map of AppErrors
const ERRORS = {
    "generic_error": {
        "id": "APP000",
        "key": "generic_error",
        "message": "Something went wrong",
    },
    "invalid_credentials": {
        "id": "AUTH001",
        "key": "invalid_credentials",
        "message": "Invalid username or password",
    },
    /* and so on */
}
