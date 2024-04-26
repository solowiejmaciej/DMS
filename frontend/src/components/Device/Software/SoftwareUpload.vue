<template>
  <div class="card">
    <Toast />
    <FileUpload
      customUpload
      @uploader="uploadHandler"
      name="firmware.bin"
      accept=".bin"
      :maxFileSize="1000000"
      :multiple="false"
      @select="onSelectedFiles"
    >
      <template
        #header="{ chooseCallback, uploadCallback, clearCallback, files }"
      >
        <div
          class="flex flex-wrap justify-content-between align-items-center flex-1 gap-2"
        >
          <div class="flex gap-2">
            <Button
              @click="chooseCallback()"
              icon="pi pi-file"
              label="Choose"
            ></Button>
            <Button
              label="Upload"
              @click="uploadCallback"
              icon="pi pi-cloud-upload"
              severity="success"
              :disabled="!files || files.length === 0"
            ></Button>
            <Button
              label="Clear"
              @click="clearCallback()"
              icon="pi pi-times"
              severity="danger"
              :disabled="!files || files.length === 0"
            ></Button>
          </div>
        </div>
      </template>
      <template
        #content="{
          files,
          uploadedFiles,
          removeUploadedFileCallback,
          removeFileCallback,
        }"
      >
        <div v-if="files.length > 0">
          <div class="flex flex-wrap p-0 sm:p-5 gap-5">
            <div
              v-for="(file, index) of files"
              :key="file.name + file.type + file.size"
              class="card m-0 px-6 flex flex-column border-1 surface-border align-items-center gap-3"
            >
              <span class="font-semibold">{{ file.name }}</span>
              <div>{{ formatSize(file.size) }}</div>
              <Badge value="Pending" severity="warning" />
            </div>
          </div>
        </div>

        <div v-if="uploadedFiles.length > 0">
          <div class="flex flex-wrap p-0 sm:p-5 gap-5">
            <div
              v-for="(file, index) of uploadedFiles"
              :key="file.name + file.type + file.size"
              class="card m-0 px-6 flex flex-column border-1 surface-border align-items-center gap-3"
            >
              <span class="font-semibold">{{ file.name }}</span>
              <div>{{ formatSize(file.size) }}</div>
              <Badge value="Completed" class="mt-3" severity="success" />
            </div>
          </div>
        </div>
      </template>
      <template #empty>
        <div class="flex align-items-center justify-content-center flex-column">
          <i
            class="pi pi-cloud-upload border-2 border-circle p-5 text-8xl text-400 border-400"
          />
          <p class="mt-4 mb-0">Drag and drop files to here to upload.</p>
        </div>
      </template>
    </FileUpload>
  </div>
</template>

<script setup>
import { ref } from "vue";
import { usePrimeVue } from "primevue/config";
import { useToast } from "primevue/usetoast";

let emit = defineEmits(["uploadFinished"]);

const $primevue = usePrimeVue();
const toast = useToast();

const totalSize = ref(0);
const files = ref([]);

const onRemoveTemplatingFile = (file, removeFileCallback, index) => {
  removeFileCallback(index);
  totalSize.value -= parseInt(formatSize(file.size));
  totalSizePercent.value = totalSize.value / 10;
};

const onClearTemplatingUpload = (clear) => {
  clear();
  totalSize.value = 0;
  totalSizePercent.value = 0;
};

const onSelectedFiles = (event) => {
  files.value = [];

  if (event.files.length > 0) {
    files.value = [event.files[0]];
    totalSize.value = parseInt(formatSize(files.value[0].size));
  } else {
    totalSize.value = 0;
  }
};

const formatSize = (bytes) => {
  const k = 1024;
  const dm = 3;
  const sizes = $primevue.config.locale.fileSizeTypes;

  if (bytes === 0) {
    return `0 ${sizes[0]}`;
  }

  const i = Math.floor(Math.log(bytes) / Math.log(k));
  const formattedSize = parseFloat((bytes / Math.pow(k, i)).toFixed(dm));

  return `${formattedSize} ${sizes[i]}`;
};

const uploadHandler = (event) => {
  toast.add({
    severity: "info",
    summary: "Success",
    detail: "File Uploaded",
    life: 3000,
  });
  emit("uploadFinished");
};
</script>
