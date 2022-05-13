import {createApp} from 'vue'
import App from './App.vue'
import installElementPlus from './plugins/element'

import router from './router'

router.beforeEach((to) => {
    document.title = `${to.meta.title} | NickyShell`;
})

const app = createApp(App)
installElementPlus(app)
app.use(router)
app.mount('#app')
