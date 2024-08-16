<template>
    <Marker :options="{ position: marker.position, label: marker.label, title: marker.title }" :clickable="true"
        :draggable="true" @click="center = marker.position">
    </Marker>
</template>

<script>
import { Marker } from 'vue3-google-map';

export default {
    name: 'TrackMarker',
    props: ['marker'],
    components: { Marker },
}
</script>