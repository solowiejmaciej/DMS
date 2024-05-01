// src/store/index.js
import apiClient from "../Api/apiClient";
import VuexPersist from "vuex-persist";
import Vuex from "vuex";
import router from "../plugins/router";

const vuexLocalStorage = new VuexPersist({
  key: "vuex",
  storage: window.sessionStorage,
});

const store = new Vuex.Store({
  state: {
    status: "",
    token: localStorage.getItem("token") || "",
    refreshToken: localStorage.getItem("refreshToken") || "",
    user: {},
    userId: localStorage.getItem("userId") || "",
    dmsApiKey: localStorage.getItem("dmsApiKey") || "",
    tokenRefreshFailed: false,
  },
  mutations: {
    auth_request(state) {
      state.status = "loading";
    },
    auth_success(state, payload) {
      console.log("auth_success mutation", payload);
      state.status = "success";
      state.token = payload.token;
      state.refreshToken = payload.refreshToken;
      state.user = payload.user;
      if (payload.user) {
        state.userId = payload.user.id;
        state.dmsApiKey = payload.user.dmsApiKey;
        console.log(state);
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
      state.refreshToken = "";
      router.push("/login?showLogoutMessage=true");
    },
    setDmsApiKey(state, dmsApiKey) {
      state.dmsApiKey = dmsApiKey.dmsApiKey;
    },
    setTokens(state, tokens) {
      state.token = tokens.token;
      state.refreshToken = tokens.refreshToken;
    },
    setTokenRefreshFailed(state, value) {
      state.tokenRefreshFailed = value;
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
              const refreshToken = resp.data.refreshToken;
              const user = resp.data.user;

              commit("auth_success", { token, refreshToken, user });
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
