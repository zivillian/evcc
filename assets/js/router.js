import { createRouter, createWebHashHistory } from "vue-router";

import Main from "./views/Main.vue";
import Config from "./views/Config.vue";
import General from "./components/Config/General.vue";
import Loadpoint from "./components/Config/Loadpoint.vue";
import PV from "./components/Config/PV.vue";
import Battery from "./components/Config/Battery.vue";
import Vehicle from "./components/Config/Vehicle.vue";
import Integration from "./components/Config/Integration.vue";

export default createRouter({
  history: createWebHashHistory(),
  routes: [
    { path: "/", component: Main, props: true },
    {
      path: "/config",
      component: Config,
      props: true,
      children: [
        { path: "general", component: General },
        { path: "loadpoint", component: Loadpoint },
        { path: "pv", component: PV },
        { path: "battery", component: Battery },
        { path: "vehicle", component: Vehicle },
        { path: "integration", component: Integration },
      ],
    },
  ],
});
