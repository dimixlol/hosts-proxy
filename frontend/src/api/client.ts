import axios, {Axios, AxiosResponse} from 'axios';

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

export class APIClient {
    private client: Axios;
    private apiVersion: string;
    constructor(config: APIClientConfig) {
        this.client = axios.create({
            baseURL: config.API_URL,
            timeout: config.CLIENT_TIMEOUT,
            headers: config.HEADERS,
            withCredentials: true,
        })
        this.apiVersion = config.API_VERSION
    }
    createSite(host: string, ip: string): Promise<APIResponse|APIError> {
        const start = Date.now();
        // const csrf = this.getCSRFToken();

        if (import.meta.env.MODE === "development") {
            return new Promise((resolve) => {
                setTimeout(() => {resolve({data: {slug: "test", host: "domain.tld", ip: "1.1.1.1"}, status: 200})}, 100);
            })
        }
        return this.client.post(`/api/${this.apiVersion}/persister/create`, {host, ip})
            .then((resp: AxiosResponse) => {
              console.log({response: resp.data, latency: Date.now() - start, status: resp.status})
              if ("error" in resp.data) {
                return resp.data as APIError
              }
              return resp.data as APIResponse
            })
            .catch((err) => {
                console.log({err: err.message, latency: Date.now() - start})
                throw err;
            })
    }

    ping(): Promise<APIResponse> {
        return this.client.get("/ping")
            .then((res: AxiosResponse) => {
                return res.data as APIResponse
            }).catch((err) => {
                console.error(err)
                throw err;
            })
    };
}
