<script>

import TrackMarker from './TrackMarker.vue';
import { MarkerCluster } from 'vue3-google-map';

export default {
    name: "TracksMap",
    components: { MarkerCluster, TrackMarker },
    props: ['markers'],

}
</script>

<template>
    <MarkerCluster>
        <TrackMarker :key="marker.label" :marker="marker" v-for="marker in markers">
        </TrackMarker>
    </MarkerCluster>
</template>