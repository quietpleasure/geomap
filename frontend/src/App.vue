<template>
  <div id="app">
    <sui-container stackable>
      <sui-header as="h1" textAlign="center">Текущее положение транспортных средств</sui-header>
      <sui-grid>
        <sui-grid-column :width="4">
          <sui-form>
            <strong>Выбор автомобилей для точек</strong>
            <sui-dropdown selection fluid multiple v-model="track" placeholder="Выбор автомобилей" :options="tracks" />
            <br />
            <sui-button fluid v-on:click="getDeviceids">Показать</sui-button>
          </sui-form>
          <br />
          <sui-divider section />
          <br />
          <sui-form>
            <strong>Выбор автомобиля для маршрута</strong>
            <sui-dropdown selection fluid v-model="trackline" placeholder="Выбор автомобиля" :options="tracks" />
            <br />
            <sui-button fluid v-on:click="getDevice">Показать</sui-button>
          </sui-form>
        </sui-grid-column>
        <sui-grid-column :width="12">
          <GoogleMap :api-key="apikey" style="width: 100%; height: 500px" :center="center" :zoom="8">
            <TracksMap :markers="markers" />
            <TrackRoute :path="path"/>
          </GoogleMap>
        </sui-grid-column>
      </sui-grid>
    </sui-container>
  </div>
</template>

<script>
import BackendService from '@/api-service/backend.service';
import { GoogleMap } from 'vue3-google-map';
import TracksMap from '@/components/TracksMap.vue';
import TrackRoute from '@/components/TrackRoute.vue';

const center = { lat: 47.012271881103516, lng: 28.860593795776367 };
const apikey = 'YOUR-API-KEY-HERE';

export default {
  name: 'App',
  components: {
    TracksMap,
    GoogleMap,
    TrackRoute
  },
  data() {
    return {
      tracks: [],
      track: null,
      trackline: null,
      markers: [],
      center: center,
      apikey: apikey,
      path: []
    }
  },
  mounted() {
    BackendService.getAllTracks()
      .then((response) => {
        this.tracks = response.data;
      })
  },
  methods: {
    getDeviceids(event) {
      event.preventDefault();
      if (this.track) {
        if (this.track.length > 0) {
          this.path= []
          BackendService.showTracks(this.track)
            .then((response) => {
              this.markers = response.data
            })
            
        }
      }
    },
    getDevice(event) {
      event.preventDefault();
      if (this.trackline) {
        BackendService.showTrack(this.trackline)
          .then((response) => {
            this.path = response.data
            this.markers = [
              {
                label: "S",
                position: this.path[0]
              },
              {
                label: "F",
                position: this.path[response.data.length-1]
              }
            ]
          })
        this.trackline = null
      }
    }
  }
};
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  /* text-align: center; */
  color: #2c3e50;
  margin-top: 20px;
}
</style>
