<script lang="ts" setup>
import HelloWorld from './components/HelloWorld.vue'
import {GetSerialPorts,OpenPort} from "../wailsjs/go/main/App";
import {onBeforeMount, onMounted, ref, shallowRef} from "vue";
import {EventsOn} from "../wailsjs/runtime";


// en 5 volts 2000 gramos newtow

const portName = ref()
const serialPorts = shallowRef<string[]>()
const voltageVal = ref()

onMounted(()=>{
    EventsOn('voltage',voltage => {

        voltageVal.value = (voltage * 2000) / 5
    })
})

const connect = () => {
    OpenPort(portName.value)
}


onBeforeMount(async () => {
  serialPorts.value = await GetSerialPorts()
})

</script>

<template>
  <div>
      <label for="port">Puerto</label>
      <select v-model="portName" name="port" id="prot">
          <option :value="s" v-for="s of  serialPorts">
              {{ s }}
          </option>
      </select>
      <button @click.prevent="connect">
          conectar
      </button>
      <h1>{{ voltageVal }} Kg/Newtown</h1>
  </div>
</template>

<style>
#logo {
  display: block;
  width: 50%;
  height: 50%;
  margin: auto;
  padding: 10% 0 0;
  background-position: center;
  background-repeat: no-repeat;
  background-size: 100% 100%;
  background-origin: content-box;
}
</style>
