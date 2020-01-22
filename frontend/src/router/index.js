import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "home",
    component: () => import("../views/Home.vue")
  }, {
    path: "/about",
    name: "about",
    component: () => import("../views/About.vue")
  }, {
    path: "/upload",
    name: "upload",
    component: () => import("../views/Upload.vue")
  }, {
    path: "/notes",
    name: "notes",
    component: () => import("../views/notes/Notes.vue")
  }, {
    path: "/notes/:id",
    name: "note",
    component: () => import("../views/notes/Note.vue")
  }, {
    path: "/login",
    name: "login",
    component: () => import("../views/Login.vue")
  }, {
    path: "/join",
    name: "join",
    component: () => import("../views/Join.vue")
  },
];

const router = new VueRouter({
  mode: "history",
  routes
});

export default router;
