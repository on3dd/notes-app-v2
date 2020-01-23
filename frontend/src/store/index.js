import Vue from "vue";
import Vuex from "vuex";

import upload from "./modules/upload";
import note from "./modules/note";
import auth from "./modules/auth"

Vue.use(Vuex);

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: {
    upload,
    note,
    auth
  }
});
