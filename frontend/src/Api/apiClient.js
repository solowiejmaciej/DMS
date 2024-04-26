import axios from "axios";
import store from "../store/store";

const apiClient = axios.create({
  baseURL: import.meta.env.VITE_APP_BASE_URL,
});

apiClient.interceptors.request.use((config) => {
  const token = store.state.token;
  const dmsApiKey = store.state.user && store.state.user.dmsApiKey;

  config.headers["Authorization"] = `${token}`;
  config.headers["X-Api-Key"] = `${dmsApiKey}`;

  return config;
});

export default apiClient;
