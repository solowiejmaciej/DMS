<template>
  <NavBar
    @open-devices="openDevices"
    @open-dashboard="openDashboard"
    @open-software="openSoftware"
  />
  <DevicesGrid v-if="isDevicesViewActive" />
  <div v-if="isDashboardViewActive">Dashboard</div>
  <Software v-if="isSoftwareViewActive" />
</template>

<script setup>
import NavBar from "../../components/NavBar.vue";
import DevicesGrid from "../../components/Device/DevicesGrid.vue";
import Software from "../../components/Device/Software/Software.vue";
import { ref } from "vue";

const isDevicesViewActive = ref(false);
const isDashboardViewActive = ref(true);
const isSoftwareViewActive = ref(false);

const openDevices = () => {
  isDevicesViewActive.value = true;
  isDashboardViewActive.value = false;
  isSoftwareViewActive.value = false;
};

const openDashboard = () => {
  isDashboardViewActive.value = true;
  isDevicesViewActive.value = false;
  isSoftwareViewActive.value = false;
};

const openSoftware = () => {
  isSoftwareViewActive.value = true;
  isDevicesViewActive.value = false;
  isDashboardViewActive.value = false;
};
</script>
