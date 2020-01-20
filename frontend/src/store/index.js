import Vue from "vue";
import Vuex from "vuex";
import upload from "./modules/upload";

Vue.use(Vuex);

export default new Vuex.Store({
  state: {},
  mutations: {},
  actions: {},
  modules: { upload }
});
