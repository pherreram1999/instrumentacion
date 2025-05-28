<script lang="ts" setup>
import {GetSerialPorts,OpenPort} from "../wailsjs/go/main/App";
import {computed, onBeforeMount, onMounted, ref, shallowRef, watch} from "vue";
import {EventsOn} from "../wailsjs/runtime";
import Chart from 'chart.js/auto'
import annotationPlugin from 'chartjs-plugin-annotation';
Chart.register(annotationPlugin);


// en 5 volts 2000 gramos newtow

const LIMIT_SUP = 2_000
const VOLT_REF = 5

const portName = ref()
const serialPorts = shallowRef<string[]>()
const voltageVal = ref<number>(0)

onMounted(()=>{
    EventsOn('voltage',voltage => {
        voltageVal.value = voltage
        //voltageVal.value = (voltage * 2000) / 5
    })
})

const convertedValue = computed(() => ((voltageVal.value * LIMIT_SUP) / VOLT_REF ))



const connect = () => {
    OpenPort(portName.value)
}


const numf = new Intl.NumberFormat('es-MX')

onBeforeMount(async () => {
  serialPorts.value = await GetSerialPorts()
})


const gaussRef = ref()

let chart: Chart;

watch(voltageVal,v => {
    if(!chart) return
    chart.data.datasets[0].data = [convertedValue.value,LIMIT_SUP-convertedValue.value]
    chart.update()
})

onMounted(()=>{

    const annotation = {
        type: 'doughnutLabel',
        content: ({chart}: any) => [
            chart.data.datasets[0].data[0].toFixed(2) + ' Kg/Nw',
            'Presión',
        ],
        drawTime: 'beforeDraw',
        position: {
            y: '-50%'
        },
        font: [{size: 50, weight: 'bold'}, {size: 20}],
        color: ({chart}: any) => ['#1e40af', 'black']
    }

    chart = new Chart(gaussRef.value,{
        type: 'doughnut',
        data: {
            datasets: [
                {
                    data: [convertedValue.value,100-convertedValue.value],
                    backgroundColor(ctx, options) {
                        return ['#1e40af','#f1f5f9']
                    },
                }
            ]
        },
        options: {
            aspectRatio: 4,
            circumference: 180,
            rotation: -90,
            plugins: {
                annotation: {
                    // @ts-ignore
                    annotations: {
                        annotation
                    }
                }
            }
        }
    })
})


</script>

<template>
    <div>
        <div class="h-full flex justify-between items-center bg-white shadow">
            <div class="flex gap-4 p-4 items-center">
                <label for="port">Puerto</label>
                <select v-model="portName" class="h-8 rounded bg-none border border-stone-300" name="port" id="prot">
                    <option :value="s" v-for="s of  serialPorts">
                        {{ s }}
                    </option>
                </select>
                <button class="px-4 py-2 bg-blue-800 text-white rounded" @click.prevent="connect">
                    conectar
                </button>
            </div>
            <div class="px-4">
                <h1 class="text-blue-800 text-2xl font-semibold">Equipo 6</h1>
            </div>
        </div>
        <div class="container mx-auto">
            <div class="py-4">
                <div class="grid grid-cols-2 bg-white rounded p-4 shadow-lg">
                    <div>
                        <span class="text-stone-500">Entrada de voltaje</span>
                        <h1 class="text-blue-800 font-semibold text-4xl">{{ numf.format(voltageVal) }}</h1>
                    </div>
                    <div>
                        <span class="text-stone-500">Conversión</span>
                        <h1 class="text-blue-800 font-semibold text-4xl">{{ numf.format(convertedValue) }} Kg/Newtown</h1>
                    </div>
                </div>

                <div class="bg-white p-4 mt-5 rounded-md shadow">
                    <canvas ref="gaussRef"></canvas>
                </div>

            </div>
        </div>
    </div>

</template>

<style>
</style>
