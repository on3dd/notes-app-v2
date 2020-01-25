import Vue from "vue";
import VueRouter from "vue-router";

Vue.use(VueRouter);

const routes = [
  {
    path: "/",
    name: "home",
    meta: {
      title: "Главная"
    },
    component: () => import("../views/Home.vue")
  }, {
    path: "/upload",
    name: "upload",
    meta: {
      title: "Загрузка"
    },
    component: () => import("../views/Upload.vue")
  }, {
    path: "/notes",
    name: "notes",
    meta: {
      title: "Последние записи"
    },
    component: () => import("../views/notes/Notes.vue")
  }, {
    path: "/notes/:id",
    name: "note",
    meta: {
      title: "Запись"
    },
    component: () => import("../views/notes/Note.vue")
  }, {
    path: "/login",
    name: "login",
    meta: {
      title: "Вход"
    },
    component: () => import("../views/Login.vue")
  }, {
    path: "/join",
    name: "join",
    meta: {
      title: "Регистрация"
    },
    component: () => import("../views/Join.vue")
  },
];

const router = new VueRouter({
  mode: "history",
  routes
});

router.beforeEach((to, from, next) => {
  if(to.matched.some(record => record.meta.requiresAuth)) {
    if (store.getters.isLoggedIn) {
      next()
      return
    }
    next('/login')
  } else {
    next()
  }
})

export default router;
