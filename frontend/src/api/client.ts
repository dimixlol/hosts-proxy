import axios, {AxiosResponse} from 'axios';
import {APIClient, APIClientConfig, APIError, APIResponse} from "./iface.ts";

export class AxiosClient implements APIClient {
    private client: axios.AxiosInstance;
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
