<template>
  <div class="flex" v-if="!shouldShowInstruction">
    <div class="w-3/5 overflow-auto card">
      <div class="flex justify-between items-center">
        <h1 class="text-2xl font-semibold">Devices</h1>
        <Button
          label="Refresh"
          icon="pi pi-refresh"
          @click="getDevices"
          class="p-button-sm p-button-text"
        />
      </div>
      <DataTable
        :loading="isDataTableLoading"
        scrollable
        scrollHeight="30rem"
        :value="devices"
        selectionMode="single"
        dataKey="deviceId"
        v-model:selection="selectedDevice"
        filterDisplay="row"
        v-model:filters="filters"
        :selection.sync="selectedDevice"
      >
        <Column header="#" :showFilterMenu="false">
          <template #body="slotProps">
            {{ slotProps.index + 1 }}
          </template>
        </Column>
        <Column
          field="deviceId"
          header="Id"
          :showFilterMenu="false"
          style="min-width: 10rem"
        >
          <template #filter="{ filterModel, filterCallback }">
            <InputText
              v-model="filterModel.value"
              type="text"
              @input="filterCallback()"
              placeholder="Search"
              class="p-column-filter p-inputtext-sm"
            /> </template
        ></Column>
        <Column
          field="deviceModel"
          header="Model"
          :showFilterMenu="false"
          style="min-width: 10rem"
        >
          <template #filter="{ filterModel, filterCallback }">
            <InputText
              v-model="filterModel.value"
              type="text"
              @input="filterCallback()"
              placeholder="Search"
              class="p-column-filter p-inputtext-sm"
            />
          </template>
        </Column>
        <Column
          field="softwareVersion"
          header="Software version"
          :showFilterMenu="false"
          style="min-width: 10rem"
        >
          <template #filter="{ filterModel, filterCallback }">
            <InputText
              v-model="filterModel.value"
              type="text"
              @input="filterCallback()"
              class="p-column-filter p-inputtext-sm"
              placeholder="Search"
            />
          </template>
        </Column>
        <Column
          field="status"
          header="Device status"
          :showFilterMenu="false"
          style="min-width: 10rem"
        >
          <template #filter="{ filterModel, filterCallback }">
            <Dropdown
              v-model="filterModel.value"
              @change="filterCallback()"
              :options="statuses"
              placeholder="Select One"
              class="p-column-filter p-inputtext-sm"
              style="max-width: 12rem"
              :showClear="true"
            >
              <template #option="slotProps">
                <Tag
                  :value="slotProps.option"
                  :severity="getSeverity(slotProps.option)"
                />
              </template>
            </Dropdown>
          </template>
          <template #body="slotProps">
            <Tag
              :value="slotProps.data.status"
              :severity="getSeverity(slotProps.data.status)"
            />
          </template>
        </Column>
      </DataTable>
    </div>
    <div class="ml-2">
      <h2 class="flex items-center justify-center">
        <i class="pi pi-microchip mr-2" style="font-size: 1.5rem"></i>
        Device: {{ selectedDevice.deviceId }}
      </h2>
      <TabView
        v-if="selectedDevice"
        style="min-width: 40rem; min-height: 20rem"
      >
        <TabPanel>
          <template #header>
            <div class="header-content">
              <i class="pi pi-cog mr-2"></i> Configuration
            </div>
          </template>
          <DeviceConfiguration :device="selectedDevice" />
        </TabPanel>
        <TabPanel>
          <template #header>
            <div class="header-content">
              <i class="pi pi-code mr-2"></i> Software
            </div>
          </template>
          <Software :device="selectedDevice" />
        </TabPanel>
      </TabView>
    </div>
  </div>
  <AddDeviceInstructions
    v-if="shouldShowInstruction"
    @setupFinished="handleSetupFinished"
  />
</template>

<script setup>
import { ref, reactive, onMounted, watch } from "vue";
import { FilterMatchMode } from "primevue/api";
import { useToast } from "primevue/usetoast";
import apiClient from "../../Api/apiClient";
import AddDeviceInstructions from "./AddDeviceInstructions.vue";
import Software from "./Software/Software.vue";

let isDataTableLoading = ref(false);

const toast = useToast();

import DeviceConfiguration from "./DeviceConfiguration.vue";
const selectedDevice = ref({});
const shouldShowInstruction = ref(false);

let devices = reactive([]);

const statuses = ref(["ONLINE", "OFFLINE"]);

const filters = ref({
  deviceId: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
  deviceModel: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
  softwareVersion: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
  status: { value: null, matchMode: FilterMatchMode.EQUALS },
});

watch(selectedDevice, (newVal, oldVal) => {
  if (newVal == null) {
    selectedDevice.value = oldVal;
  }
});

const getDevices = async () => {
  isDataTableLoading.value = true;

  try {
    const response = await apiClient.get("/device");
    if (response.data.length == 0) {
      shouldShowInstruction.value = true;
      return;
    }
    devices = response.data.map(mapDeviceData);
    try {
      await Promise.all(
        devices.map(async (device) => {
          device.status = await getDeviceStatus(device.deviceId);
        })
      );

      toast.add({
        severity: "success",
        summary: "Success",
        detail: "Devices fetched successfully",
        life: 3000,
      });
    } catch (error) {
      toast.add({
        severity: "error",
        summary: "Error",
        detail: "Failed to fetch devices status",
        life: 3000,
      });
    }
  } catch (error) {
    console.log(error);
    toast.add({
      severity: "error",
      summary: "Error",
      detail: "Failed to fetch devices",
      life: 3000,
    });
  } finally {
    isDataTableLoading.value = false;
  }
};

const getDeviceStatus = async (deviceId) => {
  const response = await apiClient.get(`/device/${deviceId}/status`);
  return response.data.status;
};

onMounted(async () => {
  await getDevices();
  selectedDevice.value = devices[0];
});

const getSeverity = (status) => {
  switch (status) {
    case "ONLINE":
      return "success";

    case "OFFLINE":
      return "danger";

    default:
      return null;
  }
};

function mapDeviceData(device) {
  return {
    deviceId: device.device_id,
    softwareVersion: device.software_version,
    deviceModel: device.device_model,
    deviceBoardType: device.device_board_type,
    status: device.status,
    codeBaseUrl: device.code_base_url,
  };
}

const handleSetupFinished = () => {
  shouldShowInstruction.value = false;
  getDevices();
};
</script>

<style scoped>
::v-deep .p-tabview-nav .p-tabview-nav-link {
  border-radius: 0 !important;
  background: none !important;
}
</style>
