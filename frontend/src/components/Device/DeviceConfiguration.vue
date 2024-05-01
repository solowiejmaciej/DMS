<template>
  <Toast></Toast>
  <ConfirmPopup></ConfirmPopup>
  <Sidebar v-model:visible="shouldShowHelp" header="Configuration">
    <p>
      We recommend using
      <a
        href="https://www.hivemq.com/products/mqtt-cloud-broker/"
        target="_blank"
        >HiveMQ</a
      >
      for your MQTT broker, as it is the platform we have tested extensively
      with our devices. HiveMQ is a MQTT broker, that has generous free tier.
      It's easy to set up and use, they provide excellent documentation to help
      you get started.
      <br />
      <br />
      To configure MQTT, you will need to set up your broker (HiveMQ) and
      configure your device with the details of the broker. Please refer to the
      <a href="https://docs.hivemq.com/hivemq-cloud/index.html" target="_blank"
        >HiveMQ documentation</a
      >
      for detailed instructions on how to do this.
    </p>
  </Sidebar>
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
      <div class="flex justify-end">
        <Button text @click="showHelp">
          <i class="pi pi-question-circle"></i>
        </Button>
      </div>
      <div class="flex justify-between">
        <div
          class="flex flex-col gap-2 mx-auto"
          :style="{ 'min-height': '16rem', 'min-width': minWidth }"
        >
          <div class="field p-fluid m-2">
            <IconField>
              <FloatLabel>
                <InputText
                  id="deviceBoardType"
                  v-model="deviceBoardType"
                  v-tooltip.left="'For example ESP32'"
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
                <InputNumber
                  :useGrouping="false"
                  id="internalLedPin"
                  v-model="internalLedPin"
                  v-tooltip.left="'Pin number for internal led'"
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
                <InputNumber
                  :useGrouping="false"
                  id="aliveInterval"
                  v-model="aliveInterval"
                  v-tooltip.left="
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
                  v-tooltip.left="'Url to the codebase of the device'"
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
                  v-tooltip.left="'When device was added to the system'"
                />
                <label for="createdAt">Created at</label>
              </FloatLabel>
              <InputIcon>
                <i class="pi pi-calendar" />
              </InputIcon>
            </IconField>
          </div>
        </div>
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
                  v-tooltip.left="'Url to the MQTT broker'"
                />
                <label for="mqttBrokerHost">Broker url</label>
              </FloatLabel>
            </IconField>
          </div>
          <div class="field p-fluid m-2">
            <IconField>
              <FloatLabel>
                <InputNumber
                  :useGrouping="false"
                  id="mqttBrokerPort"
                  v-model="mqttBrokerPort"
                  v-tooltip.left="'Port number of the MQTT broker'"
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
                  v-tooltip.left="'Username for the MQTT broker'"
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
                  v-tooltip.left="'Password for the MQTT broker'"
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
      </div>

      <div class="flex justify-between mt-4">
        <div class="field p-fluid m-2">
          <Button
            v-tooltip="'Delete device and configuration from the system'"
            @click="confirmDeleteDevice($event)"
            label="Delete"
            severity="danger"
            icon="pi pi-trash"
          ></Button>
        </div>
        <div class="field p-fluid m-2">
          <Button
            v-tooltip="'Reboot the device'"
            :loading="shouldDisableRebootButton"
            type="button"
            label="Reboot"
            @click="rebootDevice"
            icon="pi pi-power-off"
            v-if="props.device.status == 'ONLINE'"
          >
          </Button>
          <Button
            v-else
            v-tooltip="'Reboot the device'"
            :loading="shouldDisableRebootButton"
            type="button"
            label="Reboot"
            @click="confirmDeviceReboot($event)"
            icon="pi pi-power-off"
          >
          </Button>
        </div>
        <div class="field p-fluid m-2">
          <Button
            type="button"
            label="Cancel"
            @click="handleCancel"
            icon="pi pi-times
"
          />
        </div>
        <div class="field p-fluid m-2">
          <Button
            icon="pi pi-check"
            type="button"
            label="Save"
            @click="saveData"
          />
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { onMounted, ref, watch } from "vue";
import { useToast } from "primevue/usetoast";
import { useConfirm } from "primevue/useconfirm";
import apiClient from "../../Api/apiClient";

const emit = defineEmits(["refresh-data"]);

const toast = useToast();
const confirm = useConfirm();

let shouldDisableTestButton = ref(false);
let shouldDisableRebootButton = ref(false);
let shouldShowHelp = ref(false);

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
  if (deviceId == null) {
    return;
  }
  try {
    const response = await apiClient.get("/configuration/" + deviceId);
    const configuration = response.data;
    deviceBoardType.value = props.device.deviceBoardType;
    internalLedPin.value = configuration.internal_led_pin;
    aliveInterval.value = configuration.alive_interval;
    createdAt.value = configuration.created_at;
    deviceCodebaseUrl.value = props.device.codeBaseUrl;
    mqttBrokerHost.value = configuration.mqtt_broker_host;
    mqttBrokerPort.value = configuration.mqtt_broker_port;
    mqttUsername.value = configuration.mqtt_username;
    mqttPassword.value = configuration.mqtt_password;
  } catch (error) {
    toast.add({
      severity: "error",
      summary: "Error",
      detail: "Failed to fetch configuration for device " + deviceId,
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
        summary: "Reboot message sent",
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
      apiClient
        .delete("/device/" + props.device.deviceId)
        .then(() => {
          toast.add({
            severity: "success",
            summary: "Success",
            detail: "Device has been deleted",
            life: 3000,
          });
          emit("refresh-data");
        })
        .catch((error) => {
          toast.add({
            severity: "error",
            summary: "Error",
            detail: "Failed to delete device",
            life: 3000,
          });
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

const confirmDeviceReboot = (event) => {
  confirm.require({
    target: event.currentTarget,
    message:
      "Device you are about to reboot is currently OFFLINE. Action will not be executed if device is not ONLINE. Are you sure you want to reboot this device?",
    icon: "pi pi-info-circle",
    rejectClass: "p-button-secondary p-button-outlined p-button-sm",
    acceptClass: "p-button-danger p-button-sm",
    rejectLabel: "Cancel",
    acceptLabel: "Reboot",
    accept: () => {
      rebootDevice();
    },
    reject: () => {
      toast.add({
        severity: "error",
        summary: "Rejected",
        detail: "Device has not been rebooted",
        life: 3000,
      });
    },
  });
};

const handleCancel = () => {
  getDeviceConfiguration(props.device.deviceId);
};

const saveData = async () => {
  const body = JSON.stringify({
    internal_led_pin: internalLedPin.value,
    alive_interval: aliveInterval.value,
    mqtt_broker_host: mqttBrokerHost.value,
    mqtt_broker_port: mqttBrokerPort.value,
    mqtt_username: mqttUsername.value,
    mqtt_password: mqttPassword.value,
  });

  try {
    let response = await apiClient.put(
      "/configuration/" + props.device.deviceId,
      body
    );

    if (response.status === 200) {
      toast.add({
        severity: "success",
        summary: "Success",
        detail: "Device configuration saved",
        life: 3000,
      });
    } else {
      toast.add({
        severity: "error",
        summary: "Error",
        detail: "Failed to save device configuration",
        life: 3000,
      });
    }
  } catch (error) {
    toast.add({
      severity: "error",
      summary: "Error",
      detail: "Failed to save device configuration",
      life: 3000,
    });
    console.error(error);
  }
};

const showHelp = () => {
  shouldShowHelp.value = true;
};

const deviceBoardType = ref("");
const internalLedPin = ref(null);
const aliveInterval = ref(null);
const createdAt = ref("");
const deviceCodebaseUrl = ref("");

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
