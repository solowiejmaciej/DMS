<template>
  <Dialog
    modal
    :draggable="false"
    @hide="closeDialog"
    dismissableMask
    v-model:visible="isDialogVisible"
    :header="dialogHeader"
    :pt="{
      root: 'border-none',
      mask: {
        style: 'backdrop-filter: blur(2px)',
      },
    }"
  >
    <ManageSoftware v-if="showManageSoftwareDialog"></ManageSoftware>
    <SoftwareUpload
      v-if="showSoftwareUploadDialog"
      @upload-finished="closeDialog"
    ></SoftwareUpload>
  </Dialog>
  <div legend="Software">
    <div
      class="flex flex-col gap-2 mx-auto"
      style="min-height: 16rem; max-width: 20rem"
    >
      <div class="field p-fluid m-2">
        <Button
          @click="openUploadSoftwareDialog"
          type="button"
          label="Upload new software"
          icon="pi pi-upload"
          link
        >
        </Button>
      </div>
      <div class="field p-fluid m-2">
        <Button
          @click="openManageSoftwareDialog"
          type="button"
          label="Manage software updates"
          icon="pi pi-cog"
        >
        </Button>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from "vue";
import ManageSoftware from "./ManageSoftware.vue";
import SoftwareUpload from "./SoftwareUpload.vue";

const showSoftwareUploadDialog = ref(false);
const showManageSoftwareDialog = ref(false);
const isDialogVisible = ref(false);

const dialogHeader = ref("");

const closeDialog = () => {
  isDialogVisible.value = false;
  showSoftwareUploadDialog.value = false;
  showManageSoftwareDialog.value = false;
};

const openUploadSoftwareDialog = () => {
  isDialogVisible.value = true;
  showSoftwareUploadDialog.value = true;
  dialogHeader.value = "Upload new software";
};

const openManageSoftwareDialog = () => {
  isDialogVisible.value = true;
  showManageSoftwareDialog.value = true;
  dialogHeader.value = "Manage software updates";
};
</script>
