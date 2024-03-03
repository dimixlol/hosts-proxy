import {APIClient, APIClientConfig, APIError, APIResponse} from "./iface.ts";
import success from "./mock.success.json";
import ping from "./mock.ping.json"

export class MockClient implements APIClient {
    private apiVersion: string

    constructor(config: APIClientConfig) {
        this.apiVersion = config.API_VERSION
    }

    private response(data: object): Promise<any> {
        return new Promise((resolve) => {
            resolve(data);
        })
    }
    createSite(host: string, ip: string): Promise<APIResponse | APIError> {
        return this.response(success)
    }
    ping(): Promise<APIResponse> {
        return this.response(ping)
    }
}