import Vue from "vue";
import Vuex from "vuex";

import upload from "./modules/upload";
import note from "./modules/note";
import notes from "./modules/notes";
import user from "./modules/user";
import user_last_notes from "./modules/user_last_notes";
import users from "./modules/users"
import auth from "./modules/auth";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    upload,
    note,
    notes,
    user,
    user_last_notes,
    users,
    auth
  }
});
