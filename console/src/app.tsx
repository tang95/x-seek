import {TOKEN_NAME} from "./constants";
import {history} from "@umijs/max";
import {RequestConfig, RequestError} from "@@/plugin-request/request";
import {message} from "antd";

export async function getInitialState(): Promise<API.AuthUser | null> {
    const str = localStorage.getItem(TOKEN_NAME);
    if (str && str.length > 0) {
        return JSON.parse(str);
    }
    if (!window.location.pathname.startsWith("/login")) {
        history.push('/login');
    }
    return null;
}

export const layout = () => {
    return {
        logo: '/favicon.png',
        layout: "mix",
        splitMenus: true,
        menu: {
            locale: false,
        },
        logout: async () => {
            localStorage.removeItem(TOKEN_NAME);
            return history.push('/login');
        },
    };
};


export const request: RequestConfig = {
    timeout: 10000,
    errorConfig: {
        errorHandler: (error: RequestError) => {
            // @ts-ignore
            const {response} = error;
            if (response?.data?.msg) {
                message.error(response?.data?.msg);
            }
            if (response?.status === 401) {
                message.error('登录已失效，重新登录');
                return history.push('/login');
            }
            throw error;
        }
    },
    requestInterceptors: [
        (url, options) => {
            const str = localStorage.getItem(TOKEN_NAME);
            if (str && str.length > 0) {
                const token = JSON.parse(str).token;
                const authHeader = {Authorization: `Bearer ${token}`};
                return {
                    url: url,
                    options: {...options, interceptors: true, headers: authHeader},
                };
            } else if (!history.location.pathname.startsWith('/login')) {
                history.push('/login');
                throw new Error('登录已失效，重新登录');
            }
            return {
                url: url,
                options: {...options, interceptors: true},
            };
        },
    ],
};
