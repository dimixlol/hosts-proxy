import axios, {Axios, AxiosResponse} from 'axios';

export interface APIClientConfig {
    API_URL: string;
    CLIENT_TIMEOUT: number;
    HEADERS: object;

}

export interface APIResponse {
    status: number
    data: {
        slug: string;
        host: string;
        ip: string;
    }
}

export interface APIError {
    status: number;
    error: string;
}

export class APIClient {
    private client: Axios;
    constructor(config: APIClientConfig) {
        this.client = axios.create({
            baseURL: config.API_URL,
            timeout: config.CLIENT_TIMEOUT,
            withCredentials: false,
            // headers: config.HEADERS,
        })
    }
    createSite(host: string, ip: string): Promise<APIResponse|APIError> {
        const start = Date.now();
        // const csrf = this.getCSRFToken();

        if (import.meta.env.MODE === "development") {
            return new Promise((resolve) => {
                setTimeout(() => {resolve({data: {slug: "test", host: "domain.tld", ip: "1.1.1.1"}, status: 200})}, 100);
            })
        }

        return this.client.post("/persister/create/", {host, ip})// {headers: {"X-CSRF-Token": csrf}}
            .then((resp: AxiosResponse) => {
              console.log({response: resp.data, latency: Date.now() - start, status: resp.status})
              return resp.data
            })
            .catch((err) => {
                console.log({err: err.message, latency: Date.now() - start})
                throw err;
            })
    }

    createCSRFToken() {
        this.client.get("/csrf")
            .then((res: AxiosResponse) => {
                return res.data
            }).catch((err) => {
                throw err;
            })
    };

    private getCSRFToken() {
        const csrf = document.getElementById("csrf_token")
        if (csrf !== null) {
            const token = csrf.getAttribute("content")
            if (token !== null) {
                return token
            }
        }
        return null
    }
}
