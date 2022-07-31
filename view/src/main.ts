import {createApp} from 'vue'
import './style.pcss'
import "element-plus/es/components/message/style/css";
import App from './App.vue'
import {createRouter, createWebHistory} from "vue-router";
import NProgress from "nprogress";
import "nprogress/nprogress.css";
import axios from "axios";

const app = createApp(App);

const routes = [
    {
        path: "/",
        name: "Home",
        component: () => import("./pages/Home.vue"),
    },
    {
        path: "/artwork/:id",
        name: "Artwork",
        component: () => import("./pages/Artwork.vue"),
    },
    {
        path: "/artwork/edit/:id",
        name: "ArtworkEdit",
        component: () => import("./pages/ArtworkEdit.vue"),
    },
    {
        path: "/artwork/search",
        name: "ArtworkSearch",
        component: () => import("./pages/ArtworkSearch.vue"),
    },
    {
        path: "/char/:ch",
        name: "Char",
        component: () => import("./pages/Char.vue"),
    },
    {
        path: "/char/search",
        name: "CharSearch",
        component: () => import("./pages/CharSearch.vue"),
    },
    {
        path: "/char/select/:id",
        name: "CharSelect",
        component: () => import("./pages/CharSelect.vue"),
    },
    {
        path: "/help",
        name: "Help",
        component: () => import("./pages/Help.vue"),
    },
    {
        path: "/:pathMatch(.*)*",
        name: "NotFound",
        component: () => import("./pages/NotFound.vue"),
    },
];

const router = createRouter({
    history: createWebHistory(),
    routes,
});

router.beforeEach((to, from, next) => {
    NProgress.start()
    next()
})

router.afterEach(() => {
    NProgress.done()
})

// axios请求拦截器
axios.interceptors.request.use(
    config => {
        NProgress.start()
        return config
    },
    error => {
        NProgress.done()
        return Promise.reject(error)
    }
)
// axios响应拦截器
axios.interceptors.response.use(
    function (response) {
        NProgress.done()
        return response
    },
    function (error) {
        NProgress.done()
        return Promise.reject(error)
    }
)

app.use(router);

app.mount('#app');
