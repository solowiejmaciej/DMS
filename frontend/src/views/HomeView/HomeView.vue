<template>
  <div class="flex flex-col min-h-screen">
    <NavBar @open-devices="openDevices" @open-dashboard="openDashboard" />

    <div class="flex-grow flex items-center justify-center">
      <DevicesGrid v-if="isDevicesViewActive" />
      <Dashboard v-if="isDashboardViewActive" />
    </div>
  </div>
</template>

<script setup>
import NavBar from "../../components/NavBar.vue";
import DevicesGrid from "../../components/Device/DevicesGrid.vue";
import Dashboard from "../../components/Dashboard/Dashboard.vue";
import { ref } from "vue";

const isDevicesViewActive = ref(false);
const isDashboardViewActive = ref(true);

const openDevices = () => {
  isDevicesViewActive.value = true;
  isDashboardViewActive.value = false;
};

const openDashboard = () => {
  isDashboardViewActive.value = true;
  isDevicesViewActive.value = false;
};

const openSoftware = () => {
  isSoftwareViewActive.value = true;
  isDevicesViewActive.value = false;
};
</script>
