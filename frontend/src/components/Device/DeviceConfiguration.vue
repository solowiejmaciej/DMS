<template>
  <Toast></Toast>
  <ConfirmPopup></ConfirmPopup>
  <div class="flex items-center justify-center">
    <div
      class="absolute inset-0 flex items-center justify-center"
      v-if="showSpinner"
    >
      <ProgressSpinner
        style="width: 50px; height: 50px"
        strokeWidth="8"
        fill="var(--surface-ground)"
        animationDuration="1.5s"
      />
    </div>
    <div v-else>
      <h2 class="flex items-center justify-center">
        <i class="pi pi-microchip mr-2" style="font-size: 1.5rem"></i>
        Device: {{ props.device.deviceId }}
      </h2>
      <div class="flex justify-between">
        <Fieldset legend="Device">
          <div
            class="flex flex-col gap-2 mx-auto"
            :style="{ 'min-height': '16rem', 'min-width': minWidth }"
          >
            <div class="field p-fluid m-2">
              <IconField>
                <FloatLabel>
                  <InputText
                    id="deviceModel"
                    v-model="deviceModel"
                    v-tooltip="'For example SmartWeatherStation'"
                  />
                  <label for="deviceModel">Model</label>
                </FloatLabel>
                <InputIcon>
                  <InputIcon>
                    <i class="pi pi-hashtag" />
                  </InputIcon>
                </InputIcon>
              </IconField>
            </div>
            <div class="field p-fluid m-2">
              <IconField>
                <FloatLabel>
                  <InputText
                    id="deviceBoardType"
                    v-model="deviceBoardType"
                    v-tooltip="'For example ESP32'"
                  />
                  <label for="deviceBoardType">Board type</label>
                </FloatLabel>
                <InputIcon>
                  <InputIcon>
                    <i class="pi pi-box" />
                  </InputIcon>
                </InputIcon>
              </IconField>
            </div>
            <div class="field p-fluid m-2">
              <IconField>
                <FloatLabel>
                  <InputText
                    id="internalLedPin"
                    v-model="internalLedPin"
                    v-tooltip="'Pin number for internal led'"
                  />
                  <label for="internalLedPin">Internal led pin</label>
                </FloatLabel>
                <InputIcon>
                  <i class="pi pi-lightbulb" />
                </InputIcon>
              </IconField>
            </div>
            <div class="field p-fluid m-2">
              <IconField>
                <FloatLabel>
                  <InputText
                    id="aliveInterval"
                    v-model="aliveInterval"
                    v-tooltip="
                      'How often device sends alive message in milliseconds. For example 60000 for 1 minute.'
                    "
                  />
                  <label for="aliveInterval">Alive interval</label>
                </FloatLabel>
                <InputIcon>
                  <i class="pi pi-hourglass" />
                </InputIcon>
              </IconField>
            </div>
            <div class="field p-fluid m-2">
              <IconField>
                <FloatLabel>
                  <InputText
                    id="deviceCodebaseUrl"
                    v-model="deviceCodebaseUrl"
                    v-tooltip="'Url to the codebase of the device'"
                  />
                  <label for="deviceCodebaseUrl">Codebase url</label>
                </FloatLabel>
                <InputIcon>
                  <i class="pi pi-github" />
                </InputIcon>
              </IconField>
            </div>
            <div class="field p-fluid m-2">
              <IconField>
                <FloatLabel>
                  <InputText
                    id="createdAt"
                    v-model="createdAt"
                    disabled
                    v-tooltip="'When device was added to the system'"
                  />
                  <label for="createdAt">Created at</label>
                </FloatLabel>
                <InputIcon>
                  <i class="pi pi-calendar" />
                </InputIcon>
              </IconField>
            </div>
            <div class="field p-fluid m-2">
              <Button
                :loading="shouldDisableRebootButton"
                type="button"
                label="Reboot"
                @click="rebootDevice"
                icon="pi pi-power-off"
              >
              </Button>
            </div>
            <div class="field p-fluid m-2">
              <Button
                @click="confirmDeleteDevice($event)"
                label="Delete"
                severity="danger"
                icon="pi pi-trash"
              ></Button>
            </div>
          </div>
        </Fieldset>
        <Fieldset legend="MQTT">
          <div
            class="flex flex-col gap-2 mx-auto"
            :style="{
              'min-height': '16rem',
              'min-width': minWidth,
              'padding-right': 0,
            }"
          >
            <div class="field p-fluid m-2">
              <IconField>
                <FloatLabel>
                  <InputText
                    id="mqttBrokerHost"
                    v-model="mqttBrokerHost"
                    v-tooltip="'Url to the MQTT broker'"
                  />
                  <label for="mqttBrokerHost">Broker url</label>
                </FloatLabel>
                <InputIcon> </InputIcon>
              </IconField>
            </div>
            <div class="field p-fluid m-2">
              <IconField>
                <FloatLabel>
                  <InputText
                    id="mqttBrokerPort"
                    v-model="mqttBrokerPort"
                    v-tooltip="'Port number of the MQTT broker'"
                  />
                  <label for="mqttBrokerPort">Port</label>
                </FloatLabel>
                <InputIcon>
                  <InputIcon>
                    <i class="pi pi-compass" />
                  </InputIcon>
                </InputIcon>
              </IconField>
            </div>
            <div class="field p-fluid m-2">
              <IconField>
                <FloatLabel>
                  <InputText
                    id="mqttUsername"
                    v-model="mqttUsername"
                    v-tooltip="'Username for the MQTT broker'"
                  />
                  <label for="mqttUsername">Username</label>
                </FloatLabel>
                <InputIcon>
                  <i class="pi pi-user" />
                </InputIcon>
              </IconField>
            </div>
            <div class="field p-fluid m-2">
              <IconField>
                <FloatLabel>
                  <InputText
                    id="mqttPassword"
                    v-model="mqttPassword"
                    v-tooltip="'Password for the MQTT broker'"
                  />
                  <label for="mqttPassword">Password</label>
                </FloatLabel>
                <InputIcon>
                  <i class="pi pi-key" />
                </InputIcon>
              </IconField>
            </div>
            <div class="field p-fluid m-2">
              <Button
                :loading="shouldDisableTestButton"
                type="button"
                label="Test MQTT connection"
                @click="testMqttConnection"
                icon="pi pi-wifi"
              >
              </Button>
            </div>
          </div>
        </Fieldset>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref, watch } from "vue";
import { useToast } from "primevue/usetoast";
import { useConfirm } from "primevue/useconfirm";
import apiClient from "../../Api/apiClient";

const toast = useToast();
const confirm = useConfirm();

let shouldDisableTestButton = ref(false);
let shouldDisableRebootButton = ref(false);

const minWidth = ref("20rem");

const props = defineProps({
  device: {
    required: true,
  },
});

let showSpinner = ref(false);

onMounted(async () => {
  if (window.innerWidth >= 1500) {
    minWidth.value = "25rem";
  }
  await getDeviceConfiguration(props.device.deviceId);
});

watch(
  () => props.device,
  async (newVal, oldVal) => {
    await getDeviceConfiguration(newVal.deviceId);
  }
);

const getDeviceConfiguration = async (deviceId) => {
  showSpinner.value = true;
  try {
    const response = await apiClient.get("/configuration/" + deviceId);
    const configuration = response.data;
    deviceModel.value = props.device.deviceModel;
    deviceBoardType.value = props.device.deviceBoardType;
    internalLedPin.value = configuration.InternalLedPin;
    aliveInterval.value = configuration.AliveInterval;
    createdAt.value = configuration.CreatedAt;
    deviceCodebaseUrl.value = props.device.codeBaseUrl;
    softwareVersion.value = configuration.SoftwareVersion;
    mqttBrokerHost.value = configuration.MqttBrokerHost;
    mqttBrokerPort.value = configuration.MqttBrokerPort;
    mqttUsername.value = configuration.MqttUsername;
    mqttPassword.value = configuration.MqttPassword;
  } catch (error) {
    toast.add({
      severity: "error",
      summary: "Error",
      detail: "Failed to fetch device configuration",
      life: 3000,
    });
  } finally {
    showSpinner.value = false;
  }
};

const testMqttConnection = () => {
  shouldDisableTestButton.value = true;
  const mqttConfig = {
    mqtt_broker_host: mqttBrokerHost.value,
    mqtt_broker_port: mqttBrokerPort.value,
    mqtt_username: mqttUsername.value,
    mqtt_password: mqttPassword.value,
  };

  apiClient
    .post(
      "/device/" + props.device.deviceId + "/test-mqtt-connection",
      mqttConfig
    )
    .then((response) => {
      if (response.data.result) {
        toast.add({
          severity: "success",
          summary: "Success",
          detail: "MQTT connection successful",
          life: 3000,
        });
      } else {
        toast.add({
          severity: "error",
          summary: "Error",
          detail: response.data.error,
          life: 6000,
        });
      }
    })
    .catch((error) => {
      toast.add({
        severity: "error",
        summary: "Error",
        detail: "DMS error while testing MQTT connection",
        life: 6000,
      });
    })
    .finally(() => {
      shouldDisableTestButton.value = false;
    });
};
const rebootDevice = () => {
  shouldDisableRebootButton.value = true;
  apiClient
    .post("/device/" + props.device.deviceId + "/reboot")
    .then(() => {
      toast.add({
        severity: "info",
        summary: "Device rebooted",
        detail: "Device " + props.device.deviceId + " has been rebooted",
        life: 3000,
      });
    })
    .catch((error) => {
      toast.add({
        severity: "error",
        summary: "Error",
        detail: "Failed to reboot device",
        life: 3000,
      });
    })
    .finally(() => {
      shouldDisableRebootButton.value = false;
    });
};

const confirmDeleteDevice = (event) => {
  confirm.require({
    target: event.currentTarget,
    message:
      "Are you sure you want to delete this device? This action cannot be undone.",
    icon: "pi pi-info-circle",
    rejectClass: "p-button-secondary p-button-outlined p-button-sm",
    acceptClass: "p-button-danger p-button-sm",
    rejectLabel: "Cancel",
    acceptLabel: "Delete",
    accept: () => {
      toast.add({
        severity: "info",
        summary: "Device deleted",
        detail: "Device " + props.device.deviceId + " has been deleted",
        life: 3000,
      });
    },
    reject: () => {
      toast.add({
        severity: "error",
        summary: "Rejected",
        detail: "Device has not been deleted",
        life: 3000,
      });
    },
  });
};

const deviceModel = ref("");
const deviceBoardType = ref("");
const internalLedPin = ref(null);
const aliveInterval = ref(null);
const createdAt = ref(null);
const deviceCodebaseUrl = ref("");
const softwareVersion = ref("1.0.0");

const mqttBrokerHost = ref("");
const mqttBrokerPort = ref(null);
const mqttUsername = ref("");
const mqttPassword = ref("");
</script>

<style scoped>
.no-border-radius {
  border-radius: 0 !important;
}
</style>
