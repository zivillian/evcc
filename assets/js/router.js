import { createRouter, createWebHistory } from "vue-router";

import Main from "./views/Main.vue";
import Config from "./views/Config.vue";

export default createRouter({
  history: createWebHistory(),
  routes: [
    { path: "/", component: Main, props: true },
    { path: "/config", component: Config, props: true },
  ],
});
