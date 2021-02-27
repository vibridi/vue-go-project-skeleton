import Vue from 'vue'
import App from "@/App.vue";
import router from '@/router/router'
import store from '@/store/store'

// Install directives
// Vue.directive('foo', Foo)

// Install plugins
// Vue.use(MyPlugin, {...})

// Install filters
// Vue.filter('bar', Bar)

new Vue({
    el: '#app',
    router,
    store,
    render: h => h(App)
})