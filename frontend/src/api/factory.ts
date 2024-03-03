import {APIClient, APIClientConfig} from "./iface.ts";
import {MockClient} from "./mock.ts";
import {AxiosClient} from "./client.ts";

export class APIClientFactory {
    makeClient(config: APIClientConfig=this.makeConfig()): APIClient  {
        if (import.meta.env.MODE === "development") {
            return new MockClient(config)
        }
        return new AxiosClient(config)

    }
    makeConfig(): APIClientConfig {
        return {
            API_URL: import.meta.env.VITE_API_URL,
            CLIENT_TIMEOUT: import.meta.env.VITE_CLIENT_TIMEOUT,
            HEADERS: import.meta.env.VITE_HEADERS,
            API_VERSION: import.meta.env.VITE_API_VERSION
        }
    }
}