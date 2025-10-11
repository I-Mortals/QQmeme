import {createApp} from 'vue'
import App from './App.vue'
import './style.less';
import { clickOutside } from './utils/directives'
import { initializeStoreFromCache } from './store'

const app = createApp(App)

app.directive('clickOutside', clickOutside)

initializeStoreFromCache()
app.mount('#app')
