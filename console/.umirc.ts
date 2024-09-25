import {defineConfig} from '@umijs/max';

export default defineConfig({
    antd: {},
    access: {},
    model: {},
    initialState: {},
    request: {},
    layout: {
        title: 'X-Seek',
    },
    proxy: {
        '/api': {
            target: "http://localhost:8080",
            changeOrigin: true,
        },
    },
    routes: [
        {name: '首页', path: '/', component: './Home', icon: 'home'},
        {name: '权限演示', path: '/access', component: './Access'},
        {name: "登录", path: '/login', layout: false, component: "./Login"},
        {name: "Github 认证", path: "/login/github", layout: false, component: "./Login/Github"},
        {name: "DingTalk 认证", path: "/login/dingtalk", layout: false, component: "./Login/DingTalk"},
        {path: '*', layout: false, component: './404'}
    ],
    npmClient: 'npm',
});

