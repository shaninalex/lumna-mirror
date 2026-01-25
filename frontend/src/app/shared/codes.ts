export function Codes(): Promise<[string, string]> {
    const verifier = generateCodeVerifier()

    return generateCodeChallenge(verifier).then((challenge) => {
        return [verifier, challenge] as [string, string]
    })
}

function generateCodeVerifier(): string {
    const bytes = new Uint8Array(32)
    crypto.getRandomValues(bytes)
    return base64UrlEncode(bytes)
}

function generateCodeChallenge(verifier: string): Promise<string> {
    const data = new TextEncoder().encode(verifier)
    return crypto.subtle.digest("SHA-256", data).then((digest) => {
        return base64UrlEncode(new Uint8Array(digest))
    })
}

function base64UrlEncode(buffer: Uint8Array): string {
    let binary = ""
    for (const b of buffer) {
        binary += String.fromCharCode(b)
    }
    return btoa(binary).replace(/\+/g, "-").replace(/\//g, "_").replace(/=+$/, "")
}




/*

Codes().then(([verifier, challenge]) => {
    console.log("verifier:", verifier)
    console.log("challenge:", challenge)
})

*/