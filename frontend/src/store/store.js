// src/store/index.js
import apiClient from "../Api/apiClient";
import VuexPersist from "vuex-persist";
import Vuex from "vuex";

const vuexLocalStorage = new VuexPersist({
  key: "vuex",
  storage: window.sessionStorage,
});

const store = new Vuex.Store({
  state: {
    status: "",
    token: localStorage.getItem("token") || "",
    user: {},
    userId: localStorage.getItem("userId") || "",
    dmsApiKey: localStorage.getItem("dmsApiKey") || "",
  },
  mutations: {
    auth_request(state) {
      state.status = "loading";
    },
    auth_success(state, payload) {
      state.status = "success";
      state.token = payload.token;
      state.user = payload.user;
      if (payload.user) {
        state.userId = payload.user.id;
        state.dmsApiKey = payload.user.dmsApiKey;
      } else {
        console.error("User is undefined in auth_success mutation");
      }
    },
    auth_error(state) {
      state.status = "error";
    },
    logout(state) {
      state.status = "";
      state.token = "";
    },
  },
  actions: {
    login({ commit }, user) {
      return new Promise((resolve, reject) => {
        commit("auth_request");
        apiClient
          .post("/token", user)
          .then((resp) => {
            if (resp.data.user) {
              const token = resp.data.token;
              const user = resp.data.user;

              commit("auth_success", { token, user });
              resolve(resp);
            } else {
              throw new Error(
                "User data is not defined in the server response"
              );
            }
          })
          .catch((err) => {
            commit("auth_error");

            delete apiClient.defaults.headers.common["Authorization"];
            delete apiClient.defaults.headers.common["X-Api-Key"];

            reject(err);
          });
      });
    },
    logout({ commit }) {
      return new Promise((resolve) => {
        commit("logout");

        localStorage.removeItem("token");
        localStorage.removeItem("dmsApiKey");
        localStorage.removeItem("userId");

        delete apiClient.defaults.headers.common["Authorization"];
        delete apiClient.defaults.headers.common["X-Api-Key"];

        resolve();
      });
    },
  },
  getters: {
    isLoggedIn: (state) => !!state.token,
    authStatus: (state) => state.status,
  },
  plugins: [vuexLocalStorage.plugin],
});

export default store;
