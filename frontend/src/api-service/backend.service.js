import Axios from 'axios'

Axios.defaults.baseURL = 'http://127.0.0.1:5000/api';

const RESOURCE_TRACKS = 'tracks'
const RESOURCE_TRACK = 'track'
export default {
  getAllTracks () {
    return Axios.get(RESOURCE_TRACKS)
  },
  showTracks(data) {
    return Axios.post(RESOURCE_TRACKS, data)
  },
  showTrack(data) {
    return Axios.post(RESOURCE_TRACK, data)
  }
}