<template>
  <div class="card flex justify-content-cente mt-2">
    <Stepper linear v-model:activeStep="active">
      <StepperPanel header="Download">
        <template #content="{ prevCallback, nextCallback }">
          <div class="flex flex-column h-12rem">
            <div class="justify-content-center align-items-center font-medium">
              <h3>Right now you don't have any devices setup</h3>
              <p>Follow the instructions to add a new device</p>
              <p>Download the DMS-ESP32 repository</p>
              <Button
                icon="pi pi-github"
                label="Download"
                @click="openGithubRepo"
              >
              </Button>
            </div>
          </div>
          <div class="flex pt-4 justify-end">
            <Button
              label="Next"
              icon="pi pi-arrow-right"
              iconPos="right"
              @click="nextCallback"
            />
          </div>
        </template>
      </StepperPanel>
      <StepperPanel header="Install">
        <template #content="{ prevCallback, nextCallback }">
          <div class="flex flex-column h-12rem">
            <div class="justify-content-center align-items-center font-medium">
              <h3>Install the library in your project</h3>
            </div>
          </div>
          <div class="flex pt-4 justify-between">
            <Button
              label="Back"
              severity="secondary"
              icon="pi pi-arrow-left"
              @click="prevCallback"
            />
            <Button
              label="Next"
              icon="pi pi-arrow-right"
              iconPos="right"
              @click="nextCallback"
            />
          </div>
        </template>
      </StepperPanel>
      <StepperPanel header="Configure">
        <template #content="{ prevCallback, nextCallback }">
          <div class="flex flex-column h-12rem">
            <div class="justify-content-center align-items-center font-medium">
              <h3>Configure the library</h3>
              <p>Paste the code you see bellow to your project</p>
              <p>Remember to replace variables with actuall values</p>
            </div>
          </div>
          <pre v-html="highlightedCode"></pre>
          <Button label="Copy" icon="pi pi-copy" text @click="copyToClipboard">
          </Button>

          <p>Upload the firmware to the device and verify that it's working</p>
          <div class="flex pt-4 justify-between">
            <Button
              label="Back"
              severity="secondary"
              icon="pi pi-arrow-left"
              @click="prevCallback"
            />
            <Button
              label="Finish"
              icon="pi pi-arrow-right"
              iconPos="right"
              @click="finish"
            />
          </div>
        </template>
      </StepperPanel>
    </Stepper>
  </div>
</template>

<script setup>
import { ref, onMounted, defineEmits } from "vue";
import hljs from "highlight.js";
import "../../../node_modules/highlight.js/styles/atom-one-dark.css";
import cpp from "highlight.js/lib/languages/cpp";

let emit = defineEmits(["setupFinished"]);

hljs.registerLanguage("cpp", cpp);

const openGithubRepo = () => {
  window.open("https://github.com/solowiejmaciej/DMS-ESP32", "_blank");
};

let code = `
#include <DMSManager.h>

const char *DMS_URL = "https://dms.solowiejmaciej.com";
const char *DMS_API_KEY = "yourkey";
const char *SSID = "YourSSID";
const char *SSID_PASSWORD = "YourPassword";

DMSManager dmsManager;

void setup()
{
  Serial.begin(115200);
  delay(100);
  dmsManager.setup(DMSOptions{
      DMS_URL,
      DMS_API_KEY,
      SSID,
      SSID_PASSWORD,
  });
  Serial.println("Device is ready!");
  //Rest of your setup code
}
void loop()
{
  dmsManager.loop();
  //Rest of your loop code
}

//Your normal code
`;

let highlightedCode = ref("");

onMounted(() => {
  highlightedCode.value = hljs.highlight("cpp", code).value;
});

const copyToClipboard = () => {
  navigator.clipboard.writeText(code);
};

const finish = () => {
  emit("setupFinished");
};
</script>
<style scoped>
.p-stepper {
  flex-basis: 50rem;
}
</style>
