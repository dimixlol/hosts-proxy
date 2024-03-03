export interface APIClientConfig {
    API_URL: string;
    CLIENT_TIMEOUT: number;
    HEADERS: object;
    API_VERSION: string;
}

export interface APISlugDataResponse {
    slug: string;
    host: string;
    ip: string;
}

export interface APIResponse {
    status: number
    data: APISlugDataResponse
}

export interface APIError {
    status: number;
    error: string;
}

export abstract class APIClient {
    abstract createSite(host: string, ip: string): Promise<APIResponse|APIError>
    abstract ping(): Promise<APIResponse>
}
