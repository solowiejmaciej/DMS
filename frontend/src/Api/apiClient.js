import axios from "axios";
import store from "../store/store";

const apiClient = axios.create({
  baseURL: import.meta.env.VITE_APP_BASE_URL,
});

async function refreshToken() {
  const currentJwt = store.state.token;
  const currentRefreshToken = store.state.refreshToken;

  const response = await apiClient.post("/token/refresh", {
    refreshToken: currentRefreshToken,
    token: currentJwt,
  });

  return response.data;
}

apiClient.interceptors.response.use(undefined, async (error) => {
  const isRefreshTokenRequest = error.config.url.includes("/token/refresh");

  if (
    error.config &&
    error.response &&
    error.response.status === 401 &&
    !isRefreshTokenRequest
  ) {
    console.log("401 error", error);
    try {
      const response = await refreshToken();

      store.commit("setTokens", response);

      error.config.headers["Authorization"] = `${response.token}`;

      return apiClient.request(error.config);
    } catch (e) {
      console.log("Error refreshing token");
      store.commit("setTokenRefreshFailed", true);

      return Promise.reject(error);
    }
  }

  if (
    isRefreshTokenRequest &&
    error.response &&
    error.response.status === 401
  ) {
    console.log("Error refreshing token");
    await store.dispatch("logout");
  }

  return Promise.reject(error);
});

apiClient.interceptors.request.use((config) => {
  const token = store.state.token;
  const dmsApiKey = store.state.dmsApiKey;

  config.headers["Authorization"] = `${token}`;
  config.headers["X-Api-Key"] = `${dmsApiKey}`;

  return config;
});

export default apiClient;
