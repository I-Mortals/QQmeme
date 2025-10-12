import {createApp} from 'vue'
import App from './App.vue'
import './style.less';
import { clickOutside } from './utils/directives'
import { initializeStoreFromCache } from './store'
import { Environment } from '@wailsjs/runtime/runtime'

window.addEventListener('load', async () => {
  const env = await Environment()
  window.platform = env.platform
  console.log('env: ', env)
})


const app = createApp(App)

app.directive('clickOutside', clickOutside)

initializeStoreFromCache()
app.mount('#app')

