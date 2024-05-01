<template>
  <div class="card flex justify-content-center">
    <Chart type="pie" :data="chartData" :options="chartOptions" />
  </div>
</template>

<script setup>
import { ref, onMounted } from "vue";

onMounted(() => {
  chartData.value = setChartData();
  chartOptions.value = setChartOptions();
});

const chartData = ref();
const chartOptions = ref();

const setChartData = () => {
  const documentStyle = getComputedStyle(document.body);

  return {
    labels: ["TEST1", "TEST2", "TEST3"],
    datasets: [
      {
        data: [540, 325, 702],
        backgroundColor: [
          documentStyle.getPropertyValue("--cyan-500"),
          documentStyle.getPropertyValue("--orange-500"),
          documentStyle.getPropertyValue("--gray-500"),
        ],
        hoverBackgroundColor: [
          documentStyle.getPropertyValue("--cyan-400"),
          documentStyle.getPropertyValue("--orange-400"),
          documentStyle.getPropertyValue("--gray-400"),
        ],
      },
    ],
  };
};

const setChartOptions = () => {
  const documentStyle = getComputedStyle(document.documentElement);
  const textColor = documentStyle.getPropertyValue("--text-color");

  return {
    plugins: {
      legend: {
        labels: {
          usePointStyle: true,
          color: textColor,
        },
      },
    },
  };
};
</script>
