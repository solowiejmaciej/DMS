<template>
  <div class="flex">
    <div class="w-1/2 overflow-auto card mt-2">
      <div class="items-right justify-end flex mb-2">
        <Button
          icon="pi pi-refresh"
          text
          :loading="isDataTableLoading"
          raised
          rounded
          aria-label="Filter"
          v-tooltip="'Refresh table'"
          @click="handleRefreshData"
        />
      </div>
      <div
        class="absolute inset-0 flex items-center justify-center"
        v-if="isDataTableLoading"
      >
        <ProgressSpinner
          style="width: 50px; height: 50px"
          strokeWidth="8"
          fill="var(--surface-ground)"
          animationDuration="1.5s"
        />
      </div>
      <DataTable
        v-else
        stateStorage="session"
        stateKey="dt-state-demo-session"
        resizableColumns
        columnResizeMode="fit"
        showGridlines
        scrollable
        scrollHeight="50rem"
        :value="devices"
        tableStyle="min-width: 50rem text-center "
        selectionMode="single"
        dataKey="deviceId"
        v-model:selection="selectedDevice"
        filterDisplay="row"
        v-model:filters="filters"
      >
        <Column header="#" headerStyle="width:3rem " :showFilterMenu="false">
          <template #body="slotProps">
            {{ slotProps.index + 1 }}
          </template>
        </Column>
        <Column
          field="deviceId"
          header="Id"
          :showFilterMenu="false"
          class="text-center"
        >
          <template #filter="{ filterModel, filterCallback }">
            <InputText
              v-model="filterModel.value"
              type="text"
              @input="filterCallback()"
              class="p-column-filter"
              placeholder="Search"
            /> </template
        ></Column>
        <Column
          field="deviceModel"
          header="Name"
          :showFilterMenu="false"
          class="text-center"
        >
          <template #filter="{ filterModel, filterCallback }">
            <InputText
              v-model="filterModel.value"
              type="text"
              @input="filterCallback()"
              class="p-column-filter"
              placeholder="Search"
            /> </template
        ></Column>

        <Column
          field="softwareVersion"
          header="Software version"
          :showFilterMenu="false"
          class="text-center"
        >
          <template #filter="{ filterModel, filterCallback }">
            <InputText
              v-model="filterModel.value"
              type="text"
              @input="filterCallback()"
              class="p-column-filter w-full"
              placeholder="Search"
            />
          </template>
        </Column>
        <Column
          field="status"
          header="Device status"
          :showFilterMenu="false"
          class="text-center"
        >
          <template #filter="{ filterModel, filterCallback }">
            <Dropdown
              v-model="filterModel.value"
              @change="filterCallback()"
              :options="statuses"
              placeholder="Select One"
              class="p-column-filter"
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
    <div class="w-1/2 flex items-center justify-center">
      <DeviceConfiguration
        v-if="selectedDevice.deviceId"
        :device="selectedDevice"
      />
    </div>
  </div>
</template>

<script setup>
import { ref, watch, reactive, onMounted } from "vue";
import { FilterMatchMode } from "primevue/api";
import { useToast } from "primevue/usetoast";
import apiClient from "../../Api/apiClient";

let isDataTableLoading = ref(false);

const toast = useToast();

import DeviceConfiguration from "./DeviceConfiguration.vue";
const selectedDevice = ref({});

let devices = reactive([]);

const statuses = ref(["ONLINE", "OFFLINE"]);

const filters = ref({
  global: { value: null, matchMode: FilterMatchMode.CONTAINS },
  deviceId: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
  deviceModel: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
  softwareVersion: { value: null, matchMode: FilterMatchMode.STARTS_WITH },
  status: { value: null, matchMode: FilterMatchMode.EQUALS },
});

const getDevices = () => {
  apiClient
    .get("/device")
    .then((response) => {
      devices = response.data.map(mapDeviceData);
      isDataTableLoading.value = true;
    })
    .catch((error) => {
      toast.add({
        severity: "error",
        summary: "Error",
        detail: "An error occurred while fetching devices",
        life: 3000,
      });
    })
    .finally(() => {
      isDataTableLoading.value = false;
    });
};

onMounted(() => {
  getDevices();
});

const handleRefreshData = () => {
  toast.add({
    severity: "success",
    summary: "Data refreshed",
    detail: "Data has been refreshed",
    life: 3000,
  });
  getDevices();
};

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
</script>
