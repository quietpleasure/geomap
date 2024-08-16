<template>
    <Polyline :options="{path:path, strokeColor: '#FF0000', strokeOpacity: '1.0', strokeWeight: '3'}" ref="polyline" />
</template>

<script>
import { Polyline } from 'vue3-google-map';

export default {
    name: 'TrackRoute',
    components: {Polyline},
    props: ['path'],
}
</script>